package stat

import "time"

type GetStatRequest struct {
	From time.Time
	To   time.Time
	By   string
}

type GetStatResponce struct {
	Period string `json:"period"`
	Sum    int    `json:"sum"`
}
