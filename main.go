package main

import (
	"fmt"
	"gogogo/internal/service"
	"gogogo/pkg/consts"
	"reflect"
)

func main() {
	var i consts.MyInterface = &service.MyService{
		Name: "yangjun", Place: "beijing", Job: "coding",
	}

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.ValueOf(i))
	fmt.Println(i.WhoAmI())
	fmt.Println(i.WhereAmI())
	fmt.Println(i.WhatAmIDoing())

}
