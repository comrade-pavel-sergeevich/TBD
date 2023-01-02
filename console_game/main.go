package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var krot mole

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	start()
}

func start() {
	krot = Newmole()
	fmt.Println(krot)
	day()
}
func day() {
	for {
		variants()
		if check() {
			return
		}
		night()
		fmt.Println("Ах эта ночь, её не смыть годам")
		if check() {
			return
		}
	}
}
func check() bool {
	fmt.Println(krot)
	if !krot.isAlive() {
		gameover()
		return true
	}
	if krot.isRespectable() {
		win()
		return true
	}
	return false
}
func command() string {
	fmt.Println("Enter command:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("You entered:", input)
	return input
}
func night() {
	krot.changeD(-2)
	krot.changeZ(20)
	krot.changeU(-2)
	krot.changeV(-5)
}
func variants() {
	variant := 0
	flag := false
	for {
		flag = false
		switch variant {
		case 0: //день
			{
				fmt.Println(
					`0. Углублять свою шахту
1. Баловаться травкой на свежем воздухе
2. Взять кого-нибудь на гоп-стоп
3. Прыгнуть в банку и спать мордой вниз`)
				for {
					switch command() {
					case "0": //рыть шахту
						{
							variant = 1
							flag = true
							break
						}
					case "1": //баловаться травкой
						{
							variant = 4
							flag = true
							break
						}
					case "2": //спарринг
						{
							variant = 7
							flag = true
							break
						}
					case "3": //спать в банке
						{
							variant = 11
							flag = true
							break
						}
					default: //неправильный выбор
						{
							fmt.Println("Вводи правильно, криворукий дегенерат")
							break
						}
					}
					if flag {
						flag = false
						break
					}
				}
				break
			}
		case 1: //рыть шахту
			{
				fmt.Println(
					`0. Углубляться в шахту как master
1. To have your hole get bigger as a slave`)
				for {
					switch command() {
					case "0": //как master
						{
							variant = 2
							flag = true
							break
						}
					case "1": //как slave
						{
							variant = 3
							flag = true
							break
						}
					default:
						{
							fmt.Println("Вводи правильно, криворукий дегенерат")
							break
						}

					}
					if flag {
						flag = false
						break
					}
				}
			}
		case 2: //рыть как master
			{
				dig(true)
				variant = -1
				break
			}
		case 3: //рыть как slave
			{
				dig(false)
				variant = -1
				break
			}
		case 4: //баловаться травкой
			{
				fmt.Println(
					`0. Забить своей проверенной травки
1. Стрельнуть травки у гопоты`)
				for {
					switch command() {
					case "0": //забить своей проверенной травки
						{
							variant = 5
							flag = true
							break
						}
					case "1": //стрельнуть травки у гопоты
						{
							variant = 6
							flag = true
							break
						}
					default: //неправильный выбор
						{
							fmt.Println("Вводи правильно, криворукий дегенерат")
							break
						}
					}
					if flag {
						flag = false
						break
					}
				}
			}
		case 5: //забить своей проверенной травки
			{
				smokeweed(true)
				variant = -1
				break
			}
		case 6: //стрельнуть травки у гопоты
			{
				smokeweed(false)
				variant = -1
				break
			}
		case 7:
			{ //схватки
				fmt.Println(
					`0. Щемануть лохозавра
1. Щемануть терпилу
2. Щемануть авторитета`)
				for {
					switch command() {
					case "0": //щемануть лохозавра
						{
							variant = 8
							flag = true
							break
						}
					case "1": //щемануть терпилу
						{
							variant = 9
							flag = true
							break
						}
					case "2": //щемануть авторитета
						{
							variant = 10
							flag = true
							break
						}

					default: //неправильный выбор
						{
							fmt.Println("Вводи правильно, криворукий дегенерат")
							break
						}
					}
					if flag {
						flag = false
						break
					}
				}
			}
		case 8: //щемануть лохозавра
			{
				fmt.Println("Посмотрим, что ты сможешь сделать против такого лоха")
				gopstop(0)
				variant = -1
				break
			}
		case 9: //щемануть терпилу
			{
				fmt.Println("Посмотрим, что ты сможешь сделать против обычного мужика")
				gopstop(1)
				variant = -1
				break
			}
		case 10: //щемануть авторитета
			{
				fmt.Println("Посмотрим, что ты сможешь сделать против ровного пацана с района")
				gopstop(2)
				variant = -1
				break
			}
		case 11: //спать в банке вверх ногами
			{
				fmt.Println("Весь день проспал в банке")
				daysleep()
				variant = -1
				break
			}
		case -1:
			{
				return
			}
		}
	}
}
func dig(master bool) {
	if master {
		fmt.Println("Ох, как глубоко расширил нору")
		krot.changeD(5)
		krot.changeZ(-30)
		return
	}
	fmt.Println("Нора, конечно, стала глубже, но можно было бы расширять и активнее")
	krot.changeD(2)
	krot.changeZ(-10)
}
func smokeweed(myweed bool) {
	if myweed {
		fmt.Println("Травка не подвела")
		krot.changeZ(10)
		krot.changeV(15)
		return
	}
	if krot.U >= 30 {
		fmt.Println("План оказался хорошим, удался")
		krot.changeZ(30)
		krot.changeV(30)
		return
	}
	fmt.Println("Пацаны не только не дали тебе травки, но ещё и о****или")
	krot.changeZ(-30)
}
func gopstop(partner int) {
	if partner == 0 {
		if mesivo(30) {
			fmt.Println("Ну такого лоха каждый завалил бы")
			krot.changeU(10)
			krot.changeZ(-5)
			return
		}
		fmt.Println("Да ты сам лох")
		krot.changeU(-5)
		krot.changeZ(-15)
		return
	}
	if partner == 1 {
		if mesivo(50) {
			fmt.Println("Крепко ты мужика этого приложил")
			krot.changeU(20)
			krot.changeZ(-10)
			return
		}
		fmt.Println("Не на того нарвался")
		krot.changeU(-10)
		krot.changeZ(-25)
		return
	}

	if mesivo(70) {
		fmt.Println("Теперь ты главный авторитет на районе")
		krot.changeU(40)
		krot.changeZ(-25)
		return
	}
	fmt.Println("Тебя сначала от***или, а потом о****или. А нечего рамсы путать")
	krot.changeU(-15)
	krot.changeZ(-35)
}
func mesivo(zhir int) bool {
	return ((float64(krot.V) / float64(krot.V+zhir)) >= rand.Float64())
}
func daysleep() {

	night()
}
func win() {
	fmt.Println("YOU WIN")
}
func gameover() {
	if krot.D <= 0 {
		fmt.Println("Размер твоей дырки стал слишком мал")
	}
	if krot.Z <= 0 {
		fmt.Println("С таким здоровьем придётся тебе надеть берёзовый бушлат")
	}
	if krot.U <= 0 {
		fmt.Println("С таким уважением тебе дорога в петухи")
	}
	if krot.V <= 0 {
		fmt.Println("А кто это по палате летает? А это тебе, дистрофану, вентиллятор включили")
	}
	fmt.Println("ПОТРАЧЕНО")
}
