package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成Bingo卡片
func generateBingoCard() [][]int {
	numbers := rand.Perm(25) // 生成0到24的隨機排列數字

	card := make([][]int, 5)
	for i := 0; i < 5; i++ {
		card[i] = numbers[i*5 : (i+1)*5]
	}

	return card
}

// 顯示Bingo卡片
func printBingoCard(card [][]int) {
	for _, row := range card {
		for _, number := range row {
			fmt.Printf("%2d | ", number)
		}
		fmt.Println("\n-------------------------")
	}
}

// 檢查獲勝
func checkBingo(card [][]int) bool {
	for _, row := range card {
		if allZeros(row) {
			return true
		}
	}

	for col := 0; col < 5; col++ {
		if allZeros(getColumn(card, col)) {
			return true
		}
	}

	if allZeros(getDiagonal(card, true)) || allZeros(getDiagonal(card, false)) {
		return true
	}

	return false
}

// 選出的牌換圈
func checkChoose(player1Card, player2Card [][]int, chooseNum int) {
	checkChooseV2(player1Card, chooseNum)
	checkChooseV2(player2Card, chooseNum)
}

func checkChooseV2(playerCard [][]int, chooseNum int) {
	for _, row := range playerCard {
		for i, number := range row {
			if number == chooseNum {
				row[i] = 0
			}
		}
	}
}

// 主循環
func playBingo() {
	rand.Seed(time.Now().UnixNano())

	canPlayCard := make([]int, 25)
	for i := range canPlayCard {
		canPlayCard[i] = i + 1
	}

	player1Card := generateBingoCard()
	player2Card := generateBingoCard()

	fmt.Println("歡迎來玩Bingo遊戲！")
	fmt.Println("Player 1 卡片：")
	printBingoCard(player1Card)
	fmt.Println("\nPlayer 2 卡片：")
	printBingoCard(player2Card)

	for {
		fmt.Println("\n------------------------------\n")
		// Player 1 出牌
		fmt.Print("Player 1 叫號碼: ")
		var player1CalledNumber int
		fmt.Scan(&player1CalledNumber)
		canPlayCard = removeElement(canPlayCard, player1CalledNumber)
		checkChoose(player1Card, player2Card, player1CalledNumber)

		fmt.Println("Player 1 的Bingo卡片：")
		printBingoCard(player1Card)
		fmt.Println("\nPlayer 2 的Bingo卡片：")
		printBingoCard(player2Card)

		if checkBingo(player1Card) {
			fmt.Println("恭喜！Player 1 獲勝了！")
			break
		}

		if checkBingo(player2Card) {
			fmt.Println("Player 2 獲勝！")
			break
		}

		// Player 2 出牌
		fmt.Println("\n------------------------------\n")
		fmt.Print("Player 2 叫號碼: ")
		var player2CalledNumber int
		fmt.Scan(&player2CalledNumber)
		canPlayCard = removeElement(canPlayCard, player2CalledNumber)
		checkChoose(player1Card, player2Card, player2CalledNumber)

		fmt.Println("Player 2 的Bingo卡片：")
		printBingoCard(player2Card)
		fmt.Println("\nPlayer 1 的Bingo卡片：")
		printBingoCard(player1Card)

		if checkBingo(player1Card) {
			fmt.Println("恭喜！Player 1 獲勝了！")
			break
		}

		if checkBingo(player2Card) {
			fmt.Println("Player 2 獲勝！")
			break
		}
	}
}

// 輔助函數，檢查切片中的元素是否全為零
func allZeros(slice []int) bool {
	for _, num := range slice {
		if num != 0 {
			return false
		}
	}
	return true
}

// 輔助函數，獲取二維切片的列
func getColumn(card [][]int, col int) []int {
	column := make([]int, 5)
	for i := 0; i < 5; i++ {
		column[i] = card[i][col]
	}
	return column
}

// 輔助函數，獲取二維切片的對角線
func getDiagonal(card [][]int, mainDiagonal bool) []int {
	diagonal := make([]int, 5)
	for i := 0; i < 5; i++ {
		if mainDiagonal {
			diagonal[i] = card[i][i]
		} else {
			diagonal[i] = card[i][4-i]
		}
	}
	return diagonal
}

// 輔助函數，移除切片中的元素
func removeElement(slice []int, elem int) []int {
	index := -1
	for i, val := range slice {
		if val == elem {
			index = i
			break
		}
	}

	if index != -1 {
		slice = append(slice[:index], slice[index+1:]...)
	}

	return slice
}

func main() {
	playBingo()
}
