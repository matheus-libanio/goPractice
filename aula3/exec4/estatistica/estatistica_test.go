package estatistica_test

import (
	"testing"

	"github.com/matheus-libanio/goPractice/exec4/estatistica"
)

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
			got, err := estatistica.Operation(tt.args.operacao)
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
			if got := estatistica.MinFunc(tt.args.values...); got != tt.want {
				t.Errorf("MinFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAvgFunc(t *testing.T) {
	// arrange
	expected := 4.0
	// act
	result := estatistica.AvgFunc(12, 2, 2, 0)

	// assert
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
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
			if got := estatistica.MaxFunc(tt.args.values...); got != tt.want {
				t.Errorf("MaxFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
