package main

import (
	"github.com/BourhaneYounes/url_shortener/model"
	"github.com/BourhaneYounes/url_shortener/server"
)



func main(){
	model.Setup()
	server.SetupAndListen()
}
