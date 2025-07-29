package main

import (
	"iter"
	"log"
	"strings"
)
// iteration of string that you input a function and returns an iter.
func iter0(str string, d func(string) string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, v := range strings.Fields(str) {
			if !yield(d(v)) {
				return
			}
		}
	}

}
// iter that does not need data to process.
func iter1(yield func(i int) bool) {
	for i := range 3 {
		if !yield(i) {
			return
		}
	}
}
// iter2 processing strings word by word
func iterStr(yield func(int, string) bool) {
	str := "The Quick Brown Fox Jumps Over The Lazy Dog"
	for i, v := range strings.Fields(str) {
		if !yield(i, v) {
			return
		}
	}
}

// a function that returns a seq2 that returns a map that will upercase it.
func uperCaser(m map[string]string) iter.Seq2[string, string] {
	return func(yield func(k, v string) bool) {
		for key, value := range m {
			if !yield(strings.ToUpper(key), strings.ToUpper(value)) {
				return
			}
		}
	}
}
// use iter as params
func iterAsParam(k iter.Seq2[string, string]) {
	log.Println("\n\nUsing Iter As Param")
	for i, v := range k {
		log.Printf("Iterating as param map[%v]=%v", i, v)
	}
}

// custom iterator function
func main() {

	log.Println("Using Iter")
	str := "The Quick Brown Fox Jumps Over The Lazy Dog"
	_ = str


	// custom iter using a for
	log.Println("custom iter: using for")
	for i := range iter1 {
		log.Printf("iter: %d", i)
	}


	log.Println("Using Iter To Update Value and Key Pair")
	// using an iter for a real value
	for i, s := range iterStr {
		log.Printf("str[%d] = %v", i, strings.ToLower(s))
	}

	log.Println("Using Iter to make an Iter run inside a function")
	// using an iterator function
	allUpper := iter0(str, func(str string) string {
		str = strings.ToUpper(str)
		return str
	})

	for str := range allUpper {
		log.Printf("iterated inside func %s", str)
	}

	nickNames := map[string]string{
		"John": "Janjan",
		"Mike": "Miky",
		"Stad": "Tray Tray",
	}
	// using an iter as a arguments
	iterAsParam(uperCaser(nickNames))
}
