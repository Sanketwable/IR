package trie


import(
	"ir/models"
)

const (
	 ALBHABET_SIZE = 255
 )
 
 type TrieNode struct {
	 childrens [ALBHABET_SIZE]*TrieNode
	 isWordEnd bool
 }
 
 type trie struct {
	 root *TrieNode
 }
 
 func InitTrie() *trie {
	 return &trie{
		 root: &TrieNode{},
	 }
 }

 func StoreintoTrie(W models.Words) {
	trie := InitTrie()
	for _, word := range W.Words {
		trie.Insert(word)
	}
}
 
 func (t *trie) Insert(word string) {
	 wordLength := len(word)
	 current := t.root
	 for i := 0; i < wordLength; i++ {
		 index := word[i] - 'a'
		 if current.childrens[index] == nil {
			 current.childrens[index] = &TrieNode{}
		 }
		 current = current.childrens[index]
	 }
	 current.isWordEnd = true
 }
 
 func (t *trie) Find(word string) bool {
	 wordLength := len(word)
	 current := t.root
	 for i := 0; i < wordLength; i++ {
		 index := word[i] - 'a'
		 if current.childrens[index] == nil {
			 return false
		 }
		 current = current.childrens[index]
	 }
	 if current.isWordEnd {
		 return true
	 }
	 return false
 }