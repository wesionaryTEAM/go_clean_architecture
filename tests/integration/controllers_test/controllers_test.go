package controllers_test

import (
	"clean-architecture/tests/setup"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	setup.TeardownDB()
	os.Exit(code)
}
