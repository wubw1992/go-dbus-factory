// Code generated by "./generator ./com.deepin.daemon.airplanemode"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package airplanemode

import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockAirplaneMode struct {
	MockInterfaceAirplaneMode // interface com.deepin.daemon.AirplaneMode
	proxy.MockObject
}

type MockInterfaceAirplaneMode struct {
	mock.Mock
}

// method DumpState

func (v *MockInterfaceAirplaneMode) GoDumpState(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAirplaneMode) DumpState(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Enable

func (v *MockInterfaceAirplaneMode) GoEnable(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAirplaneMode) Enable(flags dbus.Flags, enabled bool) error {
	mockArgs := v.Called(flags, enabled)

	return mockArgs.Error(0)
}

// method EnableBluetooth

func (v *MockInterfaceAirplaneMode) GoEnableBluetooth(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAirplaneMode) EnableBluetooth(flags dbus.Flags, enabled bool) error {
	mockArgs := v.Called(flags, enabled)

	return mockArgs.Error(0)
}

// method EnableWifi

func (v *MockInterfaceAirplaneMode) GoEnableWifi(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAirplaneMode) EnableWifi(flags dbus.Flags, enabled bool) error {
	mockArgs := v.Called(flags, enabled)

	return mockArgs.Error(0)
}

// property Enabled b

func (v *MockInterfaceAirplaneMode) Enabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property WifiEnabled b

func (v *MockInterfaceAirplaneMode) WifiEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property BluetoothEnabled b

func (v *MockInterfaceAirplaneMode) BluetoothEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
