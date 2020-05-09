package mysubaru

import (
	"github.com/spf13/viper"
)

// Request is a struct
type Request struct {
	requestID      string
	urlPathExecute string
	urlPathStatus  string
	urlPathCancel  string
	body           map[string]string
	client         *Client

}

// StartEngine method
func (req *Request) StartEngine() []byte {
	req.setClimatParams()
	req.setPIN()
	req.setNow()
	req.urlPathExecute = "/service/g2/engineStart/execute.json"
	req.urlPathStatus = "/service/g2/remoteService/status.json"
	req.urlPathCancel = "/service/g2/engineStart/cancel.json"

	return req.client.execute(req)
}

// StopEngine method
func (req *Request) StopEngine() []byte {
	req.setPIN()
	req.setNow()
	req.urlPathExecute = "/service/g2/engineStop/execute.json"
	req.urlPathStatus = "/service/g2/remoteService/status.json"
	req.urlPathCancel = "/service/g2/engineStop/cancel.json"

	return req.client.execute(req)
}

// Lock method
func (req *Request) Lock() []byte {
	req.setPIN()
	req.setNow()
	req.body["delay"] = "0"
	req.body["unlockDoorType"] = "FRONT_LEFT_DOOR_CMD"                    // FRONT_LEFT_DOOR_CMD | ALL
	req.urlPathExecute = "/service/g2/lock/execute.json"
	req.urlPathStatus = "/service/g2/remoteService/status.json"
	req.urlPathCancel = "/service/g2/lock/cancel.json"

	return req.client.execute(req)
}

// Unlock method
func (req *Request) Unlock() []byte {
	req.setPIN()
	req.setNow()
	req.body["delay"] = "0"
	req.body["unlockDoorType"] = "FRONT_LEFT_DOOR_CMD"                    // FRONT_LEFT_DOOR_CMD | ALL
	req.urlPathExecute = "/service/g2/unlock/execute.json"
	req.urlPathStatus = "/service/g2/remoteService/status.json"
	req.urlPathCancel = "/service/g2/unlock/cancel.json"

	return req.client.execute(req)
}

// StartLighHorn method
func (req *Request) StartLightHorn() []byte {
	req.setPIN()
	req.setNow()
	req.body["delay"] = "0"
	req.body["horn"] = "true"
	req.urlPathExecute = "/service/g2/hornlight/execute.json"
	req.urlPathStatus = "/service/g2/remoteService/status.json"
	req.urlPathCancel = "/service/g2/hornlight/cancel.json"

	return req.client.execute(req)
}

// StopLightHorn method
func (req *Request) StopLightHorn() []byte {
	req.setPIN()
	req.setNow()
	req.urlPathExecute = "/service/g2/hornLights/stop.json"
	req.urlPathStatus = "/service/g2/remoteService/status.json"
	req.urlPathCancel = "/service/g2/hornlight/cancel.json"

	return req.client.execute(req)
}

// GetLocation method
func (req *Request) GetLocation() []byte {
	req.setPIN()
	req.setNow()
	req.urlPathExecute = "/service/g2/refreshVehicleStatus/execute.json"
	req.urlPathStatus = "/service/g2/locate/execute.json"

	return req.client.execute(req)
}

// GetHealth method
func (req *Request) GetHealth() []byte {
	req.setPIN()
	req.setNow()
	req.urlPathExecute = "/service/g2/remoteHealth/execute.json"
	req.urlPathStatus = "/service/g2/remoteHealth/status.json"

	return req.client.execute(req)
}

func (req *Request) GetClimateSettings() []byte {
	req.setPIN()
	req.setNow()
	req.urlPathExecute = "/service/g2/remoteStart/fetch.json"

	return req.client.execute(req)
}

// Exec is a method
//func (req *Request) Exec() ([]byte) {
//	resp := req.client.execute(req)
//	return resp
//}

func (req *Request) setClimatParams() {
	req.body["delay"] = "0"
	req.body["horn"] = "true"
	req.body["unlockDoorType"] = "FRONT_LEFT_DOOR_CMD"                    // FRONT_LEFT_DOOR_CMD | ALL
	req.body["climateSettings"] = "climateSettings"                       // climateSettings
	req.body["climateZoneFrontTemp"] = "65"                               // 60-??
	req.body["climateZoneFrontAirMode"] = "FEET_WINDOW"                   // FEET_FACE_BALANCED | FEET_WINDOW | WINDOW | FEET
	req.body["climateZoneFrontAirVolume"] = "1"                           // 1-7
	req.body["heatedSeatFrontLeft"] = "OFF"                               // OFF | LOW_HEAT | MEDIUM_HEAT | HIGH_ HEAT
	req.body["heatedSeatFrontRight"] = "OFF"                              // ---//---
	req.body["heatedRearWindowActive"] = "false"                          // boolean
	req.body["outerAirCirculation"] = "outsideAir"                        // outsideAir | recirculation
	req.body["airConditionOn"] = "true"                                   // boolean
	req.body["runTimeMinutes"] = "10"                                     // 1-10
	req.body["startConfiguration"] = "START_ENGINE_ALLOW_KEY_IN_IGNITION" // START_ENGINE_ALLOW_KEY_IN_IGNITION
}

func (req *Request) setPIN() {
	req.body["pin"] = viper.GetString("credentials.pin")
}

func (req *Request) setNow() {
	req.body["now"] = timestamp()
}
