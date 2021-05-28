// Code generated by "./generator ./com.deepin.udcp.iam"; DO NOT EDIT.

package iam

import "fmt"
import "github.com/godbus/dbus"
import "github.com/stretchr/testify/mock"

import "pkg.deepin.io/lib/dbusutil/proxy"

type MockUdcpCache struct {
	mockInterfaceUdcpCache // interface com.deepin.udcp.iam
}

type mockInterfaceUdcpCache struct {
	mock.Mock
}

// method GetUserIdList

func (v *mockInterfaceUdcpCache) GoGetUserIdList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUdcpCache) GetUserIdList(flags dbus.Flags) ([]uint32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetUserGroups

func (v *mockInterfaceUdcpCache) GoGetUserGroups(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUdcpCache) GetUserGroups(flags dbus.Flags, name string) ([]string, error) {
	mockArgs := v.Called(flags, name)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RemoveCacheFile

func (v *mockInterfaceUdcpCache) GoRemoveCacheFile(flags dbus.Flags, ch chan *dbus.Call, uId uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, uId)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUdcpCache) RemoveCacheFile(flags dbus.Flags, uId uint32) (bool, error) {
	mockArgs := v.Called(flags, uId)

	ret0, ok := mockArgs.Get(0).(bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Enable b

func (v *mockInterfaceUdcpCache) Enable() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
