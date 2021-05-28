// Code generated by "./generator ./com.deepin.udcp.iam"; DO NOT EDIT.

package iam

import "github.com/godbus/dbus"

import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"
import "errors"

type UdcpCache interface {
	udcpCache // interface com.deepin.udcp.iam
	proxy.Object
}

type objectUdcpCache struct {
	interfaceUdcpCache // interface com.deepin.udcp.iam
	proxy.ImplObject
}

func NewUdcpCache(conn *dbus.Conn, serviceName string, path dbus.ObjectPath) (UdcpCache, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectUdcpCache)
	obj.ImplObject.Init_(conn, serviceName, path)
	return obj, nil
}

type udcpCache interface {
	GoGetUserIdList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetUserIdList(flags dbus.Flags) ([]uint32, error)
	GoGetUserGroups(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call
	GetUserGroups(flags dbus.Flags, name string) ([]string, error)
	GoRemoveCacheFile(flags dbus.Flags, ch chan *dbus.Call, uId uint32) *dbus.Call
	RemoveCacheFile(flags dbus.Flags, uId uint32) (bool, error)
	Enable() proxy.PropBool
}

type interfaceUdcpCache struct{}

func (v *interfaceUdcpCache) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceUdcpCache) GetInterfaceName_() string {
	return "com.deepin.udcp.iam"
}

// method GetUserIdList

func (v *interfaceUdcpCache) GoGetUserIdList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUserIdList", flags, ch)
}

func (*interfaceUdcpCache) StoreGetUserIdList(call *dbus.Call) (uIdList []uint32, err error) {
	err = call.Store(&uIdList)
	return
}

func (v *interfaceUdcpCache) GetUserIdList(flags dbus.Flags) ([]uint32, error) {
	return v.StoreGetUserIdList(
		<-v.GoGetUserIdList(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetUserGroups

func (v *interfaceUdcpCache) GoGetUserGroups(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUserGroups", flags, ch, name)
}

func (*interfaceUdcpCache) StoreGetUserGroups(call *dbus.Call) (groups []string, err error) {
	err = call.Store(&groups)
	return
}

func (v *interfaceUdcpCache) GetUserGroups(flags dbus.Flags, name string) ([]string, error) {
	return v.StoreGetUserGroups(
		<-v.GoGetUserGroups(flags, make(chan *dbus.Call, 1), name).Done)
}

// method RemoveCacheFile

func (v *interfaceUdcpCache) GoRemoveCacheFile(flags dbus.Flags, ch chan *dbus.Call, uId uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveCacheFile", flags, ch, uId)
}

func (*interfaceUdcpCache) StoreRemoveCacheFile(call *dbus.Call) (result bool, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceUdcpCache) RemoveCacheFile(flags dbus.Flags, uId uint32) (bool, error) {
	return v.StoreRemoveCacheFile(
		<-v.GoRemoveCacheFile(flags, make(chan *dbus.Call, 1), uId).Done)
}

// property Enable b

func (v *interfaceUdcpCache) Enable() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Enable",
	}
}
