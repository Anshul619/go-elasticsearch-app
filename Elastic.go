package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var cfg = elasticsearch.Config{
	Addresses: []string{
		"http://localhost:9200",
	},
	Username: "elastic",
	Password: "XF8mbiscAFQ4jZWycx5q",
}
var esClient, _ = elasticsearch.NewClient(cfg)

func Exit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func ReadText(reader *bufio.Scanner, prompt string) string {
	fmt.Print(prompt + ": ")
	reader.Scan()
	return reader.Text()
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("0) Exit")
		fmt.Println("1) Load spacecraft")
		fmt.Println("2) Get spacecraft")
		fmt.Println("3) Search spacecraft by key and value")
		fmt.Println("4) Search spacecraft by key and prefix")
		option := ReadText(reader, "Enter option")
		if option == "0" {
			Exit()
		} else if option == "1" {
			LoadData()
		} else if option == "2" {
			GetData(reader)
			fmt.Println("Invalid option")
		} else if option == "3" {
			Search(reader, "match")
		} else if option == "4" {
			Search(reader, "prefix")
		} else {
			fmt.Println("Invalid option")
		}
	}
}

func getESInfo() {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "elastic",
		Password: "XF8mbiscAFQ4jZWycx5q",
	}
	esClient, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("error creating the elasticsearch : %s", err)
	}

	res, err := esClient.Info()

	if err != nil {
		log.Fatalf("error getting response from elasticsearch: %s", err)
	}

	log.Println(res)
	res.Body.Close()
}
