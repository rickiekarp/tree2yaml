# Tree2Yaml

Generates a file list for a given path in yaml format

### Available commands
```
Usage of tree2yaml:
  -enableArchive
    	generates archive of the generated filelist
  -enableMetadata
    	generates metadata of the generated filelist
  -filterByDate string
    	filters files by given date, e.g. -filterByDate=2022-12-24
  -filterByDateDirection string
    	direction of files to be filtered, e.g. 'old', 'new' (default "new")
  -find string
    	prints all file names that match the given arguments, grouped by occurrence, e.g. -find=foo,bar
  -findArchivedFiles
    	prints archived files
  -findFilesIn string
    	finds files by a given search path, e.g. tree2yaml -load -findFilesIn=foo/bar /foo/bar.yaml
  -findFoldersIn string
    	finds folders by a given search path, e.g. tree2yaml -load -findFoldersIn=foo/bar /foo/bar.yaml
  -generate
    	generates a filelist (default true)
  -generateMetadataFromFile
    	load a file list file
  -git
    	check git history
  -git-depth int
    	git log depth (default 3)
  -groupByRevision
    	groups the file archive results by the number of revisions
  -groupByRevisionLimit int
    	limits results by the number of revisions (0 = no limit)
  -hash string
    	calculate hash sum of each file (crc32, crc64, md5)
  -help
    	prints all available options
  -ignoreCase
    	ignore case when matching files, can be combined with -find flag
  -load
    	loads an existing filelist
  -outfile string
    	path of the output file

```


## Create file list

```
go run main.go -outfile=test/filelist_test.yaml test/rootFolder
```

### Create filelist with metadata

```
go run main.go -outfile=test/filelist_test.yaml -enableMetadata test/rootFolder
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

## Create file list archive

### List files grouped and limited by revision
Input:
```
go run main.go -outfile=test/filelist_test.yaml -load -findArchivedFiles -groupByRevision -groupByRevisionLimit=2 test/filelist_test.yaml.archive
```

Output:
```
--- Files with number of revisions: 4
fileB

--- Files with number of revisions: 3
subFileA

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

### File Events
Input
```
go run main.go -eventsEnabled -eventHost=localhost:12000 foo/
```
The above command will enable file events to be sent to the given `eventHost`.
