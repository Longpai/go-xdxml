// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs types.go

package xdxml

type Node_path struct {
	Card               [128]int8
	Hwmon              [128]int8
	Name               [256]int8
	Dbdf               [256]int8
	Current_link_speed [256]int8
	Current_link_width [256]int8
	Max_link_speed     [256]int8
	Max_link_width     [256]int8
}

type Device struct {
	Handle *_Ctype_struct_xdx_device_st
}
