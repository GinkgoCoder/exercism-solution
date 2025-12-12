package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }
    return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauces float64) {
    for _, layer := range(layers) {
        if layer == "noodles" {
            noodles += 50
        } else if layer == "sauce"{
            sauces += 0.2
        }
    }
    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendLayers []string, myLayers []string) {
    myLayers[len(myLayers) - 1] = friendLayers[len(friendLayers) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {
    newQs := make([]float64, len(quantities))
    copy(newQs, quantities)
    for i, _ := range(newQs) {
        newQs[i] = quantities[i] * float64(scale) / 2
    }
    return newQs
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
