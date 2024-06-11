## CLI tool for quick data normalization

This simple CLI tool currently only allows for quick *normalization* or *standardisation* of csv data for ML applications and saves it to your device.

### Download:
Go to assets under the most recent release. Download the appropriate binary and put it in a PATH.

#### Use -> through the CLI:
<p> To apply a standardisation process to your dataset, simply run the program and specify the process and file. If ran without -e (the export flag) the file will automatically save to the current working directory as 'data.csv'. </p>

```
norm stand data.csv
```
```
norm nm data.csv
```
<p> To specify the location to save the file to add the -e flag and export file path:
*Example within Desktop:*

```
norm stand data.csv -e ../../Desktop/data.csv
```
```
norm nm data.csv -e ../../Desktop/data.csv
```

