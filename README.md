#color

```
import (
	"image/color"
	"fmt"
	hbscolor "github.com/mengxiaozhu/color"
)

func main() {

		hbs :=hbscolor.RGBToHLS(color.RGBA{R:153, G: 153, B:153})
    	fmt.Print(hbs.Gethsb())
    
    	rgb:=hbs.HSB2RGB()
    	fmt.Println(rgb)
    	hbs =hbscolor.RGBToHLS(color.RGBA{R:255, G: 255, B:255})
    	fmt.Print(hbs.Gethsb())
    
    	rgb=hbs.HSB2RGB()
    	fmt.Println(rgb)
    
    	hbs=hbscolor.RGBToHLS(color.RGBA{R:0, G: 0, B:0})
    	fmt.Print(hbs.Gethsb())
    
    	rgb=hbs.HSB2RGB()
    	fmt.Println(rgb)
    
    	hbs=hbscolor.RGBToHLS(color.RGBA{R:127, G: 0, B:127})
    	fmt.Print(hbs.Gethsb())
    
    	rgb=hbs.HSB2RGB()
    	fmt.Println(rgb)
}
```

output
```
0 0 60{153 153 153 0}
0 0 100{255 255 255 0}
0 0 0{0 0 0 0}
300 100 49{127 0 127 0}
```