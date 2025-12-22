package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"git.rickiekarp.net/rickie/nexusform"
	"git.rickiekarp.net/rickie/tree2yaml/eventsender"
	"github.com/sirupsen/logrus"
)

func SendPreferenceUpdate(processId int64) {
	url := eventsender.EventSenderProtocol + "://" + eventsender.EventTargetHost + "/storage/v1/preferences"

	preferenceData := nexusform.StoragePreference{
		Property: "activelist",
		Value:    strconv.FormatInt(processId, 10),
	}

	jsonData, err := json.Marshal(preferenceData)
	if err != nil {
		log.Fatal(err)
	}

	// Create PATCH request
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(jsonData))
	if err != nil {
		logrus.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Println(err)
	}
	defer resp.Body.Close()

	logrus.Debug("SendPreferenceUpdate:Status ", resp.Status, " - ", url)
}
