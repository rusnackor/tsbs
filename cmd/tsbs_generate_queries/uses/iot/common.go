package iot

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/timescale/tsbs/pkg/data/usecases/iot"

	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/uses/common"
	"github.com/timescale/tsbs/pkg/query"
)

const (
	// ReadingsTableName is the name of the table where all the readings
	// time series data is stored.
	ReadingsTableName = "readings"
	// DiagnosticsTableName is the name of the table where all the diagnostics
	// time series data is stored.
	DiagnosticsTableName = "diagnostics"

	// StationaryDuration is the time duration to evaluate stationary trucks.
	StationaryDuration = 10 * time.Minute
	// LongDrivingSessionDuration is the the time duration which is considered a
	// long driving session without stopping.
	LongDrivingSessionDuration = 4 * time.Hour
	// DailyDrivingDuration is time duration of one day of driving.
	DailyDrivingDuration = 24 * time.Hour
	//TEST intervals
	Test1HourDuration = 1 * time.Hour
	Test10MinDuration = 10 * time.Minute
	Test20MinDuration = 20 * time.Minute

	// LabelLastLoc is the label for the last location query.
	LabelLastLoc = "last-loc"
	// LabelLastLocSingleTruck is the label for the last location query for a single truck.
	LabelLastLocSingleTruck = "single-last-loc"
	// LabelLowFuel is the label for the low fuel query.
	LabelLowFuel = "low-fuel"
	// LabelHighLoad is the label for the high load query.
	LabelHighLoad = "high-load"
	// LabelStationaryTrucks is the label for the stationary trucks query.
	LabelStationaryTrucks = "stationary-trucks"
	// LabelLongDrivingSessions is the label for the long driving sessions query.
	LabelLongDrivingSessions = "long-driving-sessions"
	// LabelLongDailySessions is the label for the long daily sessions query.
	LabelLongDailySessions = "long-daily-sessions"
	// LabelAvgVsProjectedFuelConsumption is the label for the avg vs projected fuel consumption query.
	LabelAvgVsProjectedFuelConsumption = "avg-vs-projected-fuel-consumption"
	// LabelAvgDailyDrivingDuration is the label for the avg daily driving durationquery.
	LabelAvgDailyDrivingDuration = "avg-daily-driving-duration"
	// LabelAvgDailyDrivingSession is the label for the avg daily driving session query.
	LabelAvgDailyDrivingSession = "avg-daily-driving-session"
	// LabelAvgLoad is the label for the avg load query.
	LabelAvgLoad = "avg-load"
	// LabelDailyActivity is the label for the daily activity query.
	LabelDailyActivity = "daily-activity"
	// LabelBreakdownFrequency is the label for the breakdown frequency query.
	LabelBreakdownFrequency = "breakdown-frequency"
	// Labels for TEST stuff
	LabelTestQuery           = "test-query"
	LabelExtract1H           = "extract-1h"
	LabelExtract10Min        = "extract-10min"
	LabelOlderThan           = "older-than"
	LabelExtract20MinWFilter = "20min-tags-extract"
	LabelStats10Min          = "stats-10min"
	LabelStats24H            = "stats-24h"
	LabelExtract1Tag         = "extract-1tag"
)

// Core is the common component of all generators for all systems.
type Core struct {
	*common.Core
}

// GetRandomFleet returns one of the fleet choices by random.
func (c Core) GetRandomFleet() string {
	return iot.FleetChoices[rand.Intn(len(iot.FleetChoices))]
}

// NewCore returns a new Core for the given time range and cardinality
func NewCore(start, end time.Time, scale int) (*Core, error) {
	c, err := common.NewCore(start, end, scale)
	return &Core{Core: c}, err

}

// GetRandomTrucks returns a random set of nTrucks from a given Core
func (c *Core) GetRandomTrucks(nTrucks int) ([]string, error) {
	return getRandomTrucks(nTrucks, c.Scale)
}

// getRandomTruckNames returns a subset of numTrucks names of a permutation of truck names,
// numbered from 0 to totalTrucks.
// Ex.: truck_12, truck_7, truck_25 for numTrucks=3 and totalTrucks=30 (3 out of 30)
func getRandomTrucks(numTrucks int, totalTrucks int) ([]string, error) {
	if numTrucks < 1 {
		return nil, fmt.Errorf("number of trucks cannot be < 1; got %d", numTrucks)
	}
	if numTrucks > totalTrucks {
		return nil, fmt.Errorf("number of trucks (%d) larger than total trucks. See --scale (%d)", numTrucks, totalTrucks)
	}

	randomNumbers, err := common.GetRandomSubsetPerm(numTrucks, totalTrucks)
	if err != nil {
		return nil, err
	}

	truckNames := []string{}
	for _, n := range randomNumbers {
		truckNames = append(truckNames, fmt.Sprintf("truck_%d", n))
	}

	return truckNames, nil
}

// LastLocFiller is a type that can fill in a last location query.
type LastLocFiller interface {
	LastLocPerTruck(query.Query)
}

// LastLocByTruckFiller is a type that can fill in a last location query for a number of trucks.
type LastLocByTruckFiller interface {
	LastLocByTruck(query.Query, int)
}

// TruckLowFuelFiller is a type that can fill in a trucks with low fuel query.
type TruckLowFuelFiller interface {
	TrucksWithLowFuel(query.Query)
}

// TruckHighLoadFiller is a type that can fill in a trucks with high load query.
type TruckHighLoadFiller interface {
	TrucksWithHighLoad(query.Query)
}

// StationaryTrucksFiller is a type that can fill in the stationary trucks query.
type StationaryTrucksFiller interface {
	StationaryTrucks(query.Query)
}

// TruckLongDrivingSessionFiller is a type that can fill in a trucks with longer driving sessions query.
type TruckLongDrivingSessionFiller interface {
	TrucksWithLongDrivingSessions(query.Query)
}

// TruckLongDailySessionFiller is a type that can fill in a trucks with longer daily driving sessions query.
type TruckLongDailySessionFiller interface {
	TrucksWithLongDailySessions(query.Query)
}

// AvgVsProjectedFuelConsumptionFiller is a type that can fill in an avg vs projected fuel consumption query.
type AvgVsProjectedFuelConsumptionFiller interface {
	AvgVsProjectedFuelConsumption(query.Query)
}

// AvgDailyDrivingDurationFiller is a type that can fill in an avg daily driving duration per driver query.
type AvgDailyDrivingDurationFiller interface {
	AvgDailyDrivingDuration(query.Query)
}

// AvgDailyDrivingSessionFiller is a type that can fill in an avg daily driving session query.
type AvgDailyDrivingSessionFiller interface {
	AvgDailyDrivingSession(query.Query)
}

// AvgLoadFiller is a type that can fill in an avg load query.
type AvgLoadFiller interface {
	AvgLoad(query.Query)
}

// DailyTruckActivityFiller is a type that can fill in the daily truck activity query.
type DailyTruckActivityFiller interface {
	DailyTruckActivity(query.Query)
}

// TruckBreakdownFrequencyFiller is a type that can fill in the truck breakdown frequency query.
type TruckBreakdownFrequencyFiller interface {
	TruckBreakdownFrequency(query.Query)
}

// TestQueryFiller is a type that is PoC to adding new query type.
type TestQueryFiller interface {
	TestQuery(query.Query)
}

// Extract1HFiller is a type that can fill in the extract 1 hour query.
type Extract1HFiller interface {
	Extract1H(query.Query)
}

// Extract10MinFiller is a type that can fill in the extract 10 min query.
type Extract10MinFiller interface {
	Extract10Min(query.Query)
}

// OlderThanFiller is a type that can fill in the older than query.
type OlderThanFiller interface {
	OlderThan(query.Query)
}

// Extract20MinWfilterFiller is a type that can fill in the extract 20 min with filter query.
type Extract20MinWfilterFiller interface {
	Extract20MinWfilter(query.Query)
}

// Statistics10MinFiller is a type that can fill in the statistics 10 min query.
type Statistics10MinFiller interface {
	Statistics10Min(query.Query)
}

// Statistics24HFiller is a type that can fill in the statistics 24 hours query.
type Statistics24HFiller interface {
	Statistics24H(query.Query)
}

// Extract1TagFiller is a type that can fill in the extract 1 tag query.
type Extract1TagFiller interface {
	Extract1Tag(query.Query)
}
