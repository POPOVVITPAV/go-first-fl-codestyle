package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func attack(player, сharacter string) string {
	if сharacter == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", player, 5+randint(3, 5))
	}

	if сharacter == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", player, 5+randint(5, 10))
	}

	if сharacter == "healer" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", player, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func defence(player, сharacter string) string {
	if сharacter == "warrior" {
		return fmt.Sprintf("%s блокировал %d урона.", player, 10+randint(5, 10))
	} else if сharacter == "mage" {
		return fmt.Sprintf("%s блокировал %d урона.", player, 10+randint(-2, 2))
	} else if сharacter == "healer" {
		return fmt.Sprintf("%s блокировал %d урона.", player, 10+randint(2, 5))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func special(player, сharacter string) string {
	if сharacter == "warrior" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", player, 80+25)
	} else if сharacter == "mage" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", player, 5+40)
	} else if сharacter == "healer" {
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", player, 10+30)
	}
	return "неизвестный класс персонажа"
}

// здесь обратите внимание на имена параметров
func start_training(player, сharacter string) string {
	if сharacter == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", player)
	}

	if сharacter == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", player)
	}

	if сharacter == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", player)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		if cmd == "attack" {
			fmt.Println(attack(player, сharacter))
		}

		if cmd == "defence" {
			fmt.Println(defence(player, сharacter))
		}

		if cmd == "special" {
			fmt.Println(special(player, сharacter))
		}
	}

	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func choiseСharacter() string {
	var key string
	var сharacter string

	for key != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &сharacter)
		if сharacter == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if сharacter == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if сharacter == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &key)
		key = strings.ToLower(key)
	}
	return сharacter
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var player string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &player)

	fmt.Printf("Здравствуй, %s\n", player)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	character := choiseСharacter()

	fmt.Println(start_training(player, character))
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
