// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tEvent struct {}
var Event tEvent


func (_ tEvent) Activate(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Event.Activate", args).Url
}

func (_ tEvent) ListenActivate(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Event.ListenActivate", args).Url
}

func (_ tEvent) Frame_tx(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Event.Frame_tx", args).Url
}

func (_ tEvent) DataFetch(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Event.DataFetch", args).Url
}


type tLogger struct {}
var Logger tLogger


func (_ tLogger) Monitor(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Logger.Monitor", args).Url
}

func (_ tLogger) MonitorWS(
		user string,
		ws interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "ws", ws)
	return revel.MainRouter.Reverse("Logger.MonitorWS", args).Url
}


type tTheory struct {}
var Theory tTheory


func (_ tTheory) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Theory.Index", args).Url
}

func (_ tTheory) AwgnBpsk(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Theory.AwgnBpsk", args).Url
}

func (_ tTheory) AwgnQpsk(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Theory.AwgnQpsk", args).Url
}


type tActivator struct {}
var Activator tActivator


func (_ tActivator) Activate(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Activator.Activate", args).Url
}

func (_ tActivator) ActivatePost(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Activator.ActivatePost", args).Url
}

func (_ tActivator) Join(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Activator.Join", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tCommand struct {}
var Command tCommand


func (_ tCommand) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Command.Index", args).Url
}

func (_ tCommand) Activity(
		user string,
		ws interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "ws", ws)
	return revel.MainRouter.Reverse("Command.Activity", args).Url
}


type tData struct {}
var Data tData


func (_ tData) Meta(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Data.Meta", args).Url
}

func (_ tData) Fetch(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Data.Fetch", args).Url
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


