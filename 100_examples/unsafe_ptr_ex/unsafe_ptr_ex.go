package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

type person struct {
	name string
	age  int
}

func main() {
	// original example manually examied the printed addressand used the value
	// updated to preserver forwared compatibility due to runtim changes shifting the address over time

	hi := "HI"

	// getting address dynamically to preserver compatibility
	address := fmt.Sprint(&hi)

	fmt.Printf("Address of var hi: %s\n", address)

	// convert to uintptr

	var adr uint64

	adr, err := strconv.ParseUint(address, 0, 64)
	if err != nil {
		panic(err)
	}
	var ptr uintptr = uintptr(adr)

	fmt.Printf("String at address: %s\n", address)
	fmt.Printf("Value: %s\n", ptrToString(ptr))

	//
	//
	//
	//
	//
	np := person{
		name: "krls",
		age:  73,
	}
	addrNP := fmt.Sprintf("%p", &np)
	fmt.Println("&np:", fmt.Sprintf("%p", &np))
	fmt.Println("addrNP:", addrNP)

	var adr1 uint64
	adr1, err = strconv.ParseUint(addrNP, 0, 64)
	if err != nil {
		panic(err)
	}
	var ptr1 uintptr = uintptr(adr1)

	ptrToPerson(ptr1)

}

func ptrToString(ptr uintptr) string {
	p := unsafe.Pointer(ptr)
	fmt.Printf("p =>, %T, v: %v\n", p, p)
	return *(*string)(p)
}

func ptrToPerson(ptr uintptr) {
	p := unsafe.Pointer(ptr)
	fmt.Printf("p =>, %T, v: %v\n", p, p)

	fmt.Println("--------->", *(*person)(p))
}
