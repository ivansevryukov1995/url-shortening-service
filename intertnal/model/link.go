package model

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	URL   string `json:"url"`
	Hash  string `json:"hash" gorm:"uniqueIndex"`
	Stats []Stat `gorm:"constraints:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		URL:  url,
		Hash: RandStringRunes(6),
	}
	link.GeneratedHash()
	return link
}

func (link *Link) GeneratedHash() {
	link.Hash = RandStringRunes(6)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
