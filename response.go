package mysubaru

// Response struct
type Response struct {
    Success		bool		`json:"success"`
    DataName		string		`json:"dataName,omitempty"`
    Data		interface{}	`json:"data"`
}

// Data struct for Unmarshal dataName "lock / engineStop"
type Data struct {
    ServiceState	string		`json:"remoteServiceState"`		// [ started | stopping | finished ]
    ServiceType		string		`json:"remoteServiceType"`		// [ hornLights | vehicleStatus | lock | unlock ]
    RequestID		string		`json:"serviceRequestId"`
    Success		bool		`json:"success"`
    Cancelled		bool		`json:"cancelled"`
    UpdateTime		string		`json:"updateTime"`
    VIN			string		`json:"vin"`
    Result		interface{}	`json:"result,omitempty"`
}

type Result struct {
    LastUpdatedTime	string		`json:"lastUpdatedTime"`
    Odometer		int		`json:"odometer"`
    VehicleStatus	[]StatusItem	`json:"vehicleStatus"`
}

type StatusItem struct {
    Key			string		`json:"key"`
    Value		string		`json:"value"`
}

type ClimateParams struct {
    FrontZoneTemp	string		`json:"climateZoneFrontTemp"`
    ClimateSettings	string		`json:"climateSettings"`
    RunTime		string		`json:"runTimeMinutes"`
    FrontAirMode	string		`json:"climateZoneFrontAirMode"`
    SeatFrontLeft	string		`json:"heatedSeatFrontLeft"`
    SeatFrontRight	string		`json:"heatedSeatFrontRight"`
    HeatedRearWindow	string		`json:"heatedRearWindowActive"`
    FrontAirVolume	string		`json:"climateZoneFrontAirVolume"`
    AirCirculation	string		`json:"outerAirCirculation"`
    AirCondition	string		`json:"airConditionOn"`
    StartConfiguration	string		`json:"startConfiguration"`
}

// Error struct for Unmarshal dataName "errorResponse"
type Error struct {
    Label		string		`json:"errorLabel"`
    Description		string		`json:"errorDescription"`
}

/*
{"success":false,
"errorCode":"InvalidInputCriteria",
"dataName":"errorResponse",
"data":{"errorLabel":"InvalidInputCriteria","errorDescription":"The input parameters are invalid"}}

{"success":false,
"errorCode":"InvalidInputCriteria",
"dataName":"errorResponse",
"data":{"errorLabel":"InvalidInputCriteria","errorDescription":"The input parameters are invalid"}}


{"data":
  {"cancelled":false,"remoteServiceState":"started","remoteServiceType":"lock","serviceRequestId":"4S4BTGND8L3137058_1588963308584_20_@NGTP","success":false,"updateTime":"2020-05-08T18:41:48.000+0000","vin":"4S4BTGND8L3137058"},
  "dataName":"remoteServiceStatus",
  "success":true}
{"data":
{"cancelled":false,"remoteServiceState":"finished","remoteServiceType":"engineStop","serviceRequestId":"4S4BTGND8L3137058_1588961508982_23_@NGTP","success":true,"updateTime":"2020-05-08T18:11:52.000+0000","vin":"4S4BTGND8L3137058"},
"dataName":"remoteServiceStatus",
"success":true}

{
  "data": {
    "cancelled": false,
    "remoteServiceState": "finished",
    "remoteServiceType": "condition",
    "result": {
      "lastUpdatedTime": "2020-05-08T18:13:08+0000",
      "odometer": 5478645,
      "vehicleStatus": [
        {
          "key": "TRANSMISSION_MODE",
          "value": "UNKNOWN"
        },
        {
          "key": "TYRE_PRESSURE_REAR_RIGHT",
          "value": "32767"
        },
        {
          "key": "SEAT_BELT_STATUS_THIRD_MIDDLE",
          "value": "UNKNOWN"
        },
        {
          "key": "DOOR_FRONT_LEFT_LOCK_STATUS",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_BELT_STATUS_FRONT_MIDDLE",
          "value": "NOT_EQUIPPED"
        },
        {
          "key": "SEAT_OCCUPATION_STATUS_FRONT_RIGHT",
          "value": "UNKNOWN"
        },
        {
          "key": "EXT_EXTERNAL_TEMP",
          "value": "-64.0"
        },
        {
          "key": "WINDOW_BACK_STATUS",
          "value": "UNKNOWN"
        },
        {
          "key": "WINDOW_REAR_LEFT_STATUS",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_BELT_STATUS_SECOND_MIDDLE",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_OCCUPATION_STATUS_THIRD_RIGHT",
          "value": "UNKNOWN"
        },
        {
          "key": "DOOR_FRONT_LEFT_POSITION",
          "value": "CLOSED"
        },
        {
          "key": "TYRE_PRESSURE_FRONT_RIGHT",
          "value": "32767"
        },
        {
          "key": "SEAT_BELT_STATUS_THIRD_RIGHT",
          "value": "UNKNOWN"
        },
        {
          "key": "DOOR_REAR_RIGHT_POSITION",
          "value": "CLOSED"
        },
        {
          "key": "DISTANCE_TO_EMPTY_FUEL",
          "value": "16383"
        },
        {
          "key": "SEAT_BELT_STATUS_FRONT_LEFT",
          "value": "BELTED"
        },
        {
          "key": "TYRE_STATUS_REAR_LEFT",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_OCCUPATION_STATUS_FRONT_MIDDLE",
          "value": "NOT_EQUIPPED"
        },
        {
          "key": "DOOR_BOOT_POSITION",
          "value": "CLOSED"
        },
        {
          "key": "AVG_FUEL_CONSUMPTION",
          "value": "16383"
        },
        {
          "key": "TYRE_PRESSURE_FRONT_LEFT",
          "value": "32767"
        },
        {
          "key": "DOOR_REAR_RIGHT_LOCK_STATUS",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_BELT_STATUS_SECOND_RIGHT",
          "value": "UNKNOWN"
        },
        {
          "key": "POSITION_SPEED_KMPH",
          "value": "0"
        },
        {
          "key": "DOOR_FRONT_RIGHT_POSITION",
          "value": "CLOSED"
        },
        {
          "key": "SEAT_OCCUPATION_STATUS_FRONT_LEFT",
          "value": "UNKNOWN"
        },
        {
          "key": "DOOR_REAR_LEFT_LOCK_STATUS",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_OCCUPATION_STATUS_SECOND_MIDDLE",
          "value": "UNKNOWN"
        },
        {
          "key": "POSITION_TIMESTAMP",
          "value": "2020-05-08T18:13:08Z"
        },
        {
          "key": "BATTERY_VOLTAGE",
          "value": "12.1"
        },
        {
          "key": "DOOR_ENGINE_HOOD_POSITION",
          "value": "CLOSED"
        },
        {
          "key": "TYRE_PRESSURE_REAR_LEFT",
          "value": "32767"
        },
        {
          "key": "WINDOW_FRONT_RIGHT_STATUS",
          "value": "UNKNOWN"
        },
        {
          "key": "WINDOW_REAR_RIGHT_STATUS",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_BELT_STATUS_SECOND_LEFT",
          "value": "UNKNOWN"
        },
        {
          "key": "TYRE_STATUS_FRONT_LEFT",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_OCCUPATION_STATUS_SECOND_LEFT",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_OCCUPATION_STATUS_SECOND_RIGHT",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_OCCUPATION_STATUS_THIRD_LEFT",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_BELT_STATUS_THIRD_LEFT",
          "value": "UNKNOWN"
        },
        {
          "key": "ODOMETER",
          "value": "5478645"
        },
        {
          "key": "POSITION_HEADING_DEGREE",
          "value": "231"
        },
        {
          "key": "WINDOW_FRONT_LEFT_STATUS",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_OCCUPATION_STATUS_THIRD_MIDDLE",
          "value": "UNKNOWN"
        },
        {
          "key": "DOOR_REAR_LEFT_POSITION",
          "value": "CLOSED"
        },
        {
          "key": "VEHICLE_STATE_TYPE",
          "value": "IGNITION_OFF"
        },
        {
          "key": "DOOR_ENGINE_HOOD_LOCK_STATUS",
          "value": "UNKNOWN"
        },
        {
          "key": "WINDOW_SUNROOF_STATUS",
          "value": "UNKNOWN"
        },
        {
          "key": "TYRE_STATUS_FRONT_RIGHT",
          "value": "UNKNOWN"
        },
        {
          "key": "SEAT_BELT_STATUS_FRONT_RIGHT",
          "value": "BELTED"
        },
        {
          "key": "TYRE_STATUS_REAR_RIGHT",
          "value": "UNKNOWN"
        },
        {
          "key": "DOOR_BOOT_LOCK_STATUS",
          "value": "UNKNOWN"
        },
        {
          "key": "DOOR_FRONT_RIGHT_LOCK_STATUS",
          "value": "UNKNOWN"
        }
      ]
    },
    "success": true,
    "vin": "4S4BTGND8L3137058"
  },
  "dataName": "remoteServiceStatus",
  "success": true
}
*/