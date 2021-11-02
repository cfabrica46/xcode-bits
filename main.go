package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	admin    = 1
	weight   = 62
	itsMayor = 64
	itsMale  = 128
	itsHuman = 256
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

	if message > 511 || message < 0 {
		log.Fatal("arg invalid")
	}

	fmt.Println(strconv.FormatInt(int64(message), 2))

	if message&admin == admin {
		fmt.Println("its admin")
	}

	fmt.Println(strconv.FormatInt(int64(message&weight), 2))
	switch message & weight {
	case 2:
		fmt.Println("your weight its 1 kg")
	case 4:
		fmt.Println("your weight its 10 kg")
	case 6:
		fmt.Println("your weight its 11 kg")
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
}
