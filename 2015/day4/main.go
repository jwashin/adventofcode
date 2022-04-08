package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	data := "yzbqklnj"

	fmt.Println(FindMd5StartingWith00000(data))
	fmt.Println(FindMd5StartingWith000000(data))

}

func FindMd5StartingWith00000(aString string) int {
	k := -1
	for {
		k += 1
		theString := []byte(aString + fmt.Sprintf("%d", k))
		hexVal := fmt.Sprintf("%x", md5.Sum(theString))
		if strings.HasPrefix(hexVal, "00000") {
			return k
		}
	}
}

func FindMd5StartingWith000000(aString string) int {
	k := -1
	for {
		k += 1
		theString := []byte(aString + fmt.Sprintf("%d", k))
		hexVal := fmt.Sprintf("%x", md5.Sum(theString))
		if strings.HasPrefix(hexVal, "000000") {
			return k
		}
	}
}
