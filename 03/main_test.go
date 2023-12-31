package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	fileContent  string
	contentArray [][]string
)

func testSetup() error {
	var err error
	fileContent, err = ReadFile(TEST_FILE)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	contentArray = ConvertToStringMatrix(fileContent)
	expectedLength := 10
	if len(contentArray) != expectedLength {
		return fmt.Errorf("Expected byte array length of %v, got %d",
			expectedLength, len(contentArray))
	}

	return nil
}

func TestMain(m *testing.M) {
	if err := testSetup(); err != nil {
		log.Fatalf("Test setup failed: %v", err)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

// ---- helper ----
func TestAppendFrontString(t *testing.T) {
	front := []string{"a", "b"}
	end := []string{"c", "d"}

	result := AppendFront(end, front...)

	expectedLen := 4
	expectedArr := []string{"a", "b", "c", "d"}

	if len(result) != expectedLen {
		t.Errorf("Length mismatch for %v. Expected %v, got %v", result, expectedLen, len(result))
	}

	for i, v := range result {
		if v != expectedArr[i] {
			t.Errorf("Elements mismatch. Expected %v, got %v", expectedArr[i], v)
		}
	}
}

func TestIsNumber(t *testing.T) {
	str1 := "a"
	str2 := "4"

	if IsNumber(str1) {
		t.Errorf("Failed. %v is not a number", str1)
	}

	if !IsNumber(str2) {
		t.Errorf("Failed. %v is a number", str2)
	}
}

func TestIsSymbol(t *testing.T) {
	str1 := "4"
	str2 := "."
	str3 := "$"

	if IsSymbol(str1) {
		t.Errorf("%v is not a symbol", str1)
	}

	if IsSymbol(str2) {
		t.Errorf("%v is not a symbol", str2)
	}

	if !IsSymbol(str3) {
		t.Errorf("%v is a symbol", str3)
	}
}

func TestSameSymbol(t *testing.T) {
	sym1 := Symbol{
		Value: "*",
		X:     4,
		Y:     2,
	}
	sym2 := Symbol{
		Value: "*",
		X:     4,
		Y:     2,
	}
	sym3 := Symbol{
		Value: "*",
		X:     2,
		Y:     3,
	}

	num1 := Number{
		Symbol: sym1,
	}
	num2 := Number{
		Symbol: sym2,
	}
	num3 := Number{
		Symbol: sym3,
	}

	if !SameSymbol(num1, num2) {
		t.Errorf("Both have the same symbol: %v and %v", num1.Symbol, num2.Symbol)
	}

	if SameSymbol(num1, num3) {
		t.Errorf("Both have different symbols: %v and %v", num1.Symbol, num3.Symbol)
	}
}

// ---- Main functions ----
func TestReadFile(t *testing.T) {
	if len(fileContent) < 1 {
		log.Fatalf("Expected some file content to be found")
	}
}

func TestHasAdjacentSymbols(t *testing.T) {
	testNum1 := Number{
		Value:      42,
		StartIndex: 1,
		EndIndex:   2,
		Line:       0,
	}
	testNum2 := Number{
		Value:      42,
		StartIndex: 1,
		EndIndex:   2,
		Line:       1,
	}

	noSymbol := [][]string{{".", "4", "2", "."}}
	hasSymbol := [][]string{{".", ".", ".", "."}, {".", "4", "2", "."}, {".", ".", ".", "$"}}

	r1, _ := HasAdjacentSymbols(&noSymbol, &testNum1)
	r2, _ := HasAdjacentSymbols(&hasSymbol, &testNum2)

	if r1 {
		t.Errorf("Failed. %v has no adjacent symbol", noSymbol)
	}

	if !r2 {
		t.Errorf("Failed. %v has one adjacent symbol", hasSymbol)
	}

}

func TestFindValidNumbers(t *testing.T) {
	result := FindValidNumbers(&contentArray)

	expectedLen := 8
	if len(result) != expectedLen {
		t.Errorf("Failed. Expected %v, got %v -> %v", 8, len(result), result)
	}
}

func TestEvaluateGamePart1(t *testing.T) {
	nums := FindValidNumbers(&contentArray)
	result := EvaluateGamePart1(&nums)

	expected := 4361

	if result != expected {
		t.Errorf("Failed. Expected %v, got %v", expected, result)
	}
}

func TestEvaluateGamePart2(t *testing.T) {
	nums := FindValidNumbers(&contentArray)
	result := EvaluateGamePart2(&nums)

	expected := 467835

	if result != expected {
		t.Errorf("Failed. Expected %v, got %v", expected, result)
	}
}
