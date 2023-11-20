package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	boardSize = 5
	maxNumber = 25
)

type Player struct {
	name     string
	selected map[int]bool
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 初始化玩家
	player1 := Player{name: "Player 1", selected: make(map[int]bool)}
	player2 := Player{name: "Player 2", selected: make(map[int]bool)}

	// 初始化Bingo板
	board := initializeBoard()

	// 遊戲開始
	for !isWinner(player1.selected) && !isWinner(player2.selected) {
		fmt.Println("---- Current Board ----")
		printBoard(board)
		fmt.Println("-----------------------")

		// 玩家1選擇號碼
		playerMove(&player1)

		// 檢查玩家1是否連成一條線
		if isWinner(player1.selected) {
			fmt.Println(player1.name, "Wins!")
			break
		}

		// 玩家2選擇號碼
		playerMove(&player2)

		// 檢查玩家2是否連成一條線
		if isWinner(player2.selected) {
			fmt.Println(player2.name, "Wins!")
			break
		}
	}

	// 最終顯示結果
	fmt.Println("---- Final Board ----")
	printBoard(board)
	fmt.Println("----------------------")
}

func initializeBoard() [][]int {
	board := make([][]int, boardSize)
	numbers := rand.Perm(maxNumber) // 隨機洗牌1~25

	for i := 0; i < boardSize; i++ {
		board[i] = numbers[i*boardSize : (i+1)*boardSize]
	}

	return board
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, num := range row {
			fmt.Printf("%2d ", num)
		}
		fmt.Println()
	}
}

func playerMove(player *Player) {
	var selectedNumber int
	fmt.Print(player.name, ", choose a number: ")
	fmt.Scan(&selectedNumber)

	// 檢查號碼的有效性
	if selectedNumber < 1 || selectedNumber > maxNumber || player.selected[selectedNumber] {
		fmt.Println("Invalid choice. Try again.")
		playerMove(player)
		return
	}

	// 標記號碼為已選擇
	player.selected[selectedNumber] = true
}

func isWinner(selected map[int]bool) bool {
	// 檢查行
	for i := 0; i < boardSize; i++ {
		if selected[i*boardSize] && selected[i*boardSize+1] && selected[i*boardSize+2] && selected[i*boardSize+3] && selected[i*boardSize+4] {
			return true
		}
	}

	// 檢查列
	for i := 0; i < boardSize; i++ {
		if selected[i] && selected[i+boardSize] && selected[i+2*boardSize] && selected[i+3*boardSize] && selected[i+4*boardSize] {
			return true
		}
	}

	// 檢查對角線
	if selected[0] && selected[6] && selected[12] && selected[18] && selected[24] {
		return true
	}
	if selected[4] && selected[8] && selected[12] && selected[16] && selected[20] {
		return true
	}

	return false
}
