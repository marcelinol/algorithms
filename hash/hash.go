package main

import "fmt"
import "errors"

// Linked List

type HashTableItem struct {
	key string
	value string
	next *HashTableItem
}

// HashTable

const hashTableSize int = 30

type HashTable struct {
	storage []interface{}
}

func NewHashTable() *HashTable {
	return &HashTable{make([]interface{}, hashTableSize)}
}

func (h *HashTable) getHashTableItem(key string) (*HashTableItem) {
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

func hashFunction(key string) int {
	value := []rune(key)
	return int(value[0])
}

func main() {
	h := NewHashTable()
	fmt.Println(h)

	fmt.Println("\n--------------------- TESTING ADDITION ----------------------\n")

	h.Add("luciano", "atlas")

	item := h.getHashTableItem("luciano")
	fmt.Println(item)

	h.Add("wagner", "gorillaz")

	h.Add("hugo", "jarvis")

	h.Add("luciano", "this should overwrite the item luciano")
	fmt.Println(h.getHashTableItem("luciano"))

	h.Add("wluciano", "this should concatenate wagner")
	h.Add("lucas", "vikings")
	h.Add("loao barbosa", "jarvis")


	item = h.getHashTableItem("wluciano")
	fmt.Println(item)
	otherItem := item.next

	fmt.Println(otherItem)

	fmt.Println("\n--------------------- TESTING REMOVE ----------------------\n")
	fmt.Println(h.getValue("lucas"))
	fmt.Println("removing lucas\n")
	h.Remove("lucas")
	fmt.Println(h.getValue("lucas"))

}

// Questions
// Why the length++ didn't work when I was not using a pointer to the hashtable?
