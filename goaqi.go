package goaqi

import (
	"errors"
	"math"
)

type Breakpoint struct {
	Low  float64
	High float64
}

var (
	// PM25Breakpoints are the breakpoints for PM2.5 in µg/m3
	PM25Breakpoints = []Breakpoint{
		{0, 12.0},      // Good
		{12.1, 35.4},   // Moderate
		{35.5, 55.4},   // Unhealthy for Sensitive Groups
		{55.5, 150.4},  // Unhealthy
		{150.5, 250.4}, // Very Unhealthy
		{250.5, 350.4}, // Hazardous
		{350.5, 500.4}, // Hazardous
	}
	// PM10Breakpoints are the breakpoints for PM10 in µg/m3
	PM100Breakpoints = []Breakpoint{
		{0, 54.9},      // Good
		{55.0, 154.9},  // Moderate
		{155.0, 254.9}, // Unhealthy for Sensitive Groups
		{255.0, 354.9}, // Unhealthy
		{355.0, 424.9}, // Very Unhealthy
		{425.0, 504.9}, // Hazardous
		{505.0, 604.9}, // Hazardous
	}
	// AQIBreakpoints are the breakpoints for AQI
	AQIBreakpoints = []Breakpoint{
		{0, 50},    // Good
		{51, 100},  // Moderate
		{101, 150}, // Unhealthy for Sensitive Groups
		{151, 200}, // Unhealthy
		{201, 300}, // Very Unhealthy
		{301, 400}, // Hazardous
		{401, 500}, // Hazardous
	}

	// AQI Name Designations
	AQIDesignations = []string{
		"Good",
		"Moderate",
		"Unhealthy for Sensitive Groups",
		"Unhealthy",
		"Very Unhealthy",
		"Hazardous",
		"Hazardous",
	}

	ErrBeyondTheScale = errors.New("beyond the scale")
)

// AQIPM25 calculates the AQI for PM2.5
//
// Requires a 24 hour average of PM2.5 concentration in µg/m3
//
// Please note that the truncation step is not performed in this function
func AQIPM25(avg float64) (int64, error) {
	return AQI(avg, PM25Breakpoints)
}

// AQI calculates the AQI for a given set of breakpoints
//
// Requires a 24 hour average of concentration in µg/m3 and a set of breakpoints
func AQI(avg float64, breakpoints []Breakpoint) (int64, error) {
	for i, bp := range breakpoints {
		if avg >= bp.Low && avg <= bp.High {
			return AQIForBreakpoint(avg, bp, AQIBreakpoints[i])
		}
	}
	return 0, ErrBeyondTheScale
}

// AQIForBreakpoint calculates the AQI from the two breakpoints and the average concentration
//
// Requires the average concentration, the breakpoint for the average concentration, and the AQI breakpoint
func AQIForBreakpoint(avg float64, bp Breakpoint, aqiBP Breakpoint) (int64, error) {
	return int64(math.Round((aqiBP.High-aqiBP.Low)/(bp.High-bp.Low)*(avg-bp.Low) + aqiBP.Low)), nil
}

func AQIDesignationFromIndex(aqi int64) (string, error) {
	for i, bp := range AQIBreakpoints {
		if aqi >= int64(bp.Low) && aqi <= int64(bp.High) {
			return AQIDesignations[i], nil
		}
	}
	return "", ErrBeyondTheScale
}
