package infrastructure_test

import (
	"os"
	"path"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInfrastructure(t *testing.T) {
	// change cwd to project root for framework.Env to work
	err := os.Chdir(path.Join("..", ".."))
	if err != nil {
		panic(err)
	}

	RegisterFailHandler(Fail)
	RunSpecs(t, "Infrastructure Suite")
}
