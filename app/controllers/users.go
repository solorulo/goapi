package controllers

import (
	"github.com/revel/revel"
	"goapi/app/api"
	"strconv"
)

type Users struct {
	*revel.Controller
}

func (c Users) Add(nombre, apellido, email, password string) revel.Result {
	if _, ok := c.Session["id_user"]; !ok { 
		return c.RenderJson(api.NotLogged())
	}
	error := api.AddUser (nombre, apellido, email, password)
	if error {
		return c.RenderJson(api.Error())
	}
	return c.RenderJson(api.MakeMap(1, "a success"))
}

func (c Users) Edit(id, nombre, apellido, email, password string) revel.Result {
	if _, ok := c.Session["id_user"]; !ok { 
		return c.RenderJson(api.NotLogged())
	}
	uid, _ := strconv.ParseInt(id, 0, 64)
	error := api.EditUser (uid, nombre, apellido, email, password)
	if error {
		return c.RenderJson(api.Error())
	}
	return c.RenderJson(api.MakeMap(1, "edit success"))
}

func (c Users) List() revel.Result {
	if _, ok := c.Session["id_user"]; !ok { 
		return c.RenderJson(api.NotLogged())
	}
	users, error := api.GetUsers ()
	if error {
		return c.RenderJson(api.Error())
	}
	arr := [](map[string]interface{}){}
	for _,user := range users {
		arr = append(arr, map[string]interface{}{
			"id"		:	user.Id,
			"nombre"	:	user.Nombre,
			"apellido"	: 	user.Apellido,
			"email"		: 	user.Email,
		})
	}
	data := api.MakeMapData(1, "list users", arr)
	return c.RenderJson(data)
}

func (c Users) Get(id string) revel.Result {
	if _, ok := c.Session["id_user"]; !ok { 
		return c.RenderJson(api.NotLogged())
	}
	uid, _ := strconv.ParseInt(id, 0, 64)
	user, error := api.GetUser (uid)
	if error {
		return c.RenderJson(api.Error())
	}
	data := api.MakeMapData(1, "get users", map[string]interface{}{
		"id"		:	user.Id,
		"nombre"	:	user.Nombre,
		"apellido"	: 	user.Apellido,
		"email"		: 	user.Email,
	})
	return c.RenderJson(data)
}
