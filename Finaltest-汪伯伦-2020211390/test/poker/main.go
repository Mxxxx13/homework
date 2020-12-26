package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const (
	//黑桃
	Spade = 0
	//红桃
	Hearts = 1
	//梅花
	Club = 2
	//方块
	Diamond = 3
)

type Poker struct {
	Num int
	Flower int
}

func (p Poker)PokerSelf()string  {
	var buffer string

	switch p.Flower {
	case Spade:
		buffer += "♤"
	case Hearts:
		buffer += "♡"
	case Club:
		buffer += "♧"
	case Diamond:
		buffer += "♢"
	}
	switch p.Num {
	case 13:
		buffer += "2"
	case 12:
		buffer += "A"
	case 11:
		buffer += "K"
	case 10:
		buffer += "Q"
	case 9:
		buffer += "J"
	default:
		buffer += strconv.Itoa(p.Num+2)
	}

	return buffer
}

func CreatePokers()(pokers Pokers)  {
	for i := 1; i < 14; i++ {
		for j := 0; j < 4; j++ {
			pokers = append(pokers, Poker{
				Num:    i,
				Flower: j,
			})
		}
	}
	return
}

type Pokers []Poker

func (p Poker) PrintPoker() string {
	return p.PokerSelf()
}

func (p Pokers)Print()  {
	for _, i2 := range p {
		fmt.Print(i2.PrintPoker()," ")
	}
	fmt.Println()
}

//总共52张牌，随机取0~52之间的数，交换牌的位置
func (p Pokers)Shuffle() Pokers{
	rand.Seed(time.Now().UnixNano())
	RandNum1 := rand.Intn(53) //[0,53)
	time.Sleep(1*time.Nanosecond)//更换随机数种子
	RandNum2 := rand.Intn(53) //[0,53)
	p[RandNum1], p[RandNum2] = p[RandNum2], p[RandNum1]
	return p
}

func (p Pokers) Len() int {
	return len(p)
}
func (p Pokers) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p Pokers) Less(i, j int) bool {
	return p[i].Num < p[j].Num
}

func main(){
	var i int
	var p Pokers
	for i =0 ;i< 1000;i++{
		p = p.Shuffle()
	}
	sort.Stable(p)
}
