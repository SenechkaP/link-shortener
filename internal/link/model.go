package link

import (
	"math/rand/v2"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url_str string) *Link {
	return &Link{
		Url:  url_str,
		Hash: getRandomHash(10),
	}
}

func getRandomHash(n int) string {
	res := make([]rune, n)
	for i := range 10 {
		alphabetPos := rand.Int32N(26)
		twoChars := [2]rune{alphabetPos + 65, alphabetPos + 97}
		res[i] = twoChars[rand.IntN(2)]
	}
	return string(res)
}
