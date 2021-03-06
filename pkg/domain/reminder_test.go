package domain_test

import (
	"testing"

	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
)

func TestAtLeastOneDay(t *testing.T) {
	tests := []struct {
		w        domain.Week
		expected bool
	}{
		{domain.Week{[7]bool{false, false, false, false, false, false, false}}, false},
		{domain.Week{[7]bool{false, true, false, false, false, false, false}}, true},
		{domain.Week{[7]bool{false, false, true, false, false, false, false}}, true},
		{domain.Week{[7]bool{false, false, false, false, true, false, false}}, true},
		{domain.Week{[7]bool{false, false, false, false, false, false, true}}, true},
	}

	for _, tt := range tests {
		got := tt.w.AtLeastOneDay()
		if got != tt.expected {
			t.Errorf("week: %v; got: %v; expected: %v", tt.w, got, tt.expected)
		}
	}
}

func TestIsValidReminder(t *testing.T) {
	w := domain.Week{[7]bool{false, false, false, false, false, false, true}}
	tests := []struct {
		r        domain.Reminder
		expected bool
	}{
		{domain.Reminder{"", w, 0, 1}, false},
		{domain.Reminder{"", w, 3, 8}, false},
		{domain.Reminder{"", w, 17, 16}, false},
		{domain.Reminder{"", w, domain.MAX_START + 1, domain.MAX_START + 2}, false},
		{domain.Reminder{"", w, 6, 7}, true},
		{domain.Reminder{"", w, 16, 21}, true},
		{domain.Reminder{"", w, 21, 22}, true},
	}

	for _, tt := range tests {
		got := tt.r.IsValid()
		if got != tt.expected {
			t.Errorf("week: %v; got: %v; expected: %v", tt.r, got, tt.expected)
		}
	}
}
