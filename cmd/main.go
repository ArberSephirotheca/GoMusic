package main

import (
	music "GoMusic/internal"
	"fmt"
)

func main() {
	songs := music.LoadDefault()
	fmt.Println(songs[0].Year())
}
