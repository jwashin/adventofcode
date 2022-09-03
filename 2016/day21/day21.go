package main

import (
	"fmt"
	"os"
	"strings"
)

// https://installmd.com/c/108/go/reverse-a-slice
func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	data := getInstructions()
	s := "abcdefgh"
	for _, v := range data {
		s = doInstruction(s, v)
	}
	fmt.Println("1.", s)
	toUnscramble := "fbgdceah"
	fmt.Println("2.", unscramble(toUnscramble, data))
}

func getInstructions() []string {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")
	return data
}

//  fbgdceah wrong for 2. hmm. !!!damn. fbgdceah is the input, so return the permutation!!!!

// https://code-maven.com/slides/golang/solution-permutations
func permutations(word string) []string {
	if word == "" {
		return []string{""}
	}
	perms := []string{}
	for i, rn := range word {
		rest := word[:i] + word[i+1:]
		//fmt.Println(rest)
		for _, result := range permutations(rest) {
			perms = append(perms, fmt.Sprintf("%c", rn)+result)
		}
		//perms = append(perms, fmt.Sprintf("%c\n", rn))
	}
	return perms
}

// fdhegbca wrong, so the answer is not the string that results in the value
func unscramble(a string, data []string) string {
	newAnswer := a
	permutation := ""
	for _, s := range permutations("abcdefgh") {
		permutation = s
		for _, v := range data {

			s = doInstruction(s, v)
		}
		if s == newAnswer {
			fmt.Println("2.", permutation)
		}
	}
	return ""
}

func doInstruction(s string, inst string) string {
	inst = strings.TrimSpace(inst)
	if strings.Index(inst, "swap") == 0 {
		if strings.Index(inst, "position") > 0 {
			var a, b int
			n, _ := fmt.Sscanf(inst, "swap position %d with position %d", &a, &b)
			if n < 2 {
				fmt.Println("swap position")
			}
			return swapPosition(s, a, b)
		}
		if strings.Index(inst, "letter") > 0 {
			var a, b string

			n, _ := fmt.Sscanf(inst, "swap letter %s with letter %s", &a, &b)
			if n < 2 {
				fmt.Println("swap letter", n, a, b, inst)
			}
			return swapLetters(s, a, b)
		}
	}

	if strings.Index(inst, "reverse") == 0 {
		var a, b int
		n, _ := fmt.Sscanf(inst, "reverse positions %d through %d", &a, &b)
		if n < 2 {
			fmt.Println("reverse")
		}
		return reversePositions(s, a, b)

	}
	if strings.Index(inst, "move") == 0 {
		var a, b int
		n, _ := fmt.Sscanf(inst, "move position %d to position %d", &a, &b)
		if n < 2 {
			fmt.Println("move position")
		}
		return movePosition(s, a, b)
	}
	if strings.Index(inst, "rotate") == 0 {
		if strings.Index(inst, "left") > 0 {
			var a int
			n, _ := fmt.Sscanf(inst, "rotate left %d step", &a)
			if n < 1 {
				fmt.Println("rotate left")
			}
			return rotateLeft(s, a)
		}

		if strings.Index(inst, "right") > 0 {
			var a int
			n, _ := fmt.Sscanf(inst, "rotate right %d step", &a)
			if n < 1 {
				fmt.Println("rotate right")
			}
			return rotateRight(s, a)
		}

		if strings.Index(inst, "based") > 0 {
			var a string
			n, _ := fmt.Sscanf(inst, "rotate based on position of letter %s", &a)
			if n < 1 {
				fmt.Println("rotate position")
			}
			return rotateBasedOn(s, a)
		}

	}
	return ""
}

func rotateBasedOn(s string, a string) string {
	ix := strings.Index(s, a)
	nrot := 1 + ix
	if ix >= 4 {
		nrot += 1
	}
	for nrot > 0 {
		s = rotateRight(s, 1)
		nrot -= 1
	}
	return s
}

func rotateRight(s string, n int) string {
	// abcd, 1 -> dabc
	return s[len(s)-n:] + s[:len(s)-n]
}

func rotateLeft(s string, n int) string {
	return s[n:] + s[:n]
}

func movePosition(s string, a, b int) string {
	ins := string(s[a])
	ns := s[:a] + s[a+1:]
	s2 := ns[:b] + ins + ns[b:]
	return s2
}

func reversePositions(s string, a, b int) string {
	start := s[:a]
	end := s[b+1:]
	toReverse := s[a : b+1]
	aList := strings.Split(toReverse, "")
	Reverse(aList)
	reversed := strings.Join(aList, "")
	return start + reversed + end

}

func swapPosition(s string, a, b int) string {
	new := strings.Split(s, "")
	astr := new[a]
	bstr := new[b]
	new[a] = bstr
	new[b] = astr
	return strings.Join(new, "")
}

func swapLetters(s string, a, b string) string {
	ai := strings.Index(s, a)
	bi := strings.Index(s, b)
	return swapPosition(s, ai, bi)
}
