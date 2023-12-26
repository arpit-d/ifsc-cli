/*
Copyright © 2023 ARPIT DESHPANDE <arpitme199@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Ifsc struct {
	IFSC string `json:"IFSC"`

	BankCode string `json:"BANKCODE"`
}

// getDetailsFromIFSCCmd represents the getDetailsFromIFSC command
var getDetailsFromIFSCCmd = &cobra.Command{
	Use:   "getDetailsFromIFSC",
	Short: "Get bank details from IFSC code",
	Long:  `Get a list of details of a bank branch including the location, bank code, MICR, bank name etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		ifscInput := ""
		fmt.Println("Enter the IFSC code")
		fmt.Scanln(&ifscInput)

		url := "https://ifsc.razorpay.com/" + ifscInput

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error while making request: ", err)
		}

		defer response.Body.Close()

		var ifsc Ifsc
		err = json.NewDecoder(response.Body).Decode(&ifsc)
		if err != nil {
			fmt.Println("Error while decoding JSON: ", err)
			return
		}
		fmt.Println("IFSC: ", ifsc.IFSC)
		fmt.Println("BANKCODE: ", ifsc.BankCode)

	},
}

func init() {
	rootCmd.AddCommand(getDetailsFromIFSCCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDetailsFromIFSCCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getDetailsFromIFSCCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
