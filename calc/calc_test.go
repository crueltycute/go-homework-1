package main

import (
	"bufio"
	"bytes"
	"testing"
)


// successful
func TestBigNumbers(t *testing.T) {
	expected := 123456 + 123456
	result, _ := calculate(bufio.NewScanner(bytes.NewBufferString("123456 123456 + =")))
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestOneAction(t *testing.T) {
	expected := 2 + 3
	result, _ := calculate(bufio.NewScanner(bytes.NewBufferString("2 3 + =")))
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestTwoActions(t *testing.T) {
	expected := 2 + 3 - 4
	result, _ := calculate(bufio.NewScanner(bytes.NewBufferString("2 3 + 4 -")))
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestThreeActions(t *testing.T) {
	expected := (2 * 3) + (4 * 5)
	result, _ := calculate(bufio.NewScanner(bytes.NewBufferString("2 3 * 4 5 * +")))
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestFourActions(t *testing.T) {
	expected := 2 / (3 - (4 + (5 * 6)))
	result, _ := calculate(bufio.NewScanner(bytes.NewBufferString("2 3 4 5 6 * + - /")))
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestNewLines(t *testing.T) {
	expected := (1 + 2) * (3 + 4)
	result, _ := calculate(bufio.NewScanner(bytes.NewBufferString("1\n2\n+\n3\n4\n+\n*\n=")))
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}


func TestNewLinesAndSpaces(t *testing.T) {
	expected := (1 + 2) * (3 + 4)
	result, _ := calculate(bufio.NewScanner(bytes.NewBufferString("1\n             2                \n + \n 3 \n 4 \n + \n*\n=")))
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

// unsuccessful
func TestManyActions(t *testing.T) {
	expected := 0
	result, _ := calculate(bufio.NewScanner(bytes.NewBufferString("2 3 + + + + + + =")))
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestNoNumbers(t *testing.T) {
	expected := 0
	result, _ := calculate(bufio.NewScanner(bytes.NewBufferString("+ * - / =")))
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestOnlyEquality(t *testing.T) {
	expected := 0
	result, _ := calculate(bufio.NewScanner(bytes.NewBufferString("=")))
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}