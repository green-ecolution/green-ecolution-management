package sensor

import (
	"time"
)

type MqttIdentifierApplicationIDResponse struct {
	ApplicationID string `json:"application_id"`
} //@Name MqttIdentifierApplicationID

type MqttIdentifierDeviceIDResponse struct {
	DeviceID       string                              `json:"device_id"`
	ApplicationIDs MqttIdentifierApplicationIDResponse `json:"application_ids"`
	DevEUI         string                              `json:"dev_eui"`
	JoinEUI        string                              `json:"join_eui"`
	DevAddr        string                              `json:"dev_addr"`
} //@Name MqttIdentifierDeviceID

type MqttIdentifierResponse struct {
	DeviceIDs MqttIdentifierDeviceIDResponse `json:"device_ids"`
} //@Name MqttIdentifier

type MqttDecodedPayloadResponse struct {
	Battery  float64 `json:"battery"`
	Humidity int     `json:"humidity"`
	Raw      int     `json:"raw"`
} //@Name MqttDecodedPayload

type MqttRxMetadataGatewayIDsResponse struct {
	GatewayID string `json:"gateway_id"`
} //@Name MqttRxMetadataGatewayIDs

type MqttRxMetadataPacketBrokerResponse struct {
	MessageID            string `json:"message_id"`
	ForwarderNetID       string `json:"forwarder_net_id"`
	ForwarderTenantID    string `json:"forwarder_tenant_id"`
	ForwarderClusterID   string `json:"forwarder_cluster_id"`
	ForwarderGatewayID   string `json:"forwarder_gateway_id"`
	ForwarderGatewayEUI  string `json:"forwarder_gateway_eui"`
	HomeNetworkNetID     string `json:"home_network_net_id"`
	HomeNetworkTenantID  string `json:"home_network_tenant_id"`
	HomeNetworkClusterID string `json:"home_network_cluster_id"`
} //@Name MqttRxMetadataPacketBroker

type MqttLocationResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
} //@Name MqttLocation

type MqttRxMetadataResponse struct {
	GatewayIDs   MqttRxMetadataGatewayIDsResponse   `json:"gateway_ids"`
	PacketBroker MqttRxMetadataPacketBrokerResponse `json:"packet_broker"`
	Time         *time.Time                         `json:"time"`
	Rssi         int                                `json:"rssi"`
	ChannelRssi  int                                `json:"channel_rssi"`
	Snr          float64                            `json:"snr"`
	Location     MqttLocationResponse               `json:"location"`
	UplinkToken  string                             `json:"uplink_token"`
	RecievedAt   *time.Time                         `json:"recieved_at"`
} //@Name MqttRxMetadata

type MqttUplinkSettingsLoraResponse struct {
	Bandwidth       int    `json:"bandwidth"`
	SpreadingFactor int    `json:"spreading_factor"`
	CodingRate      string `json:"coding_rate"`
} //@Name MqttUplinkSettingsLora

type MqttUplinkSettingsDataRateResponse struct {
	Lora MqttUplinkSettingsLoraResponse `json:"lora"`
} //@Name MqttUplinkSettingsDataRate

type MqttUplinkSettingsResponse struct {
	DataRate  MqttUplinkSettingsDataRateResponse `json:"data_rate"`
	Frequency string                             `json:"frequency"`
} //@Name MqttUplinkSettings

type MqttVersionIDsResponse struct {
	BrandID         string `json:"brand_id"`
	ModelID         string `json:"model_id"`
	HardwareVersion string `json:"hardware_version"`
	FirmwareVersion string `json:"firmware_version"`
	BandID          string `json:"band_id"`
} //@Name MqttVersionIDs

type MqttNetworkIDsResponse struct {
	NetID          string `json:"net_id"`
	NSID           string `json:"nsid"`
	TenantID       string `json:"tenant_id"`
	ClusterID      string `json:"cluster_id"`
	ClusterAddress string `json:"cluster_address"`
	TenantAddress  string `json:"tenant_address"`
} //@Name MqttNetworkIDs

type MqttUplinkMessageResponse struct {
	SessionKeyID    string                     `json:"session_key_id"`
	FPort           int                        `json:"f_port"`
	Fcnt            int                        `json:"fcnt"`
	FRMPayload      string                     `json:"frm_payload"`
	DecodedPayload  MqttDecodedPayloadResponse `json:"decoded_payload"`
	RxMetadata      []MqttRxMetadataResponse   `json:"rx_metadata"`
	Settings        MqttUplinkSettingsResponse `json:"settings"`
	ReceivedAt      *time.Time                 `json:"received_at"`
	Confirmed       bool                       `json:"confirmed"`
	ConsumedAirtime string                     `json:"consumed_airtime"`
	VersionIDs      MqttVersionIDsResponse     `json:"version_ids"`
	NetworkIDs      MqttNetworkIDsResponse     `json:"network_ids"`
} //@Name MqttUplinkMessage

type MqttDataPayloadResponse struct {
	Type           string                         `json:"type"`
	EndDeviceIDs   MqttIdentifierDeviceIDResponse `json:"end_device_ids"`
	CorrelationIDs []string                       `json:"correlation_ids"`
	ReceivedAt     *time.Time                     `json:"received_at"`
	UplinkMessage  MqttUplinkMessageResponse      `json:"uplink_message"`
} //@Name MqttDataPayload

type MqttVisibilityResponse struct {
	Rights []string `json:"rights"`
} //@Name MqttVisibility

type MqttDataResponse struct {
	Name           string                   `json:"name"`
	Time           *time.Time               `json:"time"`
	Identifiers    []MqttIdentifierResponse `json:"identifiers"`
	Data           MqttDataPayloadResponse  `json:"data"`
	CorrelationIDs []string                 `json:"correlation_ids"`
	Origin         string                   `json:"origin"`
	Context        map[string]string        `json:"context"`
	Visibility     MqttVisibilityResponse   `json:"visibility"`
	UniqueID       string                   `json:"unique_id"`
} //@Name MqttData

type MqttPayloadResponse struct {
	EndDeviceIDs   MqttIdentifierDeviceIDResponse `json:"end_device_ids"`
	CorrelationIDs []string                       `json:"correlation_ids"`
	ReceivedAt     *time.Time                     `json:"received_at"`
	UplinkMessage  MqttUplinkMessageResponse      `json:"uplink_message"`
} //@Name MqttPayload
