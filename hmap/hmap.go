package hmap

import (
	"ir/models"
)

var Dictionary map[string]bool

func StoreintoMap(W models.Words) {
	Dictionary = make(map[string]bool)
	for _, word := range W.Words {
		if word != "" {
			Dictionary[word] = true
		}
	}
	// fmt.Println("dictionary is ", Dictionary)

}

func FindWord(word string) bool {
	return Dictionary[word]
}
