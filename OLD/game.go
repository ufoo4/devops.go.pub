package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	fmt.Println("\n")
	fmt.Println("ИГРА - УГАДАЙ ЧИСЛО")
	fmt.Println("\n")
	fmt.Println("Выбираю число от 1 до 100")
	fmt.Println("Число загадано!")

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses < 10; guesses++ {

		fmt.Println("У вас осталось ", 10-guesses, " попыток")
		fmt.Println("Введите ваше число: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess > target {
			fmt.Println("Ваше значение больше, чем загаданное ")
		} else if guess < target {
			fmt.Println("Ваше значение меньше, чем загаданное ")
		} else {
			success = true
			fmt.Println("Поздравляю, вы угадали загаданное число! ")
			break
		}
	}

	if !success {
		fmt.Println("Игра окончена ! Вы проиграли! Ха-ха =)")
		fmt.Println("Загаданное число было:", target)
	}
}
