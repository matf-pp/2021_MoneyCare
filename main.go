package main

import (
	"github.com/gotk3/gotk3/gtk"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"main/src/admin"
	_ "main/src/services"
)

func main() {

	admin.SetupSeed()
	gtk.Init(nil)
	setupGui()
	gtk.Main()

}
