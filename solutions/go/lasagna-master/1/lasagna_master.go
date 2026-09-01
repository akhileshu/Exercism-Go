package lasagnamaster

import "slices"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preparationTimePerLayer int) int {
	if preparationTimePerLayer == 0 {
		preparationTimePerLayer = 2
	}
	return len(layers) * preparationTimePerLayer
}

func Count[T comparable](items []T, target T) int {
	count := 0

	for _, item := range items {
		if item == target {
			count++
		}
	}

	return count
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	return Count(layers, "noodles") * 50, float64(Count(layers, "sauce")) * .2
}

// TODO: define the 'AddSecretIngredient()' function
/*
`slices.Replace` uses a half-open range `[i:j)`.
```text
i = len(myList)-1   // start at last element
j = len(myList)     // stop after last element
```
i is inclusive
j is exclusive
*/
func AddSecretIngredient(friendsList []string, myList []string) {
	myList = slices.Replace(myList, len(myList)-1, len(myList), friendsList[len(friendsList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i := range quantities {
		scaledQuantities[i] = quantities[i]/2 * float64(portions)
	}
	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
