package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	width, height = 600, 320            //キャンパスの大きさ
	cells         = 100                 //格子のます目の数
	xyrange       = 30.0                // 軸の範囲(-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x単位　および　y単位あたりの画素数
	zscale        = height * 0.4        // z単位当たりの画素数
	angle         = math.Pi / 6         // x,y軸の角度(=30度)
	topc          = "#ff0000"
	middlec       = "#00ff00"
	bottomc       = "#0000ff"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			z := top(i, j)
			if z > 0.05 {
				fmt.Printf("<polygon points='%g %g %g %g %g %g %g %g' style='fill:%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, topc)
			} else if z < -0.05 {
				fmt.Printf("<polygon points='%g %g %g %g %g %g %g %g' style='fill:%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, bottomc)
			} else {
				fmt.Printf("<polygon points='%g %g %g %g %g %g %g %g' style='fill:%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, middlec)
			}

		}
	}
	fmt.Println("</svg>")
}

func top(i, j int) float64 {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	return f(x, y)
}

func corner(i, j int) (float64, float64, error) {
	// ます目(i,j)のかどの点(x,y)を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する
	z := f(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, errors.New("invalid value")
	}

	if math.IsNaN(z) {
		return 0, 0, errors.New("invalid value")
	}

	// (x,y,z)を2-D SVGキャンパス(sx,sy)へ等角的に投影
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0,0)からの距離
	return math.Sin(r) / r
}
