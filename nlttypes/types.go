package nlttypes

import "time"

const (
	ActivationOTAA = "OTAA"
	ActivationABP  = "ABP"

	AdrModeOn  = "on"
	AdrModeOff = "off"

	EncryptionNS = "NS"

	BandName = "LA915-928A"

	DevClassA = "A"
	DevClassC = "C"
)

// Auth

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	UserType    int    `json:"user_type"`
}

type AuthError struct {
	Detail string `json:"detail"`
}

// Device

type DeviceCreateError struct {
	Detail string `json:"detail"`
}

// Tags

type Tag struct {
	Name     string `json:"name"`
	ClientID int    `json:"client_id"`
	ID       int    `json:"id"`
}

type Tags []Tag

type TagRequest struct {
	Name     string `json:"name"`
	ClientID int    `json:"client_id"`
}

// Downlink

type DownlinkRequest struct {
	Payload   string `json:"payload"`
	Port      int    `json:"port"`
	Confirmed bool   `json:"confirmed"`
}

type DownlinkResponse struct {
	Type   string `json:"type"`
	Meta   Meta   `json:"meta"`
	Params Params `json:"params"`
}

type Meta struct {
	Network     string  `json:"network"`
	PacketHash  string  `json:"packet_hash"`
	Application string  `json:"application"`
	DeviceAddr  string  `json:"device_addr"`
	Time        float64 `json:"time"`
	Device      string  `json:"device"`
	PacketID    string  `json:"packet_id"`
	Gateway     string  `json:"gateway"`
}

type Modulation struct {
	Bandwidth int    `json:"bandwidth"`
	Coderate  string `json:"coderate"`
	Type      string `json:"type"`
	Spreading int    `json:"spreading"`
	Inverted  bool   `json:"inverted"`
}

type Hardware struct {
	Immediately bool    `json:"immediately"`
	Chain       int     `json:"chain"`
	Power       float64 `json:"power"`
}

type Radio struct {
	Modulation Modulation `json:"modulation"`
	Hardware   Hardware   `json:"hardware"`
	Freq       float64    `json:"freq"`
	Time       float64    `json:"time"`
	Datr       string     `json:"datr"`
}

type Params struct {
	Payload          string `json:"payload"`
	Radio            Radio  `json:"radio"`
	CounterDown      int    `json:"counter_down"`
	Port             int    `json:"port"`
	EncryptedPayload string `json:"encrypted_payload"`
}

// Messages

type Messages struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Type string `json:"type"`
	Meta struct {
		Network     string  `json:"network"`
		PacketHash  string  `json:"packet_hash"`
		Application string  `json:"application"`
		DeviceAddr  string  `json:"device_addr"`
		Time        float64 `json:"time"`
		Device      string  `json:"device"`
		PacketID    string  `json:"packet_id"`
		Gateway     string  `json:"gateway"`
	} `json:"meta"`
	Params struct {
		Payload   string `json:"payload"`
		Port      int    `json:"port"`
		Duplicate bool   `json:"duplicate"`
		Radio     struct {
			GpsTime  int64 `json:"gps_time"`
			Hardware struct {
				Status  int     `json:"status"`
				Chain   int     `json:"chain"`
				Tmst    int     `json:"tmst"`
				Snr     float64 `json:"snr"`
				Rssi    int     `json:"rssi"`
				Channel int     `json:"channel"`
				Gps     struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
					Alt int     `json:"alt"`
				} `json:"gps"`
			} `json:"hardware"`
			Datarate   int `json:"datarate"`
			Modulation struct {
				Bandwidth int    `json:"bandwidth"`
				Type      string `json:"type"`
				Spreading int    `json:"spreading"`
				Coderate  string `json:"coderate"`
			} `json:"modulation"`
			Delay float64 `json:"delay"`
			Time  float64 `json:"time"`
			Freq  float64 `json:"freq"`
			Size  int     `json:"size"`
		} `json:"radio"`
		CounterUp        int     `json:"counter_up"`
		RxTime           float64 `json:"rx_time"`
		EncryptedPayload string  `json:"encrypted_payload"`
	} `json:"params"`
	InsertTime string `json:"insert_time"`
}

// Connection

type ConnectionResponse struct {
	Total  int    `json:"total"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Data   []Data `json:"data"`
}

type Connectionmodel struct {
	AuthHeader     string    `json:"auth_header"`
	URL            string    `json:"url"`
	ConnectionType string    `json:"connection_type"`
	Description    string    `json:"description"`
	ID             int       `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
}

type Filtermodel struct {
	Applications []string  `json:"applications"`
	Description  string    `json:"description"`
	Devices      []string  `json:"devices"`
	Duplicate    bool      `json:"duplicate"`
	Gateways     []string  `json:"gateways"`
	Lora         bool      `json:"lora"`
	Radio        bool      `json:"radio"`
	Tags         []string  `json:"tags"`
	Types        []string  `json:"types"`
	WithTags     bool      `json:"with_tags"`
	IsDisabled   bool      `json:"is_disabled"`
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
}

type Data struct {
	Connectionmodel Connectionmodel `json:"ConnectionModel"`
	Filtermodel     Filtermodel     `json:"FilterModel"`
}

// Device

type DeviceListResponse []Device

type Device struct {
	Tags          []string    `json:"tags"`
	Activation    string      `json:"activation"`
	Adr           Adr         `json:"adr"`
	AppEui        string      `json:"app_eui"`
	AppKey        string      `json:"app_key"`
	Appskey       string      `json:"appskey"`
	Band          string      `json:"band"`
	CountersSize  int         `json:"counters_size"`
	DevAddr       string      `json:"dev_addr"`
	DevClass      string      `json:"dev_class"`
	Encryption    string      `json:"encryption"`
	Nwkskey       string      `json:"nwkskey"`
	Rx1           Rx1         `json:"rx1"`
	StrictCounter bool        `json:"strict_counter"`
	DeviceType    string      `json:"device_type"`
	ContractID    int         `json:"contract_id"`
	DevEui        string      `json:"dev_eui"`
	BlockDownlink bool        `json:"block_downlink"`
	BlockUplink   bool        `json:"block_uplink"`
	ID            int         `json:"id"`
	CounterDown   int         `json:"counter_down"`
	CounterUp     int         `json:"counter_up"`
	Geolocation   Geolocation `json:"geolocation"`
	LastActivity  string      `json:"last_activity"`
	LastJoin      string      `json:"last_join"`
	ActivatedAt   string      `json:"activated_at"`
	DeactivatedAt string      `json:"deactivated_at"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	Detail        string      `json:"detail"`
	Message       string      `json:"message"`
}

type Adr struct {
	Mode string `json:"mode"`
}

type Rx1 struct {
	Delay int `json:"delay"`
}

type Geolocation struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

// Device Create

type DeviceCreateRequest struct {
	Tags          []string `json:"tags"`
	Activation    string   `json:"activation"`
	Adr           DevAdr   `json:"adr"`
	AppEui        string   `json:"app_eui"`
	AppKey        string   `json:"app_key,omitempty"`
	Appskey       string   `json:"appskey,omitempty"`
	Band          string   `json:"band"`
	CountersSize  int      `json:"counters_size"`
	DevAddr       string   `json:"dev_addr,omitempty"`
	DevClass      string   `json:"dev_class"`
	Encryption    string   `json:"encryption"`
	Nwkskey       string   `json:"nwkskey,omitempty"`
	Rx1           DevRx1   `json:"rx1"`
	StrictCounter bool     `json:"strict_counter"`
	DeviceType    string   `json:"device_type"`
	ContractID    int      `json:"contract_id"`
	DevEui        string   `json:"dev_eui"`
	BlockDownlink bool     `json:"block_downlink"`
	BlockUplink   bool     `json:"block_uplink"`
}

type DeviceUpdateRequest struct {
	Tags          []string `json:"tags"`
	Activation    string   `json:"activation"`
	Adr           DevAdr   `json:"adr"`
	AppEui        string   `json:"app_eui"`
	AppKey        string   `json:"app_key,omitempty"`
	Appskey       string   `json:"appskey,omitempty"`
	Band          string   `json:"band"`
	CountersSize  int      `json:"counters_size"`
	DevAddr       string   `json:"dev_addr,omitempty"`
	DevClass      string   `json:"dev_class"`
	Encryption    string   `json:"encryption"`
	Nwkskey       string   `json:"nwkskey,omitempty"`
	Rx1           DevRx1   `json:"rx1"`
	StrictCounter bool     `json:"strict_counter"`
	DeviceType    string   `json:"device_type"`
	ContractID    int      `json:"contract_id"`
	DevEui        string   `json:"dev_eui"`
	BlockDownlink bool     `json:"block_downlink"`
	BlockUplink   bool     `json:"block_uplink"`
}

type DevAdr struct {
	Mode string `json:"mode"`
}

type DevRx1 struct {
	Delay int `json:"delay"`
}

// Create Connection

type CreateConnectionRequest struct {
	Connectionmodel Connectionmodel `json:"connections"`
	Filtermodel     Filtermodel     `json:"filters"`
}

type CreateConnectionResponse struct {
	Connectionmodel Connectionmodel `json:"connections"`
	Filtermodel     Filtermodel     `json:"filters"`
}

// Update Connection

type UpdateConnectionRequest struct {
	Connectionmodel Connectionmodel `json:"connections"`
	Filtermodel     Filtermodel     `json:"filters"`
}

type UpdateConnectionResponse struct {
	Connectionmodel Connectionmodel `json:"connections"`
	Filtermodel     Filtermodel     `json:"filters"`
}

// Delete Connection

type DeleteConnectionResponse struct {
	Message string `json:"message"`
}
