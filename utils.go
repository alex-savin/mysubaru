package mysubaru

import (
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
)

// isNilFixed is a function
func isNil(i interface{}) bool {
    if i == nil {
	return true
    }
    switch reflect.TypeOf(i).Kind() {
    case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
	return reflect.ValueOf(i).IsNil()
    }
    return false
}

// timeTrack is a function
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Infof("%s took %s\n", name, elapsed)
}

// timestamp is a function
func timestamp() string {
	return string(time.Now().UnixNano() / 1000000)
}
