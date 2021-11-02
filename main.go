package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cfabrica46/xcode-bits/cmd"
)

//The project handles a 9-bit scheme: 000000000
const (
	//admin. bits 9
	isAdmin = 1 << 0
	//weight. bits 4:8
	weight = 1<<5 + 1<<4 + 1<<3 + 1<<2 + 1<<1
	//admin. bits 3
	isLegal = 1 << 6
	//admin. bits 2
	isMale = 1 << 7
	//admin. bits 1
	isHuman = 1 << 8
)

func main() {

	cmd.Execute()

	if cmd.GetWillBeDecoded() {

		/* args := os.Args
		if len(args) < 2 {
			log.Fatal("arg its necesary")
		} */

		arg := cmd.GetValueToDecode()

		valueToDecode, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Original Value is %s\n", strconv.FormatInt(int64(valueToDecode), 2))

		if valueToDecode >= 1<<9 || valueToDecode <= 0 {
			log.Fatal("arg invalid")
		}

		fmt.Printf("Is a human:\t\t%t\n", valueToDecode&isHuman == isHuman)

		fmt.Printf("Is a animal:\t\t%t\n", valueToDecode&isHuman == 0)

		//---
		fmt.Printf("Is a male:\t\t%t\n", valueToDecode&isMale == isMale)

		fmt.Printf("Is a female:\t\t%t\n", valueToDecode&isMale == 0)

		//---
		fmt.Printf("Is of legal age:\t%t\n", valueToDecode&isLegal == isLegal)

		//---
		currentWeight := (weight & valueToDecode) >> 1
		fmt.Printf("Your weight is:\t\t%dkg\n", currentWeight)

		//---
		fmt.Printf("Is an administrator:\t%t\n", valueToDecode&isAdmin == isAdmin)

	} else {
		var result int

		currentHuman, currentMale, currentLegal, currentAdmin, currentWeight := cmd.GetValuesToEncode()

		// fmt.Println(currentHuman, currentMale, currentLegal, currentAdmin, currentWeight)

		if currentHuman {
			result += isHuman
		}

		if currentMale {
			result += isMale
		}

		if currentLegal {
			result += isLegal
		}

		result += currentWeight << 1

		if currentAdmin {
			result += isAdmin
		}

		fmt.Printf("Result: %d\n", result)
	}
}
