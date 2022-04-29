package main

import (
	"fmt"
	"os"
	"strings"
)

/// store all the password in the data slice
//var data []string
// create the @ function
func attherate(uData []string) {
	var pass []string
	fName := uData[0]
	for _, data := range uData {
		if data == fName {
			continue
		} else {
			password := fName + "@" + data
			pass = append(pass, password)
		}
	}
	fLatterU := strings.ToUpper(string(fName[0]))
	fNameU := fLatterU + fName[1:]
	for _, data := range uData {
		if data == fName {
			continue
		} else {
			password := fNameU + "@" + data
			pass = append(pass, password)
		}
	}
	fmt.Print(pass)

}

// create a help function ---
func help() {
	fmt.Println("\033[35;1m Use : cupp ")
	fmt.Println("\033[35;1m Cupp is a password generator for the indian user ")
	fmt.Println("\033[35;1m Argument : -h or --help for the help ")
	fmt.Println("\033[35;1m Argument : -i for intractive mode ")
}
func main() {
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		help()
	} else if os.Args[1] == "-i" || os.Args[1] == "-i " {
		// now take the input from the user and store them
		var fName, lName, dob, aadhar, num string
		var userData []string
		fmt.Print("\033[33;1m First name of the target: ")
		fmt.Scanf("%s", &fName)
		fmt.Print("\033[33;1m Last name of the target: ")
		fmt.Scanf("%s", &lName)
		fmt.Print("\033[33;1m Date of birth: ")
		fmt.Scanf("%s", &dob)
		fmt.Print("\033[33;1m Target mobile number without 91 :")
		fmt.Scanf("%s", &num)
		fmt.Print("\033[33;1m First four digit aadhar number: ")
		fmt.Scanf("%s", &aadhar)
		userData = append(userData, fName, lName, dob, aadhar, num)
		attherate(userData)

	} else {
		fmt.Println("\033[31;1m Invalid argument provided ")
		help()
	}
}
