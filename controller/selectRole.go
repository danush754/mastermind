package controller

import "fmt"

func SelectRole() int {

	var userSelection int

	fmt.Println("**************************** MasterMind ****************************")
	fmt.Println("********************************************************************")

	fmt.Println("Do you want set the code or find the code ?")
	fmt.Println("")
	fmt.Print("If you want to set the code Press 1 or If you want to find the code Press 2 :")
	fmt.Scan(&userSelection)
	fmt.Println("")

	if userSelection == 1 {
		fmt.Println("you have selected to set the code!!")
		fmt.Println("")
	} else if userSelection == 2 {
		fmt.Println("you have selected the to find the code!!")
		fmt.Println("")
	}

	return userSelection
}
