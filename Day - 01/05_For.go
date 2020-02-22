/*
for is Go’s only looping construct. Here are three basic types of for loops.
1. The most basic type, with a single condition.
2. A classic initial/condition/after for loop.
3. for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
You can also continue to the next iteration of the loop.
*/
package main

import (
    "fmt"
)

func main() {
	for {
		fmt.Println("1st Loop")
		break
	}

	i := 1 /*Short variable declarations*/
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("2nd Loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
