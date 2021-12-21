package application

import (
	"fmt"
	"github-program-rate/pkg/domain/model"
	svg "github.com/ajstarks/svgo"
	"github.com/rs/zerolog"
	"math"
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

	// NOTE: レートは400ごとに色が変わる
	percentPre := rate - (rate/400)*400
	percent := int32((float64(percentPre) / 400.0) * 100)
	d.logger.Info().Interface("percentPre", percentPre).Msg("")
	d.logger.Info().Interface("percent", percent).Msg("")

	width := 500
	height := 300

	canvas.Start(width, height)
	canvas.Gstyle("")

	// NOTE: 背景
	canvas.Gstyle("fill:white")
	canvas.Rect(0, 0, width, height)
	canvas.Gend()

	// NOTE: 外枠

	corner := 10

	//canvas.Gstyle("fill:black;stroke:black;stroke-width:5;stroke-linecap:round")
	canvas.Gstyle("fill:black")
	canvas.Roundrect(0, 0, width, height, corner, corner)
	canvas.Gend()

	canvas.Gstyle("fill:white")
	canvas.Roundrect(0+2, 0+2, width-5, height-5, corner, corner)
	canvas.Gend()

	// NOTE: cx: center-x, cy: center-y
	cx := width / 2
	cy := height / 2
	// NOTE: 半径
	r := width / 4

	// REFERENCE : https://webkatu.com/20150127-draw-pie-chart-in-svg/
	// NOTE: 円の終わりの角度
	endDeg := 360 * float64(percent) / 100.0
	d.logger.Info().Interface("endDeg", endDeg).Msg("")

	// NOTE: 円弧の始まり
	sx := cx + r*int(math.Sin(0.0))
	sy := cy - r*int(math.Cos(0.0))
	d.logger.Info().Interface("sx", sx).Interface("sy", sy).Msg("")
	// NOTE: 円弧の終わり（時計回り）
	ex := float64(cx) - float64(r)*math.Sin(endDeg*math.Pi/180.0)
	ey := float64(cy) + float64(r)*math.Cos(endDeg*math.Pi/180.0)
	d.logger.Info().Interface("eeex", math.Sin(endDeg*math.Pi/180.0)).Interface("eeey", math.Cos(endDeg*math.Pi/180.0)).Msg("")
	d.logger.Info().Interface("ex", ex).Interface("ey", ey).Msg("")
	d.logger.Info().Interface("ex", int(ex)).Interface("ey", int(ey)).Msg("")
	var largeArcFlag int
	if endDeg <= 180 {
		largeArcFlag = 0
	} else {
		largeArcFlag = 1
	}

	canvas.Gstyle(string("fill:"+d.judgeColor(rate)) + ";stroke:black;stroke-width:3")
	//canvas.TranslateRotate(cx, cy, -angle)
	// NOTE: sweepはfalseで反時計回り
	// NOTE: 第五引数は円弧の回転度なので必要ない
	//canvas.Arc(sx, sy, r, r, 0, largeArcFlag, false, int(ex), int(ey),"closepath:Z;Z:")
	move := fmt.Sprintf("M%d %d ", cx, cy)
	line := fmt.Sprintf("L%d %d ", sx, sy)
	arc := fmt.Sprintf("A%d %d 0 %d 0 %f %f", r, r, largeArcFlag, ex, ey)
	close := " Z"
	canvas.Path(move+line+arc+close, "stroke:black")
	canvas.Gend()

	// NOTE: 中心の白い円
	canvas.Gstyle("fill:white;stroke:black;stroke-width:3")
	canvas.Circle(width/2, height/2, 100)
	canvas.Gend()

	// NOTE: レート
	canvas.Gstyle("font-size:30pt")
	rateStr := strconv.Itoa(int(rate))
	canvas.Text(width/2-40, 160, rateStr, "fill:black;font-family:Meiryo;font-wight:bold")
	canvas.Gend()

	// NOTE: ユーザネーム
	canvas.Gstyle("fill:black")
	canvas.Text(width-90, height-30, name)
	canvas.Gend()

	// NOTE: AtCoder Rating
	canvas.Gstyle("fill:black;font-size:30pt")
	canvas.Text(100, 45, "AtCoder Rating")
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
		return model.LightBlue
	case 4:
		return model.Blue
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
