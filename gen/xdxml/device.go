package xdxml

import (
	"strconv"
)

// xdxml.DeviceGetCount()
func DeviceGetCount() (int, Return) {
	var DeviceCount uint32
	ret := xdxml_device_get_count(&DeviceCount)
	return int(DeviceCount), ret
}

// xdxml.DeviceGetHandleByIndex()
func DeviceGetHandleByIndex(Index int) (Device, Return) {
	var Device Device
	ret := xdxml_device_get_handle_by_index(uint32(Index), &Device)
	return Device, ret
}

// xdxml.DeviceGetUUID()
func DeviceGetUUID(Device Device) (string, Return) {
	var uuidStr string
	ret := xdxml_device_get_uuid(Device)
	for _, num := range Device.Uuid {
		uuidStr += strconv.Itoa(int(num))
	}
	return uuidStr, ret
}

func (Device Device) GetUUID() (string, Return) {
	return DeviceGetUUID(Device)
}

// xdxml.DeviceGetUUID()
func DeviceGetMinorNumber(Device Device) (int, Return) {
	var minorNumber int32
	ret := xdxml_device_get_minor_number(Device, &minorNumber)
	return int(minorNumber), ret
}

func (Device Device) GetMinorNumber() (int, Return) {
	return DeviceGetMinorNumber(Device)
}
