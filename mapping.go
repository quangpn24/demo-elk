package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type Student struct {
	Name  string  `json:"name"`
	Age   int64   `json:"age"`
	Score float64 `json:"score"`
}

func main() {

	client, _ := elasticsearch.NewDefaultClient()
	CreateIndex(client, "test")
	CreateDocument(client, "test")
	return
}
func CreateDocument(client *elasticsearch.Client, indexName string) {
	data := []Student{
		{
			Name:  "Nguyễn Văn A",
			Age:   20,
			Score: 10,
		},
		{
			Name:  "Nguyễn Thị A",
			Age:   22,
			Score: 8,
		},
		{
			Name:  "Trần D",
			Age:   30,
			Score: 7.5,
		},
		{
			Name:  "Thạch",
			Age:   32,
			Score: 8,
		},
	}
	for _, d := range data {
		body, _ := json.Marshal(d)
		client.Index(indexName, bytes.NewReader(body))
	}
}
func CreateIndex(client *elasticsearch.Client, indexName string) error {
	mapping := `{
    	"settings": {
			"number_of_shards": 3,
			"number_of_replicas": 1,
			"analysis": {
			"analyzer": {
			  "my_analyzer": {
				"type": "custom",
				"tokenizer": "my_tokenizer",
				"filter": [
				  "asciifolding",
				  "lowercase"
				]
			  }
			},
			"tokenizer": {
			  "my_tokenizer": {
				"type": "ngram",
				"min_gram": 2,
				"max_gram": 3
			  }
			}
		  }
      	},
      	"mappings": {
        	"properties": {
          		"name": {
            		"type": "text",
					"analyzer": "my_analyzer"
          		},
	 			"age": {
            		"type": "integer"
          		}
			}
        }
	}`
	res, err := client.Indices.Create(indexName, client.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	return err
}
