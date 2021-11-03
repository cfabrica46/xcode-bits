package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode will return an integer which will be the encoding of the data entered",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 0 {
			log.Fatal("arg invalid")
		}

		if weight <= 0 || weight > 31 {
			fmt.Printf("weight: [ 1 - 255 ]\n> ")
			_, err := fmt.Scan(&weight)
			if err != nil {
				log.Fatal(err)
			}
		}

		var result int

		var cliResponse string

		if human {
			result |= isHuman
		} else {
			fmt.Printf("You are human? [ Y / N ]\n> ")
			fmt.Scan(&cliResponse)
			cliResponse = strings.ToUpper(cliResponse)
			if cliResponse == "Y" {
				result |= isHuman
			}
		}

		if male {
			result |= isMale
		} else {
			fmt.Printf("You are male? [ Y / N ]\n> ")
			fmt.Scan(&cliResponse)
			cliResponse = strings.ToUpper(cliResponse)
			if cliResponse == "Y" {
				result |= isMale
			}
		}

		if legal {
			result |= isLegal
		} else {
			fmt.Printf("You are legal age? [ Y / N ]\n> ")
			fmt.Scan(&cliResponse)
			cliResponse = strings.ToUpper(cliResponse)
			if cliResponse == "Y" {
				result |= isLegal
			}
		}

		result |= weight << 1

		if administrator {
			result |= isAdmin
		} else {
			fmt.Printf("You are an administrator? [ Y / N ]\n> ")
			fmt.Scan(&cliResponse)
			cliResponse = strings.ToUpper(cliResponse)
			if cliResponse == "Y" {
				result |= isAdmin
			}
		}

		fmt.Printf("Result: %d\n", result)
	},
}
