package errcode

import (
	"net/http"
	"testing"
)

func TestToHttpStatusCode(t *testing.T) {
	m := map[int]int{
		0:        http.StatusOK,
		FAIL:     http.StatusInternalServerError,
		20000000: http.StatusOK,
	}
	for code, expectStatus := range m {
		status := ToHttpStatusCode(code)
		if status != expectStatus {
			t.Errorf("code:%d expectStatus:%d status:%d", code, expectStatus, status)
		}
	}
}
