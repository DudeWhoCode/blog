package flightdata

type LiveData struct {
	Updated         string  `json:"updated"`
	Latitude        float32 `json:"latitude"`
	Longitude       float32 `json:"longitude"`
	Altitude        float32 `json:"altitude"`
	Direction       float32 `json:"direction"`
	SpeedHorizontal float32 `json:"speed_horizontal"`
	SpeedVertical   float32 `json:"speed_vertical"`
	OnGround        bool    `json:"is_ground"`
}

type Tracker interface {
	GetLiveData(flightNumber string) (LiveData, error)
}
