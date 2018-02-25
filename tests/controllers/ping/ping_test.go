package ping_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"os"

	"github.com/efrenfuentes/imageproxy/http/routers"
	"github.com/efrenfuentes/imageproxy/http/settings"
	"github.com/efrenfuentes/imageproxy/tests/utils"

	"github.com/gorilla/mux"
)

var router *mux.Router

var _ = Describe("Ping Controller", func() {
	BeforeSuite(func() {
		os.Setenv("IMAGEPROXY_ENV", "test")
		settings.Init()
		router = routers.Init()
	})

	It("Calling /ping", func() {
		var pingResponse interface{}

		w := utils.MakeRequest("GET", "/ping", nil, router)

		Expect(w.Code).To(Equal(http.StatusOK))

		pingResponse = w.Body.String()

		Expect(pingResponse).To(Equal("pong"))
	})
})
