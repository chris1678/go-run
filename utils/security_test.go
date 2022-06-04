package utils

import (
	"fmt"
	"testing"
)

func TestGenerateRandomNum6(t *testing.T) {
	type args struct {
		width int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"a", args{width: 6}, "123456"},
		{"a", args{width: 6}, "123456"},
		{"a", args{width: 6}, "123456"},
		{"a", args{width: 6}, "123456"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRandomNum6(tt.args.width)
			fmt.Println(got)
			if len(got) != len(tt.want) {
				t.Errorf("GenerateRandomNum6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPassword(t *testing.T) {
	type args struct {
		password string
		salt     string
	}
	tests := []struct {
		name       string
		args       args
		wantVerify string
		wantErr    bool
	}{
		{name: "fdsa", args: args{password: "321321", salt: "fdsa"}},
		{name: "fdsa", args: args{password: "3221321", salt: "fdsa"}},
		{name: "fdsa", args: args{password: "3251321", salt: "fdsa"}},
		{name: "fdsa", args: args{password: "3231321", salt: "fdsa"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVerify, _ := SetPassword(tt.args.password, tt.args.salt)
			fmt.Println(gotVerify)
		})
	}
}
