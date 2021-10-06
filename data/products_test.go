package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := Product{
		Name:  "Lucas",
		Price: 1.00,
		SKU:   "avs-asa-asd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
