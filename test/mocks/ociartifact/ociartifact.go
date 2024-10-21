// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cri-o/cri-o/internal/config/ociartifact (interfaces: Impl)
//
// Generated by this command:
//
//	mockgen -package ociartifactmock -destination ./test/mocks/ociartifact/ociartifact.go github.com/cri-o/cri-o/internal/config/ociartifact Impl
//

// Package ociartifactmock is a generated GoMock package.
package ociartifactmock

import (
	context "context"
	io "io"
	fs "io/fs"
	reflect "reflect"

	reference "github.com/containers/image/v5/docker/reference"
	manifest "github.com/containers/image/v5/manifest"
	types "github.com/containers/image/v5/types"
	digest "github.com/opencontainers/go-digest"
	gomock "go.uber.org/mock/gomock"
)

// MockImpl is a mock of Impl interface.
type MockImpl struct {
	ctrl     *gomock.Controller
	recorder *MockImplMockRecorder
	isgomock struct{}
}

// MockImplMockRecorder is the mock recorder for MockImpl.
type MockImplMockRecorder struct {
	mock *MockImpl
}

// NewMockImpl creates a new mock instance.
func NewMockImpl(ctrl *gomock.Controller) *MockImpl {
	mock := &MockImpl{ctrl: ctrl}
	mock.recorder = &MockImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImpl) EXPECT() *MockImplMockRecorder {
	return m.recorder
}

// GetBlob mocks base method.
func (m *MockImpl) GetBlob(arg0 context.Context, arg1 types.ImageSource, arg2 types.BlobInfo, arg3 types.BlobInfoCache) (io.ReadCloser, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlob", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBlob indicates an expected call of GetBlob.
func (mr *MockImplMockRecorder) GetBlob(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlob", reflect.TypeOf((*MockImpl)(nil).GetBlob), arg0, arg1, arg2, arg3)
}

// GetManifest mocks base method.
func (m *MockImpl) GetManifest(arg0 context.Context, arg1 types.ImageSource, arg2 *digest.Digest) ([]byte, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManifest", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetManifest indicates an expected call of GetManifest.
func (mr *MockImplMockRecorder) GetManifest(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManifest", reflect.TypeOf((*MockImpl)(nil).GetManifest), arg0, arg1, arg2)
}

// LayerInfos mocks base method.
func (m *MockImpl) LayerInfos(arg0 manifest.Manifest) []manifest.LayerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LayerInfos", arg0)
	ret0, _ := ret[0].([]manifest.LayerInfo)
	return ret0
}

// LayerInfos indicates an expected call of LayerInfos.
func (mr *MockImplMockRecorder) LayerInfos(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayerInfos", reflect.TypeOf((*MockImpl)(nil).LayerInfos), arg0)
}

// ManifestConfigInfo mocks base method.
func (m *MockImpl) ManifestConfigInfo(arg0 manifest.Manifest) types.BlobInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManifestConfigInfo", arg0)
	ret0, _ := ret[0].(types.BlobInfo)
	return ret0
}

// ManifestConfigInfo indicates an expected call of ManifestConfigInfo.
func (mr *MockImplMockRecorder) ManifestConfigInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManifestConfigInfo", reflect.TypeOf((*MockImpl)(nil).ManifestConfigInfo), arg0)
}

// ManifestFromBlob mocks base method.
func (m *MockImpl) ManifestFromBlob(arg0 []byte, arg1 string) (manifest.Manifest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManifestFromBlob", arg0, arg1)
	ret0, _ := ret[0].(manifest.Manifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ManifestFromBlob indicates an expected call of ManifestFromBlob.
func (mr *MockImplMockRecorder) ManifestFromBlob(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManifestFromBlob", reflect.TypeOf((*MockImpl)(nil).ManifestFromBlob), arg0, arg1)
}

// MkdirAll mocks base method.
func (m *MockImpl) MkdirAll(arg0 string, arg1 fs.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkdirAll", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MkdirAll indicates an expected call of MkdirAll.
func (mr *MockImplMockRecorder) MkdirAll(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkdirAll", reflect.TypeOf((*MockImpl)(nil).MkdirAll), arg0, arg1)
}

// NewImageSource mocks base method.
func (m *MockImpl) NewImageSource(arg0 context.Context, arg1 types.ImageReference, arg2 *types.SystemContext) (types.ImageSource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewImageSource", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.ImageSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewImageSource indicates an expected call of NewImageSource.
func (mr *MockImplMockRecorder) NewImageSource(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewImageSource", reflect.TypeOf((*MockImpl)(nil).NewImageSource), arg0, arg1, arg2)
}

// NewReference mocks base method.
func (m *MockImpl) NewReference(arg0 reference.Named) (types.ImageReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReference", arg0)
	ret0, _ := ret[0].(types.ImageReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewReference indicates an expected call of NewReference.
func (mr *MockImplMockRecorder) NewReference(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReference", reflect.TypeOf((*MockImpl)(nil).NewReference), arg0)
}

// ParseNormalizedNamed mocks base method.
func (m *MockImpl) ParseNormalizedNamed(arg0 string) (reference.Named, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseNormalizedNamed", arg0)
	ret0, _ := ret[0].(reference.Named)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseNormalizedNamed indicates an expected call of ParseNormalizedNamed.
func (mr *MockImplMockRecorder) ParseNormalizedNamed(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseNormalizedNamed", reflect.TypeOf((*MockImpl)(nil).ParseNormalizedNamed), arg0)
}

// ReadAll mocks base method.
func (m *MockImpl) ReadAll(arg0 io.Reader) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockImplMockRecorder) ReadAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockImpl)(nil).ReadAll), arg0)
}

// ReadDir mocks base method.
func (m *MockImpl) ReadDir(arg0 string) ([]fs.DirEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDir", arg0)
	ret0, _ := ret[0].([]fs.DirEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDir indicates an expected call of ReadDir.
func (mr *MockImplMockRecorder) ReadDir(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDir", reflect.TypeOf((*MockImpl)(nil).ReadDir), arg0)
}

// ReadFile mocks base method.
func (m *MockImpl) ReadFile(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockImplMockRecorder) ReadFile(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockImpl)(nil).ReadFile), arg0)
}

// RemoveAll mocks base method.
func (m *MockImpl) RemoveAll(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll.
func (mr *MockImplMockRecorder) RemoveAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockImpl)(nil).RemoveAll), arg0)
}

// WriteFile mocks base method.
func (m *MockImpl) WriteFile(arg0 string, arg1 []byte, arg2 fs.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile.
func (mr *MockImplMockRecorder) WriteFile(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockImpl)(nil).WriteFile), arg0, arg1, arg2)
}
