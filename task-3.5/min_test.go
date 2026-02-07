package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloHandlerStatusCode проверяет статус код 200 OK
func TestHelloHandlerStatusCode(t *testing.T) {
	// Создаем тестовый запрос
	req := httptest.NewRequest("GET", "/hello", nil)

	rr := httptest.NewRecorder()
	helloHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}

// TestHelloHandlerContentType проверяет заголовок Content-Type
func TestHelloHandlerContentType(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	// Проверяем заголовок Content-Type
	contentType := rr.Header().Get("Content-Type")
	expectedContentType := "application/json"

	if contentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, contentType)
	}
}

// TestHelloHandlerJSONBody проверяет тело JSON ответа
func TestHelloHandlerJSONBody(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)

	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Проверяем наличие поля message
	expectedMessage := "hello, world!"
	if message, ok := response["message"]; !ok {
		t.Error("Response JSON missing 'message' field")
	} else if message != expectedMessage {
		t.Errorf("Expected message '%s', got '%s'", expectedMessage, message)
	}
}

// TestHelloHandlerWrongMethod проверяет обработку неверного метода
func TestHelloHandlerWrongMethod(t *testing.T) {
	req := httptest.NewRequest("POST", "/hello", nil) // POST вместо GET
	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d for POST, got %d",
			http.StatusMethodNotAllowed, rr.Code) // Должен вернуть 405 Method Not Allowed
	}
}

// TestHelloHandlerIntegration интеграционный тест через httptest.Server
func TestHelloHandlerIntegration(t *testing.T) {
	handler := http.HandlerFunc(helloHandler)
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL + "/hello")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Проверяем Content-Type
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	// Читаем и проверяем тело
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var response map[string]string
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if response["message"] != "hello, world!" {
		t.Errorf("Expected message 'hello, world!', got %s", response["message"])
	}
}

// TestHelloHandlerResponseStructure проверяет структуру JSON ответа
func TestHelloHandlerResponseStructure(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	// Проверяем, что ответ - валидный JSON
	var data interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &data)
	if err != nil {
		t.Fatalf("Response is not valid JSON: %v", err)
	}

	// Проверяем, что это объект (map)
	m, ok := data.(map[string]interface{})
	if !ok {
		t.Fatal("JSON response is not an object")
	}

	// Проверяем наличие только одного поля "message"
	if len(m) != 1 {
		t.Errorf("Expected 1 field in JSON, got %d", len(m))
	}

	// Проверяем, что поле "message" существует и это строка
	message, exists := m["message"]
	if !exists {
		t.Error("Missing 'message' field in JSON")
	}

	_, isString := message.(string)
	if !isString {
		t.Error("'message' field is not a string")
	}
}

// BenchmarkHelloHandler бенчмарк производительности хендлера
func BenchmarkHelloHandler(b *testing.B) {
	req := httptest.NewRequest("GET", "/hello", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		helloHandler(rr, req)
	}
}
