package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	admin    = 1
	weight   = 1<<5 + 1<<4 + 1<<3 + 1<<2 + 1<<1
	itsMayor = 1 << 6
	itsMale  = 1 << 7
	itsHuman = 1 << 8
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

	if message&admin == admin {
		fmt.Println("its admin")
	}

	if message&itsMayor == itsMayor {
		fmt.Println("is mayor de edad")
	} else {
		fmt.Println("isn't mayor de edad")
	}

	if message&itsMale == itsMale {
		fmt.Println("is male")
	} else {
		fmt.Println("is female")
	}

	if message&itsHuman == itsHuman {
		fmt.Println("is human")
	} else {
		fmt.Println("is animal")
	}

	// fmt.Println(weight & 10)
	// currentWeight := strconv.FormatInt(int64((weight&message)>>1), 2)

	currentWeight := (weight & message) >> 1
	fmt.Printf("Your weight is %dkg\n", currentWeight)
}
