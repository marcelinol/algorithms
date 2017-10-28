package main

import "fmt"
import "errors"

// Linked List

type HashTableItem struct {
	key   string
	value string
	next  *HashTableItem
}

type HashTableStorage struct {
	item *HashTableItem
	size int
}

// HashTable

const hashTableSize int = 2
const limitHashTableLength int = 2

type HashTable struct {
	storage []HashTableStorage
	length  int
}

func NewHashTable() *HashTable {
	return &HashTable{make([]HashTableStorage, hashTableSize), 0}
}

func resizeHashTable(h *HashTable) *HashTable {
	fmt.Println("resizing hash")
	newSize := len(h.storage) * 2
	newHashTable := HashTable{make([]HashTableStorage, newSize), 0}
	fmt.Println(newHashTable)

	i := 0
	for {
		if i == len(h.storage) {
			break
		}

		if h.storage[i].item != nil {
			item := h.storage[i].item
			for {
				newHashTable.Add(item.key, item.value)

				if item.next == nil {
					break
				}

				item = item.next
			}
		}

		i++
	}
	fmt.Print(">>>>>>>>>>>>>>>\n\nnew hash: ")
	fmt.Println(newHashTable)
	return &newHashTable
}

func (h *HashTable) GetHashTableItem(key string) *HashTableItem {
	index := h.GetHashTableItemIndex(key)

	if h.storage[index].size == 0 {
		return nil
	}

	s := h.storage[index]

	for {
		if s.item.key == key {
			return s.item
		}

		if s.item.next == nil {
			return nil
		}

		s.item = s.item.next
	}
}

func (h *HashTable) Add(key string, value string) {
	item := h.GetHashTableItem(key)
	newItem := &HashTableItem{key, value, nil}

	if item != nil {
		item.value = newItem.value
		return
	}

	index := h.GetHashTableItemIndex(key)

	if h.storage[index].size == 0 {
		s := HashTableStorage{newItem, 1}
		h.storage[index] = s
		return
	}

	if h.storage[index].size > limitHashTableLength {
		h = resizeHashTable(h)
	}

	newItem.next = h.storage[index].item
	h.storage[index].item = newItem
	h.storage[index].size++
	return
}

func (h *HashTable) GetHashTableItemIndex(key string) int {
	hashCode := hashFunction(key)
	index := hashCode % len(h.storage)
	return index
}

func (h *HashTable) Remove(key string) {
	index := h.GetHashTableItemIndex(key)

	if h.storage[index].item == nil {
		return
	}

	item := h.storage[index].item
	h.storage[index].size--

	if item.key == key {
		h.storage[index].item = item.next
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
}

func (h *HashTable) getValue(key string) (string, error) {
	item := h.GetHashTableItem(key)

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
	fmt.Println(h.length)
	fmt.Println(h)

	fmt.Println("\n--------------------- TESTING ADDITION ----------------------")

	h.Add("luciano", "atlas")

	item := h.GetHashTableItem("luciano")
	fmt.Println(item)

	h.Add("wagner", "gorillaz")

	h.Add("hugo", "jarvis")

	h.Add("luciano", "this should overwrite the item luciano")
	fmt.Println(h.GetHashTableItem("luciano"))

	h.Add("wluciano", "this should concatenate wagner")
	h.Add("lucas", "vikings")
	fmt.Println("\n--------------------- RESIZING IN THE NEXT ADDITION ----------------------")
	fmt.Print("old hash: ")
	fmt.Println(h)
	h.Add("loao barbosa", "jarvis")
	fmt.Println(h.length)

	fmt.Print("new hash: ")
	fmt.Println(h)
	fmt.Println(h.GetHashTableItem("luciano"))
	fmt.Println(h.GetHashTableItem("wagner"))
	fmt.Println(h.GetHashTableItem("hugo"))
	fmt.Println(h.GetHashTableItem("wagner"))

	item = h.GetHashTableItem("wagner")
	fmt.Println(item)
	otherItem := item.next

	fmt.Println(otherItem)

	fmt.Println("\n--------------------- TESTING REMOVE ----------------------")
	fmt.Println(h.getValue("lucas"))
	fmt.Println("removing lucas")
	h.Remove("lucas")
	fmt.Println(h.getValue("lucas"))

	fmt.Println(h)

}

// Questions
// Why the length++ didn't work when I was not using a pointer to the hashtable?
