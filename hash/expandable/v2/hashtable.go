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
	items []interface{}
	size  int
}

// HashTable

const hashTableSize int = 2
const limitHashTableLength int = 2

type HashTable struct {
	storage []interface{}
	length  int
}

func NewHashTable() *HashTable {
	return &HashTable{make([]interface{}, hashTableSize), 0}
}

func (h *HashTable) getHashTableItem(key string) *HashTableItem {
	index := h.GetHashTableItemIndex(key)

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

func resizeHashTable(h *HashTable) *HashTable {
	fmt.Println("resizing hash")
	newSize := len(h.storage) * 2
	newHashTable := HashTable{make([]interface{}, newSize), 0}
	fmt.Println(newHashTable)

	i := 0
	for {
		if i == len(h.storage) {
			break
		}

		if h.storage[i] != nil {
			item := h.storage[i].(*HashTableItem)
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

func (h *HashTable) newGetHashTableItem(key string) *HashTableItem {
	index := h.GetHashTableItemIndex(key)

	if h.storage[index] == nil {
		return nil
	}

	if h.storage[index].(int) == 0 {
		return nil
	}

	s := h.storage[index]
	item := s.(*HashTableItem)

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

func (h *HashTable) NewAdd(key string, value string) {
	item := h.newGetHashTableItem(key)
	newItem := HashTableItem{key, value, nil}

	if item != nil {
		item.value = newItem.value
		return
	}

	index := h.GetHashTableItemIndex(key)

	if h.storage[index] == nil {
		s := HashTableStorage{make([]interface{}, 0), 0}
		h.storage[index] = s
		s.items = &newItem
		h.length++
		return
	}

	if h.storage[index] == nil {

	}

	newItem.next = h.storage[index].(*HashTableItem)
	h.storage[index] = &newItem
	h.length++
	return
}

func (h *HashTable) Add(key string, value string) {
	item := h.getHashTableItem(key)
	newItem := HashTableItem{key, value, nil}

	if item != nil {
		item.value = newItem.value
		return
	}

	if h.length > (limitHashTableLength * len(h.storage)) {
		*h = *resizeHashTable(h)
	}

	index := h.GetHashTableItemIndex(key)

	if h.storage[index] == nil {
		h.storage[index] = &newItem
		h.length++
		return
	}

	newItem.next = h.storage[index].(*HashTableItem)
	h.storage[index] = &newItem
	h.length++
	return
}

func (h *HashTable) GetHashTableItemIndex(key string) int {
	hashCode := hashFunction(key)
	index := hashCode % len(h.storage)
	return index
}

func (h *HashTable) Remove(key string) {
	index := h.GetHashTableItemIndex(key)

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
	fmt.Println(h.length)
	fmt.Println(h)

	fmt.Println("\n--------------------- TESTING ADDITION ----------------------")

	h.Add("luciano", "atlas")

	item := h.getHashTableItem("luciano")
	fmt.Println(item)

	h.Add("wagner", "gorillaz")

	h.Add("hugo", "jarvis")

	h.Add("luciano", "this should overwrite the item luciano")
	fmt.Println(h.getHashTableItem("luciano"))

	h.Add("wluciano", "this should concatenate wagner")
	h.Add("lucas", "vikings")
	fmt.Println("\n--------------------- RESIZING IN THE NEXT ADDITION ----------------------")
	fmt.Print("old hash: ")
	fmt.Println(h)
	h.Add("loao barbosa", "jarvis")
	fmt.Println(h.length)

	fmt.Print("new hash: ")
	fmt.Println(h)
	fmt.Println(h.getHashTableItem("luciano"))
	fmt.Println(h.getHashTableItem("wagner"))
	fmt.Println(h.getHashTableItem("hugo"))
	fmt.Println(h.getHashTableItem("wagner"))

	item = h.getHashTableItem("wagner")
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
