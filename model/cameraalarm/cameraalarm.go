package cameraalarm

// EventNotificationAlert 报警事件
type EventNotificationAlert struct {
	IPAddress           string              `xml:"ipAddress"`
	ProtocolType        string              `xml:"protocolType"`
	MacAddress          string              `xml:"macAddress"`
	ChannelID           string              `xml:"channelID"`
	DateTime            string              `xml:"dateTime"`
	ActivePostCount     string              `xml:"activePostCount"`
	EventType           string              `xml:"eventType"`
	EventState          string              `xml:"eventState"`
	EventDescription    string              `xml:"eventDescription"`
	DetectionRegionList DetectionRegionList `xml:"DetectionRegionList"`
	ChannelName         string              `xml:"channelName"`
	Extensions          Extensions          `xml:"Extensions"`
}

// DetectionRegionList ...
type DetectionRegionList struct {
	DetectionRegionEntry DetectionRegionEntry `xml:"DetectionRegionEntry"`
}

// DetectionRegionEntry ...
type DetectionRegionEntry struct {
	RegionID              string                `xml:"regionID"`
	SensitivityLevel      string                `xml:"sensitivityLevel"`
	RegionCoordinatesList RegionCoordinatesList `xml:"RegionCoordinatesList"`
}

// RegionCoordinatesList 坐标
type RegionCoordinatesList struct {
	RegionCoordinates []RegionCoordinates `xml:"RegionCoordinates"`
}

// RegionCoordinates 坐标点
type RegionCoordinates struct {
	PositionX string `xml:"positionX"`
	PositionY string `xml:"positionY"`
}

// Extensions 扩展
type Extensions struct {
	SerialNumber string `xml:"serialNumber"`
	EventPush    string `xml:"eventPush"`
}

// Device ...
type Device struct {
	SubSN              string `json:"sub_sn"`           //设备标识(设备序列号)	 PcNodeID
	SubName            string `json:"sub_name"`         //设备名称				PcName
	SubType            string `json:"sub_type"`         //设备类型				PcDeviceType
	PcManufacturerID   string `json:"pcmanufacturerid"` //厂商ID					PcManufacturerID
	PcModel            string `json:"pcmodel"`          //设备型号				 PcModel
	PcManufacturerName string `json:"sub_mfrs"`         //厂商名称				 PcManufacturerName
	ServiceID          string `json:"service_id"`       //

	DateTime     string `json:"dateTime"`
	EventType    string `json:"eventType"`
	EventState   string `json:"eventState"`
	SerialNumber string `json:"serialNumber"`
	IPAddress    string `json:"ipAddress"`
	MacAddress   string `json:"macAddress"`
	CameraName   string `json:"cameraName"`
}

// PubDatas ...
type PubDatas struct {
	Devs []Device `json:"devices"`
}

// PubData ...
type PubData struct {
	Dev Device `json:"device"`
}
