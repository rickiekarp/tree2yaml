
## Create file list

```
go run main.go test/rootFolder > test/filelist_test.yaml
```

## Load file list
```
go run main.go --load test/filelist_test.yaml
```

### Find files in a specific folder
Input:
```
go run main.go --load --findFilesIn=rootFolder/folderA test/filelist_test.yaml
```

Output:
```
fileA
```

### Find folders in a specific folder
Input:
```
go run main.go --load --findFoldersIn=rootFolder test/filelist_test.yaml
```

Outut:
```
folderA
folderB
```

### Filtering

#### Filter by date
Input:
```
go run main.go --load --findFilesIn=rootFolder/folderA --filterByDate=2022-01-24 --filterByDateDirection=new test/filelist_test.yaml
```
Above command will list all files in `rootFolder/folderA` and filter by date `2022-01-24` only showing files `new`er than the given date.

Output:
```
fileA
fileB
```
