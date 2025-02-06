package matcher

import (
    "testing"
    "go.uber.org/zap"
)


// Test generated using Keploy
func TestValidateAndMarshalJSON(t *testing.T) {
    log := zap.NewNop()
    exp := `{"key":"value"}`
    act := `{"key":"value"}`
    validatedJSON, err := ValidateAndMarshalJSON(log, &exp, &act)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if !validatedJSON.IsIdentical() {
        t.Errorf("Expected validatedJSON to be identical, got false")
    }
}
