package main

import (
	"fmt"
	"strings"
)

// Linked List

type HashSetItem struct {
	value string
	next  *HashSetItem
}

// HashSet

const hashTableSize int = 30

type HashSet struct {
	storage []interface{}
}

func NewHashSet() *HashSet {
	return &HashSet{make([]interface{}, hashTableSize)}
}

func (h *HashSet) getHashSetItem(value string) *HashSetItem {
	index := GetHashSetItemIndex(value)

	if h.storage[index] == nil {
		return nil
	}

	item := h.storage[index].(*HashSetItem)

	for {
		if item.value == value {
			return item
		}

		if item.next == nil {
			return nil
		}

		item = item.next
	}
}

func (h *HashSet) Add(value string) {
	item := h.getHashSetItem(value)
	newItem := HashSetItem{value, nil}

	if item != nil {
		return
	}

	index := GetHashSetItemIndex(value)

	if h.storage[index] == nil {
		h.storage[index] = &newItem
		return
	}

	newItem.next = h.storage[index].(*HashSetItem)
	h.storage[index] = &newItem
	return
}

func GetHashSetItemIndex(value string) int {
	hashCode := hashFunction(value)
	index := hashCode % hashTableSize
	return index
}

func (h *HashSet) Remove(value string) {
	index := GetHashSetItemIndex(value)

	if h.storage[index] == nil {
		return
	}

	item := h.storage[index].(*HashSetItem)

	if item.value == value {
		h.storage[index] = item.next
		return
	}

	previous := item
	item = item.next

	for {
		if item.value == value {
			previous.next = item.next
			return
		}

		previous = item
		item = item.next
	}
	return
}

func (h *HashSet) hasValue(value string) bool {
	item := h.getHashSetItem(value)
	return !(item == nil)
}

func hashFunction(value string) int {
	val := []rune(value)
	return int(val[0])
}

type SpellChecker struct {
	words *HashSet
}

func NewSpellChecker() *SpellChecker {
	return &SpellChecker{NewHashSet()}
}

func (s *SpellChecker) AddWords(words []string) {
	i := 0
	for i < len(words) {
		s.words.Add(words[i])
		i++
	}
}

func (s *SpellChecker) CheckPhrase(sentence string) []string {
	words := strings.Split(sentence, " ")
	wrongWords := make([]string, 0)
	i := 0
	for i < len(words) {
		if !(s.words.hasValue(words[i])) {
			wrongWords = append(wrongWords, words[i])
		}
		i++
	}
	return wrongWords
}

func main() {
	checker := NewSpellChecker()
	words := []string{"Pimpolho", "is", "a", "nice", "guy"}
	checker.AddWords(words)

	fmt.Printf("Dictionary: ")
	fmt.Println(words)

	sentence := "Pimpolho is not a nice guy"
	fmt.Printf("Checking sentence: %s\n", sentence)

	fmt.Println(checker.CheckPhrase(sentence))
}
