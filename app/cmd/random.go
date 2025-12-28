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

// DogImageResponse は Dog API のレスポンスを表す構造体
type DogImageResponse struct {
	Message string `json:"message"`
	Status string `json:"status"`
}

func FetchRandomDogImageURL() (string, error) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil{
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil{
		return "", err
	}

	return result.Message, nil
}

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "ランダムに画像 URL を取得し、その URL を出力する",
	Long: `Dog API からランダムに画像 URL を取得し、その URL を出力する`,
	
	// 犬の画像 URL をランダムに取得して出力し、URLを返す処理
	RunE: func(cmd *cobra.Command, args []string) error {
		url, err := FetchRandomDogImageURL()
		if err != nil {
			return err
		}
		fmt.Println(url)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
