package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hilton-james/FetchExercise/internal/core/entities"
)

func TestReceiptHandler_ProcessReceipt(t *testing.T) {
	handler := &ReceiptHandler{}

	tests := []struct {
		name           string
		payload        entities.Receipt
		expectedStatus int
		expectedPoints int
	}{
		{
			name: "Valid receipt - All rules apply",
			payload: entities.Receipt{
				ID:           "1",
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "14:30",
				Items: []entities.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Pepsi 12PK", Price: "5.49"},
				},
				Total: "6.75",
			},
			expectedStatus: http.StatusOK,
			expectedPoints: 71, // Based on rules and data
		},
		{
			name: "Invalid payload",
			payload: entities.Receipt{
				ID:           "2",
				Retailer:     "",
				PurchaseDate: "invalid-date",
				PurchaseTime: "invalid-time",
				Items:        nil,
				Total:        "not-a-number",
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedPoints: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Serialize payload
			payloadBytes, _ := json.Marshal(test.payload)
			req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(payloadBytes))
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.ProcessReceipt(rr, req)

			// Check status code
			if rr.Code != test.expectedStatus {
				t.Errorf("expected status %d but got %d", test.expectedStatus, rr.Code)
			}

			// If status is OK, verify points
			if test.expectedStatus == http.StatusOK {
				var response map[string]interface{}
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("failed to parse response body: %v", err)
				}

				points, ok := response["points"].(float64)
				if !ok {
					t.Errorf("expected 'points' field in response")
				}

				if int(points) != test.expectedPoints {
					t.Errorf("expected points %d but got %d", test.expectedPoints, int(points))
				}
			}
		})
	}
}
