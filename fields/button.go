package fields

import (
	"github.com/kirves/go-form-it/common"
)

// Submit button type
type SubmitButtonType struct {
	Field
}

// Reset button type
type ResetButtonType struct {
	Field
}

// SubmitButton creates a default button with the provided name and text.
func SubmitButton(name string, text string) *SubmitButtonType {
	ret := &SubmitButtonType{
		FieldWithType(name, formcommon.SUBMIT),
	}
	ret.SetText(text)
	return ret
}

// ResetButton creates a default reset button with the provided name and text.
func ResetButton(name string, text string) *ResetButtonType {
	ret := &ResetButtonType{
		FieldWithType(name, formcommon.RESET),
	}
	ret.SetText(text)
	return ret
}
