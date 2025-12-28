/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

//処理のイメージ
//	犬種を指定。その犬種の画像URLを取得して出力。
//	例: dog-cli breed Akita
//	①犬種指定されていない場合のエラーハンドリング
//	Error: 犬種が指定されていません
//	②APIからのレスポンスがエラーの場合のエラーハンドリング
//	Error: 無効な犬種です
//	③正常にレスポンスが返ってきた場合、画像URLを出力
//	Output:https://dog.ceo/api/breed/akita/images/random

type DogAPIResponse struct {
	Message string `json:"message"`
	Status string `json:"status"`
}

func FetchBreedDogImageURL(breed string) (string, error) {
	url := fmt.Sprintf("https://dog.ceo/api/breed/%s/images/random", breed)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("無効な犬種です")
	}
	var result DogAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "",  err
	}
	return result.Message, nil
}

// breedCmd represents the breed command
var breedCmd = &cobra.Command{
	Use:   "breed",
	Short: "犬種を指定し、その犬種の画像 URL を取得する",
	Long: `犬種を指定し、その犬種の画像 URL を取得する`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("犬種が指定されていません")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		breed := strings.ToLower(args[0])

		url, err := FetchBreedDogImageURL(breed)
		if err != nil {
			return err
		}
		fmt.Println(url)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(breedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// breedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// breedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
