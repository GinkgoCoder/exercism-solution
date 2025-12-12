package thefarm
import "fmt"
import "errors"

// TODO: define the 'DivideFood' function
func DivideFood(cal FodderCalculator, cows int) (float64, error) {
    total, err := cal.FodderAmount(cows)
    if err != nil {
        return 0., err
    }
    factor, err := cal.FatteningFactor()
    if err != nil {
        return 0., err
    }
    return total * factor / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(cal FodderCalculator, cows int) (float64, error) {
    if (cows <= 0) {
        return 0., errors.New("invalid number of cows")
    }
    return DivideFood(cal, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
    message string
    cows int
}

func (err *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", err.cows, err. message)
}

func ValidateNumberOfCows(cows int) (error) {
	if (cows < 0) {
        return  &InvalidCowsError{
            message: "there are no negative cows",
            cows: cows,
        }
    }
    if (cows == 0) {
        return  &InvalidCowsError{
            message: "no cows don't need food",
            cows: cows,
        }
    }
    return nil
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
