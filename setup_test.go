package packageName

import (
	"testing"
)

func TestRealCommander_Run(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := RealCommander{}
			if got := c.Run(tt.args.command); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ask(t *testing.T) {
	type args struct {
		question     string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"testing Y/n", args{"confirm Y", "Y/n"}, "Y/n"},
		{"testing y/N", args{"confirm N", "y/N"}, "y/N"},
		{"testing word", args{"author name", "name"}, "name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ask(tt.args.question, tt.args.defaultValue); got != tt.want {
				t.Errorf("ask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_confirm(t *testing.T) {
	type args struct {
		question     string
		defaultValue bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"returns true when default is true", args{"test question", true}, true},
		{"returns false when default is true", args{"test question", false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := confirm(tt.args.question, tt.args.defaultValue); got != tt.want {
				t.Errorf("confirm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slugify(t *testing.T) {
	type args struct {
		subject string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"lower case test", args{"test and roll"}, "test-and-roll"},
		{"upper case test", args{"TEST AND ROLL"}, "test-and-roll"},
		{"Title Case case test", args{"Test and Roll"}, "test-and-roll"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slugify(tt.args.subject); got != tt.want {
				t.Errorf("slugify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_titleCase(t *testing.T) {
	type args struct {
		subject string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"lower case test", args{"test and roll"}, "TestAndRoll"},
		{"upper case test", args{"TEST AND ROLL"}, "TestAndRoll"},
		{"Title Case case test", args{"Test and Roll"}, "TestAndRoll"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleCase(tt.args.subject); got != tt.want {
				t.Errorf("titleCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeln(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
	}{
		{"can write", args{"test and roll"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeln(tt.args.line)
		})
	}
}
