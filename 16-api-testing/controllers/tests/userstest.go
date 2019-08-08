package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//Users : struct for set Users Dependency Injection
type Users struct {
	App http.Handler
}

//List : http handler for returning list of users
func (u *Users) List(t *testing.T) {
	req := httptest.NewRequest("GET", "/users", nil)
	resp := httptest.NewRecorder()

	u.App.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("getting: expected status code %v, got %v", http.StatusOK, resp.Code)
	}

	var list []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		t.Fatalf("decoding: %s", err)
	}

	want := []map[string]interface{}{
		{
			"id":        float64(1),
			"username":  "jackyhtg",
			"is_active": true,
		},
	}

	if diff := cmp.Diff(want, list); diff != "" {
		t.Fatalf("Response did not match expected. Diff:\n%s", diff)
	}
}

//View : http handler for retrieve user by id
func (u *Users) View(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/1", nil)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	u.App.ServeHTTP(resp, req)

	if http.StatusOK != resp.Code {
		t.Fatalf("retrieving: expected status code %v, got %v", http.StatusOK, resp.Code)
	}

	var fetched map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&fetched); err != nil {
		t.Fatalf("decoding: %s", err)
	}

	want := map[string]interface{}{
		"id":        float64(1),
		"username":  "jackyhtg",
		"is_active": true,
	}

	// Fetched product should match the one we created.
	if diff := cmp.Diff(want, fetched); diff != "" {
		t.Fatalf("Retrieved user should match created. Diff:\n%s", diff)
	}
}

//Create : http handler for create new user
func (u *Users) Create(t *testing.T) {
	var created map[string]interface{}
	jsonBody := `
		{
			"username": "peterpan",
			"email": "peterpan@gmail.com", 
			"password": "1234", 
			"re_password": "1234", 
			"is_active": true
		}
	`
	body := strings.NewReader(jsonBody)

	req := httptest.NewRequest("POST", "/users", body)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	u.App.ServeHTTP(resp, req)

	if http.StatusCreated != resp.Code {
		t.Fatalf("posting: expected status code %v, got %v", http.StatusCreated, resp.Code)
	}

	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		t.Fatalf("decoding: %s", err)
	}

	if created["id"] == "" || created["id"] == nil {
		t.Fatal("expected non-empty product id")
	}

	want := map[string]interface{}{
		"id":        created["id"],
		"username":  "peterpan",
		"is_active": true,
	}

	if diff := cmp.Diff(want, created); diff != "" {
		t.Fatalf("Response did not match expected. Diff:\n%s", diff)
	}
}

//Update : http handler for update user by id
func (u *Users) Update(t *testing.T) {
	var updated map[string]interface{}
	jsonBody := `
		{
			"id": 1,
			"username": "gatholoco",
			"email": "gatholoco@gmail.com", 
			"password": "1234", 
			"re_password": "1234", 
			"is_active": false
		}
	`
	body := strings.NewReader(jsonBody)

	req := httptest.NewRequest("PUT", "/users/1", body)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	u.App.ServeHTTP(resp, req)

	if http.StatusOK != resp.Code {
		t.Fatalf("posting: expected status code %v, got %v", http.StatusOK, resp.Code)
	}

	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		t.Fatalf("decoding: %s", err)
	}

	want := map[string]interface{}{
		"id":        float64(1),
		"username":  "gatholoco",
		"is_active": false,
	}

	if diff := cmp.Diff(want, updated); diff != "" {
		t.Fatalf("Response did not match expected. Diff:\n%s", diff)
	}
}

//Delete : http handler for delete user by id
func (u *Users) Delete(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/users/1", nil)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	u.App.ServeHTTP(resp, req)

	if http.StatusNoContent != resp.Code {
		t.Fatalf("retrieving: expected status code %v, got %v", http.StatusNoContent, resp.Code)
	}

	var deleted map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&deleted); err != nil {
		t.Fatalf("decoding: %s", err)
	}

	var want map[string]interface{}

	// Fetched product should match the one we created.
	if diff := cmp.Diff(want, deleted); diff != "" {
		t.Fatalf("Response did not match expected. Diff:\n%s", diff)
	}
}
