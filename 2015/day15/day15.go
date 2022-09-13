package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("1.", maximizeScore(string(input)))
	fmt.Println("2.", maximizeScore500(string(input)))
}

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func getIngredients(s string) []ingredient {
	r := []ingredient{}
	name := ""
	capacity, durability, flavor, texture, calories := 0, 0, 0, 0, 0
	for _, v := range strings.Split(s, "\n") {
		fmt.Sscanf(v, `%s capacity %d, durability %d, flavor %d, texture %d, calories %d`, &name, &capacity, &durability, &flavor, &texture, &calories)
		item := ingredient{name, capacity, durability, flavor, texture, calories}
		r = append(r, item)
	}
	return r
}

func getScore(ingredients []ingredient, amounts []int) int {
	capacity, durability, flavor, texture := 0, 0, 0, 0
	for i, v := range ingredients {
		capacity += amounts[i] * v.capacity
		durability += v.durability * amounts[i]
		flavor += v.flavor * amounts[i]
		texture += v.texture * amounts[i]
	}
	if capacity < 0 {
		capacity = 0
	}
	if durability < 0 {
		durability = 0
	}
	if flavor < 0 {
		flavor = 0
	}
	if texture < 0 {
		texture = 0
	}
	return capacity * durability * flavor * texture
}

func maximizeScore(s string) int {
	ingredients := getIngredients(s)
	ingredientCount := len(ingredients)
	max := 0
	amounts := next(ingredientCount, 100, []int{})
	for amounts != nil {

		score := getScore(ingredients, amounts)
		if score > max {
			max = score
		}
		amounts = next(ingredientCount, 100, amounts)
	}
	return max
}

func countCalories(ingredients []ingredient, amounts []int) int {
	calories := 0
	for i, v := range ingredients {
		calories += amounts[i] * v.calories
	}
	return calories
}

func maximizeScore500(s string) int {
	ingredients := getIngredients(s)
	ingredientCount := len(ingredients)
	max := 0
	amounts := next(ingredientCount, 100, []int{})
	for amounts != nil {

		score := getScore(ingredients, amounts)
		calories := countCalories(ingredients, amounts)
		if score > max && calories == 500 {
			max = score
		}
		amounts = next(ingredientCount, 100, amounts)
	}
	return max
}

func listSum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func countItems(elementCount int, target int) int {
	current := next(elementCount, target, []int{})
	x := 0
	for {
		x += 1
		current = next(elementCount, target, current)
		if current == nil {
			return x
		}
	}

}

// generates all lists of elementCount items with sum of target
// give it nil to get the first one. the last one gets you nil

func next(elementCount int, target int, current []int) []int {

	// first one
	if len(current) == 0 || current == nil {
		out := []int{0}
		for len(out) < elementCount {
			out = append(out, 1)
		}
		out[0] = target - listSum(out[1:])
		return out
	}
	// last one
	if current[0] == 1 && current[1] == target-elementCount+1 {
		return nil
	}

	// add 1 to the last column
	currentIndex := elementCount - 1
	current[currentIndex] += 1

	// arithmetic mod target--. carry the 1.
	for current[currentIndex] == target-elementCount+2 && currentIndex >= 1 {
		current[currentIndex] = 1
		current[currentIndex-1] += 1
		currentIndex -= 1
	}

	//  alter the first column so we sum to target
	currentValue := target - listSum(current[1:])
	current[0] = currentValue

	//  throw the list away and get another if not positive
	if currentValue <= 0 {
		return next(elementCount, target, current)
	}

	// fmt.Println(current)

	return current
}
