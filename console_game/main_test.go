package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	krot := Newmole()
	require.True(t, "D=10, Z=100, U=20, V=30" == krot.String())
}
func TestDig(t *testing.T) {
	krot := Newmole()
	krot.dig(true)
	krot.dig(false)
}
func TestSmokeweed(t *testing.T) {
	krot := Newmole()
	krot.smokeweed(true)
	krot.smokeweed(false)
	krot = mole{100, 100, 100, 100}
	krot.smokeweed(false)
}
func TestShvatka(t *testing.T) {
	krot := mole{1, 1, 1, 1}
	krot.gopstop(0)
	krot.gopstop(1)
	krot.gopstop(2)
	krot = mole{1000, 1000, 1000, 1000}
	krot.gopstop(0)
	krot.gopstop(1)
	krot.gopstop(2)
}
func TestMoleChanges(t *testing.T) {
	krot.changeD(1)
	krot.changeZ(1)
	krot.changeU(1)
	krot.changeV(1)
}
func TestSleep(t *testing.T) {
	krot := Newmole()
	krot.night()
	krot.daysleep()
}
func TestCheck(t *testing.T) {
	krot := mole{0, 0, 0, 0}
	krot.check()
	require.Equal(t, false, krot.isAlive())
	krot = mole{10, 10, 10, 10}
	krot.check()
	require.Equal(t, true, krot.isAlive())
	krot = mole{1000, 1000, 1000, 1000}
	krot.check()
	require.Equal(t, true, krot.isAlive())
}
func TestRespect(t *testing.T) {
	krot := mole{0, 0, 0, 0}
	require.Equal(t, false, krot.isRespectable())
	krot = mole{1000, 1000, 1000, 1000}
	require.Equal(t, true, krot.isRespectable())
}
