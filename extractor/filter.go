package extractor

import (
	"time"

	"git.rickiekarp.net/rickie/tree2yaml/model"
)

func FilterByModDate(folder *model.Folder, dateFilter string, filterDirection string) *model.Folder {
	date, error := time.Parse("2006-01-02", dateFilter)
	if error != nil {
		return folder
	}

	filteredFolder := []*model.File{}

	for _, file := range folder.Files {
		switch filterDirection {
		case "new":
			if file.LastModified.After(date) {
				filteredFolder = append(filteredFolder, file)
			}
		case "old":
			if file.LastModified.Before(date) {
				filteredFolder = append(filteredFolder, file)
			}
		default:
			if file.LastModified.After(date) {
				filteredFolder = append(filteredFolder, file)
			}
		}
	}

	folder.Files = filteredFolder
	return folder
}
