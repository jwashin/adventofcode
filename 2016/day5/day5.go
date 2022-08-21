package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	input := "cxdnnyjw"
	fmt.Println(findPassword(input))
}

func hash(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func findPassword(doorId string) string {
	password := ""
	idx := 0
	for len(password) < 8 {
		idx += 1
		t := hash(doorId + fmt.Sprint(idx))
		if t[:5] == "00000" {
			password += string(t[5])
			fmt.Println(password)
		}
	}
	return password
}
func findPassword2(doorId string) string {
	password := ""
	idx := 0
	for len(password) < 8 {
		idx += 1
		t := hash(doorId + fmt.Sprint(idx))
		if t[:5] == "00000" {
			password += string(t[5])
			fmt.Println(password)
		}
	}
	return password
}
