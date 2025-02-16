package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler o nome")
	}

	g.Name = strings.TrimSpace(name)

	fmt.Printf("Vamos ao jogo, %s!\n", g.Name)

}

func (g *GameState) ProcessCSV() {
	f, err := os.Open("questions.csv")
	if err != nil {
		panic("Erro ao ler o arquivo")
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler o csv")
	}

	for index, record := range records {
		if index > 0 {
			correctAnswer, _ := toInt(record[5])
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *GameState) Run() {
	for index, question := range g.Questions {

		fmt.Printf("\033[94m %d. %s \033[0m\n", index+1, question.Text)
		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Println("Escolha uma opção:")

		var answer int
		var err error

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')

			answer, err = toInt(read[:len(read)-1])

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break

		}
		if answer == question.Answer {
			fmt.Printf("\033[92mResposta correta!\033[0m\n")
			g.Points += 10
		} else {
			fmt.Printf("\033[91mResposta incorreta!\033[0m\n")
			fmt.Println("----------------------------------")
		}

	}
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("Não é permitido caractere diferente de número, por favor insira um número")
	}
	return i, nil
}

func main() {
	game := &GameState{}
	game.ProcessCSV()
	game.Init()
	game.Run()
	fmt.Printf("\033[94mFim de jogo, você fez %d pontos!\033[0m\n", game.Points)
	if game.Points >= 30 {
		fmt.Println("\033[92mParabéns, você foi aprovado!\033[0m")
	} else {
		fmt.Println("\033[91mVocê foi reprovado. Tente novamente!\033[0m")
	}
}
