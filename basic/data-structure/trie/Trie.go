package trie

// 单词前缀树
type Trie struct {
	value string
	children map[string]*Trie
	isWord bool
}

// 初始化构造器
func Constructor() Trie {
	return Trie{value: "", children: map[string]*Trie{}, isWord: false}
}

func (this *Trie) Insert(word string) {
	t := this
	for _, v := range word {
		str := string(v)

		if _, ok := t.children[str]; ok {
			t = t.children[str]
		} else {
			t.children[str] = &Trie {
				value: str,
				children: map[string]*Trie{},
				isWord: false,
			}
			t = t.children[str]
		}
	}

	t.isWord = true
}

// 查询单词是否在前缀树中
func (this *Trie) Search(word string) bool {
	t := this
	for _, v := range word {
		if _, ok := t.children[string(v)]; ok {
			t = t.children[string(v)]
		} else {
			return false
		}
	}
	return t.isWord
}

// 查询树中是否存在单词是 prefix 开头的
func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for _, v := range prefix {
		if _, ok := t.children[string(v)]; ok {
			t = t.children[string(v)]
		} else {
			return false
		}
	}

	return true
}
