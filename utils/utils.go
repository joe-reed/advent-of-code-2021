package utils

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ConvertToInt(s string) int {
	result, err := strconv.Atoi(s)
	Check(err)
	return result
}

func CheckRegexp(pattern string, value string) bool {
	isValid, err := regexp.MatchString(pattern, value)
	Check(err)
	return isValid
}

func IsInArray(needle string, haystack []string) bool {
	for _, validValue := range haystack {
		if needle == validValue {
			return true
		}
	}
	return false
}

func IntInArray(needle int, haystack []int) bool {
	for _, validValue := range haystack {
		if needle == validValue {
			return true
		}
	}
	return false
}

func GetKeys(aMap map[string]func(string) bool) []string {
	keys := make([]string, len(aMap))
	for key := range aMap {
		keys = append(keys, key)
	}
	return keys
}

func PrintTimeSince(start time.Time) {
	fmt.Println(time.Since(start))
}

func LoadFile(path string) []string {
	file, err := ioutil.ReadFile(path)
	Check(err)
	lines := strings.Split(string(file), "\n")
	return lines[:len(lines)-1]
}

func MapToInts(strings []string) []int {
	var ints []int
	for _, s := range strings {
		ints = append(ints, ConvertToInt(s))
	}
	return ints
}

func MinInt(ints []int) int {
	min := ints[0]
	for _, i := range ints {
		if i < min {
			min = i
		}
	}
	return min
}

func SortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}
