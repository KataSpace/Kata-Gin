package KataGin

import "testing"

// aTestFunction  a test function for doc
func aTestFunction() {}

// MultilineFunction
// a new line comment
func MultilineFunction() {

}
func TestFuncDescription(t *testing.T) {
	type args struct {
		f interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a single line test",
			args: args{f: aTestFunction},
			want: "aTestFunction  a test function for doc",
		},
		{
			name: "a multiple line test",
			args: args{f: MultilineFunction},
			want: "MultilineFunction\na new line comment",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FuncDescription(tt.args.f); got != tt.want {
				t.Errorf("FuncDescription() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}
