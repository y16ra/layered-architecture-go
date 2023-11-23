package model

import (
	"testing"
)

func TestNewUserName(t *testing.T) {
	testCases := []struct {
		name          string
		userName      string
		expectedError bool
		expectedValue string
	}{
		{"Empty string", "", true, ""},
		{"Less than min length", "ab", true, ""},
		{"Min length", "abc", false, "abc"},
		{"Max length", "12345678901234567890123456789012345678901234567890", false, "12345678901234567890123456789012345678901234567890"},
		{"More than max length", "123456789012345678901234567890123456789012345678901", true, ""},
		{"Only spaces", "   ", true, ""},
		{"Spaces before and after", " abc ", false, "abc"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userName, err := NewUserName(tc.userName)
			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error status %v for input '%v', got %v", tc.expectedError, tc.userName, err)
			}

			if !tc.expectedError && string(userName) != tc.expectedValue {
				t.Errorf("Expected '%v', got %v", tc.expectedValue, userName)
			}
		})
	}
}

func TestNewUser(t *testing.T) {
	type args struct {
		name UserName
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "normal case",
			args: args{
				name: "testuser",
			},
			want: &User{
				Name: "testuser",
			},
		},
		{
			name: "empty name",
			args: args{
				name: "",
			},
			want: nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUser(tt.args.name)
			if got == nil {
				if tt.want != nil {
					t.Errorf("NewUser() = nil, want %v", tt.want)
				}
				return
			}
			if got.Name != tt.want.Name {
				t.Errorf("NewUser().Name = %v, want %v", got.Name, tt.want.Name)
			}
		})
	}
}
