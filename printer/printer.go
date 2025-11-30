package printer

import (
	"fmt"

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
