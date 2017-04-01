package main

import "fmt"

func containsCamel(dict []string, camelRegex string) ([]string) {
	dictLen := len(dict)
	camelRegexLen := len(camelRegex)

	// preconditions
	if dictLen == 0 {
		return nil
	}
	if camelRegexLen == 0 {
		return nil
	}

	// on this. 02:48p
	// you are not sure on slice.
	// variable length stuff.
	var result []string
	// question. when to use range?
	for i:=0 ; i<dictLen ; i++ {
		if (camelInStr(camelRegex, dict[i])) {
			// increase the slice by 1.
			result = result[0:i+i]
			// insert into result. 02:55p
			result[i] = dict[i]
		}
	}
	return result
}

// unsure of string ops..
// so this is good practice...
// you can just pass it an array arg..
func camelInStr(camelRegex, camelStr string) (bool) {
	possibleR := false
	
	camelRegexA := strToCamel(camelRegex)
	camelStrA := strToCamel(camelStr)

	csi := 0
	cri := 0
	for ; csi<len(camelStrA) && cri<len(camelRegexA) ; csi++ {
		fmt.Println(">LOOP", camelStrA[csi], camelRegexA[cri], cri, len(camelRegexA))
		if strContainsPrefix(camelStrA[csi], camelRegexA[cri]) {
			possibleR = true
			cri++
		} else if possibleR {
			if strContainsPrefix(camelStrA[csi], camelRegexA[cri]) {
				cri++
			}
		}

		if cri == len(camelRegexA) {
			return true
		} 
	}

	return false
}

func strContainsPrefix(s1, s2 string) (bool) {
	s1Len := len(s1)
	s2Len := len(s2)
	smallStrLen := -1

	if s1Len < s2Len {
		smallStrLen = s1Len
	} else {
		smallStrLen = s2Len
	}

	possibleR := s1[0] == s2[0]
	if possibleR {
		for i:=0 ; i<smallStrLen ; i++ {
			possibleR = possibleR && s1[i] == s2[i]
		}
	}
	return possibleR
}

// takes an string and splits it up to an array of strings.
// split off an upper case.
func strToCamel(camelStr string) ([]string) {
	camelStrLen := len(camelStr)
	if camelStrLen == 0 {
		return nil
	}

	// this stores the arry of strings
	var result []string
	
	fmt.Println("result", len(result))
	// loop thru the string
	for i, j:=0, 0 ; i<camelStrLen ; i, j = i+1, j+1 {
		fmt.Println(i, j)
		startPos := i
		fmt.Println("> startPos ",startPos)
		// find the next word.
		for ; i<camelStrLen ; i++ {
			fmt.Println("inner> ", i)

			// so you didn't find one.
			if i+1 == camelStrLen {
				fmt.Println(">>>>>>>>>>>>>>>got here")
				break
			}

			// loop to the next upper case character.
			// that is the next word.
			if isUpper(camelStr[i+1]) {
				fmt.Println("> found", i)
				// back up 1 character and break.
				break
			}
		}
		
		// increment here. since we are checking i+1
		//i++

		// add a word
		fmt.Println(">1")
		newResult := make([]string, j+1)
		copy(newResult, result)
		result = newResult
		fmt.Println(">1.5", j, startPos, i)
		result[j] = camelStr[startPos:i+1]
		fmt.Println(">2", len(result))

		// decrement back
		//i--
		fmt.Println(">3", i)
		
	}

	return result
}

// not sure on what the type is.
// i think it should be a rune.
func isUpper(x byte) (bool) {
	fmt.Println(x)
	// 97 - 122
	if x >= 65 && x <= 90 {
		return true
	} else {
		return false
	}
}

func main() {
	//dictList := []string{"HelloMars", "HelloWorld", "HelloWorldMars", "HiHo"}
	//myStr := "H"
	//fmt.Println(isUpper(myStr[0]))
	//fmt.Println(strToCamel(myStr))
	//fmt.Println(strContainsPrefix("Hello", "He"))

	fmt.Println(camelInStr("HelloWorldMars", "HelloWorldMars"))
}


/*

=== NOTES:

So you are given an array of strings.
You search return an arrya of strings base of a string serarch
The string search seems to be which camel case and what letter that follows.

02:35

an naive solution.
Just up the search results into caps and stuff that follow it.
for each item in the set, search for it, if found, go to the second item.
- note that the second and all subsequent item has to follow the position of the search result.
- repeat until you get until the end of the list and get all yeses.
- if you hit a no, its a fail for that word.
- so a helper function would be helpful does a string regex contain a word in a dict...

could you optimize this?
02:38p?
also, can you assume that it starts with a capital letter?
i'm not sure...
for for the sake of this. lets say yes for now....

one thing you might be able to do is to do the following....
- you at a min can probably have some sort of helper to split camel case words yup to caps words
- ad just return a list of words. that way don't have to do the i and j and k pointer stuff...

02:44p

03:07p
struggling is good....
on the main func.....

03:37... this is fucking hard...
but this is a good lesson....
i think i'm going to do a ton of errors trying to do this...

so think data structs....
use 2 queus.
04:02
one for the string
on for a regex.


HelloMars
H

take Nst cap letter in regex.
- Find the the string
- if not found, just return.
- if found, regex until you you hit the next caps.
	- if at any point, they dont match, bail. skip ahead to next word in string.
  - if found.
- prevCamelFound && currentCamelFound...
- bail on string length over.


WP
you scan.......
- you find w. (true)
- you don't find p. (true && false)
- you do find p... (true && true)

just stuck on the thing....
might want ot anotehr solution.....
the inital thing that i wanted to try..



=== REF:

List of string that represent class names in CamelCaseNotation.

Write a function that given a List and a pattern returns the matching elements.

['HelloMars', 'HelloWorld', 'HelloWorldMars', 'HiHo']

H -> [HelloMars, HelloWorld, HelloWorldMars, HiHo]
HW -> [HelloWorld, HelloWorldMars]
Ho -> []
HeWorM -> [HelloWorldMars]


https://careercup.com/question?id=5660887265312768
*/