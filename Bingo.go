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
		playerMove(&player1, board)

		// 檢查玩家1是否連成一條線
		if isWinner(player1.selected) {
			fmt.Println(player1.name, "Wins!")
			break
		}

		// 玩家2選擇號碼
		playerMove(&player2, board)

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
	numbers := rand.Perm(maxNumber) // 隨機洗牌0~24

	for i := 0; i < boardSize; i++ {
		board[i] = numbers[i*boardSize : (i+1)*boardSize]
		for j := range board[i] {
			board[i][j]++ // 將每個數字加1，以滿足1~25的範圍
		}
	}

	return board
}

/* 原 0~24
func initializeBoard() [][]int {
	board := make([][]int, boardSize)
	numbers := rand.Perm(maxNumber) // 隨機洗牌1~25

	for i := 0; i < boardSize; i++ {
		board[i] = numbers[i*boardSize : (i+1)*boardSize]
	}

	return board
} */

func printBoard(board [][]int) {
	for _, row := range board {
		for _, num := range row {
			fmt.Printf("%2d ", num)
		}
		fmt.Println()
	}
}

/*
選擇號碼不會修改的問題

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
*/
func playerMove(player *Player, board [][]int) {
	var selectedNumber int
	fmt.Print(player.name, ", choose a number: ")
	fmt.Scan(&selectedNumber)

	// 檢查號碼的有效性
	if selectedNumber < 1 || selectedNumber > maxNumber || player.selected[selectedNumber] {
		fmt.Println("Invalid choice. Try again.")
		playerMove(player, board)
		return
	}

	// 標記號碼為已選擇
	player.selected[selectedNumber] = true

	// 更新Bingo板
	for i := range board {
		for j := range board[i] {
			if board[i][j] == selectedNumber && selectedNumber != 0 {
				board[i][j] = 0
			}
		}
	}
}

func isWinner(selected map[int]bool) bool {
	// 檢查橫線
	for i := 0; i < boardSize; i++ {
		win := true
		for j := 0; j < boardSize; j++ {
			if !selected[i*boardSize+j] {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// 檢查豎線
	for i := 0; i < boardSize; i++ {
		win := true
		for j := 0; j < boardSize; j++ {
			if !selected[j*boardSize+i] {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// 檢查對角線
	diagonal1 := true
	diagonal2 := true
	for i := 0; i < boardSize; i++ {
		if !selected[i*boardSize+i] {
			diagonal1 = false
			break
		}
	}
	for i := 0; i < boardSize; i++ {
		if !selected[i*boardSize+(boardSize-i-1)] {
			diagonal2 = false
			break
		}
	}

	return diagonal1 || diagonal2
}
