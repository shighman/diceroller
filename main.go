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
	if len(os.Args) < 2 {
		fmt.Println("Must enter a dice role such as 3d6, 4d6-l, 2d6+6 or just stats")
		return
	}
	args := os.Args[1]
	roller := complexDiceRoller()
	if args == "stats" {
		fmt.Println("Rolling AD&D 1E using 4D6 drop lowest")
		fmt.Println("STAT            VALUE")
		fmt.Println("----            ----")
		fmt.Println("STR:           ", roller(6, 4, "-", "l"))
		fmt.Println("DEX:           ", roller(6, 4, "-", "l"))
		fmt.Println("INT:           ", roller(6, 4, "-", "l"))
		fmt.Println("WIS:           ", roller(6, 4, "-", "l"))
		fmt.Println("CON:           ", roller(6, 4, "-", "l"))
		fmt.Println("CHR:           ", roller(6, 4, "-", "l"))
		fmt.Println("CMS:           ", roller(6, 4, "-", "l"))
		fmt.Println("Psionics roll: ", roller(100, 1, "", ""))
		fmt.Println()
	} else {
		reg := regexp.MustCompile(`([\d]*)d([\d]+)([\+\-]*)([\dlh]*)`)
		results := reg.FindAllStringSubmatch(args, -1)
		qty, _ := strconv.Atoi(results[0][1])
		sides, _ := strconv.Atoi(results[0][2])
		total := roller(sides, qty, results[0][3], results[0][4])
		fmt.Println("Total for", results[0][0], "=", total)
	}
}

func getModifier(mod string, min, max int) int {
	modifier, err := strconv.Atoi(mod)
	if err != nil {
		if mod == "l" {
			modifier = min
		} else if mod == "h" {
			modifier = max
		} else {
			modifier = 0
		}
	}
	return modifier
}

func rollAndTrackMinAndMax(gen *rand.Rand, sides int, qty int) (total, min, max int) {
	min = sides
	for i := 0; i < qty; i++ {
		roll := gen.Intn(sides) + 1
		if min > roll {
			min = roll
		}
		if max < roll {
			max = roll
		}
		total += roll
	}
	return
}

func complexDiceRoller() func(sides int, qty int, mod string, value string) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return func(sides int, qty int, mod string, value string) int {
		total, min, max := rollAndTrackMinAndMax(r1, sides, qty)
		modifier := getModifier(value, min, max)
		if mod == "+" {
			total += modifier
		} else if mod == "-" {
			total -= modifier
		}
		return total
	}
}
