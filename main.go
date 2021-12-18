package main

import (
	"github-program-rate/pkg/domain/model"
	svg "github.com/ajstarks/svgo"
	"log"
	"net/http"
)

func main() {
	http.Handle("/svg", http.HandlerFunc(circle))
	err := http.ListenAndServe(":82", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	canvas := svg.New(w)
	/*
		s.Start(800, 800)
		s.Circle(300, 300, 125, "fill:none;stroke:#3ab60b")
		s.Circle(305, 300, 105, "fill:none;stroke:#3ab60b")
		s.End()
	*/

	width := 500
	height := 500
	angle, cx, cy := 30.0, width/2, height/2
	r := width / 4

	canvas.Start(width, height)

	// NOTE: 背景
	canvas.Gstyle("fill:white")
	canvas.Rect(0, 0, width, height)
	canvas.Gend()

	// NOTE: 外枠
	canvas.Gstyle("fill:none;stroke:black")
	canvas.Roundrect(0, 0, width, height, 10, 10)
	canvas.Gend()

	canvas.Gstyle(string("fill:" + model.Gray))
	canvas.TranslateRotate(cx, cy, -angle)
	canvas.Arc(-r, 0, r, r, 30, false, true, r, 0)
	canvas.Gend()

	canvas.TranslateRotate(cx, cy, angle)
	canvas.Arc(-r, 0, r, r, 30, false, false, r, 0)
	canvas.Gend()

	canvas.Gstyle("fill:white")
	canvas.Circle(width/2, height/2, 90)
	canvas.Gend()

	canvas.Gstyle("fill:black")
	canvas.Text(width/2, height/2, "730")
	canvas.Gend()

	canvas.Gend()
	canvas.End()
}
