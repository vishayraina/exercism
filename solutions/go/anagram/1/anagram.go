package anagram

import (
    "unicode"
	"strings"
) 

func BoC(s string) map[rune]int {
    var m = map[rune]int{}
    for _, x := range s {
        m[unicode.ToLower(x)] += 1
    }
    return m
}

func IsMapSame(m, n map[rune]int) bool {
    if len(m) != len(n) {
        return false 
    }
    for k, v := range m {
        if v != n[k] {
            return false
        }
    }
    return true
}

func Detect(subject string, candidates []string) []string {
	subMap := BoC(subject)
    res := []string{}
    for _, x := range candidates {
        if strings.ToLower(subject) != strings.ToLower(x) {
            canMap := BoC(x)
            if IsMapSame(subMap, canMap) {
                res = append(res, x)
            }
        }
    }
    return res
}
