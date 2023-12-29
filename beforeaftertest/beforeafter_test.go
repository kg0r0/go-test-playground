package beforeaftertest

import "testing"

func TestMain(m *testing.M) {
	// before
	setup()

	// after
	defer teardown()

	m.Run()
}

func setup() {
	// before
}

func teardown() {
	// after
}

func TestHoge(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// test cases.
	}

	defer func() {
		// after
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
			}()

			got := Hoge(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("Hoge() = %v, want %v", got, tt.want)
			}
		})
	}
}
