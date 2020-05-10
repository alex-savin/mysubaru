package mysubaru

import (
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

// Client is a struct
type Client struct {
	httpClient *resty.Client
}

// R method creates a new Request instances to be use lates while the firing a request with client
func (c *Client) R() *Request {
	req := &Request{
		body:   map[string]string{},
		client: c,
	}
	return req
}

// Exec method executes a Client instance with the Request
func (c *Client) execute(req *Request) []byte {
	defer timeTrack(time.Now(), "Executing")

	resp, _ := c.httpClient.R().SetFormData(req.body).Post(req.urlPathExecute)
	jsonParsed, err := gabs.ParseJSON(resp.Body())
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Debugf("%+v", jsonParsed)

	if req.urlPathStatus != "" {
		serviceRequestID, _ := jsonParsed.Path("data.serviceRequestId").Data().(string)
		for {
			time.Sleep(5 * time.Second)
			resp, _ := c.httpClient.R().SetFormData(map[string]string{"serviceRequestId": serviceRequestID}).Post(req.urlPathStatus)
			jsonParsed, err := gabs.ParseJSON(resp.Body())
			if err != nil {
				log.Fatalf(err.Error())
			}
			log.Debugf("%+v", jsonParsed)

			if success, _ := jsonParsed.Path("data.success").Data().(bool); success == true {
				break
			}
		}
	}

	return resp.Body()
}
