package eventsender

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

var FlagEventsEnabled = flag.Bool("eventsEnabled", false, "whether to send file events")
var flagEventHost = flag.String("eventHost", "api.rickiekarp.net", "event url to send file events to")

func SendFileEvent(fileEvent FileStorageEventMessage) {
	url := "http://" + *flagEventHost + "/hub/v1/queue/push"

	tmpData, err := json.Marshal(fileEvent)
	if err != nil {
		log.Fatal(err)
	}

	eventMessage := HubQueueEventMessage{
		Event:   FilestoreAdd,
		Payload: string(tmpData),
	}

	jsonData, err := json.Marshal(eventMessage)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	logrus.Debug("SendFileEvent:Status ", resp.Status)
}
