package models

import (
	"testing"
	"time"
)

func TestEmptyFilename(t *testing.T) {
	order, err := ReadModel("")
	if err == nil {
		t.Fatal("Error is empty")
	}

	if order != nil {
		t.Fatal("Order is not empty")
	}

	if err.Error() != "filename can't be empty" {
		t.Fatalf("wrong error")
	}
}

// Testing ReadModel function
func TestReadModel(t *testing.T) {
	filename := "../../model.json"
	order, err := ReadModel(filename)
	if err != nil {
		t.Fatal(err)
	}

	TestOrder(t, order)
}
