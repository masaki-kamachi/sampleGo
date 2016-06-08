package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

type Key struct {
	Hello string

}

func (c App) Index() revel.Result {
	return c.RenderJson(Key{"World"})
}
