package main

import (
	"fmt"
	passwdqc "go-passwdqc"
)

func main() {
	if err := passwdqc.CheckPassword("zei1Ithoh0ahgeiHei0eghoo1johxu3u"); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Good password")
}
