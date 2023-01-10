package main

import (
	"fmt"
	"math/rand"
	"time"
)

type mole struct {
	D, //Donbass
	Z, //Za pobedu
	U, //Ukrain
	V /* Sila V pravde*/ int
}

func (m *mole) isAlive() bool {
	return (m.D > 0) && (m.Z > 0) && (m.U > 0) && (m.V > 0)
}
func (m *mole) isRespectable() bool {
	return m.U > 100
}
func (m *mole) changeD(value int) {
	m.D = m.D + value
}
func (m *mole) changeZ(value int) {
	m.Z = m.Z + value
}
func (m *mole) changeU(value int) {
	m.U = m.U + value
}
func (m *mole) changeV(value int) {
	m.V = m.V + value
}
func Newmole() mole {
	return mole{10, 100, 20, 30}
}
func (m mole) String() string {
	return fmt.Sprintf("D=%d, Z=%d, U=%d, V=%d", m.D, m.Z, m.U, m.V)
}
func (m *mole) Start() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(m)
	m.day()
}
func (m *mole) day() {
	for {
		m.variants()
		if m.check() {
			return
		}
		m.night()
		fmt.Println("Ах эта ночь, её не смыть годам")
		if m.check() {
			return
		}
	}
}
func (m *mole) check() bool {
	fmt.Println(m)
	if !m.isAlive() {
		m.gameover()
		return true
	}
	if m.isRespectable() {
		win()
		return true
	}
	return false
}

func (m *mole) night() {
	m.changeD(-2)
	m.changeZ(20)
	m.changeU(-2)
	m.changeV(-5)
}
func (m *mole) variants() {
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
				m.dig(true)
				variant = -1
				break
			}
		case 3: //рыть как slave
			{
				m.dig(false)
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
				m.smokeweed(true)
				variant = -1
				break
			}
		case 6: //стрельнуть травки у гопоты
			{
				m.smokeweed(false)
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
				m.gopstop(0)
				variant = -1
				break
			}
		case 9: //щемануть терпилу
			{
				fmt.Println("Посмотрим, что ты сможешь сделать против обычного мужика")
				m.gopstop(1)
				variant = -1
				break
			}
		case 10: //щемануть авторитета
			{
				fmt.Println("Посмотрим, что ты сможешь сделать против ровного пацана с района")
				m.gopstop(2)
				variant = -1
				break
			}
		case 11: //спать в банке вверх ногами
			{
				fmt.Println("Весь день проспал в банке")
				m.daysleep()
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
func (m *mole) dig(master bool) {
	if master {
		fmt.Println("Ох, как глубоко расширил нору")
		m.changeD(5)
		m.changeZ(-30)
		return
	}
	fmt.Println("Нора, конечно, стала глубже, но можно было бы расширять и активнее")
	m.changeD(2)
	m.changeZ(-10)
}
func (m *mole) smokeweed(myweed bool) {
	if myweed {
		fmt.Println("Травка не подвела")
		m.changeZ(10)
		m.changeV(15)
		return
	}
	if m.U >= 30 {
		fmt.Println("План оказался хорошим, удался")
		m.changeZ(30)
		m.changeV(30)
		return
	}
	fmt.Println("Пацаны не только не дали тебе травки, но ещё и о****или")
	m.changeZ(-30)
}
func (m *mole) gopstop(partner int) {
	if partner == 0 {
		if m.mesivo(30) {
			fmt.Println("Ну такого лоха каждый завалил бы")
			m.changeU(10)
			m.changeZ(-5)
			return
		}
		fmt.Println("Да ты сам лох")
		m.changeU(-5)
		m.changeZ(-15)
		return
	}
	if partner == 1 {
		if m.mesivo(50) {
			fmt.Println("Крепко ты мужика этого приложил")
			m.changeU(20)
			m.changeZ(-10)
			return
		}
		fmt.Println("Не на того нарвался")
		m.changeU(-10)
		m.changeZ(-25)
		return
	}

	if m.mesivo(70) {
		fmt.Println("Теперь ты главный авторитет на районе")
		m.changeU(40)
		m.changeZ(-25)
		return
	}
	fmt.Println("Тебя сначала от***или, а потом о****или. А нечего рамсы путать")
	m.changeU(-15)
	m.changeZ(-35)
}
func (m *mole) mesivo(zhir int) bool {
	return ((float64(m.V) / float64(m.V+zhir)) >= rand.Float64())
}
func (m *mole) daysleep() {

	m.night()
}
func win() {
	fmt.Println("YOU WIN")
}
func (m *mole) gameover() {
	if m.D <= 0 {
		fmt.Println("Размер твоей дырки стал слишком мал")
	}
	if m.Z <= 0 {
		fmt.Println("С таким здоровьем придётся тебе надеть берёзовый бушлат")
	}
	if m.U <= 0 {
		fmt.Println("С таким уважением тебе дорога в петухи")
	}
	if m.V <= 0 {
		fmt.Println("А кто это по палате летает? А это тебе, дистрофану, вентиллятор включили")
	}
	fmt.Println("ПОТРАЧЕНО")
}
