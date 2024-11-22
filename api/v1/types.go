package v1

type Webserver struct {
	Mac              string `json:"mac"`
	WifiChannel      int    `json:"wifi_channel"`
	WifiQuality      int    `json:"wifi_quality"`
	WifiRssi         int    `json:"wifi_rssi"`
	Interface        string `json:"interface"`
	WsFirmware       string `json:"ws_firmware"`
	LmachineFirmware string `json:"lmachine_firmware"`
	CloudConnected   int    `json:"cloud_connected"`
	WsType           string `json:"ws_type"`
}

type Version struct {
	Version string `json:"version"`
}

type System struct {
	SystemID         int           `json:"systemID"`
	McConnected      int           `json:"mc_connected"`
	SystemFirmware   string        `json:"system_firmware"`
	SystemType       int           `json:"system_type"`
	SystemTechnology int           `json:"system_technology"`
	Manufacturer     string        `json:"manufacturer"`
	Errors           []interface{} `json:"errors"`
	Qadapt           int           `json:"qadapt"`
}

type Zone struct {
	SystemID        int           `json:"systemID"`
	ZoneID          int           `json:"zoneID"`
	Name            string        `json:"name"`
	ThermosType     int           `json:"thermos_type"`
	ThermosFirmware string        `json:"thermos_firmware"`
	ThermosRadio    int           `json:"thermos_radio"`
	On              int           `json:"on"`
	DoubleSp        int           `json:"double_sp"`
	Coolsetpoint    float64       `json:"coolsetpoint"`
	Coolmaxtemp     float64       `json:"coolmaxtemp"`
	Coolmintemp     float64       `json:"coolmintemp"`
	Heatsetpoint    float64       `json:"heatsetpoint"`
	Heatmaxtemp     float64       `json:"heatmaxtemp"`
	Heatmintemp     float64       `json:"heatmintemp"`
	MaxTemp         float64       `json:"maxTemp"`
	MinTemp         float64       `json:"minTemp"`
	Setpoint        float64       `json:"setpoint"`
	RoomTemp        float64       `json:"roomTemp"`
	Sleep           int           `json:"sleep"`
	TempStep        float64       `json:"temp_step"`
	Modes           []int         `json:"modes"`
	Mode            int           `json:"mode"`
	Speed           int           `json:"speed"`
	ColdStage       int           `json:"coldStage"`
	HeatStage       int           `json:"heatStage"`
	ColdStages      int           `json:"coldStages"`
	HeatStages      int           `json:"heatStages"`
	Humidity        float64       `json:"humidity"`
	Units           int           `json:"units"`
	Errors          []interface{} `json:"errors"`
	AirDemand       int           `json:"air_demand"`
	FloorDemand     int           `json:"floor_demand"`
	ColdDemand      int           `json:"cold_demand"`
	HeatDemand      int           `json:"heat_demand"`
	Heatangle       int           `json:"heatangle"`
	Coldangle       int           `json:"coldangle"`
	MasterZoneID    int           `json:"master_zoneID"`
	EcoAdapt        string        `json:"eco_adapt"`
	Antifreeze      int           `json:"antifreeze"`
}
