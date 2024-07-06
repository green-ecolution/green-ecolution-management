package tree

type PredictedHealthResponse = string //@Name PredictedHealth

const (
	HealthGood     PredictedHealthResponse = "good"
	HealthModerate PredictedHealthResponse = "moderate"
	HealthBad      PredictedHealthResponse = "bad"
)

type SensorPredictionResponse struct {
	SensorID string                  `json:"sensor_id"`
	Tree     *TreeResponse           `json:"tree"`
	Health   PredictedHealthResponse `json:"health"`
	Humidity int                     `json:"humidity"`
} //@Name SensorPrediction
