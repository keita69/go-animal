package animal

import (
	"fmt"
)

type Cat struct{
	voice string
}

func NewCat() *Cat {
	return &Cat{"にゃーにゃー"}
}

func (c Cat) Bark(){
	fmt.Println(c.voice)
}
