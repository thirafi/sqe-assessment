// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository/database/otp.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	model "SQ-Assessment/model"
	util "SQ-Assessment/util"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOtpRepository is a mock of OtpRepository interface.
type MockOtpRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOtpRepositoryMockRecorder
}

// MockOtpRepositoryMockRecorder is the mock recorder for MockOtpRepository.
type MockOtpRepositoryMockRecorder struct {
	mock *MockOtpRepository
}

// NewMockOtpRepository creates a new mock instance.
func NewMockOtpRepository(ctrl *gomock.Controller) *MockOtpRepository {
	mock := &MockOtpRepository{ctrl: ctrl}
	mock.recorder = &MockOtpRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOtpRepository) EXPECT() *MockOtpRepositoryMockRecorder {
	return m.recorder
}

// CreateOTP mocks base method.
func (m *MockOtpRepository) CreateOTP(otp *model.Otp) util.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOTP", otp)
	ret0, _ := ret[0].(util.Error)
	return ret0
}

// CreateOTP indicates an expected call of CreateOTP.
func (mr *MockOtpRepositoryMockRecorder) CreateOTP(otp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOTP", reflect.TypeOf((*MockOtpRepository)(nil).CreateOTP), otp)
}

// GetActiveOTPByUserId mocks base method.
func (m *MockOtpRepository) GetActiveOTPByUserId(userId uint) (model.Otp, util.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveOTPByUserId", userId)
	ret0, _ := ret[0].(model.Otp)
	ret1, _ := ret[1].(util.Error)
	return ret0, ret1
}

// GetActiveOTPByUserId indicates an expected call of GetActiveOTPByUserId.
func (mr *MockOtpRepositoryMockRecorder) GetActiveOTPByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveOTPByUserId", reflect.TypeOf((*MockOtpRepository)(nil).GetActiveOTPByUserId), userId)
}

// GetOTPByUserId mocks base method.
func (m *MockOtpRepository) GetOTPByUserId(userId uint, otpNumber string) (model.Otp, util.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOTPByUserId", userId, otpNumber)
	ret0, _ := ret[0].(model.Otp)
	ret1, _ := ret[1].(util.Error)
	return ret0, ret1
}

// GetOTPByUserId indicates an expected call of GetOTPByUserId.
func (mr *MockOtpRepositoryMockRecorder) GetOTPByUserId(userId, otpNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOTPByUserId", reflect.TypeOf((*MockOtpRepository)(nil).GetOTPByUserId), userId, otpNumber)
}

// UpdateOTP mocks base method.
func (m *MockOtpRepository) UpdateOTP(otp *model.Otp) util.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOTP", otp)
	ret0, _ := ret[0].(util.Error)
	return ret0
}

// UpdateOTP indicates an expected call of UpdateOTP.
func (mr *MockOtpRepositoryMockRecorder) UpdateOTP(otp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOTP", reflect.TypeOf((*MockOtpRepository)(nil).UpdateOTP), otp)
}
