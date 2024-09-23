package main

import (
	"fmt"
	"mastermind/controller"
)

func main() {
	userSelection := controller.SelectRole()

	if userSelection == 1 {

		isvalid := controller.GetUserCode()
		fmt.Println("secret code: ", isvalid)
		if !isvalid {

			controller.GetUserCode()
		} else {
		}

	} else if userSelection == 2 {

	} else {
		fmt.Println("Invalid input!!, please choose the correct option")
		fmt.Println("")
		controller.SelectRole()
	}

}
