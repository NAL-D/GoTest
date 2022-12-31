package main

import (
	"fmt"

	"GoTest/test"
)

type Me struct {
	Name   string `json:"Gender"`
	Gender string
	Age    int
	Oh
}

type Oh struct {
	// Ccc string
	DesT map[string]interface{}
}

func main() {

	test.Test()

	// src := `{"Name": "dadf","Gender": "bbb", "age": 1, "DesT": {"abcd", "1234"}}`
	// src := `{"Name": "dadf","CCC": "bbb", "age": 1, "DesT": {"ccc": "1234"}}`

	// ch := make(chan struct{}, 100)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ch:
	// 			fmt.Println("aaa")
	// 			return
	// 		case ch <- struct{}{}:
	// 			fmt.Println("bbb")
	// 		default:
	// 			fmt.Println("default")
	// 			return
	// 		}
	// 	}
	// }()

	// var me Me
	// me := new(Me)
	// _ = json.Unmarshal([]byte(src), &me)
	// fmt.Printf("%+v/n", me)

	// <-ch
	// ch <- struct{}{}
	// close(ch)

	aaa := map[string][]string{
		"aa": {"1"},
		"bb": {"1", "2"},
		"cc": {"1", "2", "3"},
	}

	for key, value := range aaa {
		fmt.Println(key)
		fmt.Println(value)
	}

	param := []string{
		"aa",
		"bb",
		"cc",
		"dd",
		"ee",
	}

	// var yama []Me

	fmt.Println(param)
	// test(param...)
}

// func aaa(options ...string) {

// 	for _, v := range options {
// 		fmt.Println(v)
// 	}
// }
