package eventsender

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"git.rickiekarp.net/rickie/gomain"
	"git.rickiekarp.net/rickie/tree2yaml/model"
	"git.rickiekarp.net/rickie/yubase"
	"github.com/sirupsen/logrus"
)

var FlagEventsEnabled = flag.Bool("eventsEnabled", false, "whether to send file events")
var FlagFileCategory = flag.Int("eventFileCategory", 0, "category of the file")

var EventSenderProtocol = "http"        // Version set during go build using ldflags
var EventTargetHost = "localhost:12000" // Version set during go build using ldflags

func sendFileEvent(fileEvent yubase.FileListEntry) {
	url := EventSenderProtocol + "://" + EventTargetHost + yubase.ApiHubQueuePush

	fileEventPayloadBytes, err := json.Marshal(fileEvent)
	if err != nil {
		log.Fatal(err)
	}

	eventMessage := yubase.YuQueueEventMessage{
		Id:    gomain.NewUUIDv4(),
		Event: yubase.FilestoreAdd,
		Payload: yubase.YuMessagePayload{
			Body:     fileEventPayloadBytes,
			Encoding: "base64",
		},
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

func SendEventForFile(file model.File, processId *int64) {
	if *FlagEventsEnabled {

		// prepare and send FileStorage event message
		event := yubase.FileListEntry{
			Path:      file.Path,
			Name:      file.Name,
			Size:      file.Size,
			Mtime:     file.LastModified.Unix(),
			Category:  FlagFileCategory,
			ProcessId: processId,
		}

		sendFileEvent(event)
	}
}
