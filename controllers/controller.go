package controllers

type IController interface {
	RunServer()
}

func Run(i IController) {
	i.RunServer()
}
