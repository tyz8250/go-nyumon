/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// BreedListResponseは Dog API のレスポンスを表す構造体
type BreedListResponse struct {
	Message []string `json:"message"`
	Status string `json:"status"`
}

func FetchBreedListResponse() ([]string, error) {
	resp, err := http.Get("https://dog.ceo/api/breeds/list")
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Message []string `json:"message"`
		Status  string `json:"status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil{
		return nil, err
	}
	return result.Message, nil
}


// breedListCmd represents the breedList command
var breedListCmd = &cobra.Command{
	Use:   "breed-list",
	Short: "指定したURLから犬種リストを取得して表示する",
	Long: `指定したURLから犬種リストを取得して表示する`,
	Run: func(cmd *cobra.Command, args []string) {
	breeds, err := FetchBreedListResponse()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, breed := range breeds {
		fmt.Println(breed)
	}
},
}

func init() {
	rootCmd.AddCommand(breedListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// breedListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// breedListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
