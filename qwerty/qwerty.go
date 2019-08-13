package main

import (
	"fmt"
)

/*
	When typing on a touch screen, occasionally the wrong key is registered.

	For example:
	When typing "Hello", a "G" might be registered instead of "H"

	Write a function that given a string, returns all nearby words.

	You are given the following helper functions:

	// for a given character return all nearby characters
	function get_nearby_chars(string $character): Set<string>

	// for a given string, return true if it is a word
	function is_word(string $word).bool;

	implement this...
	// for a given word get all possible nearby word
	function nearby_words(string $word): Set<String>

	Tips:
	# Ask questions to clarify the problem

	Example:
	1.) What format is the word given to us?
	2.) How do we define a nearby word?
	3.) Do we need to remove duplicate words in the output?
	4.) Does the output need to be sort?

	# Explain your approach before writing code

	Example:
	1.) I would like to create mapping relation of possible word
	gi
	g => g, h, f
	i => i, o, k

	gi	hi	fi
	go	ho	fo
	gk	hk	fk

	2.) Give the word to is_word function
	false => gi fi ho fo gk hk fk
	true	=> hi, go

	Hope, the interview get an idea, they might give an idea about time and space complexity

	3.) Test by checking edge case

	4.) Analyze time and space complexity

	5.) Understand trade off between alternate solution
	Interview may asking idea from you about alternative solution

	6.) Prepare questions for the interviewer
*/

func main() {
	myWord := "gi"
	fmt.Println(nearbyWords(myWord))
}

func nearbyWords(w string) []string {
	result := []string{}

	subword := getNearBychars(string(w[1]))[string(w[1])]
	word := getNearBychars(string(w[0]))[string(w[0])]

	for _, unword := range getNearByWords(subword, word) {
		if isWord(unword) {
			result = append(result, unword)
		}
	}

	return result
}

func getNearByWords(subword, word []string) []string {
	result := []string{}

	// g
	// 	-> h
	//		-> f
	for _, w := range word {
		for _, sw := range subword { // i -> o -> k
			result = append(result, w+sw)
		}
	}

	return result
}

func getNearBychars(c string) map[string][]string {
	hMap := make(map[string][]string)
	if c == "g" {
		hMap["g"] = append(hMap["g"], "g")
		hMap["g"] = append(hMap["g"], "h")
		hMap["g"] = append(hMap["g"], "f")
	} else if c == "i" {
		hMap["i"] = append(hMap["i"], "i")
		hMap["i"] = append(hMap["i"], "o")
		hMap["i"] = append(hMap["i"], "k")
	}

	return hMap
}

func isWord(w string) bool {
	return w == "hi" || w == "go"
}
