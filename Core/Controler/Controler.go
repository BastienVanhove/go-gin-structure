package controler

type Controler struct {
	Route string
	Start func(route string)
}
