package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "cxdnnyjw"
	fmt.Println("1. " + findPassword(input))
	fmt.Println("2. " + findPassword2(input))
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
	idx := 0
	password := []string{"_", "_", "_", "_", "_", "_", "_", "_"}
	out := "_"
	for strings.Contains(out, "_") {
		idx += 1
		t := hash(doorId + fmt.Sprint(idx))
		if t[:5] == "00000" {
			pos := string(t[5])
			position, _ := strconv.ParseInt("0x"+pos, 0, 64)
			if position >= 0 && position <= 7 {
				if password[position] == "_" {
					password[position] = string(t[6])
					out = catPass(password)
					fmt.Println(out)
				}
			}
		}
	}
	return out
}

func catPass(a []string) string {
	return strings.Join(a, "")
}
