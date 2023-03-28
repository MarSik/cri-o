package hostport

import (
	"bytes"
	"net"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

	utiliptables "github.com/cri-o/cri-o/internal/iptables"
)

var _ = t.Describe("MetaHostportManager", func() {
	It("should succeed", func() {
		// ipv4
		iptables := newFakeIPTables()
		iptables.protocol = utiliptables.ProtocolIPv4
		// ipv6
		ip6tables := newFakeIPTables()
		ip6tables.protocol = utiliptables.ProtocolIPv6

		manager := metaHostportManager{
			ipv4HostportManager: &hostportManager{
				iptables: iptables,
			},
			ipv6HostportManager: &hostportManager{
				iptables: ip6tables,
			},
		}

		testCases := []struct {
			mapping     *PodPortMapping
			expectError bool
		}{
			{
				mapping: &PodPortMapping{
					Name:        "pod1",
					Namespace:   "ns1",
					IP:          net.ParseIP("192.168.2.7"),
					HostNetwork: false,
					PortMappings: []*PortMapping{
						{
							HostPort:      8080,
							ContainerPort: 80,
							Protocol:      v1.ProtocolTCP,
						},
						{
							HostPort:      8081,
							ContainerPort: 81,
							Protocol:      v1.ProtocolUDP,
						},
						{
							HostPort:      8083,
							ContainerPort: 83,
							Protocol:      v1.ProtocolSCTP,
						},
						{
							HostPort:      8084,
							ContainerPort: 84,
							Protocol:      v1.ProtocolTCP,
							HostIP:        "127.0.0.1",
						},
					},
				},
				expectError: false,
			},
			// same pod and portmappings,
			// but different IP must work
			{
				mapping: &PodPortMapping{
					Name:        "pod1",
					Namespace:   "ns1",
					IP:          net.ParseIP("2001:beef::3"),
					HostNetwork: false,
					PortMappings: []*PortMapping{
						{
							HostPort:      8080,
							ContainerPort: 80,
							Protocol:      v1.ProtocolTCP,
						},
						{
							HostPort:      8081,
							ContainerPort: 81,
							Protocol:      v1.ProtocolUDP,
						},
						{
							HostPort:      8083,
							ContainerPort: 83,
							Protocol:      v1.ProtocolSCTP,
						},
						{
							HostPort:      8084,
							ContainerPort: 84,
							Protocol:      v1.ProtocolTCP,
							HostIP:        "127.0.0.1",
						},
					},
				},
				expectError: false,
			},
			{
				mapping: &PodPortMapping{
					Name:        "pod3",
					Namespace:   "ns1",
					IP:          net.ParseIP("2001:beef::4"),
					HostNetwork: false,
					PortMappings: []*PortMapping{
						{
							HostPort:      8443,
							ContainerPort: 443,
							Protocol:      v1.ProtocolTCP,
						},
					},
				},
				expectError: false,
			},
			// port already taken by other pod
			// but using another IP family
			{
				mapping: &PodPortMapping{
					Name:        "pod4",
					Namespace:   "ns2",
					IP:          net.ParseIP("192.168.2.2"),
					HostNetwork: false,
					PortMappings: []*PortMapping{
						{
							HostPort:      8443,
							ContainerPort: 443,
							Protocol:      v1.ProtocolTCP,
						},
					},
				},
				expectError: false,
			},
		}

		// Add Hostports
		for _, tc := range testCases {
			err := manager.Add("id", tc.mapping, "")
			if tc.expectError {
				Expect(err).To(HaveOccurred())
				continue
			}
			Expect(err).NotTo(HaveOccurred())
		}

		// Check IPv4 Iptables-save result after adding hostports
		raw := bytes.NewBuffer(nil)

		err := iptables.SaveInto(utiliptables.TableNAT, raw)
		Expect(err).NotTo(HaveOccurred())

		lines := strings.Split(raw.String(), "\n")
		expectedLines := map[string]bool{
			`*nat`:                                true,
			`:KUBE-HOSTPORTS - [0:0]`:             true,
			`:CRIO-HOSTPORTS-MASQ - [0:0]`:        true,
			`:OUTPUT - [0:0]`:                     true,
			`:PREROUTING - [0:0]`:                 true,
			`:POSTROUTING - [0:0]`:                true,
			`:KUBE-HP-IJHALPHTORMHHPPK - [0:0]`:   true,
			`:CRIO-MASQ-IJHALPHTORMHHPPK - [0:0]`: true,
			`:KUBE-HP-63UPIDJXVRSZGSUZ - [0:0]`:   true,
			`:CRIO-MASQ-63UPIDJXVRSZGSUZ - [0:0]`: true,
			`:KUBE-HP-WFBOALXEP42XEMJK - [0:0]`:   true,
			`:CRIO-MASQ-WFBOALXEP42XEMJK - [0:0]`: true,
			`:KUBE-HP-XU6AWMMJYOZOFTFZ - [0:0]`:   true,
			`:CRIO-MASQ-XU6AWMMJYOZOFTFZ - [0:0]`: true,
			`:KUBE-HP-CHN66X54O4WXZ5CW - [0:0]`:   true,
			`:CRIO-MASQ-CHN66X54O4WXZ5CW - [0:0]`: true,
			"-A KUBE-HOSTPORTS -m comment --comment \"pod1_ns1 hostport 8081\" -m udp -p udp --dport 8081 -j KUBE-HP-63UPIDJXVRSZGSUZ":                                                                     true,
			"-A CRIO-HOSTPORTS-MASQ -m comment --comment \"pod1_ns1 hostport 8081\" -j CRIO-MASQ-63UPIDJXVRSZGSUZ":                                                                                         true,
			"-A KUBE-HOSTPORTS -m comment --comment \"pod1_ns1 hostport 8080\" -m tcp -p tcp --dport 8080 -j KUBE-HP-IJHALPHTORMHHPPK":                                                                     true,
			"-A CRIO-HOSTPORTS-MASQ -m comment --comment \"pod1_ns1 hostport 8080\" -j CRIO-MASQ-IJHALPHTORMHHPPK":                                                                                         true,
			"-A KUBE-HOSTPORTS -m comment --comment \"pod1_ns1 hostport 8083\" -m sctp -p sctp --dport 8083 -j KUBE-HP-XU6AWMMJYOZOFTFZ":                                                                   true,
			"-A CRIO-HOSTPORTS-MASQ -m comment --comment \"pod1_ns1 hostport 8083\" -j CRIO-MASQ-XU6AWMMJYOZOFTFZ":                                                                                         true,
			"-A KUBE-HOSTPORTS -m comment --comment \"pod1_ns1 hostport 8084\" -m tcp -p tcp --dport 8084 -j KUBE-HP-CHN66X54O4WXZ5CW":                                                                     true,
			"-A CRIO-HOSTPORTS-MASQ -m comment --comment \"pod1_ns1 hostport 8084\" -j CRIO-MASQ-CHN66X54O4WXZ5CW":                                                                                         true,
			"-A KUBE-HOSTPORTS -m comment --comment \"pod4_ns2 hostport 8443\" -m tcp -p tcp --dport 8443 -j KUBE-HP-WFBOALXEP42XEMJK":                                                                     true,
			"-A CRIO-HOSTPORTS-MASQ -m comment --comment \"pod4_ns2 hostport 8443\" -j CRIO-MASQ-WFBOALXEP42XEMJK":                                                                                         true,
			"-A OUTPUT -m comment --comment \"kube hostport portals\" -m addrtype --dst-type LOCAL -j KUBE-HOSTPORTS":                                                                                      true,
			"-A PREROUTING -m comment --comment \"kube hostport portals\" -m addrtype --dst-type LOCAL -j KUBE-HOSTPORTS":                                                                                  true,
			"-A POSTROUTING -m comment --comment \"kube hostport masquerading\" -m conntrack --ctstate DNAT -j CRIO-HOSTPORTS-MASQ":                                                                        true,
			"-A CRIO-HOSTPORTS-MASQ -m comment --comment \"SNAT for localhost access to hostports\" -o cbr0 -s ::1/128 -j MASQUERADE":                                                                      true,
			"-A CRIO-MASQ-IJHALPHTORMHHPPK -m comment --comment \"pod1_ns1 hostport 8080\" -m conntrack --ctorigdstport 8080 -m tcp -p tcp --dport 80 -s 192.168.2.7/32 -d 192.168.2.7/32 -j MASQUERADE":   true,
			"-A KUBE-HP-IJHALPHTORMHHPPK -m comment --comment \"pod1_ns1 hostport 8080\" -m tcp -p tcp -j DNAT --to-destination 192.168.2.7:80":                                                            true,
			"-A CRIO-MASQ-63UPIDJXVRSZGSUZ -m comment --comment \"pod1_ns1 hostport 8081\" -m conntrack --ctorigdstport 8081 -m udp -p udp --dport 81 -s 192.168.2.7/32 -d 192.168.2.7/32 -j MASQUERADE":   true,
			"-A KUBE-HP-63UPIDJXVRSZGSUZ -m comment --comment \"pod1_ns1 hostport 8081\" -m udp -p udp -j DNAT --to-destination 192.168.2.7:81":                                                            true,
			"-A CRIO-MASQ-XU6AWMMJYOZOFTFZ -m comment --comment \"pod1_ns1 hostport 8083\" -m conntrack --ctorigdstport 8083 -m sctp -p sctp --dport 83 -s 192.168.2.7/32 -d 192.168.2.7/32 -j MASQUERADE": true,
			"-A KUBE-HP-XU6AWMMJYOZOFTFZ -m comment --comment \"pod1_ns1 hostport 8083\" -m sctp -p sctp -j DNAT --to-destination 192.168.2.7:83":                                                          true,
			"-A CRIO-MASQ-CHN66X54O4WXZ5CW -m comment --comment \"pod1_ns1 hostport 8084\" -m conntrack --ctorigdstport 8084 -m tcp -p tcp --dport 84 -s 192.168.2.7/32 -d 192.168.2.7/32 -j MASQUERADE":   true,
			"-A KUBE-HP-CHN66X54O4WXZ5CW -m comment --comment \"pod1_ns1 hostport 8084\" -m tcp -p tcp -d 127.0.0.1/32 -j DNAT --to-destination 192.168.2.7:84":                                            true,
			"-A CRIO-MASQ-WFBOALXEP42XEMJK -m comment --comment \"pod4_ns2 hostport 8443\" -m conntrack --ctorigdstport 8443 -m tcp -p tcp --dport 443 -s 192.168.2.2/32 -d 192.168.2.2/32 -j MASQUERADE":  true,
			"-A KUBE-HP-WFBOALXEP42XEMJK -m comment --comment \"pod4_ns2 hostport 8443\" -m tcp -p tcp -j DNAT --to-destination 192.168.2.2:443":                                                           true,
			`COMMIT`: true,
		}
		for _, line := range lines {
			GinkgoWriter.Printf("Line: %s\n", line)
			if strings.TrimSpace(line) != "" {
				_, ok := expectedLines[strings.TrimSpace(line)]
				Expect(ok).To(BeTrue())
			}
		}

		// Remove all added hostports
		for _, tc := range testCases {
			if !tc.expectError {
				err := manager.Remove("id", tc.mapping)
				Expect(err).NotTo(HaveOccurred())
			}
		}

		// Check IPv6 Iptables-save result after deleting hostports
		raw.Reset()
		err = iptables.SaveInto(utiliptables.TableNAT, raw)
		Expect(err).NotTo(HaveOccurred())
		lines = strings.Split(raw.String(), "\n")
		remainingChains := make(map[string]bool)
		for _, line := range lines {
			if strings.HasPrefix(line, ":") {
				remainingChains[strings.TrimSpace(line)] = true
			}
		}
		expectDeletedChains := []string{"KUBE-HP-4YVONL46AKYWSKS3", "KUBE-HP-7THKRFSEH4GIIXK7", "KUBE-HP-5N7UH5JAXCVP5UJR", "KUBE-HP-CHN66X54O4WXZ5CW"}
		for _, chain := range expectDeletedChains {
			_, ok := remainingChains[chain]
			Expect(ok).To(BeFalse())
		}

		// Check Iptables-save result after adding hostports
		rawv6 := bytes.NewBuffer(nil)

		err = ip6tables.SaveInto(utiliptables.TableNAT, rawv6)
		Expect(err).NotTo(HaveOccurred())

		linesv6 := strings.Split(rawv6.String(), "\n")
		expectedv6Lines := map[string]bool{
			`*nat`:                                true,
			`:KUBE-HOSTPORTS - [0:0]`:             true,
			`:CRIO-HOSTPORTS-MASQ - [0:0]`:        true,
			`:OUTPUT - [0:0]`:                     true,
			`:PREROUTING - [0:0]`:                 true,
			`:POSTROUTING - [0:0]`:                true,
			`:KUBE-HP-IJHALPHTORMHHPPK - [0:0]`:   true,
			`:CRIO-MASQ-IJHALPHTORMHHPPK - [0:0]`: true,
			`:KUBE-HP-63UPIDJXVRSZGSUZ - [0:0]`:   true,
			`:CRIO-MASQ-63UPIDJXVRSZGSUZ - [0:0]`: true,
			`:KUBE-HP-WFBOALXEP42XEMJK - [0:0]`:   true,
			`:CRIO-MASQ-WFBOALXEP42XEMJK - [0:0]`: true,
			`:KUBE-HP-XU6AWMMJYOZOFTFZ - [0:0]`:   true,
			`:CRIO-MASQ-XU6AWMMJYOZOFTFZ - [0:0]`: true,
			"-A KUBE-HOSTPORTS -m comment --comment \"pod3_ns1 hostport 8443\" -m tcp -p tcp --dport 8443 -j KUBE-HP-WFBOALXEP42XEMJK":                                                                      true,
			"-A CRIO-HOSTPORTS-MASQ -m comment --comment \"pod3_ns1 hostport 8443\" -j CRIO-MASQ-WFBOALXEP42XEMJK":                                                                                          true,
			"-A KUBE-HOSTPORTS -m comment --comment \"pod1_ns1 hostport 8081\" -m udp -p udp --dport 8081 -j KUBE-HP-63UPIDJXVRSZGSUZ":                                                                      true,
			"-A CRIO-HOSTPORTS-MASQ -m comment --comment \"pod1_ns1 hostport 8081\" -j CRIO-MASQ-63UPIDJXVRSZGSUZ":                                                                                          true,
			"-A KUBE-HOSTPORTS -m comment --comment \"pod1_ns1 hostport 8080\" -m tcp -p tcp --dport 8080 -j KUBE-HP-IJHALPHTORMHHPPK":                                                                      true,
			"-A CRIO-HOSTPORTS-MASQ -m comment --comment \"pod1_ns1 hostport 8080\" -j CRIO-MASQ-IJHALPHTORMHHPPK":                                                                                          true,
			"-A KUBE-HOSTPORTS -m comment --comment \"pod1_ns1 hostport 8083\" -m sctp -p sctp --dport 8083 -j KUBE-HP-XU6AWMMJYOZOFTFZ":                                                                    true,
			"-A CRIO-HOSTPORTS-MASQ -m comment --comment \"pod1_ns1 hostport 8083\" -j CRIO-MASQ-XU6AWMMJYOZOFTFZ":                                                                                          true,
			"-A OUTPUT -m comment --comment \"kube hostport portals\" -m addrtype --dst-type LOCAL -j KUBE-HOSTPORTS":                                                                                       true,
			"-A PREROUTING -m comment --comment \"kube hostport portals\" -m addrtype --dst-type LOCAL -j KUBE-HOSTPORTS":                                                                                   true,
			"-A POSTROUTING -m comment --comment \"kube hostport masquerading\" -m conntrack --ctstate DNAT -j CRIO-HOSTPORTS-MASQ":                                                                         true,
			"-A CRIO-HOSTPORTS-MASQ -m comment --comment \"SNAT for localhost access to hostports\" -o cbr0 -s ::1/128 -j MASQUERADE":                                                                       true,
			"-A CRIO-MASQ-IJHALPHTORMHHPPK -m comment --comment \"pod1_ns1 hostport 8080\" -m conntrack --ctorigdstport 9999 -m tcp -p tcp --dport 443 -s 2001:beef::2/32 -d 2001:beef::2/32 -j MASQUERADE": true,
			"-A KUBE-HP-IJHALPHTORMHHPPK -m comment --comment \"pod1_ns1 hostport 8080\" -m tcp -p tcp -j DNAT --to-destination [2001:beef::2]:80":                                                          true,
			"-A CRIO-MASQ-63UPIDJXVRSZGSUZ -m comment --comment \"pod1_ns1 hostport 8081\" -m conntrack --ctorigdstport 9999 -m tcp -p tcp --dport 443 -s 2001:beef::2/32 -d 2001:beef::2/32 -j MASQUERADE": true,
			"-A KUBE-HP-63UPIDJXVRSZGSUZ -m comment --comment \"pod1_ns1 hostport 8081\" -m udp -p udp -j DNAT --to-destination [2001:beef::2]:81":                                                          true,
			"-A CRIO-MASQ-XU6AWMMJYOZOFTFZ -m comment --comment \"pod1_ns1 hostport 8083\" -m conntrack --ctorigdstport 9999 -m tcp -p tcp --dport 443 -s 2001:beef::2/32 -d 2001:beef::2/32 -j MASQUERADE": true,
			"-A KUBE-HP-XU6AWMMJYOZOFTFZ -m comment --comment \"pod1_ns1 hostport 8083\" -m sctp -p sctp -j DNAT --to-destination [2001:beef::2]:83":                                                        true,
			"-A CRIO-MASQ-WFBOALXEP42XEMJK -m comment --comment \"pod3_ns1 hostport 8443\" -m conntrack --ctorigdstport 9999 -m tcp -p tcp --dport 443 -s 2001:beef::4/32 -d 2001:beef::4/32 -j MASQUERADE": true,
			"-A KUBE-HP-WFBOALXEP42XEMJK -m comment --comment \"pod3_ns1 hostport 8443\" -m tcp -p tcp -j DNAT --to-destination [2001:beef::4]:443":                                                         true,
			`COMMIT`: true,
		}
		for _, line := range linesv6 {
			if strings.TrimSpace(line) != "" {
				_, ok := expectedv6Lines[strings.TrimSpace(line)]
				Expect(ok).To(BeTrue())
			}
		}

		// Remove all added hostports
		for _, tc := range testCases {
			if !tc.expectError {
				err := manager.Remove("id", tc.mapping)
				Expect(err).NotTo(HaveOccurred())
			}
		}

		// Check Iptables-save result after deleting hostports
		rawv6.Reset()
		err = ip6tables.SaveInto(utiliptables.TableNAT, rawv6)
		Expect(err).NotTo(HaveOccurred())
		linesv6 = strings.Split(rawv6.String(), "\n")
		remainingv6Chains := make(map[string]bool)
		for _, line := range linesv6 {
			if strings.HasPrefix(line, ":") {
				remainingv6Chains[strings.TrimSpace(line)] = true
			}
		}
		expectv6DeletedChains := []string{"KUBE-HP-4YVONL46AKYWSKS3", "KUBE-HP-7THKRFSEH4GIIXK7", "KUBE-HP-5N7UH5JAXCVP5UJR"}
		for _, chain := range expectv6DeletedChains {
			_, ok := remainingv6Chains[chain]
			Expect(ok).To(BeFalse())
		}
	})
})
