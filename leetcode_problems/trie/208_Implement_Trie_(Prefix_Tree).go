package leetcode

/*
	Tried a harder version that don't quite work, keeping this here as a save
*/
// type Trie struct {
// 	Val string
// 	// Is the current node the end of a word
// 	IsWord bool
// 	Childs []*Trie
// }

// func Constructor() Trie {
// 	return Trie{IsWord: true}
// }

// func (this *Trie) Insert(word string) {
// 	// The word got no link to any of the childs
// 	found := false

// 	// For each child
// 	for _, child := range this.Childs {
// 		// If the word is the same as the child, then we reached the end of the word
// 		if word == child.Val {
// 			child.IsWord = true
// 			return
// 		}

// 		i := 0
// 		// We compare the child to the current word
// 		for i := 0; i < int(math.Min(float64(len(word)), float64(len(child.Val)))); i++ {
// 			if word[i] != child.Val[i] {
// 				// If their first letter is already different we go to the next child
// 				if i == 1 {
// 					goto nextVal
// 				} else { // Else it means that word and this child got the same prefix
// 					wasChild := child
// 					// So we update the value of the child to be this prefix
// 					child.Val = child.Val[:i]
// 					// And we give two childs to this new node, one which is the rest of our word,
// 					// and the other which is the rest of the tree that was previously here
// 					child.Childs = append(this.Childs, &Trie{Val: wasChild.Val[i:], IsWord: wasChild.IsWord, Childs: wasChild.Childs})
// 					child.Childs = append(this.Childs, &Trie{Val: word[i:], IsWord: true})
// 					return
// 				}
// 			}
// 		}

// 		// If we exited the loop because of the len of the child it means that the word's spot is lower in the tree
// 		if i == len(child.Val) {
// 			child.Insert(word[len(child.Val):])
// 		} else if i == len(word) { // If the current word is a prefix of the child
// 			// We change the child value to the word
// 			// The other part of the word becomes a child of the child
// 			tempChilds := child.Childs
// 			child.Childs = append(this.Childs, &Trie{Val: child.Val[len(word):], Childs: tempChilds})
// 			child.Val = child.Val[:len(word)]
// 			this.IsWord = true
// 			return
// 		}

// 	nextVal:
// 	}

// 	// Then we append the rest of the word as a child
// 	if !found {
// 		this.Childs = append(this.Childs, &Trie{Val: word, IsWord: true})
// 	}
// }

// func (this *Trie) Search(word string) bool {
// 	if word == "" || word == this.Val {
// 		return this.IsWord
// 	}

// 	for _, val := range this.Childs {
// 		if strings.HasPrefix(word, val.Val) {
// 			return val.Search(word[len(val.Val):])
// 		}
// 	}

// 	return false
// }

// func (this *Trie) StartsWith(word string) bool {
// 	if word == this.Val {
// 		return true
// 	}

// 	for _, val := range this.Childs {
// 		if val.Val == word {
// 			return true
// 		}

// 		i := 0
// 		for i = 0; i < int(math.Min(float64(len(word)), float64(len(val.Val)))); i++ {
// 			if word[i] != val.Val[i] {
// 				goto nextVal
// 			}
// 		}

// 		if i == len(val.Val) {
// 			return val.StartsWith(word[len(val.Val):])
// 		} else if i == len(word) {
// 			return true
// 		}

// 	nextVal:
// 	}

// 	return false
// }

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, l := range word {
		index := l - 'a'
		if curr.children[index] == nil {
			curr.children[index] = &Trie{}
		}
		curr = curr.children[index]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, l := range word {
		index := l - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, l := range prefix {
		index := l - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return true
}
