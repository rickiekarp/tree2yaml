
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

### Find files and list their number of occurrences

Used to find files by given parts of their file name.
Lists all files and the probability of how many parts of the file name are found in a file name.

Input:
```
go run main.go --find=file,A --load test/filelist_test.yaml
```

Output:
```
Results (Probability: 100%)
fileA

Results (Probability: 50%)
fileB
fileB
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
