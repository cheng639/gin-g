package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var registerHandlers = make(map[string]*HandlerFunc)

func Handlers() map[string]*HandlerFunc {
	return registerHandlers
}

type HandlerFunc struct {
	Path   string
	Method string
	F      gin.HandlerFunc
}

func RegisterHandler(method, path string, f func(c *gin.Context)) {
	registerHandlers[method+path] = &HandlerFunc{
		Path:   path,
		Method: method,
		F:      f,
	}
}

func RegisterRouters(g *gin.Engine, basePath string) {
	var path string

	for _, v := range registerHandlers {
		path = basePath + v.Path
		switch strings.ToUpper(v.Method) {
		case http.MethodGet:
			g.GET(path, v.F)
		case http.MethodPost:
			g.POST(path, v.F)
		case http.MethodPut:
			g.PUT(path, v.F)
		case http.MethodDelete:
			g.DELETE(path, v.F)
		case http.MethodOptions:
			g.OPTIONS(path, v.F)
		case http.MethodPatch:
			g.PATCH(path, v.F)
		}
	}
}
