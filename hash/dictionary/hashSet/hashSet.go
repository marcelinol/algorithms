// package main
//
// import "fmt"
//
// // Linked List
//
// type HashSetItem struct {
// 	value string
// 	next  *HashSetItem
// }
//
// // HashSet
//
// const hashTableSize int = 30
//
// type HashSet struct {
// 	storage []interface{}
// }
//
// func NewHashSet() *HashSet {
// 	return &HashSet{make([]interface{}, hashTableSize)}
// }
//
// func (h *HashSet) getHashSetItem(value string) *HashSetItem {
// 	index := GetHashSetItemIndex(value)
//
// 	if h.storage[index] == nil {
// 		return nil
// 	}
//
// 	item := h.storage[index].(*HashSetItem)
//
// 	for {
// 		if item.value == value {
// 			return item
// 		}
//
// 		if item.next == nil {
// 			return nil
// 		}
//
// 		item = item.next
// 	}
// }
//
// func (h *HashSet) Add(value string) {
// 	item := h.getHashSetItem(value)
// 	newItem := HashSetItem{value, nil}
//
// 	if item != nil {
// 		return
// 	}
//
// 	index := GetHashSetItemIndex(value)
//
// 	if h.storage[index] == nil {
// 		h.storage[index] = &newItem
// 		return
// 	}
//
// 	newItem.next = h.storage[index].(*HashSetItem)
// 	h.storage[index] = &newItem
// 	return
// }
//
// func GetHashSetItemIndex(value string) int {
// 	hashCode := hashFunction(value)
// 	index := hashCode % hashTableSize
// 	return index
// }
//
// func (h *HashSet) Remove(value string) {
// 	index := GetHashSetItemIndex(value)
//
// 	if h.storage[index] == nil {
// 		return
// 	}
//
// 	item := h.storage[index].(*HashSetItem)
//
// 	if item.value == value {
// 		h.storage[index] = item.next
// 		return
// 	}
//
// 	previous := item
// 	item = item.next
//
// 	for {
// 		if item.value == value {
// 			previous.next = item.next
// 			return
// 		}
//
// 		previous = item
// 		item = item.next
// 	}
// 	return
// }
//
// func (h *HashSet) hasValue(value string) bool {
// 	item := h.getHashSetItem(value)
// 	return !(item == nil)
// }
//
// func hashFunction(value string) int {
// 	val := []rune(value)
// 	return int(val[0])
// }
//
// func main() {
// 	h := NewHashSet()
// 	fmt.Println(h)
//
// 	fmt.Println("\n--------------------- TESTING ADDITION ----------------------")
//
// 	h.Add("luciano")
//
// 	item := h.getHashSetItem("luciano")
// 	fmt.Println(item)
//
// 	h.Add("wagner")
//
// 	h.Add("hugo")
//
// 	h.Add("luciano")
// 	fmt.Println(h.getHashSetItem("luciano"))
//
// 	h.Add("wluciano")
// 	h.Add("lucas")
// 	h.Add("loao barbosa")
//
// 	item = h.getHashSetItem("wluciano")
// 	fmt.Println(item)
// 	otherItem := item.next
//
// 	fmt.Println(otherItem)
//
// 	fmt.Println("\n--------------------- TESTING REMOVE ----------------------")
// 	fmt.Println(h.hasValue("lucas"))
// 	fmt.Println("removing lucas")
// 	h.Remove("lucas")
// 	fmt.Println(h.hasValue("lucas"))
//
// }
//
// Questions
// Why the length++ didn't work when I was not using a pointer to the hashtable?
