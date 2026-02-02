package perfect

import (
	"errors"
	"math"
)

// Define the Classification type here.
type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
	ClassficationUnknown
)

var ErrOnlyPositive = errors.New("Only support positive number")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return ClassficationUnknown, ErrOnlyPositive
	}
	if n == 1 {
		return ClassificationDeficient, nil
	}
	var sum int64 = 1
	var i int64
	for i = 2; float64(i) < math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			sum += i
			if i*i != n {
				sum += n / i
			}
		}
	}
	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
