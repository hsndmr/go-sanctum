package config

import "testing"

// TestIsExistFile tests the IsExistFile function
func TestIsExistFile(t *testing.T) {
	isExistFile := isExistFile("config.go")
	if !isExistFile {
		t.Errorf("failed finding file: config.go",)
	}
}

// TestFindEnv tests the FindEnv function
func TestFindEnv(t *testing.T) {
	_, err := findEnv()
	if err != nil {
		t.Errorf("failed finding env: %s", err)
	}
}

func TestIsRunningOnTest(t *testing.T) {
	isRunningOnTest := isRunningOnTest()
	if !isRunningOnTest {
		t.Errorf("not running on test")
	}
}