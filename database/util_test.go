package database

import (
	"testing"
)

func TestReplace(t *testing.T) {
	testCases := []string {
		replace("INSERT INTO ?(?, ?, ?, ?, ?) VALUES(?, ?, ?, ?, ?)", "?", 
				1,[]string { 
				"SERVER",
						"ID", 
						"NAME", 
						"ACTIVE", 
						"IP", 
						"SERVICE", 
						"ABC", 
						"XYZ", 
						"ACTIVE", 
						"192.168.9.100",
						"NOTIFICATION"}),

		replace("replace the string using letter", "e", 1,
				[]string {"x", "x"}),

		replace("abc?xyz", "?", 1, []string {}),
			}

	expected := []string {
		"INSERT INTO SERVER(ID, NAME, ACTIVE, IP, SERVICE) VALUES(ABC, XYZ, ACTIVE, 192.168.9.100, NOTIFICATION)",
		"rxplacx the string using letter",
		"abc?xyz",
	}

	for i := 0; i < len(testCases); i++ {
		if testCases[i] != expected[i] {
			t.Error(testCases[i], " should equal ", expected[i])
		}
	}
}