/*package main

import "fmt"

var a int = 1
var k int = 0

func main() {
	for i := 0; i < 5; i++ {
		for x := k; x < 5; x++ {
			fmt.Print(" ")
		}

		for j := 0; j < a; j++ {
			fmt.Print("*")
		}
		a += 2
		k++
		fmt.Println()
	}
}
*/

package main

import "fmt"

var a int = 5

func main() {
	for i := 0; i < a; i++ {
		for j := 0; j < (2*a - 1); j++ {
			if j < (a-i) || j > (a+i) {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()

	}

}
