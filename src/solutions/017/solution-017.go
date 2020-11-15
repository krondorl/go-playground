package main

import (
	"fmt"
)

func isEqualList(a, b []int) bool {
	if len(a) == len(b) {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	return false
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func isSubList(subList, list []int) bool {
	if len(subList) < len(list) {
		for i := 0; i < len(subList); i++ {
			if contains(list, subList[i]) == false {
				return false
			}
		}

		return true
	}

	return false
}

func isSuperList(superList, list []int) bool {
	if len(superList) > len(list) {
		if isSubList(list, superList) {
			return true
		}
	}

	return false
}

func checkListRelationType(a, b []int) string {
	if isEqualList(a, b) {
		return "equal"
	} else if isSubList(a, b) {
		return "sublist"
	} else if isSuperList(a, b) {
		return "superlist"
	}

	return "not equal, superlist or sublist"
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3}
	c := []int{1, 4}
	d := []int{2, 5}

	fmt.Println("Sublist")
	fmt.Println("")
	fmt.Println(`a := []int{1, 2, 3, 4, 5}
b := []int{1, 2, 3}
c := []int{1, 4}
d := []int{2, 5}`)
	fmt.Println("")
	fmt.Println("checkListRelationType(a,b) = ", checkListRelationType(a, b))
	fmt.Println("checkListRelationType(b,b) = ", checkListRelationType(b, b))
	fmt.Println("checkListRelationType(b,a) = ", checkListRelationType(b, a))
	fmt.Println("checkListRelationType(a,c) = ", checkListRelationType(a, c))
	fmt.Println("checkListRelationType(c,d) = ", checkListRelationType(c, d))
	fmt.Println("")
	fmt.Println("isEqualList(c,c) = ", isEqualList(c, c))
	fmt.Println("isEqualList(c,d) = ", isEqualList(c, d))
	fmt.Println("")
	fmt.Println("isSubList(b,a) = ", isSubList(b, a))
	fmt.Println("isSubList(a,b) = ", isSubList(a, b))
	fmt.Println("")
	fmt.Println("isSuperList(a,b) = ", isSuperList(a, b))
	fmt.Println("isSuperList(b,a) = ", isSuperList(b, a))
	fmt.Println("")
}
