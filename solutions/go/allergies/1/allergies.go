package allergies

/**
- eggs (1)
- peanuts (2)
- shellfish (4)
- strawberries (8)
- tomatoes (16)
- chocolate (32)
- pollen (64)
- cats (128)
**/

var allergens = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}
var allergensMap = map[string]int{
	"eggs":         0,
	"peanuts":      1,
	"shellfish":    2,
	"strawberries": 3,
	"tomatoes":     4,
	"chocolate":    5,
	"pollen":       6,
	"cats":         7,
}

func Allergies(allergies uint) []string {
	res := []string{}
	for k, v := range allergens {
		if allergies&(1<<k) != 0 {
			res = append(res, v)
		}
	}
	return res
}

func AllergicTo(allergies uint, allergen string) bool {
	return allergies&(1<<allergensMap[allergen]) != 0
}
