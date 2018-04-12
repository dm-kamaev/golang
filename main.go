//usr/local/bin/go run $0 $@ ; exit

package main

import "fmt"






// func main() {
// 	// init hash
// 	elements := map[string]map[string]string{
// 		"H": map[string]string{
// 			"name":  "Hydrogen",
// 			"state": "gas",
// 		},
// 		"He": map[string]string{
// 			"name":  "Helium",
// 			"state": "gas",
// 		},
// 		"Li": map[string]string{
// 			"name":  "Lithium",
// 			"state": "solid",
// 		},
// 		"Be": map[string]string{
// 			"name":  "Beryllium",
// 			"state": "solid",
// 		},
// 		"B": map[string]string{
// 			"name":  "Boron",
// 			"state": "solid",
// 		},
// 		"C": map[string]string{
// 			"name":  "Carbon",
// 			"state": "solid",
// 		},
// 		"N": map[string]string{
// 			"name":  "Nitrogen",
// 			"state": "gas",
// 		},
// 		"O": map[string]string{
// 			"name":  "Oxygen",
// 			"state": "gas",
// 		},
// 		"F": map[string]string{
// 			"name":  "Fluorine",
// 			"state": "gas",
// 		},
// 		"Ne": map[string]string{
// 			"name":  "Neon",
// 			"state": "gas",
// 		},
// 	}

// 	if el, ok := elements["Li"]; ok {
// 		fmt.Println(el["name"], el["state"])
// 	}
}

// func main() {
// 	hash := make(map[string]string) // OR var hash = make(map[string]string)
// 	hash["test"] = "hello world"
// 	fmt.Println(hash)
// 	key := "something"
// 	if val, ok := hash[key]; !ok {
// 		fmt.Println("Not exist val '" + val + "' ")
// 	}
// }

// func main() {
// 	slice_1 := []int{1, 2, 3}
// 	slice_2 := make([]int, 2)
// 	// create new array, but not modification exist array
// 	copy(slice_2, slice_1)
// 	slice_1[0] = 100
// 	fmt.Println(slice_1)
// 	fmt.Println(slice_2)
// }

// func main() {
// 	slice_1 := []int{1, 2, 3}
// 	// create new array, but not modification exist array
// 	slice_2 := append(slice_1, 4, 5, 6)
// 	slice_1[0] = 100
// 	fmt.Println(slice_1)
// 	fmt.Println(slice_2)
// }

// func main() {
// 	// slice uses memory, which use original array
// 	ar := [5]int{1, 2, 3, 4, 5}
// 	slice := ar[0:3]
// 	fmt.Println(ar)
// 	fmt.Println(slice)
// 	ar[0] = 200
// 	fmt.Println()
// 	fmt.Println(ar)
// 	fmt.Println(slice)
// 	slice[10] = 2
// }

// func main() {
// 	ar := [5]float64{1, 2, 3, 4, 5}
// 	var total float64 = 0
// 	for _, val := range ar {
// 		total += val
// 	}
// 	fmt.Println(total / float64(len(ar)))
// }

// func main() {
// 	var ar [5]float64
// 	ar[0] = 1
// 	ar[1] = 2
// 	ar[2] = 3
// 	ar[3] = 4
// 	ar[4] = 5
// 	var total float64 = 0
// 	for _, val := range ar {
// 		total += val
// 	}
// 	fmt.Println(total / float64(len(ar)))
// }

// func main() {
// 	var ar [5]float64
// 	ar[0] = 1
// 	ar[1] = 2
// 	ar[2] = 3
// 	ar[3] = 4
// 	ar[4] = 5
// 	var total float64 = 0
// 	for i := 0; i < len(ar); i++ {
// 		total += ar[i]
// 	}
// 	fmt.Println(total / float64(len(ar)))
// }

// func main() {
// 	var ar [5]float64
// 	ar[0] = 1
// 	ar[1] = 2
// 	ar[2] = 3
// 	ar[3] = 4
// 	ar[4] = 5
// 	for i := 0; i < len(ar); i++ {
// 		fmt.Println(ar[i])
// 	}
// }
