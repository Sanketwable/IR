package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ir/configs"
	"ir/hmap"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
)

type response struct {
	Words []string `json:"words"`
}
type request struct {
	Word string `json:"word"`
}

func GetWord(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	word := request{}
	json.Unmarshal(body, &word)

	w := word.Word
	var words []string
	res := response{}
	if hmap.FindWord(word.Word) {
		res.Words = append(res.Words, word.Word)
		return c.JSON(http.StatusAccepted, res)
	} else if w[0] == '*' && w[len(w)-1] == '*' {
		wrd := w[1:len(w)-1]
		words = getSubstr(wrd)
		
	} else if w[0] == '*' {
		
		wrd := w[1:]
		words = getSufixstr(wrd)
	} else if w[len(w)-1] == '*' {
		
		wrd := w[0:len(w)-1]
		fmt.Println("here", wrd)
		words = getPrefixstr(wrd)
	} else {
		countstar := 0
		var index int
		for i, c := range w {
			if c == '*' {
				countstar++
				index = i
				break
			}
		}
		if countstar == 0 {
			words = getSimilarWord(word.Word)
		} else {
			startword := w[0:index]
			endword := w[index+1:]
			fmt.Println("startword = ", startword)
			fmt.Println("endword = ", endword)
			prefixword := getPrefixstr(startword)
			suffixword := getSufixstr(endword)
			preffixmap := make(map[string]bool)
			suffixmap := make(map[string]bool)
			for _, word := range prefixword {
				preffixmap[word] = true
			}
			for _, word := range suffixword {
				suffixmap[word] = true
			}
			for word := range preffixmap {
				if suffixmap[word] {
					words = append(words, word)
				}
			}
		}
 	}
	
	sort.Strings(words)
	
	res.Words = append(res.Words, words...)
	c.Request().Response.Header.Add("Access-Control-Allow-Origin", "*")
	c.Request().Header.Add("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusAccepted, res)
}

func getSimilarWord(word string) []string {
	mp := make(map[int]([]string))
	for _, w := range configs.W.Words {
		edistance := Editdistance(word, w)
		mp[edistance] = append(mp[edistance], w)
	}
	var words []string
	var ints []int
	for editd := range mp {
		ints = append(ints, editd)
	}
	sort.Ints(ints)
	eD := ints[0]
	words = append(words, mp[ints[0]]...)
	if len(ints) >=2 {
		words = append(words, mp[ints[1]]...)
	}
	var finalwords []string
	for _, w := range words {
		if w[0] == word[0] ||  eD == 1 {
			finalwords = append(finalwords, w)
		}
	}

	return finalwords
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
	configs.W.Words = append(configs.W.Words, str)
	configs.StoreWord()
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
	for _, word := range configs.W.Words {
		if strStr(word, str) != -1 {
			words = append(words, word)
		}
	}
	return words
}

func getPrefixstr(str string) []string {
	var words []string
	for _, word := range configs.W.Words {
		match := true
		for i := range str {
			if i >= len(word) {
				match = false
				break
			}
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
	for _, word := range configs.W.Words {
		match := true
		lenword := len(word)
		for i := range str {
			if i >= lenword {
				match = false
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

func Mins(value int, values ...int) int {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

func Editdistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			offset := 1
			if word1[i-1] == word2[j-1] {
				offset = 0
			}
			dp[i][j] = Mins(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+offset)
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
