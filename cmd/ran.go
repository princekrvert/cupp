/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// make a funciton to check if password is exist in a slice or not
func isAvailable(passlist []string, pass string) bool {
	for _, password := range passlist {
		if password == pass {
			return true
		} else {
			continue
		}
	}
	return false
}

// create a function to write the file
func writePass(passl []string) {
	for _, data := range passl {
		file, err := os.OpenFile("ranpass.txt", os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Fprintln(file, data)
		}

		file.Close()
	}
	fmt.Println("\033[32;1m [~]Password saved in ranpass.txt file")
}

// ranCmd represents the ran command
var ranCmd = &cobra.Command{
	Use:   "ran",
	Short: "This will genarate a random password list",
	Long:  `Enter the length of the password list and the number of password to generate the password.`,
	Run: func(cmd *cobra.Command, args []string) {
		// first check the lenght of the password ..
		var maxnum string
		minnum := "1"
		passlen, _ := strconv.Atoi(args[0])
		for i := 1; i <= passlen; i++ {
			if i == 1 {
				maxnum = ""
			} else {
				minnum = minnum + "0"
			}

			maxnum = maxnum + "9"

		}
		// convert max num and min num to integer
		maxNum, _ := strconv.Atoi(maxnum)
		minNum, _ := strconv.Atoi(minnum)
		passMax, _ := strconv.Atoi(args[1])
		if passMax > maxNum-minNum {
			fmt.Println("\033[31;1m Kiya hacker banaga re tu")
			os.Exit(1)
		}
		// now we got the max and min number now generate the random number ...
		rand.Seed(time.Now().UnixNano())
		// now create a slice to store the pass and keep tack of previous pass
		var passlist = []string{}

		ranpass := rand.Intn(maxNum-minNum) + minNum
		// convert ranpass to string and push into the passlit slice
		ranPass := strconv.Itoa(ranpass)
		passlist = append(passlist, ranPass)
		for I := 1; I < passMax; I++ {
			ranpass := rand.Intn(maxNum-minNum) + minNum
			ranPass := strconv.Itoa(ranpass)
			if isAvailable(passlist, ranPass) {
				continue
			} else {
				passlist = append(passlist, ranPass)
			}
		}
		file, err := os.Create("ranpass.txt")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Creating file ranpass.txt")
		}
		file.Close()
		writePass(passlist)

	},
}

func init() {
	rootCmd.AddCommand(ranCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ranCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ranCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
