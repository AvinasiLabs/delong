package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const testSecret = "test-secret-key-123"

// createTestToken creates a test JWT token with HS256 algorithm
func createTestToken(username string, expiresAt time.Time) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Use HS256 algorithm explicitly
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(testSecret))
}

// setupTestRouter creates a test router with JWT middleware
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	jwtMiddleware := NewJwtMiddleware(testSecret)

	// Protected route that requires JWT authentication
	protected := router.Group("/api")
	protected.Use(jwtMiddleware.Auth())
	protected.GET("/test", func(c *gin.Context) {
		username, _ := GetUsername(c)
		c.JSON(200, gin.H{
			"message":  "success",
			"username": username,
		})
	})

	return router
}

// TestJwtMiddleware_ValidToken tests JWT middleware with a valid HS256 token
func TestJwtMiddleware_ValidToken(t *testing.T) {
	router := setupTestRouter()

	// Create a valid token that expires in 1 hour
	token, err := createTestToken("testuser", time.Now().Add(time.Hour))
	if err != nil {
		t.Fatalf("Failed to create test token: %v", err)
	}

	// Create request with valid token
	req, _ := http.NewRequest("GET", "/api/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// Execute request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verify response
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Check if response contains the username
	body := w.Body.String()
	if !containsString(body, "testuser") {
		t.Errorf("Expected response to contain 'testuser', got %s", body)
	}
}

// TestJwtMiddleware_NoToken tests JWT middleware without token
func TestJwtMiddleware_NoToken(t *testing.T) {
	router := setupTestRouter()

	// Create request without Authorization header
	req, _ := http.NewRequest("GET", "/api/test", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should return unauthorized error
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	body := w.Body.String()
	if !containsString(body, "UNAUTHORIZED") {
		t.Errorf("Expected response to contain 'UNAUTHORIZED', got %s", body)
	}
}

// TestJwtMiddleware_InvalidToken tests JWT middleware with invalid token
func TestJwtMiddleware_InvalidToken(t *testing.T) {
	router := setupTestRouter()

	// Create request with invalid token
	req, _ := http.NewRequest("GET", "/api/test", nil)
	req.Header.Set("Authorization", "Bearer invalid-token-string")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should return unauthorized error
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	body := w.Body.String()
	if !containsString(body, "UNAUTHORIZED") {
		t.Errorf("Expected response to contain 'UNAUTHORIZED', got %s", body)
	}
}

// TestJwtMiddleware_ExpiredToken tests JWT middleware with expired token
func TestJwtMiddleware_ExpiredToken(t *testing.T) {
	router := setupTestRouter()

	// Create expired token (expired 1 hour ago)
	token, err := createTestToken("testuser", time.Now().Add(-time.Hour))
	if err != nil {
		t.Fatalf("Failed to create expired test token: %v", err)
	}

	req, _ := http.NewRequest("GET", "/api/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should return unauthorized error for expired token
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	body := w.Body.String()
	if !containsString(body, "UNAUTHORIZED") {
		t.Errorf("Expected response to contain 'UNAUTHORIZED', got %s", body)
	}
}

// TestJwtMiddleware_MissingBearerPrefix tests JWT middleware without Bearer prefix
func TestJwtMiddleware_MissingBearerPrefix(t *testing.T) {
	router := setupTestRouter()

	token, err := createTestToken("testuser", time.Now().Add(time.Hour))
	if err != nil {
		t.Fatalf("Failed to create test token: %v", err)
	}

	// Set Authorization header without "Bearer " prefix
	req, _ := http.NewRequest("GET", "/api/test", nil)
	req.Header.Set("Authorization", token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should return unauthorized error
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	body := w.Body.String()
	if !containsString(body, "UNAUTHORIZED") {
		t.Errorf("Expected response to contain 'UNAUTHORIZED', got %s", body)
	}
}

// Helper function to check if string contains substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && findSubstring(s, substr)
}

// Helper function to find substring without using strings package
func findSubstring(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}

	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			if s[i+j] != substr[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
