package main

import (
  "fmt"
  "strings"
  "regexp"
)

func strCount(s string) map[string]int{
  allLower := strings.ToLower(s)
  word := strings.Split(allLower, "")
  hashTable := make(map[string]int)
  r := regexp.MustCompile("^[a-zA-Z0-9_]*$")
  for _, el := range word {
    if r.MatchString(el) {
        _, ok := hashTable[el]
        if ok {
            hashTable[el] += 1
        } else {
            hashTable[el] = 1
        }
    }
    }
  fmt.Println(hashTable)
  return hashTable
}

func main() {
  strCount("Hello WOrld!")
}