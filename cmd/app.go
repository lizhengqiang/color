package main

import (
	"image/color"
	"fmt"
	hbscolor "github.com/mengxiaozhu/color"
)

func main() {
	count,correct:=0,0
	for r:=0;r<256;r++{
		for g:=0;g<256;g++{
			for b:=0;b<256;b++{
				c:=color.RGBA{R:uint8(r), G: uint8(g), B:uint8(b)}
				hbs :=hbscolor.RGBToHLS(c)
				count++
				rgb:=hbs.HSB2RGB()
				if(c.R!=rgb.R||c.G!=rgb.G||c.B!=rgb.B){
					correct++
				}
			}
		}
	}
	fmt.Println("count:",count,"correct",correct,float64(correct)/float64(count),"%")

}
