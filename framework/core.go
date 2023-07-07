package framework

import (
	"net/http"
	"strings"
)
const (
	GET="GET"
	POST="POST"
	PUT="PUT"
	DELETE="DELETE"
)
type Core struct {
	router map[string]map[string]ControllerHandler
}

func NewCore() *Core {
getRouter := map[string]ControllerHandler{}
postRouter := map[string]ControllerHandler{}
putRouter := map[string]ControllerHandler{}
deleteRouter := map[string]ControllerHandler{}

router := map[string]map[string]ControllerHandler{}
router[GET] = getRouter
router[POST] = postRouter
router[PUT] = putRouter
router[DELETE] = deleteRouter

return &Core{router: router}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router[GET][upperUrl] = handler
}
func (c *Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router[POST][upperUrl] = handler
}

func (c *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router[PUT][upperUrl] = handler
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router[DELETE][upperUrl] = handler
}


func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {

	upperMethod := strings.ToUpper(request.Method)
	upperUri := strings.ToUpper( request.URL.Path)
	if methodHandlers, ok := c.router[upperMethod]; ok {
		if handler, ok := methodHandlers[upperUri]; ok {
			return handler
		}
	}
	return nil

}
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(request, response)
	router := c.FindRouteByRequest(request)
	if router == nil {
		ctx.Json(404, "not found")
		return
	}
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}

}