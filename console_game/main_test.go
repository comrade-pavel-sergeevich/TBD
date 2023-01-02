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
	dig(true)
	dig(false)
}
func TestSmokeweed(t *testing.T) {
	smokeweed(true)
	smokeweed(false)
}
func TestShvatka(t *testing.T) {
	gopstop(0)
	gopstop(1)
	gopstop(2)
}
func TestMoleChanges(t *testing.T) {
	krot.changeD(1)
	krot.changeZ(1)
	krot.changeU(1)
	krot.changeV(1)
}
func TestSleep(t *testing.T) {
	night()
	daysleep()
}
func TestCheck(t *testing.T) {
	krot := mole{0, 0, 0, 0}
	check()
	require.Equal(t, false, krot.isAlive())
}
func TestRespect(t *testing.T) {
	krot := mole{0, 0, 0, 0}
	require.Equal(t, false, krot.isRespectable())
}
