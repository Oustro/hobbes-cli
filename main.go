package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ApiResponse struct {
	Response string `json:"response"`
}

// sudo mv hobbes /usr/local/bin

func main() {
	fmt.Printf("\033[38;2;%d;%d;%dm", 104, 81, 255)
	fmt.Println()

	if len(os.Args) == 1 {
		fmt.Println("Howdy I am Hobbes, the Official CLI tool for Oustro, LLC.")
		fmt.Println()
		return
	}

	url := ""

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	if os.Args[1] == "-zu" {

		url = os.Getenv("API_SECRET_1")

	} else if os.Args[1] == "-zi" {

		url = os.Getenv("API_SECRET_2")

	} else if os.Args[1] == "-zo" {

		url = os.Getenv("API_SECRET_3")

	} else {
		fmt.Printf("\033[38;2;%d;%d;%dm", 56, 182, 255)
		fmt.Println("Hobbes Commands are:")
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("| 1) -zu: Wondering how many users Ziggy has?                        |")
		fmt.Println("|                                                                    |")
		fmt.Println("| 2) -zi: Wondering how many interviews have been made using Ziggy?  |")
		fmt.Println("|                                                                    |")
		fmt.Println("| 3) -zo: Wondering how many organizaitons are using Ziggy?          |")
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println()
		fmt.Printf("\033[38;2;%d;%d;%dm", 255, 255, 255)
		fmt.Println("example: '$ hobbes -zu'")
		fmt.Println()
		return
	}

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\033[38;2;%d;%d;%dm", 56, 182, 255)
	fmt.Println(apiResponse.Response)
	fmt.Println()
}
