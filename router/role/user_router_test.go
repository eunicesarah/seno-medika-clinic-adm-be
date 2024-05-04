package role

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserRouter(t *testing.T) {
	r := gin.New()

	UserRouter(r)

	if r.Routes() == nil {
		t.Errorf("Expected routes, got nil")
	}

	if r.Routes()[0].Path != "/user" {
		t.Errorf("Expected /user, got %v", r.Routes()[0].Path)
	}

	if r.Routes()[1].Path != "/user" {
		t.Errorf("Expected /user, got %v", r.Routes()[1].Path)
	}

	if r.Routes()[2].Path != "/user" {
		t.Errorf("Expected /user, got %v", r.Routes()[2].Path)
	}

	if r.Routes()[3].Path != "/user" {
		t.Errorf("Expected /user, got %v", r.Routes()[3].Path)
	}

	if r.Routes()[4].Path != "/user" {
		t.Errorf("Expected /user, got %v", r.Routes()[4].Path)
	}

	if r.Routes()[0].Method != "GET" {
		t.Errorf("Expected GET, got %v", r.Routes()[0].Method)
	}

	if r.Routes()[1].Method != "POST" {
		t.Errorf("Expected POST, got %v", r.Routes()[1].Method)
	}

	if r.Routes()[2].Method != "PUT" {
		t.Errorf("Expected PUT, got %v", r.Routes()[2].Method)
	}

	if r.Routes()[3].Method != "PATCH" {
		t.Errorf("Expected PATCH, got %v", r.Routes()[3].Method)
	}

	if r.Routes()[4].Method != "DELETE" {
		t.Errorf("Expected DELETE, got %v", r.Routes()[4].Method)
	}

	t.Run("GetUser", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/user", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code == http.StatusNotFound {
			t.Errorf("expected status Exists; got %v", w.Code)
		}
	})

	t.Run("AddUser", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/user", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code == http.StatusNotFound {
			t.Errorf("expected status Exists; got %v", w.Code)
		}
	})

	t.Run("PutUser", func(t *testing.T) {
		jsonStr := []byte(`{"username": "testuser", "password": "testpassword"}`)
		req, _ := http.NewRequest("PUT", "/user", bytes.NewBuffer(jsonStr))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code == http.StatusNotFound {
			t.Errorf("expected status Exists; got %v", w.Code)
		}
	})

	t.Run("PatchUser", func(t *testing.T) {
		jsonStr := []byte(`{"username": "testuser", "password": "testpassword"}`)
		req, _ := http.NewRequest("PATCH", "/user", bytes.NewBuffer(jsonStr))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code == http.StatusNotFound {
			t.Errorf("expected status Exists; got %v", w.Code)
		}
	})

	t.Run("DeleteUser", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/user", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code == http.StatusNotFound {
			t.Errorf("expected status Exists; got %v", w.Code)
		}
	})
}
