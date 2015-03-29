package main

import (
	"./Router"
	"./Models"
	"./Config"
)

func main(){
	Config.Load()
	Models.Init()
	Router.Init()
}