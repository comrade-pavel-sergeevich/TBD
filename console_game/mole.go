package main

import "fmt"

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
