package stat

import (
	"advpractice/pkg/res"
	"net/http"
	"time"
)

const (
	ErrByValue     = `"By" CAN BE ONLY "day" OR "month"`
	ErrQueryParams = "INVALID QUERY PARAMETERS"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps *StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.HandleFunc("GET /stat", handler.getStat())
}

func (handler *StatHandler) getStat() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		timeFrom, err1 := time.Parse("2006-01-02", q.URL.Query().Get("from"))
		timeTo, err2 := time.Parse("2006-01-02", q.URL.Query().Get("to"))
		by := q.URL.Query().Get("by")

		if by != GroupByDay && by != GroupByMonth {
			res.JsonDump(w, ErrByValue, http.StatusBadRequest)
			return
		}
		if err1 != nil || err2 != nil {
			res.JsonDump(w, ErrQueryParams, http.StatusBadRequest)
			return
		}

		stats := handler.StatRepository.GetStat(GetStatRequest{
			From: timeFrom,
			To:   timeTo,
			By:   by,
		})
		res.JsonDump(w, stats, http.StatusOK)
	}
}
