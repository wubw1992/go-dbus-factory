// Code generated by "./generator ./com.deepin.api.pinyin"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package pinyin

import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Pinyin interface {
	pinyin // interface com.deepin.api.Pinyin
	proxy.Object
}

type objectPinyin struct {
	interfacePinyin // interface com.deepin.api.Pinyin
	proxy.ImplObject
}

func NewPinyin(conn *dbus.Conn) Pinyin {
	obj := new(objectPinyin)
	obj.ImplObject.Init_(conn, "com.deepin.api.Pinyin", "/com/deepin/api/Pinyin")
	return obj
}

type pinyin interface {
	GoQuery(flags dbus.Flags, ch chan *dbus.Call, hans string) *dbus.Call
	Query(flags dbus.Flags, hans string) ([]string, error)
	GoQueryList(flags dbus.Flags, ch chan *dbus.Call, hansList []string) *dbus.Call
	QueryList(flags dbus.Flags, hansList []string) (string, error)
}

type interfacePinyin struct{}

func (v *interfacePinyin) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfacePinyin) GetInterfaceName_() string {
	return "com.deepin.api.Pinyin"
}

// method Query

func (v *interfacePinyin) GoQuery(flags dbus.Flags, ch chan *dbus.Call, hans string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Query", flags, ch, hans)
}

func (*interfacePinyin) StoreQuery(call *dbus.Call) (pinyin []string, err error) {
	err = call.Store(&pinyin)
	return
}

func (v *interfacePinyin) Query(flags dbus.Flags, hans string) ([]string, error) {
	return v.StoreQuery(
		<-v.GoQuery(flags, make(chan *dbus.Call, 1), hans).Done)
}

// method QueryList

func (v *interfacePinyin) GoQueryList(flags dbus.Flags, ch chan *dbus.Call, hansList []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".QueryList", flags, ch, hansList)
}

func (*interfacePinyin) StoreQueryList(call *dbus.Call) (json string, err error) {
	err = call.Store(&json)
	return
}

func (v *interfacePinyin) QueryList(flags dbus.Flags, hansList []string) (string, error) {
	return v.StoreQueryList(
		<-v.GoQueryList(flags, make(chan *dbus.Call, 1), hansList).Done)
}
