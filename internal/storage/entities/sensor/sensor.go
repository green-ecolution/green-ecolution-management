package sensor

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MqttIdentifierApplicationIDEntity struct {
	ApplicationID string `bson:"application_id"`
}

type MqttIdentifierDeviceIDEntity struct {
	DeviceID       string                            `bson:"device_id"`
	ApplicationIDs MqttIdentifierApplicationIDEntity `bson:"application_ids"`
	DevEUI         string                            `bson:"dev_eui"`
	JoinEUI        string                            `bson:"join_eui"`
	DevAddr        string                            `bson:"dev_addr"`
}

type MqttIdentifierEntity struct {
	DeviceIDs MqttIdentifierDeviceIDEntity `bson:"device_ids"`
}

type MqttDecodedPayloadEntity struct {
	Battery  float64 `bson:"battery"`
	Humidity int     `bson:"humidity"`
	Raw      int     `bson:"raw"`
}

type MqttRxMetadataGatewayIDsEntity struct {
	GatewayID string `bson:"gateway_id"`
}

type MqttRxMetadataPacketBrokerEntity struct {
	MessageID            string `bson:"message_id"`
	ForwarderNetID       string `bson:"forwarder_net_id"`
	ForwarderTenantID    string `bson:"forwarder_tenant_id"`
	ForwarderClusterID   string `bson:"forwarder_cluster_id"`
	ForwarderGatewayID   string `bson:"forwarder_gateway_id"`
	ForwarderGatewayEUI  string `bson:"forwarder_gateway_eui"`
	HomeNetworkNetID     string `bson:"home_network_net_id"`
	HomeNetworkTenantID  string `bson:"home_network_tenant_id"`
	HomeNetworkClusterID string `bson:"home_network_cluster_id"`
}

type MqttLocationEntity struct {
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
	Altitude  float64 `bson:"altitude"`
}

type MqttRxMetadataEntity struct {
	GatewayIDs   MqttRxMetadataGatewayIDsEntity   `bson:"gateway_i_ds"`
	PacketBroker MqttRxMetadataPacketBrokerEntity `bson:"packet_broker"`
	Time         *time.Time                       `bson:"time"`
	Rssi         int                              `bson:"rssi"`
	ChannelRssi  int                              `bson:"channel_rssi"`
	Snr          float64                          `bson:"snr"`
	Location     MqttLocationEntity               `bson:"location"`
	UplinkToken  string                           `bson:"uplink_token"`
	RecievedAt   *time.Time                       `bson:"recieved_at"`
}

type MqttUplinkSettingsLoraEntity struct {
	Bandwidth       int    `bson:"bandwidth"`
	SpreadingFactor int    `bson:"spreading_factor"`
	CodingRate      string `bson:"coding_rate"`
}

type MqttUplinkSettingsDataRateEntity struct {
	Lora MqttUplinkSettingsLoraEntity `bson:"lora"`
}

type MqttUplinkSettingsEntity struct {
	DataRate  MqttUplinkSettingsDataRateEntity `bson:"data_rate"`
	Frequency string                           `json:"frequency" bson:"frequency"`
}

type MqttVersionIDsEntity struct {
	BrandID         string `bson:"brand_id"`
	ModelID         string `bson:"model_id"`
	HardwareVersion string `bson:"hardware_version"`
	FirmwareVersion string `bson:"firmware_version"`
	BandID          string `bson:"band_id"`
}

type MqttNetworkIDsEntity struct {
	NetID          string `bson:"net_id"`
	NSID           string `bson:"nsid"`
	TenantID       string `bson:"tenant_id"`
	ClusterID      string `bson:"cluster_id"`
	ClusterAddress string `bson:"cluster_address"`
	TenantAddress  string `bson:"tenant_address"`
}

type MqttUplinkMessageEntity struct {
	SessionKeyID    string                   `bson:"session_key_id"`
	FPort           int                      `bson:"f_port"`
	Fcnt            int                      `bson:"fcnt"`
	FRMPayload      string                   `bson:"frm_payload"`
	DecodedPayload  MqttDecodedPayloadEntity `bson:"decoded_payload"`
	RxMetadata      []MqttRxMetadataEntity   `bson:"rx_metadata"`
	Settings        MqttUplinkSettingsEntity `bson:"settings"`
	ReceivedAt      *time.Time               `bson:"received_at"`
	Confirmed       bool                     `bson:"confirmed"`
	ConsumedAirtime string                   `bson:"consumed_airtime"`
	VersionIDs      MqttVersionIDsEntity     `bson:"version_ids"`
	NetworkIDs      MqttNetworkIDsEntity     `bson:"network_ids"`
}

type MqttDataPayloadEntity struct {
	Type           string                       `bson:"type"`
	EndDeviceIDs   MqttIdentifierDeviceIDEntity `bson:"end_device_ids"`
	CorrelationIDs []string                     `bson:"correlation_ids"`
	ReceivedAt     *time.Time                   `bson:"received_at"`
	UplinkMessage  MqttUplinkMessageEntity      `bson:"uplink_message"`
}

type MqttVisibilityEntity struct {
	Rights []string `bson:"rights"`
}

type MqttDataEntity struct {
	Name           string                 `bson:"name"`
	Time           *time.Time             `bson:"time"`
	Identifiers    []MqttIdentifierEntity `bson:"identifiers"`
	Data           MqttDataPayloadEntity  `bson:"data"`
	CorrelationIDs []string               `bson:"correlation_ids"`
	Origin         string                 `bson:"origin"`
	Context        map[string]string      `bson:"context"`
	Visibility     MqttVisibilityEntity   `bson:"visibility"`
	UniqueID       string                 `bson:"unique_id"`
}

type MqttPayloadEntity struct {
	EndDeviceIDs   MqttIdentifierDeviceIDEntity `bson:"end_device_ids"`
	CorrelationIDs []string                     `bson:"correlation_ids"`
	ReceivedAt     *time.Time                   `bson:"received_at"`
	UplinkMessage  MqttUplinkMessageEntity      `bson:"uplink_message"`
}

type MqttEntity struct {
	ID     primitive.ObjectID `bson:"_id"`
	TreeID string             `bson:"tree_id"`
	Data   MqttPayloadEntity  `bson:"data"`
}

func (m *MqttEntity) GetID() string {
  return m.ID.Hex()
}

func (m *MqttEntity) SetID(id string) error {
  objID, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return err
  }
  m.ID = objID
  return nil
}
