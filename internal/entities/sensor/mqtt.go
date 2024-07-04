package sensor

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MqttIdentifierDeviceID struct {
	DeviceID       string `json:"device_id" bson:"device_id"`
	ApplicationIDs struct {
		ApplicationID string `json:"application_id" bson:"application_id"`
	} `json:"application_ids" bson:"application_ids"`
	DevEUI  string `json:"dev_eui" bson:"dev_eui"`
	JoinEUI string `json:"join_eui" bson:"join_eui"`
	DevAddr string `json:"dev_addr" bson:"dev_addr"`
}

type MqttIdentifier struct {
	DeviceIDs MqttIdentifierDeviceID `json:"device_ids" bson:"device_ids"`
}

type MqttDecodedPayload struct {
	Battery  float64 `json:"battery" bson:"battery"`
	Humidity int     `json:"humidity" bson:"humidity"`
	Raw      int     `json:"raw" bson:"raw"`
}

type MqttRxMetadata struct {
	GatewayIDs struct {
		GatewayID string `json:"gateway_id" bson:"gateway_id"`
	} `json:"gateway_ids" bson:"gateway_ids"`
	PacketBroker struct {
		MessageID            string `json:"message_id" bson:"message_id"`
		ForwarderNetID       string `json:"forwarder_net_id" bson:"forwarder_net_id"`
		ForwarderTenantID    string `json:"forwarder_tenant_id" bson:"forwarder_tenant_id"`
		ForwarderClusterID   string `json:"forwarder_cluster_id" bson:"forwarder_cluster_id"`
		ForwarderGatewayID   string `json:"forwarder_gateway_id" bson:"forwarder_gateway_id"`
		ForwarderGatewayEUI  string `json:"forwarder_gateway_eui" bson:"forwarder_gateway_eui"`
		HomeNetworkNetID     string `json:"home_network_net_id" bson:"home_network_net_id"`
		HomeNetworkTenantID  string `json:"home_network_tenant_id" bson:"home_network_tenant_id"`
		HomeNetworkClusterID string `json:"home_network_cluster_id" bson:"home_network_cluster_id"`
	} `json:"packet_broker" bson:"packet_broker"`
	Time        *time.Time `json:"time" bson:"time"`
	Rssi        int        `json:"rssi" bson:"rssi"`
	ChannelRssi int        `json:"channel_rssi" bson:"channel_rssi"`
	Snr         float64    `json:"snr" bson:"snr"`
	Location    struct {
		Latitude  float64 `json:"latitude" bson:"latitude"`
		Longitude float64 `json:"longitude" bson:"longitude"`
		Altitude  float64 `json:"altitude" bson:"altitude"`
	} `json:"location" bson:"location"`
	UplinkToken string     `json:"uplink_token" bson:"uplink_token"`
	RecievedAt  *time.Time `json:"received_at" bson:"received_at"`
}

type MqttUplinkSettings struct {
	DataRate struct {
		Lora struct {
			Bandwidth       int    `json:"bandwidth" bson:"bandwidth"`
			SpreadingFactor int    `json:"spreading_factor" bson:"spreading_factor"`
			CodingRate      string `json:"coding_rate" bson:"coding_rate"`
		} `json:"lora" bson:"lora"`
	} `json:"data_rate" bson:"data_rate"`
	Frequency string `json:"frequency" bson:"frequency"`
}

type MqttVersionIDs struct {
	BrandID         string `json:"brand_id" bson:"brand_id"`
	ModelID         string `json:"model_id" bson:"model_id"`
	HardwareVersion string `json:"hardware_version" bson:"hardware_version"`
	FirmwareVersion string `json:"firmware_version" bson:"firmware_version"`
	BandID          string `json:"band_id" bson:"band_id"`
}

type MqttNetworkIDs struct {
	NetID          string `json:"net_id" bson:"net_id"`
	NSID           string `json:"ns_id" bson:"ns_id"`
	TenantID       string `json:"tenant_id" bson:"tenant_id"`
	ClusterID      string `json:"cluster_id" bson:"cluster_id"`
	ClusterAddress string `json:"cluster_address" bson:"cluster_address"`
	TenantAddress  string `json:"tenant_address" bson:"tenant_address"`
}

type MqttUplinkMessage struct {
	SessionKeyID    string             `json:"session_key_id" bson:"session_key_id"`
	FPort           int                `json:"f_port" bson:"f_port"`
	Fcnt            int                `json:"fcnt" bson:"fcnt"`
	FRMPayload      string             `json:"frm_payload" bson:"frm_payload"`
	DecodedPayload  MqttDecodedPayload `json:"decoded_payload" bson:"decoded_payload"`
	RxMetadata      []MqttRxMetadata   `json:"rx_metadata" bson:"rx_metadata"`
	Settings        MqttUplinkSettings `json:"settings" bson:"settings"`
	ReceivedAt      *time.Time         `json:"received_at" bson:"received_at"`
	Confirmed       bool               `json:"confirmed" bson:"confirmed"`
	ConsumedAirtime string             `json:"consumed_airtime" bson:"consumed_airtime"` // time.Duration
	VersionIDs      MqttVersionIDs     `json:"version_ids" bson:"version_ids"`
	NetworkIDs      MqttNetworkIDs     `json:"network_ids" bson:"network_ids"`
}

type MqttDataPayload struct {
	Type           string                 `json:"type" bson:"@type"`
	EndDeviceIDs   MqttIdentifierDeviceID `json:"end_device_ids" bson:"end_device_ids"`
	CorrelationIDs []string               `json:"correlation_ids" bson:"correlation_ids"`
	ReceivedAt     *time.Time             `json:"received_at" bson:"received_at"`
	UplinkMessage  MqttUplinkMessage      `json:"uplink_message" bson:"uplink_message"`
}

type MqttVisibility struct {
	Rights []string `json:"rights" bson:"rights"`
}

type MqttData struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	TreeID         string             `json:"tree_id" bson:"tree_id"`
	Name           string             `json:"name" bson:"name"`
	Time           *time.Time         `json:"time" bson:"time"`
	Identifiers    []MqttIdentifier   `json:"identifiers" bson:"identifiers"`
	Data           MqttDataPayload    `json:"data" bson:"data"`
	CorrelationIDs []string           `json:"correlation_ids" bson:"correlation_ids"`
	Origin         string             `json:"origin" bson:"origin"`
	Context        map[string]string  `json:"context" bson:"context"`
	Visibility     MqttVisibility     `json:"visibility" bson:"visibility"`
	UniqueID       string             `json:"unique_id" bson:"unique_id"`
}
