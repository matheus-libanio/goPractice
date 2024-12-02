package media

import "testing"

func TestMediaNotas(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name      string
		args      args
		wantMedia float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMedia := MediaNotas(tt.args.values...); gotMedia != tt.wantMedia {
				t.Errorf("MediaNotas() = %v, want %v", gotMedia, tt.wantMedia)
			}
		})
	}
}
