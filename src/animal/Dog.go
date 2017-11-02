package animal

import (
	"fmt"
)

type Dog struct{
    voice string
}

func NewDog() *Dog {
	return &Dog{"わんわん"}
}

func (d Dog) Bark(){
	fmt.Println(d.voice)
}