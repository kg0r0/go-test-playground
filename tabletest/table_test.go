package tabletest

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		src []string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "found", args: args{src: []string{"a", "b", "c"}, dst: "b"}, want: 1, wantErr: false},
		{name: "not found", args: args{src: []string{"a", "b", "c"}, dst: "d"}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Contains(tt.args.src, tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contains() error = %v, hasErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalc(t *testing.T) {
	type args struct {
		a        int
		b        int
		operator string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "add",
			args: args{
				a:        1,
				b:        2,
				operator: "+",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Invalid operator",
			args: args{
				a:        1,
				b:        2,
				operator: "?",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calc(tt.args.a, tt.args.b, tt.args.operator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, hasErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
