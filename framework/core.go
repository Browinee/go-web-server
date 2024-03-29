package framework

import (
	"log"
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
	router map[string]*Tree
}

func NewCore() *Core {

router := map[string]*Tree{}
router[GET] = NewTree()
router[POST] = NewTree()
router[PUT] = NewTree()
router[DELETE] = NewTree()

return &Core{router: router}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
    log.Fatal("add router error: ", err)
}
}
func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
    log.Fatal("add router error: ", err)
}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
    log.Fatal("add router error: ", err)
}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
    log.Fatal("add router error: ", err)
}
}


func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
 uri := request.URL.Path
 method := request.Method
 upperMethod := strings.ToUpper(method)

 if methodHandlers, ok := c.router[upperMethod]; ok {
 	return methodHandlers.FindHandler(uri)
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