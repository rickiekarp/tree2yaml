package eventsender

type QueueEventType string

const (
	FilestoreAdd       QueueEventType = "filestore_add"
	FilestoreAddExtras QueueEventType = "filestore_add_extras"
)

type HubQueueEventMessage struct {
	Event   QueueEventType `json:"event,omitempty"`
	Payload string         `json:"payload,omitempty"`
}

type FileStorageEventMessage struct {
	Id             *int64                                   `json:"id,omitempty"`
	Path           string                                   `json:"path,omitempty"`
	Name           string                                   `json:"name,omitempty"`
	Size           int64                                    `json:"size,omitempty"`
	Mtime          int64                                    `json:"mtime,omitempty"`
	Checksum       *string                                  `json:"checksum,omitempty"`
	Owner          *string                                  `json:"owner,omitempty"`
	Inserttime     *int64                                   `json:"inserttime,omitempty"`
	Lastupdate     *int64                                   `json:"lastupdate,omitempty"`
	AdditionalData *[]FileStorageAdditionalDataEventMessage `json:"additional_data,omitempty"`
}

type FileStorageAdditionalDataEventMessage struct {
	FilesId  *int64 `json:"file_id,omitempty"`
	Property string `json:"property,omitempty"`
	Value    string `json:"value,omitempty"`
}
