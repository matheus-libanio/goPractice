package estatistica

import "testing"

func TestOperation(t *testing.T) {
	type args struct {
		operacao int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Operation(tt.args.operacao)
			if (err != nil) != tt.wantErr {
				t.Errorf("Operation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinFunc(t *testing.T) {
	type args struct {
		values []int
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
			if got := MinFunc(tt.args.values...); got != tt.want {
				t.Errorf("MinFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAvgFunc(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AvgFunc(tt.args.values...); got != tt.want {
				t.Errorf("AvgFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxFunc(t *testing.T) {
	type args struct {
		values []int
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
			if got := MaxFunc(tt.args.values...); got != tt.want {
				t.Errorf("MaxFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
