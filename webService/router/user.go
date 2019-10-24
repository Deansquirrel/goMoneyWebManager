package router

import "github.com/kataras/iris"

type user struct {
	app *iris.Application
	c   common
}

func NewRouterUser(app *iris.Application) *user {
	return &user{
		app: app,
		c:   common{},
	}
}

func (user *user) check(ctx iris.Context) {

}
