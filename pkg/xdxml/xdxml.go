// WARNING: This file has automatically been generated on Fri, 21 Jul 2023 13:47:56 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package xdxml

/*
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files
#include "xdxml.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// xdxml_init function as declared in xdxml/xdxml.h:133
func xdxml_init() Return {
	__ret := C.xdxml_init()
	__v := (Return)(__ret)
	return __v
}

// xdxml_shutdown function as declared in xdxml/xdxml.h:145
func xdxml_shutdown() Return {
	__ret := C.xdxml_shutdown()
	__v := (Return)(__ret)
	return __v
}

// xdxml_device_get_count function as declared in xdxml/xdxml.h:158
func xdxml_device_get_count(Device_count *uint32) Return {
	cDevice_count, cDevice_countAllocMap := (*C.uint)(unsafe.Pointer(Device_count)), cgoAllocsUnknown
	__ret := C.xdxml_device_get_count(cDevice_count)
	runtime.KeepAlive(cDevice_countAllocMap)
	__v := (Return)(__ret)
	return __v
}

// xdxml_device_get_handle_by_index function as declared in xdxml/xdxml.h:176
func xdxml_device_get_handle_by_index(Index uint32, Device *Device) Return {
	cIndex, cIndexAllocMap := (C.uint)(Index), cgoAllocsUnknown
	cDevice, cDeviceAllocMap := (*C.xdx_device_t)(unsafe.Pointer(Device)), cgoAllocsUnknown
	__ret := C.xdxml_device_get_handle_by_index(cIndex, cDevice)
	runtime.KeepAlive(cDeviceAllocMap)
	runtime.KeepAlive(cIndexAllocMap)
	__v := (Return)(__ret)
	return __v
}

// xdxml_device_get_uuid function as declared in xdxml/xdxml.h:186
func xdxml_device_get_uuid(Device Device) Return {
	cDevice, cDeviceAllocMap := *(*C.xdx_device_t)(unsafe.Pointer(&Device)), cgoAllocsUnknown
	__ret := C.xdxml_device_get_uuid(cDevice)
	runtime.KeepAlive(cDeviceAllocMap)
	__v := (Return)(__ret)
	return __v
}
