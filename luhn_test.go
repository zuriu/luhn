package luhn

import (
	"testing"
)

func TestLuhn(t *testing.T) {
	sampleString := "559350484135234"

	if isValid, err := IsValid(sampleString); isValid || err != nil {
		t.Errorf("Should have failed validation since no Luhn was appended.")
	}

	var err error
	if sampleString, err = Append(sampleString); err != nil {
		t.Errorf("Failed append: %v", err)
	}

	if isValid, err := IsValid(sampleString); !isValid || err != nil {
		t.Errorf("Failed validation after appending luhn.")
	}
}

func TestPredefinedLuhn(t *testing.T) {
	ccNumber := "4916201457134830"

	if isValid, err := IsValid(ccNumber); !isValid || err != nil {
		t.Errorf("Failed CC Number check: %v", isValid)
	}
}