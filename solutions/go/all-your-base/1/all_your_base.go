package allyourbase

import (
	"errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}
	res := 0
	for i := 0; i < len(inputDigits); i++ {
		if inputDigits[i] >= 0 && inputDigits[i] < inputBase {
			res = res*inputBase + inputDigits[i]
		} else {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}
	if res == 0 {
		return []int{0}, nil
	}
	final := []int{}
	for res > 0 {
		r, remain := res/outputBase, res%outputBase
		final = append([]int{remain}, final...)
		res = r
	}
	return final, nil
}
