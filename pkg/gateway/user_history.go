package gateway

import (
	"encoding/json"
	"fmt"
	"github-program-rate/pkg/domain/application"
	"github-program-rate/pkg/gateway/data"
	svg "github.com/ajstarks/svgo"
	"github.com/rs/zerolog"
	"io/ioutil"
	"net/http"
)

type UserHistoryHandler struct {
	drawSVGUseCase *application.DrawSVGUseCase
	logger         *zerolog.Logger
}

func NewUserHistoryHandler(drawSVGUseCase *application.DrawSVGUseCase, logger *zerolog.Logger) *UserHistoryHandler {
	return &UserHistoryHandler{
		drawSVGUseCase: drawSVGUseCase,
		logger:         logger,
	}
}

func (u *UserHistoryHandler) drawRating(writer http.ResponseWriter, request *http.Request) {

	u.logger.Info().Msg("/svg Access")
	query := request.URL.Query()
	name := query.Get("name")
	u.logger.Info().Interface("name", name).Msg(name)

	if name == "" {
		fmt.Fprintf(writer, "query parameters are required ex:) svg?name=XXX")
	} else {

		writer.Header().Set("Content-Type", "image/svg+xml")
		res, err := http.Get("https://atcoder.jp/users/" + name + "/history/json")
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			u.logger.Fatal().Err(err).Msg("Error")
		}
		info := string(body)
		var atCoderHistories []data.History

		// NOTE: atCoderHistoriesにJSONをパースしたバイト列を格納する
		if err := json.Unmarshal([]byte(info), &atCoderHistories); err != nil {
			fmt.Println(err)
		}

		u.logger.Info().Interface("history", atCoderHistories).Msg("")
		u.logger.Info().Interface("rate", atCoderHistories[len(atCoderHistories)-1].NewRating).Msg("")
		u.drawSVGUseCase.Draw(name, atCoderHistories[len(atCoderHistories)-1].NewRating, svg.New(writer))
	}

}
