package main

import "fmt"

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Name      string
	Points    int
	Questions []Question
}

func (g *GameState) Init() {
	fmt.Println("Seja bem vindo(o) ao quiz!")
	fmt.Println("Qual é o seu nome?")
}

func main() {
	game := &GameState{}
	game.Init()
}
