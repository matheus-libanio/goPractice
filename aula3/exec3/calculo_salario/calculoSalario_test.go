package calculosalario

import "testing"

func TestCalculoSalario(t *testing.T) {
	type args struct {
		minutos   int
		categoria string
	}
	tests := []struct {
		name        string
		args        args
		wantSalario float64
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSalario, err := CalculoSalario(tt.args.minutos, tt.args.categoria)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculoSalario() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSalario != tt.wantSalario {
				t.Errorf("CalculoSalario() = %v, want %v", gotSalario, tt.wantSalario)
			}
		})
	}
}
