package handler_test

import (
	"testing"
	"time"

	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
	"github.com/nextuponstream/workoutReminderBot/pkg/handler"
)

func TestRemindMe(t *testing.T) {
	tests := []struct {
		reminder domain.Reminder
		expected string
	}{
		{domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{true, true, true, false, false, false, false}}, From: 16, To: 17}, "You will be reminded on: Monday, Tuesday, Wednesday"},
		{domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, false, false, false, false, true}}, From: 16, To: 17}, "You will be reminded on: Sunday"},
		{domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, true, true, false, false, false, true}}, From: 16, To: 17}, "You will be reminded on: Tuesday, Wednesday, Sunday"},
		{domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, true, false, false, false, false, false}}, From: 16, To: 17}, "You will be reminded on: Tuesday"},
		{domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, true, true, false, false, false, false}}, From: 16, To: 17}, "You will be reminded on: Tuesday, Wednesday"},
	}

	for _, tt := range tests {
		got := handler.RemindMessage(tt.reminder)
		if got != tt.expected {
			t.Errorf("reminder: %v; got:%v; want: %v", tt.reminder, got, tt.expected)
		}
	}
}

func TestRemainingDayBeforeReminder(t *testing.T) {
	tests := []struct {
		today    int
		reminder domain.Reminder
		expected int
	}{
		{0, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{true, true, true, false, false, false, false}}, From: 16, To: 17}, 1},
		{1, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, false, false, false, false, true}}, From: 16, To: 17}, 5},
		{3, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, true, true, false, false, false, true}}, From: 16, To: 17}, 3},
		{6, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, true, false, false, false, false, false}}, From: 16, To: 17}, 2},
		{6, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, false, false, false, false, true}}, From: 16, To: 17}, 7},
	}

	for _, tt := range tests {
		got := handler.RemainingDayBeforeReminder(tt.today, tt.reminder)
		if got != tt.expected {
			t.Errorf("today: %v; reminder: %v; got:%v; want: %v", tt.today, tt.reminder, got, tt.expected)
		}
	}
}

func TestTimeUntil(t *testing.T) {
	day1 := time.Date(2020, 2, 3, 15, 0, 0, 0, time.Local)
	day2 := time.Date(2020, 2, 3, 15, 30, 0, 0, time.Local)

	tests := []struct {
		now      time.Time
		hour     int
		expected time.Duration
	}{
		{day1, 16, time.Hour},
		{day1, 17, 2 * time.Hour},
		{day1, 15, 0},
		{day1, 13, time.Hour*24*7 - 2*time.Hour},
		{day2, 16, time.Minute * 30},
		{day2, 20, 4*time.Hour + time.Minute*30},
	}

	for _, tt := range tests {
		got := handler.TimeUntil(tt.hour, tt.now)
		if got != tt.expected {
			t.Errorf("now: %v; hour: %v; got %v; want %v;", tt.now, tt.hour, got, tt.expected)
		}
	}
}

func TestAddDays(t *testing.T) {
	tests := []struct {
		from     int
		to       int
		expected time.Duration
	}{
		{0, 1, time.Hour * 24},
		{3, 6, 3 * time.Hour * 24},
		{1, 1, 0},
		{2, 1, 6 * time.Hour * 24},
		{6, 5, 6 * time.Hour * 24},
	}

	for _, tt := range tests {
		got := handler.AddDays(tt.from, tt.to)
		if got != tt.expected {
			t.Errorf("from: %v; to: %v; got %v; want %v;", tt.from, tt.to, got, tt.expected)
		}
	}
}

func TestGetRemainingTime(t *testing.T) {
	// 2021 feb 24: wednesday
	day1 := time.Date(2021, 2, 24, 15, 0, 0, 0, time.Local)
	day2 := time.Date(2021, 2, 24, 15, 30, 0, 0, time.Local)
	// 2021 feb 24: sunday
	day3 := time.Date(2021, 2, 28, 15, 0, 0, 0, time.Local)
	day4 := time.Date(2021, 2, 28, 15, 30, 0, 0, time.Local)
	tests := []struct {
		now      time.Time
		reminder domain.Reminder
		expected time.Duration
		wantErr  bool
	}{
		{day1, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, true, false, false, false, false}}, From: 15, To: 18}, 0, false},
		{day1, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, true, false, false, false, false}}, From: 16, To: 18}, time.Hour, false},
		{day1, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, true, false, false, false, false}}, From: 17, To: 18}, 2 * time.Hour, false},
		{day2, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, true, false, false, false, false}}, From: 17, To: 18}, time.Hour + 30*time.Minute, false},
		{day1, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, false, true, false, false, false}}, From: 17, To: 18}, 26 * time.Hour, false},
		{day1, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, false, false, true, false, false}}, From: 17, To: 18}, 50 * time.Hour, false},
		{day2, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, false, false, true, false, false}}, From: 17, To: 18}, 49*time.Hour + 30*time.Minute, false},
		{day1, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, true, false, false, false, false}}, From: 14, To: 18}, 167 * time.Hour, false},
		{day2, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, true, false, false, false, false}}, From: 15, To: 18}, 168*time.Hour - 30*time.Minute, false},
		{day3, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, false, false, false, false, true}}, From: 15, To: 18}, 0, false},
		{day4, domain.Reminder{RoutineName: "", When: domain.Week{Week: [7]bool{false, false, false, false, false, false, true}}, From: 15, To: 18}, 168*time.Hour - 30*time.Minute, false},
	}

	for _, tt := range tests {
		got, err := handler.GetRemainingTime(tt.now, tt.reminder)
		if tt.wantErr {
			if err == nil {
				t.Errorf("now: %v; reminder: %v; got %v; want error but got none;", tt.now, tt.reminder, got)
			}
		} else if got != tt.expected {
			t.Errorf("now: %v; reminder: %v; got %v; want %v;", tt.now, tt.reminder, got, tt.expected)
		}
	}
}
