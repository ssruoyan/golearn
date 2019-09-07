package main

type Trie struct {
	list []string
}

func Constructor() Trie {
	return Trie{list: []string{}}
}

func (t Trie) Insert(s string) {
	t.list = append(t.list, s)
}

func (t Trie) Search(s string) bool {

}

func (t Trie) StartsWith(s string) bool {

}