package stringset

import "strings"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
    m map[string]bool
    els []string
}

func New() Set {
	return Set{
        m: map[string]bool{},
        els: []string{},
    }
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, str := range l {
        if !s.m[str] {
            s.m[str] = true
            s.els = append(s.els, str)
        }
    }
    return s
}

func (s Set) String() string {
    res := []string{}
    for _, el := range s.els {
        res = append(res, `"` + el + `"`)
    }
	return `{` + strings.Join(res, ", ") + `}`
}

func (s Set) IsEmpty() bool {
	return len(s.m) == 0
}

func (s Set) Has(elem string) bool {
	return s.m[elem]
}

func (s Set) Add(elem string) {
	if !s.m[elem] {
        s.m[elem] = true
        s.els = append(s.els, elem)
    }
}

func Subset(s1, s2 Set) bool {
	for _, el := range s1.els {
        if !s2.Has(el) {
            return false
        }
    }
    return true
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

func Intersection(s1, s2 Set) Set {
	s := New()
    for _, el := range s1.els {
        if s2.Has(el) {
            s.Add(el)
        }
    }
    return s
}

func Difference(s1, s2 Set) Set {
	intersection := Intersection(s1, s2)
    diff := New()
    for _, el := range s1.els {
        if !intersection.Has(el) {
            diff.Add(el)
        }
    }
    for _, el := range s2.els {
        if !intersection.Has(el) {
            diff.Add(el)
        }
    }
    return diff
}

func Union(s1, s2 Set) Set {
	s := New()
    for _, el := range s1.els {
        s.Add(el)
    }
    for _, el := range s2.els {
        s.Add(el)
    }
    return s
}
