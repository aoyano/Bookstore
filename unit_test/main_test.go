package main

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := 8
	result := Max(numbers)

	if result != expected {
		t.Errorf("Expected %d got %d", expected, result)
	}
}

func TestMin(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	expected := 0
	result := Min(numbers)

	if result != expected {
		t.Errorf("Expected %d got %d", expected, result)
	}
}

func TestGetName(t *testing.T) {
	expected := GetName()
	if expected != "Ayan" {
		t.Errorf("Expected %s fot %s", expected, "Ayan")
	}
}

func TestUbdateName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"ReturnName(\"Ayan\") must return Ayanat", "Ayanat", "Ayanat"},
		{"ReturnName(\"Aibek\") must return Aigul", "Aigul", "Aigul"},
		{"ReturnName(\"Eldar\") must return Karl", "Karl", "Karl"},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			result := UbdateName(tst.input)
			if result != tst.expected {
				t.Errorf("Expected %s got %s", tst.expected, result)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"ReturnEmail(\"ayan.nurgaliyev.123456789@gmail.com\") must return ayan.nurgaliyev.123456789@gmail.com", "ayan.nurgaliyev.123456789@gmail.com", true},
		{"ReturnEmail(\"ayan.ayanat.123@gmail.com\") must return ayan.ayanat.123@gmail.com", "ayan.ayanat.123@gmail.com", true},
		{"ReturnEmail(\"nurgaliyev.ayan@mail.ru\") must return nurgaliyev.ayan@mail.ru", "nurgaliyev.ayan@mail.ru", true},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			if login(tst.input) != tst.input {
				t.Errorf("Expected %s got", tst.input)
			}
		})
	}
}

func TestCalculateDiscount(t *testing.T) {
	tests := []struct {
		price    float64
		discount float64
		expected float64
	}{
		{
			price:    120,
			discount: 0.8,
			expected: 96,
		},
		{
			price:    80,
			discount: 0.8,
			expected: 64,
		},
		{
			price:    95,
			discount: 0.8,
			expected: 76,
		},
	}

	for _, tst := range tests {
		t.Run(fmt.Sprintf("price = %.2f, discount = %.2f", tst.price, tst.discount), func(t *testing.T) {
			got := CalculateDiscount(tst.price, tst.discount)
			if got != tst.expected {
				t.Errorf("CalculateDiscount() = %.2f expected %.2f", got, tst.expected)
			}
		})
	}
}

func TestEmptyField(t *testing.T) {
	if !EmptyField("") {
		t.Errorf("EmptyField(\"\") returned false, expected true")
	}

	if !EmptyField("  ") {
		t.Errorf("EmptyField(\"  \") returned false, expected true")
	}

	if EmptyField("hello") {
		t.Errorf("EmptyField(\"hello\") returned true, expected false")
	}
}

func TestTwoPlusOne(t *testing.T) {
	expected := TwoPlusOne(120, 80, 95, 80)
	if expected != 215 {
		t.Errorf("Expected %d got %d", expected, 215)
	}
}

func TestGetAccountsByOrder(t *testing.T) {
	tests := []struct {
		model    string
		input    string
		expected string
	}{
		{"ReturnModel(\"Jordan\") must return Jordan", "Jordan", "Jordan"},
		{"ReturnModel(\"KD\") must return KD", "KD", "KD"},
		{"ReturnModel(\"Lebron\") must return Lebron", "Lebron", "Lebron"},
	}

	for _, tst := range tests {
		t.Run(tst.model, func(t *testing.T) {
			result := getAccountsByOrder(tst.input)
			if result != tst.expected {
				t.Errorf("Expected %s got %s", tst.expected, result)
			}
		})
	}
}

func TestVolume(t *testing.T) {
	expected := Volume(1, 2, 3)
	if expected != 6 {
		t.Errorf("Expected %d got %d", expected, 6)
	}
}
