// Code generated by "./generator ./com.deepin.system.network"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package network

import "errors"
import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Network interface {
	network // interface com.deepin.system.Network
	proxy.Object
}

type objectNetwork struct {
	interfaceNetwork // interface com.deepin.system.Network
	proxy.ImplObject
}

func NewNetwork(conn *dbus.Conn) Network {
	obj := new(objectNetwork)
	obj.ImplObject.Init_(conn, "com.deepin.system.Network", "/com/deepin/system/Network")
	return obj
}

type network interface {
	GoEnableDevice(flags dbus.Flags, ch chan *dbus.Call, pathOrIface string, enabled bool) *dbus.Call
	EnableDevice(flags dbus.Flags, pathOrIface string, enabled bool) (dbus.ObjectPath, error)
	GoIsDeviceEnabled(flags dbus.Flags, ch chan *dbus.Call, pathOrIface string) *dbus.Call
	IsDeviceEnabled(flags dbus.Flags, pathOrIface string) (bool, error)
	GoPing(flags dbus.Flags, ch chan *dbus.Call, host string) *dbus.Call
	Ping(flags dbus.Flags, host string) error
	GoToggleWirelessEnabled(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ToggleWirelessEnabled(flags dbus.Flags) (bool, error)
	ConnectDeviceEnabled(cb func(devPath dbus.ObjectPath, enabled bool)) (dbusutil.SignalHandlerId, error)
	VpnEnabled() proxy.PropBool
}

type interfaceNetwork struct{}

func (v *interfaceNetwork) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceNetwork) GetInterfaceName_() string {
	return "com.deepin.system.Network"
}

// method EnableDevice

func (v *interfaceNetwork) GoEnableDevice(flags dbus.Flags, ch chan *dbus.Call, pathOrIface string, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableDevice", flags, ch, pathOrIface, enabled)
}

func (*interfaceNetwork) StoreEnableDevice(call *dbus.Call) (cpath dbus.ObjectPath, err error) {
	err = call.Store(&cpath)
	return
}

func (v *interfaceNetwork) EnableDevice(flags dbus.Flags, pathOrIface string, enabled bool) (dbus.ObjectPath, error) {
	return v.StoreEnableDevice(
		<-v.GoEnableDevice(flags, make(chan *dbus.Call, 1), pathOrIface, enabled).Done)
}

// method IsDeviceEnabled

func (v *interfaceNetwork) GoIsDeviceEnabled(flags dbus.Flags, ch chan *dbus.Call, pathOrIface string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsDeviceEnabled", flags, ch, pathOrIface)
}

func (*interfaceNetwork) StoreIsDeviceEnabled(call *dbus.Call) (enabled bool, err error) {
	err = call.Store(&enabled)
	return
}

func (v *interfaceNetwork) IsDeviceEnabled(flags dbus.Flags, pathOrIface string) (bool, error) {
	return v.StoreIsDeviceEnabled(
		<-v.GoIsDeviceEnabled(flags, make(chan *dbus.Call, 1), pathOrIface).Done)
}

// method Ping

func (v *interfaceNetwork) GoPing(flags dbus.Flags, ch chan *dbus.Call, host string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Ping", flags, ch, host)
}

func (v *interfaceNetwork) Ping(flags dbus.Flags, host string) error {
	return (<-v.GoPing(flags, make(chan *dbus.Call, 1), host).Done).Err
}

// method ToggleWirelessEnabled

func (v *interfaceNetwork) GoToggleWirelessEnabled(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ToggleWirelessEnabled", flags, ch)
}

func (*interfaceNetwork) StoreToggleWirelessEnabled(call *dbus.Call) (enabled bool, err error) {
	err = call.Store(&enabled)
	return
}

func (v *interfaceNetwork) ToggleWirelessEnabled(flags dbus.Flags) (bool, error) {
	return v.StoreToggleWirelessEnabled(
		<-v.GoToggleWirelessEnabled(flags, make(chan *dbus.Call, 1)).Done)
}

// signal DeviceEnabled

func (v *interfaceNetwork) ConnectDeviceEnabled(cb func(devPath dbus.ObjectPath, enabled bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceEnabled", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceEnabled",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var devPath dbus.ObjectPath
		var enabled bool
		err := dbus.Store(sig.Body, &devPath, &enabled)
		if err == nil {
			cb(devPath, enabled)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property VpnEnabled b

func (v *interfaceNetwork) VpnEnabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "VpnEnabled",
	}
}
