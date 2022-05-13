package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8

	pic = make([][]uint8, dx)

	for i := 0; i < dx; i++ {
		pic[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			pic[i][j] = uint8((i + j) / 2) //每次改变这里
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)

}
