package mysubaru

import (
	//	"errors"
	"encoding/json"
	"strconv"
	"time"

	"github.com/Jeffail/gabs/v2"
	log "github.com/sirupsen/logrus"
)

// Response struct
type Response struct {
	Success   bool   `json:"success"`
	DataName  string `json:"dataName,omitempty"`
	Data      Data   `json:"data"`
	ErrorCode string `json:"errorCode"`
}

// Data struct for Unmarshal dataName "lock / engineStop"
type Data struct {
	ServiceState string      `json:"remoteServiceState,omitempty"` // [ started | stopping | finished ]
	ServiceType  string      `json:"remoteServiceType,omitempty"`  // [ hornLights | vehicleStatus | lock | unlock ]
	RequestID    string      `json:"serviceRequestId,omitempty"`
	Success      bool        `json:"success,omitempty"`
	Cancelled    bool        `json:"cancelled,omitempty"`
	UpdateTime   time.Time   `json:"updateTime,omitempty"`
	VIN          string      `json:"vin,omitempty"`
	Result       interface{} `json:"result,omitempty"`
}

// Result struct
type Result struct {
	LastUpdatedTime   string       `json:"lastUpdatedTime,omitempty"`
	Odometer          int          `json:"odometer,omitempty"`
	VehicleStatus     []StatusItem `json:"vehicleStatus,omitempty"`
	Latitude          float32      `json:"latitude,omitempty"`
	Longitude         float32      `json:"longitude,omitempty"`
	LocationTimestamp string       `json:"locationTimestamp"`
}

// StatusItem struct
type StatusItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ClimateParams struct
type ClimateParams struct {
	FrontZoneTemp      string `json:"climateZoneFrontTemp"`
	ClimateSettings    string `json:"climateSettings"`
	RunTime            string `json:"runTimeMinutes"`
	FrontAirMode       string `json:"climateZoneFrontAirMode"`
	SeatFrontLeft      string `json:"heatedSeatFrontLeft"`
	SeatFrontRight     string `json:"heatedSeatFrontRight"`
	HeatedRearWindow   string `json:"heatedRearWindowActive"`
	FrontAirVolume     string `json:"climateZoneFrontAirVolume"`
	AirCirculation     string `json:"outerAirCirculation"`
	AirCondition       string `json:"airConditionOn"`
	StartConfiguration string `json:"startConfiguration"`
}

// Error struct for Unmarshal dataName "errorResponse"
type Error struct {
	Label       string `json:"errorLabel"`
	Description string `json:"errorDescription"`
}

// Parse method
func (res *Response) Parse(reps []byte) {
	respParsed, err := gabs.ParseJSON(reps)
	if err != nil {
		log.Fatalf(err.Error())
	}

	success, _ := strconv.ParseBool(respParsed.Path("success").String())
	if success == true {
		if serviceType, ok := respParsed.Search("data", "remoteServiceType").Data().(string); ok == true {
			switch serviceType {
			case "condition":
				log.Infof("condition")

				var result Result
				err = json.Unmarshal([]byte(respParsed.Search("data", "result").String()), &result)
				if err != nil {
					log.Warnf(err.Error())
				}
				log.Infof("%+v", result)
			case "locate":
				log.Infof("locate")

				var result Result
				err = json.Unmarshal([]byte(respParsed.Search("data", "result").String()), &result)
				if err != nil {
					log.Warnf(err.Error())
				}
				log.Infof("%+v", result)
			case "hornLights":
				log.Infof("hornLights")
			case "unlock":
				log.Infof("unlock")
			case "lock":
				log.Infof("lock")
			case "engineStart":
				log.Infof("engineStart")
			case "engineStop":
				log.Infof("engineStop")
			}
		}

		// return resp
	}

	//	return nil, errors.New(respParsed.Path("data.errorDescription").String())
}
