package gateway

import (
	"fmt"
	"github-program-rate/pkg/domain/model"
	svg "github.com/ajstarks/svgo"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

func NewServer(logger *zerolog.Logger) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/svg", circle)
	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "HELLOOOOOOO")
		logger.Info().Msg(" / Access is Successful")
	})
	mux.HandleFunc("/stun", func(writer http.ResponseWriter, r *http.Request) {
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		fmt.Fprintf(writer, "stun:")
		logger.Info().Msg(r.RemoteAddr)
		fmt.Fprintf(writer, r.RemoteAddr)
		logger.Info().Msg(" /stun Access is Successful")
	})
	server := &http.Server{
		// NOTE: ここ変えるならクライアントも変えなければならない
		// NOTE: 127.0.0.1 では繋がらないが、localhostは繋がる
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server
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