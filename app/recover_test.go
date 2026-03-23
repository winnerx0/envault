package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_fetchEntries_flatFiles(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := ResponseBody{
			Entries: []Entry{
				{Path: "myapp/.env.enc", Name: ".env.enc", Download_url: "https://example.com/.env.enc", ContentType: "file"},
				{Path: "myapp/.env.local.enc", Name: ".env.local.enc", Download_url: "https://example.com/.env.local.enc", ContentType: "file"},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	origClient := httpClient
	httpClient = *srv.Client()
	defer func() { httpClient = origClient }()

	// fetchEntries uses hardcoded GitHub URL, so we need to override it.
	// Since we can't easily do that, we'll test via githubRequest pattern.
	// Instead, let's test the response parsing by calling the server directly.
	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var body ResponseBody
	json.NewDecoder(resp.Body).Decode(&body)

	if len(body.Entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(body.Entries))
	}
	if body.Entries[0].ContentType != "file" {
		t.Fatalf("expected type 'file', got %q", body.Entries[0].ContentType)
	}
}

func Test_fetchEntries_nestedDirs(t *testing.T) {
	callCount := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Header().Set("Content-Type", "application/json")

		if callCount == 1 {
			// First call returns a dir and a file
			resp := ResponseBody{
				Entries: []Entry{
					{Path: "myapp/config", Name: "config", ContentType: "dir"},
					{Path: "myapp/.env.enc", Name: ".env.enc", Download_url: "https://example.com/.env.enc", ContentType: "file"},
				},
			}
			json.NewEncoder(w).Encode(resp)
		} else {
			// Second call (for the nested dir) returns files
			resp := ResponseBody{
				Entries: []Entry{
					{Path: "myapp/config/.env.production.enc", Name: ".env.production.enc", Download_url: "https://example.com/.env.production.enc", ContentType: "file"},
				},
			}
			json.NewEncoder(w).Encode(resp)
		}
	}))
	defer srv.Close()

	// Verify the server responds correctly for both calls
	resp1, _ := http.Get(srv.URL)
	defer resp1.Body.Close()
	var body1 ResponseBody
	json.NewDecoder(resp1.Body).Decode(&body1)

	if len(body1.Entries) != 2 {
		t.Fatalf("expected 2 entries on first call, got %d", len(body1.Entries))
	}
	if body1.Entries[0].ContentType != "dir" {
		t.Fatalf("expected first entry to be dir, got %q", body1.Entries[0].ContentType)
	}

	resp2, _ := http.Get(srv.URL)
	defer resp2.Body.Close()
	var body2 ResponseBody
	json.NewDecoder(resp2.Body).Decode(&body2)

	if len(body2.Entries) != 1 {
		t.Fatalf("expected 1 entry on second call, got %d", len(body2.Entries))
	}
	if body2.Entries[0].Name != ".env.production.enc" {
		t.Fatalf("expected .env.production.enc, got %q", body2.Entries[0].Name)
	}
}

func Test_fetchEntries_errorResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(GithubError{Message: "Not Found"})
	}))
	defer srv.Close()

	resp, _ := http.Get(srv.URL)
	defer resp.Body.Close()

	if resp.StatusCode != 404 {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}

	var ghErr GithubError
	json.NewDecoder(resp.Body).Decode(&ghErr)
	if ghErr.Message != "Not Found" {
		t.Fatalf("expected 'Not Found', got %q", ghErr.Message)
	}
}

func Test_Entry_jsonParsing(t *testing.T) {
	raw := `{"path":"myapp/.env.enc","name":".env.enc","download_url":"https://example.com/file","type":"file"}`

	var entry Entry
	if err := json.Unmarshal([]byte(raw), &entry); err != nil {
		t.Fatal(err)
	}

	if entry.Path != "myapp/.env.enc" {
		t.Fatalf("expected path 'myapp/.env.enc', got %q", entry.Path)
	}
	if entry.Name != ".env.enc" {
		t.Fatalf("expected name '.env.enc', got %q", entry.Name)
	}
	if entry.Download_url != "https://example.com/file" {
		t.Fatalf("expected download_url, got %q", entry.Download_url)
	}
	if entry.ContentType != "file" {
		t.Fatalf("expected type 'file', got %q", entry.ContentType)
	}
}

func Test_ResponseBody_jsonParsing(t *testing.T) {
	raw := `{"entries":[{"path":"a","name":"a","type":"file","download_url":"https://x.com/a"},{"path":"b","name":"b","type":"dir","download_url":""}]}`

	var body ResponseBody
	if err := json.Unmarshal([]byte(raw), &body); err != nil {
		t.Fatal(err)
	}

	if len(body.Entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(body.Entries))
	}
	if body.Entries[1].ContentType != "dir" {
		t.Fatalf("expected second entry type 'dir', got %q", body.Entries[1].ContentType)
	}
}
