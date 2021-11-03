package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode allows you to decode an integer and give you your data in strings",
	Run: func(cmd *cobra.Command, args []string) {
		var valueToDecode int
		var valuePipe string
		var err error

		scanner := bufio.NewScanner(os.Stdin)

		if len(args) == 0 {
			fmt.Println("a")
			if scanner.Scan() {
				fmt.Println("c")
				valuePipe = scanner.Text()
				valueToDecode, err = strconv.Atoi(valuePipe)
				if err != nil {
					log.Fatal(err)
				}
			}
			fmt.Println("b")
			if err := scanner.Err(); err != nil {
				fmt.Printf("Value: [ 1 - 4095 ]\n> ")
				_, err = fmt.Scan(&valueToDecode)
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			valueToDecode, err = strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
		}

		if valueToDecode >= 1<<12 || valueToDecode <= 0 {
			log.Fatal("arg invalid")
		}

		fmt.Printf("Original Value is %s\n", strconv.FormatInt(int64(valueToDecode), 2))

		if valueToDecode&isHuman == isHuman {
			fmt.Printf("Is a human or animal?\t%s\n", "HUMAN")
		} else {
			fmt.Printf("Is a human or animal?\t%s\n", "ANIMAL")
		}

		if valueToDecode&isMale == isMale {
			fmt.Printf("Is a male or female?\t%s\n", "MALE")
		} else {
			fmt.Printf("Is a male or female?\t%s\n", "FEMALE")
		}

		fmt.Printf("Is of legal age:\t%t\n", valueToDecode&isLegal == isLegal)

		currentWeight := yourWeight & valueToDecode >> 1
		fmt.Printf("Your weight is:\t\t%dkg\n", currentWeight)

		fmt.Printf("Is an administrator:\t%t\n", valueToDecode&isAdmin == isAdmin)
	},
}
