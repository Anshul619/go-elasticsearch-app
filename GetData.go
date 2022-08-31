package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/esapi"
)

func Print(spacecraft map[string]interface{}) {
	name := spacecraft["name"]
	status := ""
	if spacecraft["status"] != nil {

		status = "- " + spacecraft["status"].(string)
	}
	registry := ""
	if spacecraft["registry"] != nil {

		registry = "- " + spacecraft["registry"].(string)
	}
	class := ""
	if spacecraft["spacecraftClass"] != nil {

		class = "- " + spacecraft["spacecraftClass"].(map[string]interface{})["name"].(string)
	}
	fmt.Println(name, registry, class, status)
}

func GetData(reader *bufio.Scanner) {
	id := ReadText(reader, "Enter spacecraft ID")
	request := esapi.GetRequest{Index: "stsc", DocumentID: id}
	response, _ := request.Do(context.Background(), esClient)
	var results map[string]interface{}
	json.NewDecoder(response.Body).Decode(&results)
	Print(results["_source"].(map[string]interface{}))
}
