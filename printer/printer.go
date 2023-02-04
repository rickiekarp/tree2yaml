package printer

import (
	"fmt"
	"sort"

	"git.rickiekarp.net/rickie/tree2yaml/model"
	"gopkg.in/yaml.v2"
)

func PrintFilelist(filelist *model.FileTree) {
	out, err := yaml.Marshal(filelist)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}

func PrintFileListWithOccurrences(fileOccurrenceMap map[int][]model.File) {
	keys := make([]int, 0, len(fileOccurrenceMap))
	for k := range fileOccurrenceMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, k := range keys {
		resultText := fmt.Sprintf("Results (Probability: %d%%)", k)
		fmt.Println(resultText)

		fileSlice := fileOccurrenceMap[k]

		for _, i := range fileSlice {
			fmt.Println(i.Name)
		}
		fmt.Println()
	}
}
