package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Сколько вам лет? ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	old, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}

	resault := ""

	if old <= 18 {
		resault = "Чо правда меньше 18? Не бзди!"
	} else if old >= 19 && old <= 30 {
		resault = "От 19 до 30 - самолет. Красавчег"
	} else if old >= 31 && old <= 49 {
		resault = "Старпер! Больше 30! Уходи..."
	} else {
		resault = "Хуеть, ты старый.."
	}

	fmt.Println(resault)
}
