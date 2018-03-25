# Raptor Management

## Purpose
This service uses Go and InfluxDB to keep track of time-traveling raptors to ensure
safe travels for human time-travelers.

## Setup 
0. Ensure go is installed. If not, follow the steps provided here: https://golang.org/doc/install
1. Ensure the main.go has been extracted and is in a location in your GOPATH
- Check your Go environment settings by: `go env`
2. Install InfluxDB

Resource: https://portal.influxdata.com/downloads (select stable version of InfluxDB to see options)
- On OSX: After install `homebrew`, run `brew update && brew install influxdb`
- Docker: `docker pull influxdb`
- Ubuntu & Debian: `wget https://dl.influxdata.com/influxdb/releases/influxdb_1.5.1_amd64.deb && sudo dpkg -i influxdb_1.5.1_amd64.deb`

### Running the Service
3. Build the service using the `build.sh` or use `go build` or `go install`
4. Ensure InfluxDB is running and that you have its hostname, port, and databasename.
- if installed, you can just run quickly using defaults: `> influxd`
5. Run the service:
-- You can use the binary to load raptor data into InfluxDB using the following flags:

|Flag|Type|Description|
|----|----|-----------|
|-raptors|String|The name of the CSV with raptor data.|

-- You can use the binary to run the temporal querying service by using the following flags:

|Flag|Type|Description|
|----|----|-----------|
|-raptors|String|The name of the CSV with raptor data.|
|-port|String|The port for the service to run on. Default is `9090`.|
|-influx-host|String|The hostname for the InfluxDB instance. Default is `0.0.0.0`.|
|-influx-post|String|The port for the InfluxDB instance. Default is `8086`.|
|-database-name|String|The name of the database to use within InfluxDB. Default is `raptorTs`.|
|-serve|Boolean|When present, directs the program to run as a service. Default and absence are `false`.|

You can load and serve simply by combining flags. For example, if you've named your binary `raptor-service`:
> `./raptor-service -raptors timeline_data.csv -serve -port 9091 -database-name raptorTs`

## Querying
### Request
URL: /raptors

Method: GET

Parameters: `time=YYYY-MM-DD`

### Response
Body:
```json
{
    "count": 5
}
```
Example:
```bash
curl -XGET "localhost:9091/raptors?time=2000-01-01"
```
```json
{
      "count": 100
}
```

## Updating
### Request
URL: /raptors

Method: GET

Parameters:
- `time=YYYY-MM-DD`
- `contained=XX` where XX is a positive integer

### Response
Body:
```json
{
    "count": 5,
    "time": <time>
}
Example:

```bash
curl -XGET "localhost:9091/raptors?time=2000-01-01&contained=11"
```

```json
{
      "count": 0,
      "time": "2000-01-01T00:00:00Z"
}
```

### Errors
Body:
```json
{
    "error": "<error message>"
}
```
