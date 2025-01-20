package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

const (
	width  = 40
	height = 10
)

var (
	dinoPos   int
	obstacles []int
	score     int
	gameOver  bool
)

func game() {
	rand.Seed(time.Now().UnixNano())
	initGame()
	go processInput()

	for !gameOver {
		updateGame()
		drawGame()
		time.Sleep(150 * time.Millisecond)
	}

	fmt.Printf("\nGame Over! Final Score: %d\n", score)
}

func initGame() {
	dinoPos = height - 2
	obstacles = []int{}
	score = 0
	gameOver = false
}

func updateGame() {
	// 移动障碍物
	for i := range obstacles {
		obstacles[i]--
	}

	// 移除屏幕外的障碍物
	if len(obstacles) > 0 && obstacles[0] < 0 {
		obstacles = obstacles[1:]
		score++
	}

	// 随机生成新障碍物
	if rand.Intn(10) < 2 {
		obstacles = append(obstacles, width-1)
	}

	// 检测碰撞
	for _, pos := range obstacles {
		if pos == 1 && dinoPos >= height-2 {
			gameOver = true
			return
		}
	}

	// 重力
	if dinoPos < height-2 {
		dinoPos++
	}
}

func drawGame() {
	// 使用字符串构建器来减少屏幕刷新次数
	var screen strings.Builder

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y == height-1 {
				screen.WriteString("_")
				continue
			}

			if y == dinoPos && x == 1 {
				screen.WriteString("R")
				continue
			}

			isObstacle := false
			for _, pos := range obstacles {
				if pos == x && y >= height-2 {
					isObstacle = true
					break
				}
			}

			if isObstacle {
				screen.WriteString("X")
			} else {
				screen.WriteString(" ")
			}
		}
		screen.WriteString("\n")
	}
	screen.WriteString(fmt.Sprintf("Score: %d\n", score))

	clearScreen()
	fmt.Print(screen.String())
}

func clearScreen() {
	// 跨平台清屏
	fmt.Print("\033[H\033[2J") // Linux/Mac
	fmt.Print("\033[2J\033[H") // Windows
}

func processInput() {
	// 设置终端为原始模式
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	reader := bufio.NewReader(os.Stdin)
	for {
		// 设置读取超时
		char, _, err := reader.ReadRune()
		if err != nil {
			continue
		}

		// 处理跳跃
		if (char == ' ' || char == 'w' || char == 'W') && dinoPos >= height-3 {
			dinoPos -= 2
		}

		// 处理退出
		if char == 'q' || char == 'Q' {
			gameOver = true
			return
		}
	}
}
