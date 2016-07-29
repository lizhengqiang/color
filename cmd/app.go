package main

import (
	"image/color"
	"fmt"
	hbscolor "github.com/mengxiaozhu/color"
)

func main() {
	count,incorrect:=0,0
	for r:=0;r<256;r++{
		for g:=0;g<256;g++{
			for b:=0;b<256;b++{
				c:=color.RGBA{R:uint8(r), G: uint8(g), B:uint8(b)}
				hbs :=hbscolor.RGBToHLS(c)
				count++
				rgb:=hbs.HSB2RGB()
				if(c.R!=rgb.R||c.G!=rgb.G||c.B!=rgb.B){
					incorrect++
				}
			}
		}
	}
	fmt.Println("count:",count,"incorrect",incorrect,float64(incorrect)/float64(count),"%")
}
