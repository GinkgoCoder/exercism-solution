package resistorcolortrio

import (
	"fmt"
	"math"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	color2num := map[string]int64{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}
	nums := []int64{}
	for _, color := range colors {
		if v, ok := color2num[color]; ok {
			nums = append(nums, v)
		}
	}
	firstTwo := nums[0]*10 + nums[1]
	third := int64(math.Pow(10, float64(nums[2])))
	res := firstTwo * third
	degree := []string{"ohms", "kiloohms", "megaohms", "gigaohms"}
	index := 0
	for res >= 1000 {
		index++
		res /= 1000
	}
	return fmt.Sprintf("%d %s", res, degree[index])
}
