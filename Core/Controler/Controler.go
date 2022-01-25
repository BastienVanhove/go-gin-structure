package controler

type Controler struct {
	Name  string
	New   func()
	Route func(route string)
	Index func()
}
