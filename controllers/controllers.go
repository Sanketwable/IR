package controllers

import (
	"IR/config"
	"IR/hashmap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Words []string `json:"words"`
}
type request struct {
	Word string `json:"word"`
}

func Substr(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	word := request{}
	json.Unmarshal(body, &word)
	words := getSubstr(word.Word)
	res := response{}
	res.Words = append(res.Words, words...)
	return c.JSON(http.StatusAccepted, res)
}

func PrefixStr(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	word := request{}
	json.Unmarshal(body, &word)
	words := getPrefixstr(word.Word)
	res := response{}
	res.Words = append(res.Words, words...)
	return c.JSON(http.StatusAccepted, res)
}

func Suffixstr(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	word := request{}
	json.Unmarshal(body, &word)
	words := getSufixstr(word.Word)
	res := response{}
	res.Words = append(res.Words, words...)
	return c.JSON(http.StatusAccepted, res)
}

func Findstr(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	word := request{}
	res := response{}
	json.Unmarshal(body, &word)
	fmt.Println(hashmap.Dictionary)
	if hashmap.FindWord(word.Word) {
		res.Words = append(res.Words, word.Word)
		return c.JSON(http.StatusAccepted, res)
	}

	return c.JSON(http.StatusAccepted, res)

}

func Addstr(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	word := request{}
	res := response{}
	json.Unmarshal(body, &word)
	AddWord(word.Word)
	res.Words = append(res.Words, word.Word)

	return c.JSON(http.StatusAccepted, res)
}

func AddWord(str string) {
	config.W.Words = append(config.W.Words, str)
	config.StoreWord()
}
func strStr(haystack string, needle string) int {
	l := len(needle)
	if l == 0 {
		return 0
	}
	next := make([]int, l)
	next[0] = -1
	k, i := -1, 1
	for i < l {
		if k == -1 || needle[i-1] == needle[k] {
			k++
			next[i] = k
			i++
		} else {
			k = next[k]
		}
	}
	m, n := 0, 0
	for m < len(haystack) && n < l {
		if haystack[m] == needle[n] {
			m++
			n++
		} else if n == 0 {
			m++
		} else {
			n = next[n]
		}
	}
	if n == l {
		return m - l
	}
	return -1
}

func getSubstr(str string) []string {
	var words []string
	for _, word := range config.W.Words {
		if strStr(word, str) != -1 {
			words = append(words, word)
		}
	}
	return words
}

func getPrefixstr(str string) []string {
	var words []string
	for _, word := range config.W.Words {
		match := true
		for i := range str {
			ch := str[i]
			wch := word[i]
			if ch != wch {
				match = false
			}
		}
		if match {
			words = append(words, word)
		}

	}
	return words
}

func getSufixstr(str string) []string {
	var words []string
	lenstr := len(str)
	for _, word := range config.W.Words {
		match := true
		lenword := len(word)
		for i := range str {
			if i >= lenword {
				break
			}
			ch := str[lenstr-1-i]
			wch := word[lenword-i-1]
			if ch != wch {
				match = false
			}
		}
		if match {
			words = append(words, word)
		}

	}
	return words
}
