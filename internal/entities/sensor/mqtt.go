package sensor

import (
	"time"
)

type MqttIdentifierApplicationID struct {
	ApplicationID string
}

type MqttIdentifierDeviceID struct {
	DeviceID       string
	ApplicationIDs MqttIdentifierApplicationID
	DevEUI         string
	JoinEUI        string
	DevAddr        string
}

type MqttIdentifier struct {
	DeviceIDs MqttIdentifierDeviceID
}

type MqttDecodedPayload struct {
	Battery  float64
	Humidity int
	Raw      int
}

type MqttRxMetadataGatewayIDs struct {
	GatewayID string
}

type MqttRxMetadataPacketBroker struct {
	MessageID            string
	ForwarderNetID       string
	ForwarderTenantID    string
	ForwarderClusterID   string
	ForwarderGatewayID   string
	ForwarderGatewayEUI  string
	HomeNetworkNetID     string
	HomeNetworkTenantID  string
	HomeNetworkClusterID string
}

type MqttLocation struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

type MqttRxMetadata struct {
	GatewayIDs   MqttRxMetadataGatewayIDs
	PacketBroker MqttRxMetadataPacketBroker
	Time         *time.Time
	Rssi         int
	ChannelRssi  int
	Snr          float64
	Location     MqttLocation
	UplinkToken  string
	RecievedAt   *time.Time
}

type MqttUplinkSettingsLora struct {
	Bandwidth       int
	SpreadingFactor int
	CodingRate      string
}

type MqttUplinkSettingsDataRate struct {
	Lora MqttUplinkSettingsLora
}

type MqttUplinkSettings struct {
	DataRate  MqttUplinkSettingsDataRate
	Frequency string
}

type MqttVersionIDs struct {
	BrandID         string
	ModelID         string
	HardwareVersion string
	FirmwareVersion string
	BandID          string
}

type MqttNetworkIDs struct {
	NetID          string
	NSID           string
	TenantID       string
	ClusterID      string
	ClusterAddress string
	TenantAddress  string
}

type MqttUplinkMessage struct {
	SessionKeyID    string
	FPort           int
	Fcnt            int
	FRMPayload      string
	DecodedPayload  MqttDecodedPayload
	RxMetadata      []MqttRxMetadata
	Settings        MqttUplinkSettings
	ReceivedAt      *time.Time
	Confirmed       bool
	ConsumedAirtime string
	VersionIDs      MqttVersionIDs
	NetworkIDs      MqttNetworkIDs
}

type MqttDataPayload struct {
	Type           string
	EndDeviceIDs   MqttIdentifierDeviceID
	CorrelationIDs []string
	ReceivedAt     *time.Time
	UplinkMessage  MqttUplinkMessage
}

type MqttVisibility struct {
	Rights []string
}

type MqttData struct {
	Name           string
	Time           *time.Time
	Identifiers    []MqttIdentifier
	Data           MqttDataPayload
	CorrelationIDs []string
	Origin         string
	Context        map[string]string
	Visibility     MqttVisibility
	UniqueID       string
}

type MqttPayload struct {
	EndDeviceIDs   MqttIdentifierDeviceID
	CorrelationIDs []string
	ReceivedAt     *time.Time
	UplinkMessage  MqttUplinkMessage
}

func (m *MqttPayload) GetHumidity() int {
  return m.UplinkMessage.DecodedPayload.Humidity
}

func (m *MqttPayload) GetBattery() float64 {
  return m.UplinkMessage.DecodedPayload.Battery
}
