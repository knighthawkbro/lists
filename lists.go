package main

import (
	"fmt"
	"lists/array"
	"lists/list"
)

// list (Public) - had to create a public interface because
// map was already taken.
type lists interface {
	Add(item interface{})
	Remove(item interface{}) bool
	Contains(item interface{}) bool
	Size() int

	AddIndex(index int, item interface{})
	Get(index int) interface{}
	RemoveIndex(index int) interface{}
	Set(index int, item interface{}) interface{}
	IndexOf(item interface{}) int
}

func main() {
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as an array...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	arr := array.New()
	driver(arr)
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as a list...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	lst := list.New()
	driver(lst)

	// Don't mind this code, testing out closures.
	/* list.Add("world")
	list.Add("hello")
	list.Add("hi")
	list.Add("there")
	list.Add("well")
	list.Add("fun")
	list.Add("having")
	list.Add("am")
	list.Add("i")
	item := list.Iterate()
	for i := 0; i < list.Size(); i++ {
		fmt.Println(item())
	}
	list = nil*/
}

func driver(words lists) {
	// adding items
	fmt.Println("\nadding 'hello' to the list")
	words.Add("hello")
	fmt.Println(words)

	fmt.Println("\nadding 'world' to the list")
	words.Add("world")
	fmt.Println(words)

	fmt.Println("\nadding 'today' to the list")
	words.Add("today")
	fmt.Println(words)

	// adding at specific position
	fmt.Println("\nadding 'there' at position 1")
	words.AddIndex(1, "there")
	fmt.Println(words)

	fmt.Println("\nadding 'well' at position 0")
	words.AddIndex(0, "well")
	fmt.Println(words)

	// replacing
	fmt.Println("\nreplacing item at position 1 with 'hi'")
	fmt.Println("replaced:", words.Set(1, "hi"))
	fmt.Println(words)

	// retrieving
	fmt.Println("\ngetting item at position 3")
	fmt.Println(words.Get(3))

	// Contains and IndexOf
	fmt.Print("\ndoes the list contain 'hello'? ")
	if words.Contains("hello") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Println("what is its index?")
	fmt.Println(words.IndexOf("hello"))

	fmt.Print("\ndoes the list contain 'hi'? ")
	if words.Contains("hi") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Println("what is its index?")
	fmt.Println(words.IndexOf("hi"))

	// removing item
	fmt.Println("\n--> testing remove methods")
	fmt.Println(words)

	fmt.Println("removing 'hi'")
	words.Remove("hi")
	fmt.Println(words)

	// removing at specified index
	fmt.Println("\nremoving item at index 2")
	fmt.Println("removed:", words.RemoveIndex(2))
	fmt.Println(words)

	fmt.Println("\nremoving item at index 0")
	fmt.Println("removed:", words.RemoveIndex(0))
	fmt.Println(words)

	fmt.Println("\nremoving item at index 1")
	fmt.Println("removed:", words.RemoveIndex(1))
	fmt.Println(words)

	fmt.Println("\nremoving item at index 0")
	fmt.Println("removed:", words.RemoveIndex(0))
	fmt.Println(words)

	fmt.Println("\nadding 'there' back in")
	words.Add("there")
	fmt.Println(words)

	// item not in list
	fmt.Println("\ntrying to remove 'hello'")
	fmt.Print("was it removed? ")
	if words.Remove("hello") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	fmt.Println("\ntrying to remove 'there'")
	fmt.Print("was it removed? ")
	if words.Remove("there") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	fmt.Println("\ntrying to remove 'world'")
	fmt.Print("was it removed? ")
	if words.Remove("world") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	fmt.Println("\ntrying to remove item at index 0")
	words.RemoveIndex(0)
}
