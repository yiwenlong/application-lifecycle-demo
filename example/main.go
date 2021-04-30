package main

import (
	"fmt"
	"github.com/yiwenlong/lifecycle"
)

func main() {
	app := lifecycle.New()
	if err := app.Run(); err != nil {
		fmt.Printf("Finish with error: %+v\n", err)
		return
	}
	fmt.Printf("Done!")
}