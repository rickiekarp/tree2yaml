# Tree2Yaml

Generates a file list for a given path in yaml format

### Available commands
```
Usage of tree2yaml:
  -enableMetadata
    	generates metadata of the generated filelist
  -eventFilelistOwner string
    	owner of the filelist event entry (default "default")
  -eventsEnabled
    	whether to send file events
  -filterByDate string
    	filters files by given date, e.g. -filterByDate=2022-12-24
  -filterByDateDirection string
    	direction of files to be filtered, e.g. 'old', 'new' (default "new")
  -findFilesIn string
    	finds files by a given search path, e.g. tree2yaml -load -findFilesIn=foo/bar /foo/bar.yaml
  -findFoldersIn string
    	finds folders by a given search path, e.g. tree2yaml -load -findFoldersIn=foo/bar /foo/bar.yaml
  -generate
    	generates a filelist (default true)
  -generateMetadataFromFile
    	load a file list file
  -hash string
    	calculate hash sum of each file (crc32, crc64, md5)
  -help
    	prints all available options
  -load
    	loads an existing filelist
  -outfile string
    	path of the output file
  -v	prints version
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

## Create file list archive

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
