package main

import (
	"fmt"
	"github.com/yiwenlong/lifecycle"
)

type HelloWorldHttpService struct {

}

func (service *HelloWorldHttpService) Start() error {
	return nil
}

func (service *HelloWorldHttpService) Stop() error {
	return nil
}

func main() {
	app := lifecycle.New()
	if err := app.Run(); err != nil {
		fmt.Printf("Finish with error: %+v\n", err)
		return
	}
	fmt.Printf("Done!")
}