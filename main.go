package main
	
import (
	// "log"
	"os"
	"github.com/svgo"
	"net/http"
)

func main() {
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


func showJpg(w http.ResponseWriter, req *http.Request) {

}

func circle(w http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "image/svg+xml")
	// s := svg.New(w)
	// s.Start(500, 500)
	// s.Circle(250, 250, 125, "fill:none;stroke:black")
	// s.End()
	width, height := 500, 500
	rsize := 100
	csize := rsize / 2
	duration := 5.0
	repeat := 2
	imw, imh := 100, 144
	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Circle(csize, csize, csize, `fill="red"`, `id="circle"`)
	canvas.Image((width/2)-(imw/2), 0, imw, imh, "haystack.jpg", `id="gopher"`)
	canvas.Square(width-rsize, 0, rsize, `fill="blue"`, `id="square"`)
	canvas.Animate("#circle", "cx", 0, width, duration, repeat)
	canvas.Animate("#circle", "cy", 0, height, duration, repeat)
	canvas.Animate("#square", "x", width, 0, duration, repeat)
	canvas.Animate("#square", "y", height, 0, duration, repeat)
	canvas.Animate("#gopher", "y", 0, height, duration, repeat)
	canvas.End()
}
