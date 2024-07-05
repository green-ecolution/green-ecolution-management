package tree


type TreeSensorPredictionResponse struct {
	Tree             *TreeResponse             `json:"tree,omitempty"`
	SensorPrediction *SensorPredictionResponse `json:"sensor_prediction,omitempty"`
	SensorData       []MqttDataResponse      `json:"sensor_data,omitempty"`
} //@Name TreeSensorPredictionResponse

type TreeSensorDataResponse struct {
	Tree       *TreeResponse        `json:"tree,omitempty"`
	SensorData []MqttDataResponse `json:"sensor_data,omitempty"`
} //@Name TreeSensorDataResponse
