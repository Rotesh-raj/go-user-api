package utils

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "Birthday passed this year",
			dob:      now.AddDate(-35, -1, 0),
			expected: 35,
		},
		{
			name:     "Birthday not passed this year",
			dob:      now.AddDate(-35, 1, 0),
			expected: 34,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if age := CalculateAge(tt.dob); age != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, age)
			}
		})
	}
}
