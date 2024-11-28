package calcimposto

import "testing"

func TestCalcImposto(t *testing.T) {
	type args struct {
		salario int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcImposto(tt.args.salario); got != tt.want {
				t.Errorf("CalcImposto() = %v, want %v", got, tt.want)
			}
		})
	}
}
