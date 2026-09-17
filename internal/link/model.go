package link

import (
	"math/rand"

	"github.com/AndroDeMohawk/link-cutter/internal/stat"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `json:"stats" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash()
	return link
}

func (l *Link) GenerateHash() string {
	l.Hash = RandStringRunes(10)
	return l.Hash
}

var letterRunes = []rune("abcdefghijklmnoprstuvqwxyzABCDEFGHIJKLMNOPRSTUVQWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
