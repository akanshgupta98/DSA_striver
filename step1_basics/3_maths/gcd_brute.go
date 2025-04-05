package main

import "fmt"

/*
a = 9, b = 12.
GCD =
Factors of a : 12  = 2x2x3
Factors of b : 18 = 2*3*3
GCD is 6.
Intuition:
1. Find factors of A. [2,2,3]
2. Find factors of B. [2,3,3]
{2,3}

3. Intersection of F(A),F(B) = GCD.
Factors ->
*/

func factors(n int) (factors []int) {
	i := 2
	fmt.Printf("Factors of %d are: ", n)
	// O(Log N)
	for n != 1 {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
		} else {
			i++
		}
	}
	factors = append(factors, 1)
	fmt.Printf("%v\n", factors)

	return
}

func intersection(f1, f2 []int) (i1 []int) {
	//O(N)
	tmpMap := make(map[int]int)

	for _, ele := range f1 {
		tmpMap[ele]++
	}

	for _, ele := range f2 {
		if val, ok := tmpMap[ele]; ok && val != 0 {
			i1 = append(i1, ele)
			tmpMap[ele]--
		}
	}
	fmt.Printf("Intersection of %v and %v is: %v\n", f1, f2, i1)

	return

}
func main() {
	n1 := 5
	n2 := 7
	f1 := factors(n1)
	f2 := factors(n2)
	i1 := intersection(f1, f2)
	gcd := 1
	for _, val := range i1 {
		gcd = gcd * val
	}
	fmt.Printf("GCD of %d and %d is: %d\n", n1, n2, gcd)

}
