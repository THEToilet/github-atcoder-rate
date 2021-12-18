package application

import (
	"github-program-rate/pkg/domain/model"
	svg "github.com/ajstarks/svgo"
	"github.com/rs/zerolog"
	"strconv"
)

type DrawSVGUseCase struct {
	logger *zerolog.Logger
}

func NewDrawSVGUseCase(logger *zerolog.Logger) *DrawSVGUseCase {
	return &DrawSVGUseCase{
		logger: logger,
	}
}

func (d *DrawSVGUseCase) Draw(name string, rate int32, canvas *svg.SVG) {
	d.logger.Info().Interface("name", name).Msg("")
	d.logger.Info().Interface("rate", rate).Msg("")

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

	canvas.Gstyle(string("fill:" + d.judgeColor(rate)))
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
	rateStr := strconv.Itoa(int(rate))
	canvas.Text(width/2, height/2, rateStr)
	canvas.Gend()

	canvas.Gstyle("fill:black")
	canvas.Text(width/4, height/4, name)
	canvas.Gend()

	canvas.Gend()
	canvas.End()
}

func (d *DrawSVGUseCase) judgeColor(rating int32) model.Color {
	switch rating / 400 {
	case 0:
		return model.Gray
	case 1:
		return model.Brown
	case 2:
		return model.Green
	case 3:
		return model.Blue
	case 4:
		return model.LightBlue
	case 5:
		return model.Yellow
	case 6:
		return model.Orange
	case 7:
		return model.Red
	default:
		return model.Gray
	}
}
