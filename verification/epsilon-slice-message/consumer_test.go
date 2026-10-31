package consumer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFailureContext(t *testing.T) {
	for _, tc := range []struct {
		name string
		expected, actual interface{}
	}{
		{"type", []float64{1}, []float32{1}},
		{"length", []float64{1, 2}, []float64{1}},
		{"element", []float64{1, 2}, []float64{1, 3}},
	} {
		for _, entry := range []string{"assert", "assert-method", "assertf", "require", "require-method", "requiref"} {
			t.Run(entry+"/"+tc.name, func(t *testing.T) {
				switch entry {
				case "assert":
					assert.InEpsilonSlice(t, tc.expected, tc.actual, .01, "batch %d", 7)
				case "assert-method":
					assert.New(t).InEpsilonSlice(tc.expected, tc.actual, .01, "batch %d", 7)
				case "assertf":
					assert.InEpsilonSlicef(t, tc.expected, tc.actual, .01, "batch %d", 7)
				case "require":
					require.InEpsilonSlice(t, tc.expected, tc.actual, .01, "batch %d", 7)
				case "require-method":
					require.New(t).InEpsilonSlice(tc.expected, tc.actual, .01, "batch %d", 7)
				case "requiref":
					require.InEpsilonSlicef(t, tc.expected, tc.actual, .01, "batch %d", 7)
				}
			})
		}
	}
}

type messageProbe struct{ calls int }
func (m *messageProbe) String() string { m.calls++; return "batch 7" }

func TestPassingAssertionDoesNotFormatMessage(t *testing.T) {
	m := new(messageProbe)
	assert.InEpsilonSlice(t, []float64{1, 2, 3}, []float64{1, 2, 3}, .01, m)
	if m.calls != 0 { t.Fatalf("message formatted %d times", m.calls) }
}
