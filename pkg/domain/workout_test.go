package domain_test

import (
	"testing"

	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
)

func TestIsValidWorkout(t *testing.T) {
	a := domain.CreateActivity("pushups")
	tests := []struct {
		workout  domain.Workout
		expected bool
	}{
		{domain.Workout{"", []domain.Activity{a}}, false},
		{domain.Workout{"gitgud", []domain.Activity{}}, false},
		{domain.Workout{"gitgud", []domain.Activity{a, a, a}}, true},
	}

	for _, tt := range tests {
		got := tt.workout.IsValid()
		if got != tt.expected {
			t.Errorf("workout: %v; got: %v; expected: %v", tt.workout, got, tt.expected)
		}
	}
}
