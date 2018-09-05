package main

import(
	"fmt"
	"math"
	"strconv"
	"os"
)

const(
	width, height = 600, 320
	cells 	= 100
	xyrange = 30.0
	xyscale = width / 2 / xyrange
	zscale = height * 0.4
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main(){
	result := ""
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	result = "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='" + strconv.Itoa(width) + "' height='" + strconv.Itoa(height) +"'>"
	
	for i :=0; i < cells; i++{
		for j := 0; j < cells; j++{
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polugon points='%g,%g %g,%g %g,%g %g,%g'/>", ax,ay,bx,by,cx,cy,dx,dy)
			result = result + "<polugon points="+ "'" + strconv.FormatFloat(ax, 'f',-1, 64)+ "," +
			strconv.FormatFloat(ay, 'f',-1, 64) + " " +
			strconv.FormatFloat(bx, 'f',-1, 64)+ "," + 
			strconv.FormatFloat(by, 'f',-1, 64) + " " +
			strconv.FormatFloat(cx, 'f',-1, 64)+ "," + 
			strconv.FormatFloat(cy, 'f',-1, 64) + " " +
			strconv.FormatFloat(dx, 'f',-1, 64)+ "," + 
			strconv.FormatFloat(dy, 'f',-1, 64) +  "'/>"
		}
	}

	fmt.Println("</svg>")
	result = result + "</svg>"

	f,err := os.Create("result.svg")	
	if err!=nil {
		fmt.Println(err)
	}

	f.WriteString(result)
	f.Close()
}

func corner(i, j int)(float64, float64){
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x,y)

	sx := width/2 + (x - y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64{
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}