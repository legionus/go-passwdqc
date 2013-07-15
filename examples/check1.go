package main

import (
	"fmt"
	passwdqc "go-passwdqc"
)

func main() {
	login   := "legion"
	newpass := "qwertyui"
	oldpass := "foobar"

	if err := passwdqc.CheckAccount(login,newpass,oldpass); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Good password")
}
