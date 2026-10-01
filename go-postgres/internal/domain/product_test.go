package domain

import (
	"errors"
	"testing"
)

func TestNewProductRejectsNegativePrice(t *testing.T) {
	_, err := NewProduct("Keyboard", -1)
	if err == nil {
		t.Fatal("expected negative price to be rejected")
	}
	if !errors.Is(err, ErrInvalid) {
		t.Fatalf("expected ErrInvalid, got %v", err)
	}
}

func TestProductUpdateChangesDetails(t *testing.T) {
	product, err := NewProduct("Keyboard", 10)
	if err != nil {
		t.Fatal(err)
	}

	if err := product.Update("Mouse", 15); err != nil {
		t.Fatal(err)
	}

	if product.Name != "Mouse" || product.Price != 15 {
		t.Fatalf("unexpected product: %+v", product)
	}
}
