package main

import (
	"fmt"
	passwdqc "go-passwdqc"
)

func main() {
	pass, err := passwdqc.Generate()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Password:", pass)
}
