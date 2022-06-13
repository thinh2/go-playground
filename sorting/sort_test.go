package sorting

import "testing"

func TestSortSlice(t *testing.T) {
	input := []int{-1, 10, 2, 9, 5}
	expected := []int{-1, 2, 5, 9, 10}
	output := sortSlice(input)
	for i := 0; i < len(output); i++ {
		if expected[i] != output[i] {
			t.Errorf("failed to sort, expected %v, output %v", expected, output)
			return
		}
	}
}

func TestSortStructSlice(t *testing.T) {
	input := []Person{
		{
			age:  30,
			name: "Eric",
		},
		{
			age:  14,
			name: "Xavi",
		},
		{
			age:  25,
			name: "Neymar",
		},
		{
			age:  9,
			name: "Messi",
		},
	}
	expected := []Person{
		{
			age:  9,
			name: "Messi",
		},
		{
			age:  14,
			name: "Xavi",
		},
		{
			age:  25,
			name: "Neymar",
		},
		{
			age:  30,
			name: "Eric",
		},
	}
	output := sortStructSlice(input)
	for i := 0; i < len(output); i++ {
		if output[i].name != expected[i].name {
			t.Errorf("sorting error, expected %v, output %v", expected, output)
		}
	}
}

func TestSortInterfaceSlice(t *testing.T) {
	input := []Person{
		{
			age:  30,
			name: "Eric",
		},
		{
			age:  14,
			name: "Xavi",
		},
		{
			age:  25,
			name: "Neymar",
		},
		{
			age:  9,
			name: "Messi",
		},
	}
	expected := []Person{
		{
			age:  9,
			name: "Messi",
		},
		{
			age:  14,
			name: "Xavi",
		},
		{
			age:  25,
			name: "Neymar",
		},
		{
			age:  30,
			name: "Eric",
		},
	}
	output := sortInterface(input)
	for i := 0; i < len(output); i++ {
		if output[i].name != expected[i].name {
			t.Errorf("sorting error, expected %v, output %v", expected, output)
		}
	}
}
