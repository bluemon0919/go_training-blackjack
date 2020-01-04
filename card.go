package main

import (
	"math/rand"
	"time"
)

const (
	spade = iota + 1
	club
	heart
	diamond
)

// トランプカードオブジェクト
type Card struct {
	num  int // カードナンバー(1〜13)
	kind int // 種類(スペード、クラブ、ハート、ダイヤ)
}

// トランプカードのセット
type CardSet struct {
	array [52]int
	cnt   int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewCardSet prepares shuffled card.
func NewCardSet() CardSet {
	var cs CardSet
	for i := 0; i < 52; i++ {
		cs.array[i] = i + 1
	}

	// Fisher–Yates shuffle
	for i := 52 - 1; i > 0; i-- {
		j := rand.Intn(52 - 1)
		cs.array[j], cs.array[i] = cs.array[i], cs.array[j]
	}
	return cs
}

// Mark gets the picture of the card.
func Mark(kind int) string {
	var mark string
	switch kind {
	case spade:
		mark = "スペード"
	case club:
		mark = "クラブ"
	case heart:
		mark = "ハート"
	case diamond:
		mark = "ダイヤ"
	}
	return mark
}

// DrawCard draw card from deck.
func (cs *CardSet) DrawCard() Card {
	var c Card
	i := (*cs).cnt
	switch {
	case (*cs).array[i] <= 13*spade:
		c.num = (*cs).array[i]
		c.kind = spade
	case (*cs).array[i] <= 13*club:
		c.num = (*cs).array[i] - (club-1)*13
		c.kind = club
	case (*cs).array[i] <= 13*heart:
		c.num = (*cs).array[i] - (heart-1)*13
		c.kind = heart
	case (*cs).array[i] <= 13*diamond:
		c.num = (*cs).array[i] - (diamond-1)*13
		c.kind = diamond
	}
	(*cs).cnt++
	return c
}
