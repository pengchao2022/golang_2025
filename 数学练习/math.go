package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var a = 4.0
	var b = 2.0
	pingfang := math.Pow(a, 2)
	fmt.Println(pingfang)

	lifang := math.Pow(a, 3)
	fmt.Println(lifang)

	sicifang := math.Pow(a, 4)
	fmt.Println(sicifang)

	wucifang := math.Pow(a, 5)
	fmt.Println(wucifang)

	IPv4 := math.Pow(b, 32)
	fmt.Println(IPv4)
	fmt.Println(reflect.TypeOf(IPv4))

}
