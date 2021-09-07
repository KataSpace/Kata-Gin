package KataGin

import "testing"

func Test_defaultGetMethods(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantMethod string
		wantName   string
	}{
		{
			name:       "lower char function",
			args:       args{s: "getMethodsFunction"},
			wantMethod: "get",
			wantName:   "MethodsFunction",
		},
		{
			name:       "upper char function",
			args:       args{s: "GetMethodsFunction"},
			wantMethod: "Get",
			wantName:   "MethodsFunction",
		},
		{
			name:       "Multiple upper name function",
			args:       args{s: "getAMethodsFunction"},
			wantMethod: "get",
			wantName:   "AMethodsFunction",
		},
		{
			name:       "Multiple upper name function with slash",
			args:       args{s: "Get/AMethods/Function"},
			wantMethod: "Get",
			wantName:   "/AMethods/Function",
		},
		{
			name:       "Multiple upper name function with more slash",
			args:       args{s: "Get/A/P/I/All/Name"},
			wantMethod: "Get",
			wantName:   "/A/P/I/All/Name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMethod, gotName := defaultGetMethods(tt.args.s)
			if gotMethod != tt.wantMethod {
				t.Errorf("defaultGetMethods() gotMethod = %v, want %v", gotMethod, tt.wantMethod)
			}
			if gotName != tt.wantName {
				t.Errorf("defaultGetMethods() gotName = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}

func TestSlashConvert(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Use Normal Upper Char",
			args: args{s: "GetAllName", n: -1},
			want: "Get/All/Name",
		},
		{
			name: "Use Multiple Upper Chars",
			args: args{s: "AAAGetAllName", n: -1},
			want: "Get/All/Name",
		},
		{
			name: "Use Multiple Upper Chars",
			args: args{s: "GetAAAAllName", n: -1},
			want: "Get/All/Name",
		},
		{
			name: "Use Multiple Upper Chars With N = 1",
			args: args{s: "GetAPIAllName", n: 1},
			want: "Get/A/P/I/All/Name",
		},
		{
			name: "Use Multiple Upper Chars With N = 2",
			args: args{s: "GetAPIAllName", n: 2},
			want: "Get/AP/IA/Name",
		},
		{
			name: "Use Multiple Upper Chars With N = 3",
			args: args{s: "GetAPIAllName", n: 3},
			want: "Get/API/All/Name",
		},
		{
			name: "Use Multiple Upper Chars With N = 10",
			args: args{s: "GetAPIAllName", n: 10},
			want: "Get/All/Name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slashConvert(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("SlashConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}
