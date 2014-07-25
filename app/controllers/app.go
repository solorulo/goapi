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
	data := map[string]interface{} {
		"code":		1,
		"msg":		"index",
	}
	return c.RenderJson(data)
}

func (c App) Login(email, pass string) revel.Result {
	user, error := api.UserAuth (email, pass)
	if error {
		return c.RenderJson(api.Error())
	}
	c.Session["id_user"] = fmt.Sprintf("%v",user.Id);
	data := map[string]interface{} {
		"code":		1,
		"msg":		"login succesful",
	}
	return c.RenderJson(data)
}

func (c App) Logout() revel.Result {
	if _, ok := c.Session["id_user"]; !ok{ 
		return c.RenderJson(api.NotLogged())
	}
	delete(c.Session, "id_user")
	data := map[string]interface{} {
		"code":		1,
		"msg":		"logout succesful",
	}
	return c.RenderJson(data)
}
