package main

import (
	"animal"
)

func main(){
	c := animal.NewCat()
	d := animal.NewDog()

	c.Bark()
	d.Bark()
}