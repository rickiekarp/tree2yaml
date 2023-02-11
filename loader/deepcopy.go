package loader

import "git.rickiekarp.net/rickie/tree2yaml/model"

func deepCopyFileMap(src, dst map[int][]model.File) map[int][]model.File {

	var tmp = make(map[model.File]int)
	var newMap = make(map[int][]model.File)

	for srcProbability, srcFiles := range src {
		newMap[srcProbability] = []model.File{}
		for _, file := range srcFiles {
			tmp[file] = srcProbability
		}
	}

	for srcProbability, srcFiles := range dst {
		newMap[srcProbability] = []model.File{}
		for _, file := range srcFiles {
			tmp[file] = srcProbability
		}
	}

	for file, probability := range tmp {
		var entry = newMap[probability]
		entry = append(entry, file)
		newMap[probability] = entry
	}

	return newMap
}
