package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var willBeDecoded bool
var valueToDecode string

var isHuman bool
var isMale bool
var isLegal bool
var weight int
var isAdministrator bool

var weightString string

func init() {
	rootCmd.AddCommand(decodeCmd)
	rootCmd.AddCommand(encodeCmd)

	encodeCmd.PersistentFlags().BoolVarP(&isHuman, "isHuman", "H", false, "is a human or animal?")
	encodeCmd.PersistentFlags().BoolVarP(&isMale, "isMale", "M", false, "is a male or female?")
	encodeCmd.PersistentFlags().BoolVarP(&isLegal, "isLegal", "L", false, "is of legal age?")
	encodeCmd.Flags().StringVarP(&weightString, "weight", "W", "0", "how much it weighs (max weight is 31)")
	// encodeCmd.MarkFlagRequired("weight")
	encodeCmd.PersistentFlags().BoolVarP(&isAdministrator, "isAdministrator", "A", false, "is an administrator?")
}

var rootCmd = &cobra.Command{Use: "app"}

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode allows you to decode an integer and give you your data in strings",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Decode needs value")
		}
		willBeDecoded = true
		valueToDecode = args[0]
	},
}

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode will return an integer which will be the encoding of the data entered",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		weight, err = strconv.Atoi(weightString)
		if err != nil {
			log.Fatal(err)
		}
		if weight <= 0 || weight > 31 {
			log.Fatal("weight invalid")
		}

		/* fmt.Println(isHuman)
		fmt.Println(isMale)
		fmt.Println(isLegal)
		fmt.Println(weight)
		fmt.Println(isAdministrator) */
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
