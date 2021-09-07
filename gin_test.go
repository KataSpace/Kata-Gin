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
