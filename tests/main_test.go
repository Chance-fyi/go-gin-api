package tests

import (
	. "github.com/smartystreets/goconvey/convey"
	"go-gin-api/boot"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := boot.SetRouter()

	Convey("api ping", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)
		router.ServeHTTP(w, req)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(w.Body.String(), ShouldEqual, "{\"message\":\"pong\"}")
	})

	Convey("404 request", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/404", nil)
		router.ServeHTTP(w, req)

		So(w.Code, ShouldEqual, http.StatusNotFound)
		So(w.Body.String(), ShouldEqual, "{\"code\":404,\"message\":\"not found\"}")
	})
}
