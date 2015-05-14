/*==================================================
	文章

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package article

import (
	"newWoku/controllers"
	"newWoku/lib/response"
	"newWoku/models/user"
)

type Controller struct {
	controllers.Base
}

func New() *Controller {
	controller := &Controller{}
	controller.NewModel(user.New())
	return controller
}

// @router /aa [Get]
func (this *Controller) Other() (int, []byte) {
	return response.Success("Delete success!")
}
