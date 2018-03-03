package array

import (
	"fmt"
	"log"
)

// Array (Public) - Structure that defines
type Array struct {
	size       int
	collection []interface{}
}

// Init (Public) - Initializes the array with whatever size is provided, This is what can be overrided by the user.
func (a *Array) Init(capacity int) *Array {
	if capacity < 0 {
		return nil
	}
	a.collection = make([]interface{}, capacity)
	a.size = 0
	return a
}

// New (Public) - Returns an initialized array with default size of 10.
func New() *Array { return new(Array).Init(10) }

// Add (Public) - Adds item to list
func (a *Array) Add(item interface{}) {
	err := checkForNil(item)
	if err != nil {
		log.Println(err)
		return
	}

	a.ensureSpace()
	a.collection[a.size] = item
	a.size++
}

// Remove (Public) - Removes item from list, returns true if item found in list
func (a *Array) Remove(item interface{}) bool {
	err := checkForNil(item)
	if err != nil {
		log.Println(err)
		return false
	}

	index := a.IndexOf(item)
	if index < 0 {
		return false
	}

	a.shiftLeft(index)
	a.size--
	return true
}

// Contains (Public) - Checks to see if item is in array, returns true or false
func (a *Array) Contains(item interface{}) bool {
	for i := range a.collection {
		if i == item {
			return true
		}
	}
	return false
}

// Size (Public) - returns the size of the Array
func (a *Array) Size() int {
	return a.size
}

// String (Public) - formats the array when fmt.Print is called.
func (a *Array) String() string {
	if a.size == 0 {
		return "[ ]"
	}
	result := "[ "
	for x := 0; x < a.size; x++ {
		result += fmt.Sprintf("%v ", a.collection[x])
	}
	return result + "]"
}

// AddIndex (Public) - Adds item at specified index
func (a *Array) AddIndex(index int, item interface{}) {
	err := checkForNil(item)
	if err != nil {
		log.Println(err)
		return
	}
	err = a.checkIndex(index)
	if err != nil {
		log.Println(err)
		return
	}
	a.ensureSpace()

	a.shiftRight(index)
	a.collection[index] = item
	a.size++
}

// Set (Public) - Replaces item at specified index with new value
// returns original value
func (a *Array) Set(index int, item interface{}) interface{} {
	err := checkForNil(item)
	if err != nil {
		log.Println(err)
		return nil
	}
	err = a.checkIndex(index)
	if err != nil {
		log.Println(err)
		return nil
	}

	removed := a.collection[index]
	a.collection[index] = item
	return removed
}

// RemoveIndex (Public) - Removes item at specified index, returns removed item
func (a *Array) RemoveIndex(index int) interface{} {
	err := a.checkIndex(index)
	if err != nil {
		log.Println(err)
		return nil
	}
	removed := a.collection[index]
	a.shiftLeft(index)
	a.size--
	return removed
}

// Get (Public) - Returns removed item
func (a *Array) Get(index int) interface{} {
	if a.size == 0 {
		return nil
	}
	err := a.checkIndex(index)
	if err != nil {
		log.Println(err)
		return nil
	}
	return a.collection[index]
}

// IndexOf (Public) -
func (a *Array) IndexOf(item interface{}) int {
	for index, x := range a.collection {
		if item == x {
			return index
		}
	}
	return -1
}

// ensureSpace (Private) - Sees if the size and capacity of the array are the same. If so,
// It creates a new array with double the capacity and overwrites the old array with a new
// array, then clears the new array for the GC.
func (a *Array) ensureSpace() {
	if a.size == cap(a.collection) {
		new := new(Array).Init(cap(a.collection) * 2)
		new.size = a.size
		for i := 0; i < a.size; i++ {
			new.collection[i] = a.collection[i]
		}
		*a = *new
		new = nil
	}
}

func checkForNil(item interface{}) error {
	if item == nil {
		return fmt.Errorf("nil not a possible value")
	}
	return nil
}

func (a *Array) checkIndex(index int) error {
	if index < 0 || index >= a.size {
		return fmt.Errorf("index outside of list")
	}
	return nil
}

// shiftLeft (Private) - Moves all the items left after index (Destructive)
func (a *Array) shiftLeft(index int) {
	for i := index; i < a.size-1; i++ {
		a.collection[i] = a.collection[i+1]
	}
}

// shiftRight (Private) - Moves all the items to the right after the index (non-destructive)
func (a *Array) shiftRight(index int) {
	for i := a.size; i > index; i-- {
		a.collection[i] = a.collection[i-1]
	}
}
