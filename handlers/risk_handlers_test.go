package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/dongdongjssy/risk-api/constants"
	"github.com/dongdongjssy/risk-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var server *gin.Engine

func setupServer() {
	server = gin.Default()
	server.GET(constants.ENDPOINT_GET_RISKS, GetRisks)
	server.GET(constants.ENDPOINT_GET_RISK, GetRisk)
	server.POST(constants.ENDPOINT_POST_RISKS, CreateRisk)
}

func TestGetRisksHandler(t *testing.T) {
	setupServer()

	t.Run("it gets an empty list of risk when server initial runs", func(t *testing.T) {
		// send a GET risks request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, constants.ENDPOINT_GET_RISKS, nil)
		server.ServeHTTP(rec, req)

		// verify response code
		assert.Equal(t, http.StatusOK, rec.Code, "should response 200 code")
		assert.Equal(t, "[]", rec.Body.String(), "should returns an empty list")
	})

	t.Run("it gets a list of risks after inserting risks", func(t *testing.T) {
		// insert 10 risks
		for i := range 10 {
			// build a valid risk
			s := fmt.Sprintf("%s%d", "Risk", i)
			expectedRisk := models.Risk{
				State:       "open",
				Title:       s,
				Description: s,
			}

			// send a POST risks request
			riskJson, _ := json.Marshal(expectedRisk)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, constants.ENDPOINT_POST_RISKS, strings.NewReader(string(riskJson)))
			server.ServeHTTP(rec, req)

			// verify response code
			assert.Equal(t, http.StatusCreated, rec.Code, "should response 201 code")
		}

		// send a GET risks request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, constants.ENDPOINT_GET_RISKS, nil)
		server.ServeHTTP(rec, req)

		// verify response code
		assert.Equal(t, http.StatusOK, rec.Code, "should response 200 code")

		// verify response body
		body, _ := io.ReadAll(rec.Result().Body)
		defer rec.Result().Body.Close()

		var risks []models.Risk
		json.Unmarshal(body, &risks)

		assert.Equal(t, 10, len(risks), "should returns a list of 10 risks")
	})
}

func TestPostRisksHandler(t *testing.T) {
	setupServer()

	t.Run("it creates a new risk with valid risk data", func(t *testing.T) {
		// build a valid risk
		expectedRisk := models.Risk{
			State:       "open",
			Title:       "Risk A",
			Description: "Risk A",
		}
		riskJson, _ := json.Marshal(expectedRisk)

		// send a POST risks request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, constants.ENDPOINT_POST_RISKS, strings.NewReader(string(riskJson)))
		server.ServeHTTP(rec, req)

		// verify response code
		assert.Equal(t, http.StatusCreated, rec.Code, "should response 201 code")

		body, _ := io.ReadAll(rec.Result().Body)
		defer rec.Result().Body.Close()

		var actualRisk models.Risk
		json.Unmarshal(body, &actualRisk)

		// verify response body
		expectedRisk.ID = actualRisk.ID
		assert.True(t, reflect.DeepEqual(expectedRisk, actualRisk), "Unexpected risk returned")
	})

	t.Run("it returns 400 with invalid risk state", func(t *testing.T) {
		// build a valid risk
		expectedRisk := models.Risk{
			State:       "invalid_state",
			Title:       "Risk A",
			Description: "Risk A",
		}
		riskJson, _ := json.Marshal(expectedRisk)

		// send a POST risks request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, constants.ENDPOINT_POST_RISKS, strings.NewReader(string(riskJson)))
		server.ServeHTTP(rec, req)

		// verify response code
		assert.Equal(t, http.StatusBadRequest, rec.Code, "should response 400 code")
		assert.Equal(
			t,
			"\"Risk.State value must be one of [open closed accepted investigating]\"",
			rec.Body.String(),
			"incorrect error message",
		)
	})

	t.Run("it returns 400 with missing risk title", func(t *testing.T) {
		// build a valid risk
		expectedRisk := models.Risk{
			State:       "accepted",
			Description: "Risk A",
		}
		riskJson, _ := json.Marshal(expectedRisk)

		// send a POST risks request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, constants.ENDPOINT_POST_RISKS, strings.NewReader(string(riskJson)))
		server.ServeHTTP(rec, req)

		// verify response code
		assert.Equal(t, http.StatusBadRequest, rec.Code, "should response 400 code")
		assert.Equal(
			t,
			"\"Risk.Title is required\"",
			rec.Body.String(),
			"incorrect error message",
		)
	})

	t.Run("it returns 400 with missing risk state", func(t *testing.T) {
		// build a valid risk
		expectedRisk := models.Risk{
			Title:       "Risk B",
			Description: "Risk A",
		}
		riskJson, _ := json.Marshal(expectedRisk)

		// send a POST risks request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, constants.ENDPOINT_POST_RISKS, strings.NewReader(string(riskJson)))
		server.ServeHTTP(rec, req)

		// verify response code
		assert.Equal(t, http.StatusBadRequest, rec.Code, "should response 400 code")
		assert.Equal(
			t,
			"\"Risk.State is required\"",
			rec.Body.String(),
			"incorrect error message",
		)
	})

	t.Run("it returns 400 with duplicated risk title", func(t *testing.T) {
		// build a valid risk
		expectedRisk := models.Risk{
			State:       "accepted",
			Title:       "Risk A",
			Description: "Risk A",
		}
		riskJson, _ := json.Marshal(expectedRisk)

		// send a POST risks request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, constants.ENDPOINT_POST_RISKS, strings.NewReader(string(riskJson)))
		server.ServeHTTP(rec, req)

		// verify response code
		assert.Equal(t, http.StatusBadRequest, rec.Code, "should response 400 code")
		assert.Equal(
			t,
			"\"duplicate risk title, please choose another title\"",
			rec.Body.String(),
			"incorrect error message",
		)
	})

	t.Run("it returns 400 with risk title exceeds max length of 128", func(t *testing.T) {
		// build a valid risk
		var longTitle = make([]byte, 129)
		for i := range 129 {
			longTitle[i] = 'a'
		}

		expectedRisk := models.Risk{
			State:       "accepted",
			Title:       string(longTitle),
			Description: "Risk A",
		}
		riskJson, _ := json.Marshal(expectedRisk)

		// send a POST risks request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, constants.ENDPOINT_POST_RISKS, strings.NewReader(string(riskJson)))
		server.ServeHTTP(rec, req)

		// verify response code
		assert.Equal(t, http.StatusBadRequest, rec.Code, "should response 400 code")
		assert.Equal(
			t,
			"\"Risk.Title exceeds the max length of 128\"",
			rec.Body.String(),
			"incorrect error message",
		)
	})
}

func TestGetRiskByIdHandler(t *testing.T) {
	setupServer()

	t.Run("it returns a risk by its id", func(t *testing.T) {
		// 1. first insert a risk
		expectedRisk := models.Risk{
			State:       "open",
			Title:       "Risk B",
			Description: "Risk B",
		}
		riskJson, _ := json.Marshal(expectedRisk)

		postRec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, constants.ENDPOINT_POST_RISKS, strings.NewReader(string(riskJson)))
		server.ServeHTTP(postRec, req)

		// verify response code
		assert.Equal(t, http.StatusCreated, postRec.Code, "should response 201 code")

		// get new inserted risk id
		body, _ := io.ReadAll(postRec.Result().Body)
		defer postRec.Result().Body.Close()

		var newInsertedRisk models.Risk
		json.Unmarshal(body, &newInsertedRisk)
		newId := newInsertedRisk.ID

		// 2. then send a GET risk by id request
		getRec := httptest.NewRecorder()
		endpointWithId := strings.Replace(constants.ENDPOINT_GET_RISK, ":id", newId, -1)
		req, _ = http.NewRequest(http.MethodGet, endpointWithId, nil)
		server.ServeHTTP(getRec, req)

		// verify response code
		assert.Equal(t, http.StatusOK, getRec.Code, "should response 200 code")

		// verify response body
		body, _ = io.ReadAll(getRec.Result().Body)
		defer getRec.Result().Body.Close()

		var risk models.Risk
		json.Unmarshal(body, &risk)

		assert.True(t, reflect.DeepEqual(newInsertedRisk, risk), "should return the new inserted risk")
	})

	t.Run("it returns 404 if no risk found by id", func(t *testing.T) {
		// send a GET risk by a random id request
		rec := httptest.NewRecorder()
		endpointWithId := strings.Replace(constants.ENDPOINT_GET_RISK, ":id", uuid.New().String(), -1)
		req, _ := http.NewRequest(http.MethodGet, endpointWithId, nil)
		server.ServeHTTP(rec, req)

		// verify response code
		assert.Equal(t, http.StatusNotFound, rec.Code, "should response 404 code")
	})
}
