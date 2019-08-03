package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Must enter a dice role such as 3d6 or 4d6-l or 2d6+6")
		return
	}
	args := os.Args[1]
	if args == "stats" {
		fmt.Println("Rolling AD&D 1E using 4D6 drop lowest")
		fmt.Println("STR: ", complexDiceRoller(6, 4, "-", "l"))
		fmt.Println("DEX: ", complexDiceRoller(6, 4, "-", "l"))
		fmt.Println("INT: ", complexDiceRoller(6, 4, "-", "l"))
		fmt.Println("WIS: ", complexDiceRoller(6, 4, "-", "l"))
		fmt.Println("CON: ", complexDiceRoller(6, 4, "-", "l"))
		fmt.Println("CHR: ", complexDiceRoller(6, 4, "-", "l"))
		fmt.Println("CMS: ", complexDiceRoller(6, 4, "-", "l"))
		fmt.Println("Psionics roll: ", complexDiceRoller(100, 1, "", ""))
	} else {
		reg := regexp.MustCompile(`([\d]*)d([\d]+)([\+\-]*)([\dlh]*)`)
		results := reg.FindAllStringSubmatch(args, -1)
		qty, _ := strconv.Atoi(results[0][1])
		sides, _ := strconv.Atoi(results[0][2])
		total := complexDiceRoller(sides, qty, results[0][3], results[0][4])
		fmt.Println("Total for", results[0][0], "=", total)
	}
}

func complexDiceRoller(sides int, qty int, mod string, value string) int {
	total := 0
	min := sides
	max := 0
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < qty; i++ {
		roll := r1.Intn(sides) + 1
		if min > roll {
			min = roll
		}
		if max < roll {
			max = roll
		}
		total += roll
	}
	modifier, err := strconv.Atoi(value)
	if err != nil {
		if value == "l" {
			modifier = min
		} else if value == "h" {
			modifier = max
		} else {
			modifier = 0
		}
	}
	if mod == "+" {
		total += modifier
	} else if mod == "-" {
		total -= modifier
	}
	return total
}
