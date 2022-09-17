package main

import "fmt"

// 3597903 too high pt 1
// 1275120 too high part2
// 776160 too low part2

func main() {
	fmt.Println("1.", part1())
	fmt.Println("2.", part2())
}

// part1. # presents seems to be the sum of all factors of the house number * 10
func part1() int {
	n := 0
	// desired value is 36000000
	theList := []int{}
	for listSum(theList)*10 <= 36000000 {
		n += 1
		theList = factors(n)
		if n%100000 == 0 {
			fmt.Println(n)
		}
	}
	return n
}

// part1. # presents seems to be the sum of all factors of the house number * 10
func part2() int {
	n := 0
	// desired value is 36000000
	theList := []int{}
	for listSum(theList)*11 <= 36000000 {
		n += 1
		theList = factors(n)
		newList := []int{}
		for _, v := range theList {
			// no more of 2s when n > 50 * 2
			// no more 1s when n > 50 * 1
			if !(n > 50*v) {
				newList = append(newList, v)
			}
		}
		theList = newList

		if n%100000 == 0 {
			fmt.Println(n)
		}
	}
	return n
}

func howManyPresents(houseNumber int) int {
	j := factors(houseNumber)
	return listSum(j) * 10
}

func listSum(items []int) int {
	total := 0
	for _, v := range items {
		total += v
	}
	return total
}

// func sim(nElves, maxHouses, gifts int) {
// 	houses := map[int]int{}
// 	elf := 1
// 	for elf < nElves {
// 		count := 0
// 		for count < maxHouses {
// 			count += 1
// 			houseNumber := elf * count
// 			houses[houseNumber] += elf * gifts
// 		}
// 		elf += 1
// 	}
// 	max := 0
// 	for k := range houses {
// 		if k > max {
// 			max = k
// 		}
// 	}
// 	i := 0
// 	for i <= max {
// 		i += 1
// 		fmt.Println(i, houses[i])
// 	}
// }

// https://rosettacode.org/wiki/Factors_of_an_integer#Go
func factors(nr int) []int {
	fs := []int{}
	if nr < 1 {
		// fmt.Println("\nFactors of", nr, "not computed")
		return fs
	}
	// fmt.Printf("\nFactors of %d: ", nr)
	// fs := make([]int64, 1)
	fs = append(fs, 1)
	apf := func(p int, e int) {
		n := len(fs)
		for i, pp := 0, p; i < e; i, pp = i+1, pp*p {
			for j := 0; j < n; j++ {
				fs = append(fs, fs[j]*pp)
			}
		}
	}
	e := 0
	for ; nr&1 == 0; e++ {
		nr >>= 1
	}
	apf(2, e)
	for d := 3; nr > 1; d += 2 {
		if d*d > nr {
			d = nr
		}
		for e = 0; nr%d == 0; e++ {
			nr /= d
		}
		if e > 0 {
			apf(d, e)
		}
	}
	// fmt.Println("Number of factors =", len(fs))
	return fs
}

// https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}
