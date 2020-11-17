package main

import (
  "fmt"
  "strings"
)

func strCount(s string) map[string]int{
  word := strings.Split(s, "")
  hashTable := make(map[string]int)
  for _, el := range word {
        _, ok := hashTable[el]
        if ok {
            hashTable[el] += 1
        } else {
            hashTable[el] = 1
        }
    }
  fmt.Println(hashTable)
  return hashTable
}

func main() {
  strCount("apple")
}