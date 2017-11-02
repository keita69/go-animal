package animal

import (
	"fmt"
)

type Barker interface {
	Bark(string)
}

func Bark(voice string) {
	fmt.Println(voice)
}