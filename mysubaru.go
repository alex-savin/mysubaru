package mysubaru

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Version is constant for the package
const Version = "0.0.0-alpha"

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

// New function creates a New Hassky client
func New() (*Client, error) {
	defer timeTrack(time.Now(), "NEW CLIENT")

	logLevel := "debug"
	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
		log.Warnf("[MY SUBARU] Invalid log level supplied: '%s'", logLevel)
	}

	viper.SetConfigName("mysubaru") // name of config file (without extension)
	viper.SetConfigType("yml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(err)
	}

	switch viper.GetString("log_level") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
		log.Warnf("[MY SUBARU] Invalid log level supplied: '%s'", logLevel)
	}

	httpClient := resty.New()
	httpClient.SetHostURL("https://www.mysubaru.com")
	httpClient.R().Get("/")

	httpClient.R().SetFormData(map[string]string{"username": viper.GetString("credentials.username"), "password": viper.GetString("credentials.password")}).Post("/login")
	resp, _ := httpClient.R().SetQueryString("now=" + timestamp()).Get("/profile/getSecurityQuestion.json")

	dec := json.NewDecoder(bytes.NewReader(resp.Body()))
	dec.UseNumber()
	respParsed, err := gabs.ParseJSONDecoder(dec)
	if err != nil {
		log.Fatalf(err.Error())
	}

	intI64, _ := respParsed.Search("challengeQuestionKey").Data().(json.Number).Int64()
	intSTR := strconv.FormatInt(intI64, 10)

	resp, _ = httpClient.R().SetFormData(map[string]string{"questionId": intSTR, "answer": viper.GetString("credentials.questions." + intSTR + ".answer"), "deviceId": timestamp(), "deviceName": "My App"}).Post("/account/securityAnswer.json")
	log.Debugf(string(resp.Body()))
	if string(resp.Body()) != "true" {
		return nil, errors.New("Not Authorized! Please check your credentials ")
	}

	client := &Client{
		httpClient: httpClient,
	}

	return client, nil
}
