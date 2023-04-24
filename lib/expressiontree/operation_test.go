package expressiontree

import "testing"

func TestOperation_String(t *testing.T) {
	type testCase[T any] struct {
		name string
		o    Operation[T]
		want string
	}
	tests := []testCase[float64]{
		{
			name: "Basic Add",
			o: func(a float64, b float64) float64 {
				return a + b
			},
			want: "func(float64, float64) float64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
