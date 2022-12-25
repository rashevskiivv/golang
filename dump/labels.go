package dump

import "fmt"

func labelsFunc() {
outerLoopLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			break outerLoopLabel //break, continue, goto
		}
	}
	fmt.Println("End")
}
