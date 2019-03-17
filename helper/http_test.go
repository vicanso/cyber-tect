package helper

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/h2non/gock"
	"github.com/vicanso/cod"
	"github.com/vicanso/hes"
)

func TestHTTPRequest(t *testing.T) {
	defer gock.Off()

	t.Run("normal", func(t *testing.T) {
		gock.New("http://aslant.site").
			Get("/").
			Reply(200).
			JSON(map[string]string{
				"name": "tree.xie",
			})
		req := httptest.NewRequest("GET", "/users/me", nil)
		ipAddr := "1.1.1.1"
		req.Header.Set(xForwardedForHeader, ipAddr)
		c := cod.NewContext(nil, req)
		cid := "abcd"
		c.ID = cid
		d := GetWithContext("http://aslant.site/", c)
		if d.GetValue(contextID).(string) != cid {
			t.Fatalf("get value fail")
		}
		d.SetClient(http.DefaultClient)
		resp, body, err := d.Do()
		if err != nil {
			t.Fatalf("get request fail, %v", err)
		}
		if d.Request.Header.Get(xForwardedForHeader) != ipAddr ||
			resp.StatusCode != 200 ||
			strings.TrimSpace(string(body)) != `{"name":"tree.xie"}` {
			t.Fatalf("get request fail")
		}
	})

	t.Run("error", func(t *testing.T) {
		gock.New("http://aslant.site").
			Get("/").
			Reply(500).
			JSON(map[string]string{
				"message": "get data fail",
			})
		d := GetWithContext("http://aslant.site/", nil)
		d.SetClient(http.DefaultClient)
		_, _, err := d.Do()
		he, ok := err.(*hes.Error)
		if !ok {
			t.Fatalf("error should convert to hes error")
		}
		if he.Category != errCategoryHTTPRequest ||
			he.Extra["uri"] != "/" ||
			he.Extra["method"] != "GET" ||
			he.Extra["host"] != "aslant.site" {
			t.Fatalf("covert error fail")
		}

		resp := d.Response
		if resp.StatusCode != 500 {
			t.Fatalf("get request fail")
		}
	})

	t.Run("timeout", func(t *testing.T) {
		d := GetWithContext("https://aslant.site/", nil)
		d.Timeout(time.Millisecond)
		_, _, err := d.Do()
		he, ok := err.(*hes.Error)
		if !ok || he.StatusCode != http.StatusRequestTimeout {
			t.Fatalf("should return timeout error")
		}
	})
}
