package router

import (
	"web-server/controller"
	"web-server/framework"
)

func RegisterRouter(core *framework.Core) {
	core.Get("/user/login", controller.UserLoginController)
	subjectApi := core.Group("/subject")
	{
			 subjectApi.Delete("/:id", controller.SubjectListController)
			 subjectApi.Put("/:id", controller.SubjectListController)
			 subjectApi.Get("/:id", controller.SubjectListController)
			subjectApi.Get("/list/all", controller.SubjectListController)
	 }
}