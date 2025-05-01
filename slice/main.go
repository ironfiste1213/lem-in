package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 1 {
		if nbrs[0] > len(a)-1 {
			return nil
		}
		if nbrs[0] < 0 {
			nbrs[0] = len(a)+nbrs[0]
			
		}
		return a[nbrs[0]:]
	}
	if nbrs[1] < nbrs[0] {
		return nil
	}
	if nbrs[0] < 0 {
		nbrs[0] = len(a) + nbrs[0]
	}
	if nbrs[1] < 0 {
		nbrs[1] = len(a) + nbrs[1]
	}
	return a[nbrs[0]:nbrs[1]]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Println(Slice(a, 1))
	fmt.Println(Slice(a, 2, 4))
	fmt.Println(Slice(a, -3))
	fmt.Println(Slice(a, 18 , -1))
	fmt.Println(Slice(a, 2, 0))
}
