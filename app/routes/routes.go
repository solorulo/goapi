// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Login(
		email string,
		pass string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "pass", pass)
	return revel.MainRouter.Reverse("App.Login", args).Url
}

func (_ tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).Url
}


type tUsers struct {}
var Users tUsers


func (_ tUsers) Add(
		nombre string,
		apellido string,
		email string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "nombre", nombre)
	revel.Unbind(args, "apellido", apellido)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Users.Add", args).Url
}

func (_ tUsers) Edit(
		id string,
		nombre string,
		apellido string,
		email string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "nombre", nombre)
	revel.Unbind(args, "apellido", apellido)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Users.Edit", args).Url
}

func (_ tUsers) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.List", args).Url
}

func (_ tUsers) Get(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Users.Get", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


