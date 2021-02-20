package exercise_test

import (
	"testing"

	"github.com/nextuponstream/workoutReminderBot/pkg/entities"
	"github.com/nextuponstream/workoutReminderBot/pkg/handlers/exercise"
)

type Ex = entities.Exercise

func TestGetE(t *testing.T) {

	tests := []struct {
		msg         string
		expected    entities.Exercise
		wantError   bool
		expectedErr error
	}{
		{"/activity pushups", Ex{Reps: 0, Set: 0, Length: 0, Duration: "", Notes: ""}, false, nil},
		{"/activity pushups\nr 3", Ex{Reps: 3, Set: 0, Length: 0, Duration: "", Notes: ""}, false, nil},
		{"/activity pushups\nr 4", Ex{Reps: 4, Set: 0, Length: 0, Duration: "", Notes: ""}, false, nil},
		{"/activity pushups\ns 5", Ex{Reps: 0, Set: 5, Length: 0, Duration: "", Notes: ""}, false, nil},
		{"/activity pushups\nl 5.5", Ex{Reps: 0, Set: 0, Length: 5.5, Duration: "", Notes: ""}, false, nil},
		{"/activity pushups\nd 15m", Ex{Reps: 0, Set: 0, Length: 0, Duration: "15m", Notes: ""}, false, nil},
		{"/activity pushups\nn c'est chaud", Ex{Reps: 0, Set: 0, Length: 0, Duration: "", Notes: "c'est chaud"}, false, nil},
		{"/activity pushups\nr abd", Ex{Reps: 3, Set: 0, Length: 0, Duration: "", Notes: ""}, true, nil},
		{"/activity pushups\nr 4\ns 5\nl 5.5\nd 15m\nn c'est chaud", Ex{Reps: 4, Set: 5, Length: 5.5, Duration: "15m", Notes: "c'est chaud"}, false, nil},
		{"/activity pushups\nr 4\ns 5\nl 5.aaa\nd 15m\nn c'est chaud", Ex{Reps: 4, Set: 5, Length: 5.5, Duration: "15m", Notes: "c'est chaud"}, true, nil},
	}

	for _, tt := range tests {
		got, err := exercise.GetExercice(tt.msg)
		if tt.wantError {
			if tt.expectedErr == nil {
				if err == nil {
					t.Errorf("msg: %s; expected error but got none", tt.msg)
				}
			} else if err != tt.expectedErr {
				t.Errorf("msg: %s; got: %v; want %v", tt.msg, err, tt.expectedErr)
			}
		} else if got != tt.expected {
			t.Errorf("msg: %s; got: %v; want %v", tt.msg, got, tt.expected)
		}
	}
}
