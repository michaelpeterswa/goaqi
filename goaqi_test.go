package goaqi_test

import (
	"testing"

	"676f.dev/goaqi"
	"github.com/stretchr/testify/assert"
)

func TestPM25AQICalculation(t *testing.T) {
	tests := []struct {
		Name        string
		Average     float64
		ExpectedAQI int64
	}{
		{
			Name:        "Good",
			Average:     10.5,
			ExpectedAQI: 44,
		},
		{
			Name:        "Moderate",
			Average:     30.2,
			ExpectedAQI: 89,
		},
		{
			Name:        "Unhealthy for Sensitive Groups",
			Average:     51.4,
			ExpectedAQI: 140,
		},
		{
			Name:        "Unhealthy",
			Average:     93.7,
			ExpectedAQI: 171,
		},
		{
			Name:        "Very Unhealthy",
			Average:     155.3,
			ExpectedAQI: 206,
		},
		{
			Name:        "Hazardous 1",
			Average:     255.6,
			ExpectedAQI: 306,
		},
		{
			Name:        "Hazardous 2",
			Average:     355.6,
			ExpectedAQI: 404,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := goaqi.AQIPM25(tc.Average)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			assert.Equal(t, tc.ExpectedAQI, result)
		})
	}
}

func TestPM25AQICalculationError(t *testing.T) {
	tests := []struct {
		Name    string
		Average float64
	}{
		{
			Name:    "Negative",
			Average: -10.5,
		},
		{
			Name:    "Super Hazardous",
			Average: 1000.5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			_, err := goaqi.AQIPM25(tc.Average)
			if assert.Error(t, err) {
				assert.Equal(t, "beyond the scale", err.Error())
			}
		})
	}
}

func TestAQIDesignationFromIndex(t *testing.T) {
	tests := []struct {
		Name                   string
		Index                  int64
		ExpectedAQIDesignation string
	}{
		{
			Name:                   "Good",
			Index:                  44,
			ExpectedAQIDesignation: "Good",
		},
		{
			Name:                   "Moderate",
			Index:                  89,
			ExpectedAQIDesignation: "Moderate",
		},
		{
			Name:                   "Unhealthy for Sensitive Groups",
			Index:                  140,
			ExpectedAQIDesignation: "Unhealthy for Sensitive Groups",
		},
		{
			Name:                   "Unhealthy",
			Index:                  171,
			ExpectedAQIDesignation: "Unhealthy",
		},
		{
			Name:                   "Very Unhealthy",
			Index:                  206,
			ExpectedAQIDesignation: "Very Unhealthy",
		},
		{
			Name:                   "Hazardous 1",
			Index:                  306,
			ExpectedAQIDesignation: "Hazardous",
		},
		{
			Name:                   "Hazardous 2",
			Index:                  404,
			ExpectedAQIDesignation: "Hazardous",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := goaqi.AQIDesignationFromIndex(tc.Index)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			assert.Equal(t, tc.ExpectedAQIDesignation, result)
		})
	}
}

func TestAQIDesignationFromIndexError(t *testing.T) {
	tests := []struct {
		Name  string
		Index int64
	}{
		{
			Name:  "Negative",
			Index: -10,
		},
		{
			Name:  "Super Hazardous",
			Index: 1000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			_, err := goaqi.AQIDesignationFromIndex(tc.Index)
			if assert.Error(t, err) {
				assert.Equal(t, "beyond the scale", err.Error())
			}
		})
	}
}
