package triangletree

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTreeSum_WithValidArray(t *testing.T) {
	input := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}

	expected := 237
	result, err := TreeSum(input)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestTreeSum_WithInvalidArray(t *testing.T) {
	invalidInput := [][]int{
		{1},
		{2, 3},
		{4, 5},
	}

	_, err := TreeSum(invalidInput)

	if err == nil {
		t.Error("expected error for invalid triangle, got nil")
	}
}

func TestTreeSum_WithValidURL(t *testing.T) {
	expected := 237

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[[59],[73,41],[52,40,53],[26,53,6,34]]`))
	}))

	defer ts.Close()

	result, err := TreeSum(ts.URL)

	if err != nil {
		t.Fatalf("คาดว่าไม่ควรมี error แต่พบว่า: %v", err)
	}

	if result != expected {
		t.Errorf("คาดว่าได้ %d แต่ได้ %d", expected, result)
	}
}

func TestTreeSum_WithInvalidURL(t *testing.T) {
	url := "https://example.com/nonexistent.json"

	_, err := TreeSum(url)

	if err == nil {
		t.Error("expected error from invalid URL, got nil")
	}
}

func TestTreeSum_WithUnsupportedInput(t *testing.T) {
	_, err := TreeSum(12345)

	if err == nil {
		t.Error("expected error for unsupported input type, got nil")
	}
}
