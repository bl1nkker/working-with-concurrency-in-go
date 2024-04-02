package dining

import (
	"testing"
	"time"
)

func TestDine(t *testing.T) {
	EatTime = 0 * time.Second
	ThinkTime = 0 * time.Second
	SleepTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		order = []string{}
		Dine()
		if len(order) != 5 {
			t.Errorf("Incorrect length of slice. Expected 5, but got %d", len(order))
		}
	}
}

func TestDineWithVaryingDelay(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", 0 * time.Second},
		{"quarter second delay", 250 * time.Millisecond},
		{"half second delay", 500 * time.Millisecond},
	}

	for _, e := range theTests {
		EatTime = e.delay
		ThinkTime = e.delay
		SleepTime = e.delay
		order = []string{}
		Dine()
		if len(order) != 5 {
			t.Errorf("%s: Incorrect length of slice. Expected 5, but got %d", e.name, len(order))
		}
	}
}
