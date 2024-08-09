package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name           string
		args           args
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Invalid path",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/invalid", nil),
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HomeHandler(tt.args.w, tt.args.r)

			resp := tt.args.w.(*httptest.ResponseRecorder)

			if resp.Code != tt.expectedStatus {
				t.Errorf("HomeHandler() status = %v, want %v", resp.Code, tt.expectedStatus)
				t.Logf("Response Body: %s", resp.Body.String())
			}

			// if tt.expectedBody != "" && resp.Body.String() != tt.expectedBody {
			// 	t.Errorf("HomeHandler() body = %v, want %v", resp.Body.String(), tt.expectedBody)
			// }
		})
	}
}
