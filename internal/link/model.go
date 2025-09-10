package link

import (
	"http_server/internal/stat"
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraints:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}

	link.GenerateHash()

	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RandStringRunes(6)
}

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

func RandStringRunes(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes)-1)]
	}

	return string(b)
}
