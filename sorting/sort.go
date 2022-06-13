package sorting

import "sort"

func sortSlice(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

func sortIntSlice(arr []int) []int {
	sort.Ints(arr)
	return arr
}

type Person struct {
	age  int
	name string
}

type People []Person

func (p People) Len() int           { return len(p) }
func (p People) Less(i, j int) bool { return p[i].age < p[j].age }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sortStructSlice(people []Person) []Person {
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	return people
}

func sortInterface(people []Person) []Person {
	sort.Sort(People(people))
	return people
}
