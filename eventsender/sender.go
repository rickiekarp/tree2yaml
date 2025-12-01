package eventsender

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"git.rickiekarp.net/rickie/nexuscore"
	"git.rickiekarp.net/rickie/nexusform"
	"git.rickiekarp.net/rickie/tree2yaml/model"
	"github.com/sirupsen/logrus"
)

var FlagEventsEnabled = flag.Bool("eventsEnabled", false, "whether to send file events")
var FlagFileEventOwner = flag.String("eventFilelistOwner", "default", "owner of the filelist event entry")

var EventSenderProtocol = "http"        // Version set during go build using ldflags
var EventTargetHost = "localhost:12000" // Version set during go build using ldflags

func sendFileEvent(fileEvent nexusform.FileListEntry) {
	url := EventSenderProtocol + "://" + EventTargetHost + nexuscore.ApiHubQueuePush

	fileEventPayloadBytes, err := json.Marshal(fileEvent)
	if err != nil {
		log.Fatal(err)
	}

	eventMessage := nexusform.HubQueueEventMessage{
		Event:   nexusform.FilestoreAdd,
		Payload: string(fileEventPayloadBytes),
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
		event := nexusform.FileListEntry{
			Path:      file.Path,
			Name:      file.Name,
			Size:      file.Size,
			Mtime:     file.LastModified.Unix(),
			Category:  FlagFileEventOwner,
			ProcessId: processId,
		}

		sendFileEvent(event)
	}
}
