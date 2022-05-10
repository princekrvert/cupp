package main

import (
	"fmt"
	"os"
	"strings"
)

/// store all the password in the data slice
//var data []string
// create the @ function
func underAtthe(uData []string) []string {
	var pass []string
	// want to add or remove years in password chenge the years slice
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
	return pass

} //create a special function --
func specialD(data []string, mob string) []string {
	var localpass []string
	// 123 pass
	fName := data[0]
	lName := data[1]
	localpass = append(localpass, fName+"123")
	localpass = append(localpass, fName+"123456")
	localpass = append(localpass, fName+"iloveyou")
	localpass = append(localpass, fName+"abc")
	localpass = append(localpass, fName+"")
	localpass = append(localpass, mob)
	localpass = append(localpass, "password")
	localpass = append(localpass, "abc123")
	localpass = append(localpass, "1234567890")
	localpass = append(localpass, "qwerty")
	localpass = append(localpass, string(fName[0]))
	localpass = append(localpass, string(fName[0])+string(lName[0]))
	return localpass
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
		var fName, lName, dob, aadhar, num, pLName, pFName string
		var userData []string
		fmt.Print("\033[33;1m First name of the target: ")
		fmt.Scanf("%s", &fName)
		if isEmpty(fName) {
			panic("Sorry first name can not be empty ")
		}
		fmt.Print("\033[33;1m Last name of the target: ")
		fmt.Scanf("%s", &lName)
		if isEmpty(lName) {
			panic("Sorry last name can not be empty ")
		}
		fmt.Print("\033[33;1m Date of birth(ddmmyyyy): ")
		fmt.Scanf("%s", &dob)
		fmt.Print("\033[33;1m Target mobile number without 91 :")
		fmt.Scanf("%s", &num)
		fmt.Print("\033[33;1m First four digit aadhar number: ")
		fmt.Scanf("%s", &aadhar)
		fmt.Print("\033[33;1m First name of the target partner: ")
		fmt.Scanf("%s", &pFName)
		fmt.Print("\033[33;1m Last name of the target partner: ")
		fmt.Scanf("%s", &pLName)
		// create the dob
		if isEmpty(dob) {
			dobyear := ""
			userData = append(userData, fName, lName, dobyear, aadhar, num)
		} else {
			dobyear := dob[4:]
			userData = append(userData, fName, lName, dobyear, aadhar, num)
		}
		// call the underscore and at the rate function ----
		passlist1 := underAtthe(userData)
		// call the special function
		var sData []string
		sData = append(sData, fName, lName)
		fmt.Print(passlist1)

	} else {
		fmt.Println("\033[31;1m Invalid argument provided ")
		help()
	}
}

// create a function to check empty var
func isEmpty(variable string) bool {
	if variable == "" {
		return true
	} else {
		return false
	}
}
