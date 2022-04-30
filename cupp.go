package main

import (
	"fmt"
	"os"
	"strings"
)

/// store all the password in the data slice
//var data []string
// create the @ function
func underAtthe(uData []string) {
	var pass []string
	var years = []string{"1990", "1991", "1992", "1993", "1994", "1995", "1996", "1997", "1998", "1999", "2000", "2001", "2002", "2003", "2004", "2005", "2006", "2007", "2008", "2009", "2010", "2011", "2012", "2013", "2014", "2015", "2016", "2017", "2018", "2019", "2020", "2021", "2022"}
	fName := uData[0]
	for _, data := range uData {
		if data == fName {
			continue
		} else {
			password := fName + "@" + data
			pass = append(pass, password)
			passunder := fName + "_" + data
			pass = append(pass, passunder)
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
			passunder := fNameU + "_" + data
			pass = append(pass, passunder)
		}
	}
	for _, year := range years {
		password := fName + "@" + year
		pass = append(pass, password)
		passunder := fName + "_" + year
		pass = append(pass, passunder)
		passwords := fNameU + "@" + year //second password for the years
		pass = append(pass, passwords)
		passunders := fNameU + "_" + year
		pass = append(pass, passunders)
	}
	fmt.Print(pass)

} //create a special function --
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
		underAtthe(userData)

	} else {
		fmt.Println("\033[31;1m Invalid argument provided ")
		help()
	}
}
