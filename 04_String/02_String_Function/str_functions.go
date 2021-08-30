package main

import (
	"fmt"
	s "strings"
)

func main() {
	var str = "hi! goLang"
	var p = fmt.Println
	fmt.Println("Length :", len(str))

	fmt.Println("UpperCase :", s.ToUpper(str))
	fmt.Println("LowerCase :", s.ToLower(str))
	fmt.Println("Title :", s.Title(str))

	p("Contains:  ", s.Contains("test", "es"))                 //right to left true
	p("ContainsAny:	", s.ContainsAny("test", "ez"))            //true check for each character in string from substring
	p("Contains: ", s.Contains("", ""))                        //true
	p("ContainsAny for blank string:	", s.ContainsAny("", "")) //false

	p("Count:     ", s.Count("test", "t"))              //case sensitive
	p("Count with blank string	:", s.Count("test", "")) // length of string + 1
	p("Count with space:	", s.Count("test", " "))

	p("HasPrefix: ", s.HasPrefix("test", "te"))                      //case sensitive
	p("HasPrefix with Blank string: ", s.HasPrefix("test", ""))      //true
	p("HasSuffix with blank string	:", s.HasSuffix("Australia", "")) //true
	p("HasSuffix: ", s.HasSuffix("test", "st"))                      //true

	//case -sensitive
	p("Index:     ", s.Index("test", "e"))                       //1
	p("IndexAny:     ", s.IndexAny("test", "pe"))                //1
	p("LastIndex:     ", s.LastIndex("Australia", "a"))          //8
	p("LastIndexAny:     ", s.LastIndexAny("Australia", "papa")) //8

	p("Join:      ", s.Join([]string{"a", "b"}, "-")) //"a-b"

	p("Repeat:    ", s.Repeat("a", 5)) //"aaaaa"
	p("Repeat:    ", s.Repeat(" B ", 5))
	p("Repeat:    ", "ohh"+s.Repeat("la ", 3))

	//case sensitive
	p("Replace:   ", s.Replace("foOoOo", "o", "e", -1))
	p("Replace:   ", s.Replace("Banana", "na", "sa", 2))
	p("Replace:   ", s.Replace("Banana", "na", "sa", 0))
	p("Replace:   ", s.Replace("Banana", "na", "sa", 1))

	//o/p will be slice only
	p("Split:     ", s.Split("a-b-c-d-e", "-")) //[a b c d e]
	p("Split:     ", s.Split("a-b-c-d-e", ".")) //[a-b-c-d-e]

	p("SplitN:     ", s.SplitN("a,b,c,d", ",", -1))
	p("SplitN:     ", s.SplitN("a,b,c,d", ",", 0))
	p("SplitN:     ", s.SplitN("a,b,c,d", ",", 1))
	p("SplitN:     ", s.SplitN("a,b,c,d", ",", 3))

	p("SplitAfter:     ", s.SplitAfter("a-b-c-d-e", "-"))
	p("SplitAfterN:     ", s.SplitAfterN("a-b-c-d-e", "-", -1))
	p("SplitAfterN:     ", s.SplitAfterN("a-b-c-d-e", "-", 0))
	p("SplitAfterN:     ", s.SplitAfterN("a-b-c-d-e", "-", 2))

	p("Compare:		", s.Compare("Germany", "GERMANY"))
	p("EqualFold:	", s.EqualFold("JAPAN-1254", "japan-1254")) //case- insensitive
	p("EqualFold:	", s.EqualFold(" ", " "))
	//case -sensitive
	p("Trim:	", s.Trim("aabcdeaa", "a"))
	p("TrimLeft:	", s.TrimLeft("aaAabaaafbcdeaa", "a"))
	p("TrimRight:	", s.TrimRight("abcdeaaaaaaaaaa", "a"))
	p("TrimSpace:	", s.TrimSpace("\t\n\rabcdea\t\n\r  "))

}
