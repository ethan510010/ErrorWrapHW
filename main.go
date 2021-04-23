package main

import (
	"fmt"
	userdao "github.com/ethan510010/week02_hw/dao/user"
)

func main() {
	 user, err := userdao.GetUserBy(6)
	 if err != nil {
	 	fmt.Printf( "err: %+v", err)
		 return
	 }
	 fmt.Println("get user", user)
}
