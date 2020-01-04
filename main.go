package main

import (
	"fmt"
)

// BjPlayer is player object.
type BjPlayer struct {
	name  string
	score int
}

func main() {
	cs := NewCardSet()

	fmt.Println("ゲームをはじめます.")
	blackjack(cs)
}

// ブラックジャックを遊ぶ
func blackjack(cs CardSet) {
	var you = BjPlayer{name: "あなた", score: 0}
	you.draw(&cs)
	you.draw(&cs)
	fmt.Printf("%sの現在のスコアは%dです.\n", you.name, you.score)

	var dealer = BjPlayer{name: "ディーラー", score: 0}
	dealer.draw(&cs)
	dealer.draw(&cs)
	fmt.Printf("%sの現在のスコアは%dです.\n", dealer.name, dealer.score)

	// プレイヤーターン.
	for {
		fmt.Println("カードを引きますか？(y/n)")

		var s string
		fmt.Scanf("%s", &s)
		if s != "y" {
			break
		}

		b := you.draw(&cs)
		fmt.Printf("%sの現在のスコアは%dです.\n", you.name, you.score)

		if !b {
			fmt.Println("あなたの負けです.")
			return
		}
	}

	// ディーラーターン
	for dealer.score < 17 {
		dealer.draw(&cs)
	}

	// 勝敗判定
	fmt.Printf("あなたのスコアは%dです\n", you.score)
	fmt.Printf("ディーラーのスコアは%dです\n", dealer.score)

	if dealer.score > 21 {
		fmt.Printf("%sの勝ちです\n", you.name)
	} else if you.score == dealer.score {
		fmt.Println("ドローです")
	} else if you.score > dealer.score {
		fmt.Printf("%sの勝ちです\n", you.name)
	} else {
		fmt.Printf("%sの負けです\n", you.name)
	}
}

// プレイヤーがカードを引く
// 引いたカードの合計値を計算し、合計が21より大きければfalseを返す
// J,Q,Kは10として計算する
func (p *BjPlayer) draw(cs *CardSet) bool {
	c := cs.DrawCard()
	fmt.Printf("%sの引いたカードは%sの%dです.\n", p.name, Mark(c.kind), c.num)
	if c.num > 10 {
		c.num = 10
	}
	p.score += c.num

	if p.score > 21 {
		return false
	} else {
		return true
	}
}
