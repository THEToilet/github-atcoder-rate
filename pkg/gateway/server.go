package gateway

import (
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)

func NewServer(logger *zerolog.Logger, userHistoryHandler *UserHistoryHandler) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/svg", userHistoryHandler.drawRating)
	logger.Info().Msg(os.Getenv("PORT"))
	server := &http.Server{
		// NOTE: ここ変えるならクライアントも変えなければならない
		// NOTE: 127.0.0.1 では繋がらないが、localhostは繋がる
		// NOTE: Heroku用に変更
		Addr:           os.Getenv("PORT"),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server
}
