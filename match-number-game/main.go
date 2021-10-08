package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Game struct {
	RandNum int
	TryCnt  int
}

type result int

const (
	small result = iota
	big
	equal
)

var g *Game

func GenerateRandNumber() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	return n
}

func init() {
	g = &Game{}
	g.RandNum = GenerateRandNumber()
}

func inputValue() (int, error) {
	var input int
	stdin := bufio.NewReader(os.Stdin)
	fmt.Print("숫자값을 입력하세요:")
	_, err := fmt.Scanln(&input)
	if err != nil {
		stdin.ReadString('\n')
	}
	return input, err
}

func (g *Game) CompareNumber(input int) result {
	if g.RandNum > input {
		return small
	} else if g.RandNum < input {
		return big
	} else {
		return equal
	}
}

func (g *Game) PlayGame() {

GameLoop:
	for {
		g.TryCnt++
		input, err := inputValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			r := g.CompareNumber(input)
			switch r {
			case small:
				fmt.Println("입력하신 숫자가 더 작습니다.")
			case big:
				fmt.Println("입력하신 숫자가 더 큽니다.")
			case equal:
				fmt.Printf("숫자를 맞췄습니다. 축하합니다. 시도한 횟수: %d", g.TryCnt)
				break GameLoop
			}
		}
	}
}

func main() {
	g.PlayGame()
}
