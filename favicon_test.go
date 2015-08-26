package favicon

import (
	"github.com/goforgery/forgery2"
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestCreate(t *testing.T) {

	Describe("Create()", func() {

		var app *f.Application
		var req *f.Request
		var res *f.Response

		BeforeEach(func() {
			app = f.CreateApp()
			req = f.CreateRequestMock(app)
			res, _ = f.CreateResponseMock(app, false)
		})

		It("should return [false]", func() {
			app.Use("", Create())
			app.Handle(req, res, 0)
			AssertNotEqual(res.Writer.Header().Get("content-type"), "image/x-icon")
		})

		It("should return [false]", func() {
			app.Use("", Create(map[string]string{}))
			app.Handle(req, res, 0)
			AssertNotEqual(res.Writer.Header().Get("content-type"), "image/x-icon")
		})

		It("should return [false]", func() {
			req.OriginalUrl = "/favicon.ic"
			app.Use("", Create())
			app.Handle(req, res, 0)
			AssertNotEqual(res.Writer.Header().Get("content-type"), "image/x-icon")
		})

		It("should return [text/plain] from not found", func() {
			req.OriginalUrl = "/favicon.ico"
			app.Use("", Create())
			app.Handle(req, res, 0)
			AssertEqual(res.Writer.Header().Get("content-type"), "text/plain")
		})

		It("should return [image/x-icon]", func() {
			req.OriginalUrl = "/favicon.ico"
			app.Use("", Create(map[string]string{"path": "./fixtures/favicon.ico"}))
			app.Handle(req, res, 0)
			AssertEqual(res.Writer.Header().Get("content-type"), "image/x-icon")
		})

		It("should return [image/x-icon] from cache", func() { // checked on coverage report
			req.OriginalUrl = "/favicon.ico"
			app.Use("", Create(map[string]string{"path": "./fixtures/favicon.ico"}))
			app.Handle(req, res, 0)
			res, _ = f.CreateResponseMock(app, false)
			app.Handle(req, res, 0)
			AssertEqual(res.Writer.Header().Get("content-type"), "image/x-icon")
		})

		It("should return [public, max-age=1]", func() {
			req.OriginalUrl = "/favicon.ico"
			app.Use("", Create(map[string]string{"path": "./fixtures/favicon.ico", "maxage": "1000"}))
			app.Handle(req, res, 0)
			AssertEqual(res.Writer.Header().Get("cache-control"), "public, max-age=1")
		})
	})

	Report(t)
}
