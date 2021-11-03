package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var human bool
var male bool
var legal bool
var weight float32
var administrator bool

//The project handles a 9-bit scheme: 000000000
const (
	//admin. bits 9
	isAdmin = 1 << 0
	//weight. bits 4:8
	yourWeight = 1<<5 + 1<<4 + 1<<3 + 1<<2 + 1<<1
	//admin. bits 3
	isLegal = 1 << 6
	//admin. bits 2
	isMale = 1 << 7
	//admin. bits 1
	isHuman = 1 << 8
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(decodeCmd)
	rootCmd.AddCommand(encodeCmd)

	encodeCmd.Flags().BoolVarP(&human, "isHuman", "H", false, "is a human or animal?")
	encodeCmd.Flags().BoolVarP(&male, "isMale", "M", false, "is a male or female?")
	encodeCmd.Flags().BoolVarP(&legal, "isLegal", "L", false, "is of legal age?")
	encodeCmd.Flags().Float32VarP(&weight, "weight", "W", 0, "how much it weighs (max weight is 31)")
	encodeCmd.PersistentFlags().BoolVarP(&administrator, "isAdministrator", "A", false, "is an administrator?")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
