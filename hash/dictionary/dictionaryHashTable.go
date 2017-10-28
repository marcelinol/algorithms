package main

import (
	"errors"
	"fmt"
	"strings"
)

type HashTableItem struct {
	key   string
	value string
	next  *HashTableItem
}

// HashTable

const hashTableSize int = 30

type HashTable struct {
	storage []interface{}
}

func NewHashTable() *HashTable {
	return &HashTable{make([]interface{}, hashTableSize)}
}

func (h *HashTable) getHashTableItem(key string) *HashTableItem {
	index := GetHashTableItemIndex(key)

	if h.storage[index] == nil {
		return nil
	}

	item := h.storage[index].(*HashTableItem)

	for {
		if item.key == key {
			return item
		}

		if item.next == nil {
			return nil
		}

		item = item.next
	}
}

func (h *HashTable) Add(key string, value string) {
	item := h.getHashTableItem(key)
	newItem := HashTableItem{key, value, nil}

	if item != nil {
		item.value = newItem.value
		return
	}

	index := GetHashTableItemIndex(key)

	if h.storage[index] == nil {
		h.storage[index] = &newItem
		return
	}

	newItem.next = h.storage[index].(*HashTableItem)
	h.storage[index] = &newItem
	return
}

func GetHashTableItemIndex(key string) int {
	hashCode := hashFunction(key)
	index := hashCode % hashTableSize
	return index
}

func (h *HashTable) Remove(key string) {
	index := GetHashTableItemIndex(key)

	if h.storage[index] == nil {
		return
	}

	item := h.storage[index].(*HashTableItem)

	if item.key == key {
		h.storage[index] = item.next
		return
	}

	previous := item
	item = item.next

	for {
		if item.key == key {
			previous.next = item.next
			return
		}

		previous = item
		item = item.next
	}
	return
}

func (h *HashTable) getValue(key string) (string, error) {
	item := h.getHashTableItem(key)

	if item == nil {
		return "", errors.New("ERROR: key not found")
	}

	return item.value, nil
}

func (h *HashTable) hasKey(key string) bool {
	item := h.getHashTableItem(key)
	return !(item == nil)
}

func hashFunction(key string) int {
	value := []rune(key)
	return int(value[0])
}

type SpellChecker struct {
	words *HashTable
}

func NewSpellChecker() *SpellChecker {
	return &SpellChecker{NewHashTable()}
}

func (s *SpellChecker) AddWords(words []string) {
	i := 0
	for i < len(words) {
		s.words.Add(words[i], "")
		i++
	}
}

func (s *SpellChecker) CheckPhrase(sentence string) []string { //O(m+n) -> O(n)
	words := strings.Split(sentence, " ") // O(n)
	wrongWords := make([]string, 0)
	i := 0
	for i < len(words) { // O(m)
		if !(s.words.hasKey(words[i])) { // O(1)
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
	fmt.Println(checker.CheckPhrase("Pimpolho is not a nice guy"))
}

// Questions
// Why the length++ didn't work when I was not using a pointer to the hashtable?
