package test

import (
	"fmt"
	"task2/client"
)

func RunTests() {
	newClient := client.NewClient("http://localhost:8080")

	_, err := newClient.GetVersion()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	decodedString, err := newClient.PostDecode("aGVsbG8gd29ybGQ=") // hello world
	if err != nil {
		fmt.Printf("PostDecode request failed: %s\n\n", err.Error())
		return
	}
	fmt.Println("PostDecode request. decodedString :", decodedString)

	status, code, err := newClient.GetHardOp()
	if err != nil {
		fmt.Printf("GetHardOp request failed: %s\n\n", err.Error())
		return
	}
	fmt.Println("GetHardOp request. status :", status, "; code :", code)
}
