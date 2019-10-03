package trie

// 单词前缀树
type Trie struct {
	value byte
	children []*Trie
	isWord bool
}

// 初始化构造器
func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	t := this
	for i := 0; i < len(word); i++ {
		n := t.Match(word[i])
		if n != nil {
			t = n
		} else {
			newNode := Trie {
				value: word[i],
				children: []*Trie{},
				isWord: false,
			}
			t.children = append(t.children, &newNode)
			t = &newNode
		}
	}

	t.isWord = true
}

// 查询单词是否在前缀树中
func (this *Trie) Search(word string) bool {
	t := this

	for i := 0; i < len(word); i++ {
		n := t.Match(word[i])

		if n != nil {
			t = n
		} else {
			return false
		}
	}

	return t.isWord
}

// 查询树中是否存在单词是 prefix 开头的
func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for i := 0; i < len(prefix); i++ {
		n := t.Match(prefix[i])

		if n != nil {
			t = n
		} else {
			return false
		}
	}

	return true
}

func (this *Trie) Match(s byte) *Trie {
	t := this
	for _, v := range t.children {
		if v.value == s {
			return v
		}
	}
	return nil
}