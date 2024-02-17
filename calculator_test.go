package main

import "testing"

func TestAddShouldBeSuccess(t *testing.T) {
	// Arrange
	numbers := []int{1, 2, 3}
	expected := 6

	// Act
	actual := add(numbers...)

	// Assert
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestAddShouldFail(t *testing.T) {
	// Arrange
	numbers := []int{1, 2, -3} // Include a negative number for expected failure
	expected := 4

	// Act
	actual := add(numbers...)

	// Assert
	if actual == expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSubtractShouldBeSuccess(t *testing.T) {
	// Arrange
	numbers := []int{5, 10}
	expected := 5

	// Act
	actual := subtract(numbers...)

	// Assert
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSubtractShouldFail(t *testing.T) {
	// Arrange
	numbers := []int{10, 5, 2} // Add an extra number for expected failure
	expected := -5

	// Act
	actual := subtract(numbers...)

	// Assert
	if actual == expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestMultiplyShouldBeSuccess(t *testing.T) {
	// Arrange
	numbers := []int{10, 10}
	expected := 100

	// Act
	actual := multiply(numbers...)

	// Assert
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestMultiplyShouldFail(t *testing.T) {
	// Arrange
	numbers := []int{0, 10} // Include a zero for expected failure
	expected := 100

	// Act
	actual := multiply(numbers...)

	// Assert
	if actual == expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDivideShouldBeSuccess(t *testing.T) {
	// Arrange
	numbers := []int{20}
	expected := 20

	// Act
	actual := divide(numbers...)

	// Assert
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDivideShouldFail(t *testing.T) {
	// Arrange
	numbers := []int{20, 0} // Include a zero for expected failure
	expected := 20

	// Act
	actual := divide(numbers...)

	// Assert
	if actual == expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
