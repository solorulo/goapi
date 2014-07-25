package controllers

import (
	"github.com/revel/revel"
	"goapi/app/api"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	if _, ok := c.Session["id_user"]; !ok{ 
		return c.RenderJson(api.NotLogged())
	}
	return c.RenderJson(api.MakeMap(1, "index"))
}

func (c App) Login(email, pass string) revel.Result {
	user, error := api.UserAuth (email, pass)
	if error {
		return c.RenderJson(api.Error())
	}
	c.Session["id_user"] = fmt.Sprintf("%v",user.Id);
	return c.RenderJson(api.MakeMap(1, "login succesful"))
}

func (c App) Logout() revel.Result {
	if _, ok := c.Session["id_user"]; !ok{ 
		return c.RenderJson(api.NotLogged())
	}
	delete(c.Session, "id_user")
	return c.RenderJson(api.MakeMap(1, "logout succesful"))
}
