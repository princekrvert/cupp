/*
Copyright © 2023 prince kumar

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/princekrvert/cupp/word"
	"github.com/spf13/cobra"
)

// create a function to check if variable is empty
func isEmpty(data string) bool {
	if data == "" {
		return true
	} else {
		return false
	}

}

//Create a banner for this tool..
func banner() {
	myFigure := figure.NewColorFigure("cupp", "", "red", true)
	myFigure.Print()
	fmt.Println("\033[32;1m MADE BY PRINCE")

}

// create a function to write the file
func writepass(passl []string, value string) {
	for _, data := range passl {
		file, err := os.OpenFile(value+".txt", os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Fprintln(file, data)
		}
		file.Close()
	}

}

//crea a similar function ro namesufiix but don't create a file
func namePsuffix(fName, lName string) []string {
	passList := []string{}
	// want to add or remove years in password chenge the years slice
	years, _ := word.Eachword("files/year.txt")
	for _, data := range years {
		passList = append(passList, fName+data)
		passList = append(passList, fName+"_"+data)
		passList = append(passList, fName+"@"+data)
	}
	ranword, _ := word.Eachword("files/indian.txt")
	for _, word := range ranword {
		passList = append(passList, fName+word)
	}
	passList = append(passList, fName+lName)
	passList = append(passList, lName+fName)

	return passList
}

// make a function with the pername ...
func Petpass(Petname, other string) []string {
	pass9 := []string{}
	// write the combination here ...
	petpass, _ := word.Eachword("files/pet.txt")
	for _, petword := range petpass {
		pass9 = append(pass9, Petname+petword)
		pass9 = append(pass9, petword+Petname)
	}
	pass9 = append(pass9, Petname+other)
	pass9 = append(pass9, other+Petname)
	pass9 = append(pass9, Petname+"and"+other)
	pass9 = append(pass9, other+"and"+Petname)
	return pass9
}

// cteate a function related to village
func village(vName, fname, lname string) []string {
	passlist := []string{}
	// now make some combination ...
	passlist = append(passlist, vName+fname)
	passlist = append(passlist, vName+lname)
	passlist = append(passlist, fname+vName)
	passlist = append(passlist, lname+vName)
	passlist = append(passlist, fname+"@"+vName)
	passlist = append(passlist, lname+"@"+vName)
	passlist = append(passlist, fname+"_"+vName)
	passlist = append(passlist, lname+"_"+vName)
	return passlist
}

// create a function to make a function to make the list of all suffix with first name and last name , and also with the year
func nameSuffix(fName, lName string) []string {
	passList := []string{}
	// here open the file as read mode
	file, err := os.Create(fName + ".txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Creating file...")
	}
	file.Close()
	//file is closed here
	// want to add or remove years in password chenge the years slice
	years, _ := word.Eachword("files/year.txt")
	for _, data := range years {
		passList = append(passList, fName+data)
		passList = append(passList, fName+"_"+data)
		passList = append(passList, fName+"@"+data)
	}
	randword, _ := word.Eachword("files/indian.txt")
	for _, word := range randword {
		passList = append(passList, fName+word)
	}
	passList = append(passList, fName+lName)
	passList = append(passList, lName+fName)

	return passList
}

// create a special function to make the passlist 2
func special(fName, lName, mob, dob string) []string {
	var passlist2 = []string{}
	passlist2 = append(passlist2, mob)
	passlist2 = append(passlist2, fName+mob[:5])
	passlist2 = append(passlist2, fName+dob)
	passlist2 = append(passlist2, lName+dob)
	passlist2 = append(passlist2, fName+dob[4:])
	return passlist2
}

// intractCmd represents the intract command
var intractCmd = &cobra.Command{
	Use:   "intract",
	Short: "Start cupp in classic mode to create user based passlist.",
	Long:  `Classic version of cupp taken name dob and some other information to create user based password make sure you have it. Thanks`,
	Run: func(cmd *cobra.Command, args []string) {
		banner()
		// ask for the requre options about the target
		// pronpt for the fname
		var fName, lName, dob, place, pFname, pLname, mob, pDob, petname string
		fmt.Printf("Enter the target first name : ")
		fmt.Scanf("%s", &fName)
		if isEmpty(fName) {
			fmt.Println("\033[31;1m First name can not be empty")
			os.Exit(1)
		}
		fmt.Printf("Enter the target lname name : ")
		fmt.Scanf("%s", &lName)
		if isEmpty(lName) {
			fmt.Println("\033[31;1m Last name can not be empty")
			os.Exit(1)
		}

		fmt.Printf("Enter phone number of target : ")
		fmt.Scanf("%s", &mob)
		fmt.Printf("Enter the target dob (ddmmyyyyy) : ")
		fmt.Scanf("%s", &dob)
		fmt.Printf("Address of the terget like village/city/town : ")
		fmt.Scanf("%s", &place)
		fmt.Printf("Enter first name of the target partner : ")
		fmt.Scanf("%s", &pFname)
		fmt.Printf("Enter last name of the target partner : ")
		fmt.Scanf("%s", &pLname)
		fmt.Printf("Enter dob of target partner : ")
		fmt.Scanf("%s", &pDob)
		fmt.Printf("Enter pet name of the target : ")
		fmt.Scanf("%s", &petname)
		pass1 := nameSuffix(fName, lName)
		writepass(pass1, fName)
		// check if number and dob is provided
		if isEmpty(mob) {
		} else {
			if isEmpty(dob) {

			} else {
				//now call the special function
				pass2 := special(fName, pLname, mob, dob)
				writepass(pass2, fName)
			}
		}
		// check if the partner first name,last name dob is provided
		if isEmpty(pFname) || isEmpty(pLname) || isEmpty(pDob) {
			// do nothing
		} else {
			pass3 := namePsuffix(pFname, pLname)
			writepass(pass3, fName)
			pass4 := special(pFname, pLname, mob, pDob)
			writepass(pass4, fName)
		}
		// check if village name is empty
		if isEmpty(place) {
			//do nothing
		} else {
			pass5 := village(place, fName, lName)
			writepass(pass5, fName)
			pass6 := village(place, pFname, pLname)
			writepass(pass6, fName)
		}
		// create a slice for indian password and write down into the file ..
		indianPass, _ := word.Eachword("files/indianpass.txt")
		writepass(indianPass, fName)
		// check if petname is empty..
		if isEmpty(petname) {
			//do nothing
		} else {
			pass := Petpass(petname, fName)
			writepass(pass, fName)
			//check if partner is empty..
			if isEmpty(pFname) {
				// do nothing
			} else {
				psss10 := Petpass(petname, pFname)
				writepass(psss10, fName)
			}
		}
		fmt.Print("Password file saved as ", fName, ".txt")

	},
}

func init() {
	rootCmd.AddCommand(intractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// intractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// intractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
