package KataGin

import "testing"

// aTestFunction  a test function for doc
func aTestFunction() {}

// MultilineFunction
// a new line comment
func MultilineFunction() {

}

type commentTest struct {
}

// StructMethod test for function in struct.
func (ct commentTest) StructMethod() {

}

func TestFuncDescription(t *testing.T) {

	type args struct {
		f interface{}
		n string
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
			if got := funcDescription(tt.args.f, tt.args.n); got != tt.want {
				t.Errorf("FuncDescription() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}
