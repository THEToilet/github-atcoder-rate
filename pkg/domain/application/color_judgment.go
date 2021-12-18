package application

import "github-program-rate/pkg/domain/model"

func judge(rating uint) model.Color {
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
