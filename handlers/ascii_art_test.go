package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAsciiArtHandler(t *testing.T) {
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
			name: "Invalid method",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/ascii-art", nil),
			},
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method Not Allowed\n",
		},
		{
			name: "Missing text",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader("banner=standard")),
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Bad Request: Missing text\n",
		},
		{
			name: "Non-ASCII characters",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader("text=Hello世界&banner=standard")),
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Bad Request: Input contains non-ASCII characters\n",
		},
		{
			name: "Invalid banner",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader("text=Hello&banner=invalid")),
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not Found: Invalid banner\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the correct Content-Type for POST requests
			if tt.args.r.Method == http.MethodPost {
				tt.args.r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			}

			AsciiArtHandler(tt.args.w, tt.args.r)

			resp := tt.args.w.(*httptest.ResponseRecorder)

			if resp.Code != tt.expectedStatus {
				t.Errorf("AsciiArtHandler() status = %v, want %v", resp.Code, tt.expectedStatus)
			}

			if tt.expectedBody != "" && resp.Body.String() != tt.expectedBody {
				t.Errorf("AsciiArtHandler() body = %v, want %v", resp.Body.String(), tt.expectedBody)
			}
		})
	}
}

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

			if tt.expectedBody != "" && resp.Body.String() != tt.expectedBody {
				t.Errorf("HomeHandler() body = %v, want %v", resp.Body.String(), tt.expectedBody)
			}
		})
	}
}
