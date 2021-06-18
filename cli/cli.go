package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ajstarks/svgo"
)

func Run(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 1
	}
	return 0
}

type appEnv struct {
	// JULIE, I DON'T LIKE THIS ANYMORE
	filters	filters
	in		string
	out		string
}


func (app *appEnv) fromArgs(args []string) error {
	flag.StringVar(&app.out, "out", "./out/img.xml", "Path to store SVG output")
	flag.StringVar(&app.in, "in", "./img/img.png", "Path to access input image")
	flag.Var(&app.filters, "filters", "Filter names from the svg-fun.rc file to apply to the input file. Multiple filter names can be provided as a single string with spaces between them, enclosed in double quotes")
	flag.Parse()
	fmt.Println(app.out)
	fmt.Println(app.in)
	fmt.Println(app.filters)
	return nil
}

func (app *appEnv) addElement() {
	// Does this need to be on app??
}

func (app *appEnv) writeOutSvg() {
	// Do the thing
	// I/O (O) to app.output
}

type filters []string

func (fl *filters) String() string {
	return "YOU DON'T DESERVE TO KNOW"
}

func (fl *filters) Set(s string) error {
	sp := strings.Split(s, " ")
	*fl = append(*fl, sp...)
	return nil
}

func otherStuff() {
	width, height := 500, 500
	rsize := 100
	csize := rsize / 2
	duration := 5.0
	repeat := 5
	imw, imh := 100, 144
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Circle(csize, csize, csize, `fill="red"`, `id="circle"`)
	canvas.Image((width/2)-(imw/2), 0, imw, imh, "gopher.jpg", `id="gopher"`)
	canvas.Square(width-rsize, 0, rsize, `fill="blue"`, `id="square"`)
	canvas.Animate("#circle", "cx", 0, width, duration, repeat)
	canvas.Animate("#circle", "cy", 0, height, duration, repeat)
	canvas.Animate("#square", "x", width, 0, duration, repeat)
	canvas.Animate("#square", "y", height, 0, duration, repeat)
	canvas.Animate("#gopher", "y", 0, height, duration, repeat)
	canvas.End()
}
