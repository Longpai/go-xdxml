package main

import (
	"fmt"
	"log"

	"github.com/chen-mao/go-xdxml/pkg/xdxml"
)

func main() {
	ret := xdxml.Init()
	if ret != xdxml.SUCCESS {
		log.Fatalf("Unable to initialize XDXML: %v", ret)
	}
	defer func() {
		ret := xdxml.Shutdown()
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to shutdown XDXML: %v", ret)
		}
	}()

	count, ret := xdxml.DeviceGetCount()
	if ret != xdxml.SUCCESS {
		log.Fatalf("Unable to get device count: %v", ret)
	}
	fmt.Printf("GPU Count: %v\n", count)

	for i := 0; i < count; i++ {
		device, ret := xdxml.DeviceGetHandleByIndex(i)
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get device at index %d: %v", i, ret)
		}

		ID, ret := device.GetMinorNumber()
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get id of device at index %v: %v", ID, ret)
		}
		fmt.Printf("ID: %v\n", ID)

		uuid, ret := device.GetUUID()
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get uuid of device at index %d: %v", i, ret)
		}

		fmt.Printf("UUID: %v\n", uuid)
	}
}
