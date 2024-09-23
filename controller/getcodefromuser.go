package controller

import (
	"fmt"
	"regexp"
	"strconv"
)

func GetUserCode() bool {

	fmt.Println("OK!, Now lets set the five digit secret code and the let the computer find it ;)")

	var secretCode int

	fmt.Print("Enter the code: ")
	fmt.Scan(&secretCode)

	var secretCodeString = strconv.Itoa(secretCode)

	return ValidateSecretCode(secretCodeString)

}

func ValidateSecretCode(secretCode string) bool {

	regeEx, err := regexp.Compile(`^[1-6][1-6]{0,4}$`)
	if err != nil {
		return false
	}

	if regeEx.MatchString(secretCode) {
		return true
	}

	return false
}
