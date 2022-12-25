package dump

import (
	"fmt"
	"os"
)

func ExitFunc() {
	defer fmt.Println("!")
	os.Exit(3)
}
