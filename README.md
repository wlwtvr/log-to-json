# Log to JSON

this command turn log text into json.

## Example
input:
```
name:"Fikar" age:10 address:{city:"Jakarta Selatan" state:"DKI Jakarta"} gender:"male"
```
output:
```
{
  "address": {
    "city": "Jakarta Selatan",
    "state": "DKI Jakarta"
  },
  "age": 10,
  "gender": "male",
  "name": "Fikar"
}
```

## How to use

### Prerequisites

Build the binary
```
$ go build -o logtojson main.go
```

#### Pass text as argument
```
$ ./logtojson "name:\"Fikar\" age:10 address:{city:\"Jakarta Selatan\" state:\"DKI Jakarta\"} gender:\"male\""
```

#### Pass text using pipe
1. copy the text (log) that you want to turn into json into clipboard
2. run
```
$ pbpaste | ./logjson
```

