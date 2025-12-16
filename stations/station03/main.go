package main

import "fmt"

func main() {
	books := []Book{
		{
		Title:  "Go入門",
		Author: "山田",
		Price:  3000,
		Categories: []string{"programming", "go"},
	},
	{
		Title:  "ネットワーク基礎",
		Author: "佐藤",
		Price:  2500,
		Categories: []string{"network"},
	},
} 

	totalPrice := Total(books)
	fmt.Println("合計", totalPrice, "円")

	authorBooksMap := ToMap(books)
	authorName := "山田"
	fmt.Println("著者: ", authorName, ", タイトル: ", authorBooksMap[authorName].Title)
}

// 構造体 Book を定義 ---
type Book struct {
	Title string // Title ... 文字列、本のタイトル
	Author string // Author ... 文字列、著者
	Price int // Price ... 整数、価格
	Categories []string // Categories ... 文字列のスライス、カテゴリ
}

// Bookの配列を引数とし、本の合計価格を戻り値とする関数 `Total` を実装する
func Total(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.Price
	}
	return  total

}

// 3. キーを「著者名 (Author)」、値を構造体 Book とするマップを戻り値とする関数 `ToMap` を実装する
func ToMap(books []Book) map[string]Book {
	result := make(map[string]Book)
	for _, book := range books {
		result[book.Author] = book
	}
	return result
}
