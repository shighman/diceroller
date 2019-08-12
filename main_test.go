package main

import "testing"

func TestPadLeft2Left(t *testing.T) {
	val := padLeft(10, 2)
	if val != "10" {
		t.Errorf("%d padded left by %d should give value %s but has given %s", 10, 2, "10", val)
	}
}

func TestPadLeft3Left(t *testing.T) {
	val := padLeft(10, 3)
	if val != " 10" {
		t.Errorf("%d padded left by %d should give value %s but has given %s", 10, 2, " 10", val)
	}
}

func TestPadLeft1Left(t *testing.T) {
	val := padLeft(10, 1)
	if val != "10" {
		t.Errorf("%d padded left by %d should give value %s but has given %s", 10, 1, "10", val)
	}
}

func TestGetModifierNum(t *testing.T) {
	val := getModifier("6", 1, 10)
	if val != 6 {
		t.Errorf("getModifier returned %d and should return %d", val, 6)
	}
}

func TestGetModifierLow(t *testing.T) {
	val := getModifier("l", 1, 10)
	if val != 1 {
		t.Errorf("getModifier returned %d and should return %d", val, 1)
	}
}

func TestGetModifierHigh(t *testing.T) {
	val := getModifier("h", 1, 10)
	if val != 10 {
		t.Errorf("getModifier returned %d and should return %d", val, 10)
	}
}
