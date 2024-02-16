package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	Routes(r)
	return r
}

func TestCreditAssignmentRouteValidInvestment(t *testing.T) {
	router := setupRouter()
	assert := assert.New(t)

	jsonBody := []byte(`{"investment": 6700}`)
	bodyReader := bytes.NewReader(jsonBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/credit-assignment", bodyReader)
	router.ServeHTTP(w, req)
	assert.Equal(200, w.Code)

	var response map[string]int32
	err := json.Unmarshal([]byte(w.Body.Bytes()), &response)

	credit1, exists1 := response["credit_type_300"]
	credit2, exists2 := response["credit_type_500"]
	credit3, exists3 := response["credit_type_700"]

	var ex1, ex2, ex3 int32 = 2, 1, 8

	assert.Nil(err)
	assert.True(exists1)
	assert.True(exists2)
	assert.True(exists3)
	assert.Equal(credit1, ex1)
	assert.Equal(credit2, ex2)
	assert.Equal(credit3, ex3)
}

func TestCreditAssignmentRouteNotValidInvestment(t *testing.T) {
	router := setupRouter()
	assert := assert.New(t)

	jsonBody := []byte(`{"investment": 400}`)
	bodyReader := bytes.NewReader(jsonBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/credit-assignment", bodyReader)
	router.ServeHTTP(w, req)
	assert.Equal(400, w.Code)
}
