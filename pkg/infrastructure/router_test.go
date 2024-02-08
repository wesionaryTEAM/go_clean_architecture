package infrastructure_test

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	"github.com/steinfletcher/apitest"

	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
)

var (
	t      GinkgoTInterface
	router infrastructure.Router
	env    *framework.Env
	logger framework.Logger
)

var _ = Describe("Router", func() {
	BeforeEach(func() {
		t = GinkgoT()
		logger = framework.GetLogger()
		env = framework.NewEnv(logger)
		router = infrastructure.NewRouter(env, logger)
	})

	Describe("HealthCheck", func() {
		Context("when health-check endpoint is called", func() {
			It("should return 200", func() {
				apitest.New().Handler(router.Engine).Get("/health-check").Expect(t).Status(http.StatusOK).End()
			})
		})
	})
})
