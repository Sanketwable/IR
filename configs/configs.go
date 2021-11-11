package configs

import (
	"ir/hmap"
	"ir/models"
	"ir/trie"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var Size int
var W models.Words

func ReadConfig() {
	fmt.Printf("Reading a file\n")
	fileName := "config.json"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	json.Unmarshal(data, &W)
	hmap.StoreintoMap(W)
	trie.StoreintoTrie(W)
	Size = len(W.Words)
}

func StoreWord() {
	file, _ := json.MarshalIndent(W, "", " ")
	_ = ioutil.WriteFile("config.json", file, 0644)

	ReadConfig()
}
