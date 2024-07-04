package sensor

type Data struct {
	EndDeviceIDs struct {
		DeviceID       string `json:"device_id"`
		ApplicationIDs struct {
			ApplicationID string `json:"application_id"`
		} `json:"application_ids"`

		DevEUI  string `json:"dev_eui"`
		JoinEUI string `json:"join_eui"`
		DevAddr string `json:"dev_addr"`
	} `json:"end_device_ids"`

	CorrelationIDs []string `json:"correlation_ids"`
	ReceivedAt     string   `json:"received_at"`

	UplinkMessage struct {
		SessionKeyID string `json:"session_key_id"`
		FPort        int    `json:"f_port"`
		FCnt         int    `json:"f_cnt"`
		FrmPayload   string `json:"frm_payload"`

		DecodedPayload struct {
			Battery  float64 `json:"battery"`
			Humidity int     `json:"humidity"`
			Raw      int     `json:"raw"`
		} `json:"decoded_payload"`

		RxMetadata []struct {
			GatewayIDs struct {
				GatewayID string `json:"gateway_id"`
				EUI       string `json:"eui"`
			} `json:"gateway_ids"`
			Time            string  `json:"time"`
			Timestamp       int64   `json:"timestamp"`
			RSSI            int     `json:"rssi"`
			ChannelRSSI     int     `json:"channel_rssi"`
			SNR             float64 `json:"snr"`
			FrequencyOffset string  `json:"frequency_offset"`
			Location        struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Altitude  int     `json:"altitude"`
				Source    string  `json:"source"`
			} `json:"location"`

			UplinkToken  string `json:"uplink_token"`
			ChannelIndex int    `json:"channel_index"`
			ReceivedAt   string `json:"received_at"`
		} `json:"rx_metadata"`

		Settings struct {
			DataRate struct {
				Lora struct {
					Bandwidth       int    `json:"bandwidth"`
					SpreadingFactor int    `json:"spreading_factor"`
					CodingRate      string `json:"coding_rate"`
				} `json:"lora"`
			} `json:"data_rate"`

			Frequency string `json:"frequency"`
			Timestamp int64  `json:"timestamp"`
			Time      string `json:"time"`
		} `json:"settings"`

		ReceivedAt      string `json:"received_at"`
		Confirmed       bool   `json:"confirmed"`
		ConsumedAirtime string `json:"consumed_airtime"`

		Locations struct {
			User struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Source    string  `json:"source"`
			} `json:"user"`
		} `json:"locations"`

		VersionIDs struct {
			BrandID         string `json:"brand_id"`
			ModelID         string `json:"model_id"`
			HardwareVersion string `json:"hardware_version"`
			FirmwareVersion string `json:"firmware_version"`
			BandID          string `json:"band_id"`
		} `json:"version_ids"`

		NetworkIDs struct {
			NetID          string `json:"net_id"`
			NsID           string `json:"ns_id"`
			TenantID       string `json:"tenant_id"`
			ClusterID      string `json:"cluster_id"`
			ClusterAddress string `json:"cluster_address"`
		} `json:"network_ids"`
	} `json:"uplink_message"`
}
