package controler

import "fmt"

type Controler struct {
	Name  string
	Start func(route string)
}

func (c *Controler) AddController(Controller Controler) {
	fmt.Println("u try to add a controller")
}