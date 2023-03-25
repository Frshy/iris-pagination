package users

import (
  "github.com/kataras/iris/v12"
)

type ControllerT struct {}
var Controller ControllerT

func (c *ControllerT) Init(app *iris.Application) {
  app.Get("/", Service.Dashboard)
}
