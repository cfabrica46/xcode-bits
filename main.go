package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//The project handles a 9-bit scheme: 000000000
const (
	//admin. bits 9
	admin = 1 << 0
	//weight. bits 4:8
	weight = 1<<5 + 1<<4 + 1<<3 + 1<<2 + 1<<1
	//admin. bits 3
	isMayor = 1 << 6
	//admin. bits 2
	isMale = 1 << 7
	//admin. bits 1
	isHuman = 1 << 8
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("arg its necesary")
	}

	message, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Original Value is %s\n", strconv.FormatInt(int64(message), 2))

	if message >= 1<<9 || message < 0 {
		log.Fatal("arg invalid")
	}

	fmt.Printf("Is a human:\t\t%t\n", message&isHuman == isHuman)

	fmt.Printf("Is a animal:\t\t%t\n", message&isHuman == 0)

	//---
	fmt.Printf("Is a male:\t\t%t\n", message&isMale == isMale)

	fmt.Printf("Is a female:\t\t%t\n", message&isMale == 0)

	//---
	fmt.Printf("Is of legal age:\t%t\n", message&isMayor == isMayor)

	//---
	currentWeight := (weight & message) >> 1
	fmt.Printf("Your weight is:\t\t%dkg\n", currentWeight)

	//---
	fmt.Printf("Is an administrator:\t%t\n", message&admin == admin)

}
