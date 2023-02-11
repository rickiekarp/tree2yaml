# Tree2Yaml

Generates a file list for a given path in yaml format

### Available commands
```
Usage of tree2yaml:
  -calcMd5
    	calculate md5 sum of each file
  -filterByDate string
    	filters files by given date, e.g. -filterByDate=2022-12-24
  -filterByDateDirection string
    	direction of files to be filtered, e.g. 'old', 'new' (default "new")
  -find string
    	prints all file names that match the given arguments, grouped by occurrence, e.g. -find=foo,bar
  -findFilesIn string
    	finds files by a given search path, e.g. tree2yaml -load -findFilesIn=foo/bar /foo/bar.yaml
  -findFoldersIn string
    	finds folders by a given search path, e.g. tree2yaml -load -findFoldersIn=foo/bar /foo/bar.yaml
  -generate
    	generates a filelist (default true)
  -git
    	check git history
  -git-depth int
    	git log depth (default 3)
  -help
    	prints all available options
  -ignoreCase
    	ignore case when matching files, can be combined with -find flag
  -load
    	loads an existing filelist
```


## Create file list

```
go run main.go test/rootFolder > test/filelist_test.yaml
```

## Load file list
```
go run main.go -load test/filelist_test.yaml
```

### Find files in a specific folder
Input:
```
go run main.go -load -findFilesIn=rootFolder/folderA test/filelist_test.yaml
```

Output:
```
fileA
```

### Find folders in a specific folder
Input:
```
go run main.go -load -findFoldersIn=rootFolder test/filelist_test.yaml
```

Outut:
```
folderA
folderB
```

### Find files and list their number of occurrences

Used to find files by given parts of their file name.
Lists all files and the probability of how many parts of the file name are found in a file name.

Input:
```
go run main.go -find=file,A -load test/filelist_test.yaml

# you can also toggle case sensitivity by providing -ignoreCase flag
go run main.go -find=file,A -ignoreCase -load test/filelist_test.yaml
```

Output:
```
Results (Probability: 100%)
fileA

Results (Probability: 50%)
fileB
fileB
```

#### Git Mode

Used to find all occurrences of `fileA` and `fileB` by checking all file lists `test/filelist_test.yaml` 
of the last n commits in the git history.

Input
```
go run main.go -find=fileA,fileB -git -git-depth=3 -load test/filelist_test.yaml
```

Output:
```
Results (Probability: 50%)
fileA
fileB
```

### Filtering

#### Filter by date
Input:
```
go run main.go -load -findFilesIn=rootFolder/folderA -filterByDate=2022-01-24 -filterByDateDirection=new test/filelist_test.yaml
```
Above command will list all files in `rootFolder/folderA` and filter by date `2022-01-24` only showing files `new`er than the given date.

Output:
```
fileA
fileB
```
