package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func handleGETRequest(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("user_id")
	fmt.Printf(id)
}

func TestAuthenticationMiddleware_Success(t *testing.T) {
    req, err := http.NewRequest("GET", "/validate", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handleGETRequest)
    handler.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusOK, rr.Code)
}
