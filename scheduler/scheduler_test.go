package scheduler

import (
	"testing"
	"time"
)

func TestAllow(t *testing.T) {
	tests := []struct {
		name       string
		expected   bool
		tm         time.Time
		publishDay time.Weekday
		maxPublish int
	}{
		{
			name:     "Should return false if time (hour) is not in correct region",
			expected: false,
			tm:       time.Date(2017, time.November, 26, 23, 0, 0, 0, time.UTC),
		},
		{
			name:       "Should return false if day is different from specified",
			expected:   false,
			tm:         time.Date(2017, time.November, 26, 11, 0, 0, 0, time.UTC), // Sunday
			publishDay: time.Monday,
		},
		{
			name:       "If exceeded maximum number to publish, return false",
			tm:         time.Date(2017, time.November, 26, 11, 0, 0, 0, time.UTC), // Sunday
			publishDay: time.Sunday,
			expected:   false,
		},
		{
			name:       "Return true, if all the above are false",
			expected:   true,
			tm:         time.Date(2017, time.November, 26, 11, 0, 0, 0, time.UTC), // Sunday
			publishDay: time.Sunday,
			maxPublish: 1,
		},
	}

	for _, test := range tests {

		s := &Scheduler{
			publishDay:   test.publishDay,
			maxToPublish: test.maxPublish,
		}

		actual := s.allow(test.tm)
		if test.expected != actual {
			t.Errorf("[%s] Expected %v, got %v", test.name, test.expected, actual)
		}
	}
}
