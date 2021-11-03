package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode allows you to decode an integer and give you your data in strings",
	Run: func(cmd *cobra.Command, args []string) {
		var valueToDecode int
		var err error

		if len(args) == 0 {
			fmt.Printf("Value: [ 1 - 511 ]\n> ")
			_, err = fmt.Scan(&valueToDecode)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			valueToDecode, err = strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
		}

		if valueToDecode >= 1<<9 || valueToDecode <= 0 {
			log.Fatal("arg invalid")
		}

		fmt.Printf("Original Value is %s\n", strconv.FormatInt(int64(valueToDecode), 2))

		fmt.Printf("Is a human:\t\t%t\n", valueToDecode&isHuman == isHuman)

		fmt.Printf("Is a animal:\t\t%t\n", valueToDecode&isHuman == 0)

		fmt.Printf("Is a male:\t\t%t\n", valueToDecode&isMale == isMale)

		fmt.Printf("Is a female:\t\t%t\n", valueToDecode&isMale == 0)

		fmt.Printf("Is of legal age:\t%t\n", valueToDecode&isLegal == isLegal)

		currentWeight := yourWeight & valueToDecode >> 1
		fmt.Printf("Your weight is:\t\t%dkg\n", currentWeight)

		fmt.Printf("Is an administrator:\t%t\n", valueToDecode&isAdmin == isAdmin)

	},
}
