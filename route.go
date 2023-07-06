package main

import "web-server/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}