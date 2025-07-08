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

func NewLink(url_str string, linkRepo *LinkRepository) *Link {
	return &Link{
		Url:  url_str,
		Hash: generateUniqueHash(linkRepo),
	}
}

func generateUniqueHash(linkRepo *LinkRepository) string {
	hash := getRandomHash(5)
	for {
		link, _ := linkRepo.GetByHash(hash)
		if link == nil {
			break
		}
		hash = getRandomHash(5)
	}
	return hash
}

func getRandomHash(n int) string {
	res := make([]rune, n)
	for i := range n {
		alphabetPos := rand.Int32N(26)
		twoChars := [2]rune{alphabetPos + 65, alphabetPos + 97}
		res[i] = twoChars[rand.IntN(2)]
	}
	return string(res)
}
