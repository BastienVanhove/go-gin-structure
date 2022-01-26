package controler

import "github.com/gin-gonic/gin"

type Controler struct {
	Route string
	Start func(string)
}
