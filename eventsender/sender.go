package eventsender

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"git.rickiekarp.net/rickie/tree2yaml/model"
	"github.com/sirupsen/logrus"
)

var FlagEventsEnabled = flag.Bool("eventsEnabled", false, "whether to send file events")
var FlagFileEventOwner = flag.String("eventFilelistOwner", "", "owner of the filelist event entry")

var EventSenderProtocol = "http"        // Version set during go build using ldflags
var EventTargetHost = "localhost:12000" // Version set during go build using ldflags

func SendFileEvent(fileEvent FileStorageEventMessage) {
	url := EventSenderProtocol + "://" + EventTargetHost + "/hub/v1/queue/push"

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

	logrus.Debug("SendFileEvent:Status ", resp.Status, " - ", url)
}

func SendEventForFile(file model.File) {
	if *FlagEventsEnabled {

		// the modifiedTime can be < 0 for old files, so we make sure here it fits into an unsigned integer
		modifiedTime := file.LastModified.Unix()
		if modifiedTime < 0 {
			modifiedTime = 0
		}

		// prepare and send FileStorage event message
		event := FileStorageEventMessage{
			Path:  file.Path,
			Name:  file.Name,
			Size:  file.Size,
			Mtime: modifiedTime,
		}

		if len(*FlagFileEventOwner) > 0 {
			event.Owner = FlagFileEventOwner
		}

		SendFileEvent(event)
	}
}
