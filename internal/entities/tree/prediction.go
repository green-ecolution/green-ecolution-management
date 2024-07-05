package tree

type PredictedHealth = string

const (
	HealthGood     PredictedHealth = "good"
	HealthModerate PredictedHealth = "moderate"
	HealthBad      PredictedHealth = "bad"
)

type SensorPrediction struct {
	SensorID string
	Tree     *Tree
	Health   PredictedHealth
	Humidity int
}
