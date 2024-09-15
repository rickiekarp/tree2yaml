package eventsender

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"git.rickiekarp.net/rickie/tree2yaml/hash"
	"git.rickiekarp.net/rickie/tree2yaml/model"
	"github.com/sirupsen/logrus"
)

var FlagEventsEnabled = flag.Bool("eventsEnabled", false, "whether to send file events")
var flagEventHost = flag.String("eventHost", "api.rickiekarp.net", "event host to send file events to")
var flagEventProtocol = flag.String("eventProtocol", "https", "event port to send file events to")

func SendFileEvent(fileEvent FileStorageEventMessage) {
	url := *flagEventProtocol + "://" + *flagEventHost + "/hub/v1/queue/push"

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
		// prepare and send FileStorage event message
		filePathDir := file.Path
		pathHash := hash.CalcSha1(filePathDir)
		fileChecksum := string(file.Sha1())

		modifiedTime := file.LastModified.Unix()
		if modifiedTime < 0 {
			modifiedTime = 0
		}

		event := FileStorageEventMessage{
			Path:     filePathDir,
			Name:     file.Name,
			Size:     file.Size,
			Mtime:    modifiedTime,
			Checksum: hash.CalcSha1(pathHash + "/" + fileChecksum),
		}
		SendFileEvent(event)
	}
}
