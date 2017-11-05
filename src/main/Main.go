package main

import (
	"animal"
)

func main(){
	var c animal.Barker = animal.NewCat()
	var d animal.Barker = animal.NewDog()

	c.Bark()
	d.Bark()
}