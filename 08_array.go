package main

import "fmt"

func main() {

    var a [5]int
    fmt.Println("Initial Stage:", a)

    a[4] = 100
    fmt.Println("After Set:", a)
    fmt.Println("Specific Index:", a[4])

    fmt.Println("Length of Array:", len(a))
	
	/*Use this syntax to declare and initialize an array in one line.*/
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("New Array:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("DD Array: ", twoD)
}