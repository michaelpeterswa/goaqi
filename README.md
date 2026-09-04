# goaqi

A Go library that computes the US EPA Air Quality Index (AQI) from pollutant
concentrations, following the [AirNow Technical Assistance Document for the
Reporting of Daily Air Quality – the Air Quality Index (AQI)](https://www.airnow.gov/sites/default/files/2020-05/aqi-technical-assistance-document-sept2018.pdf)
(September 2018).

```sh
go get github.com/michaelpeterswa/goaqi
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/michaelpeterswa/goaqi"
)

func main() {
	// 24 hour average PM2.5 concentration in µg/m3.
	index, err := goaqi.AQIPM25(30.2)
	if err != nil {
		log.Fatal(err)
	}

	designation, err := goaqi.AQIDesignationFromIndex(index)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("AQI %d (%s)\n", index, designation) // AQI 89 (Moderate)
}
```

`AQIPM25` and `AQIPM100` take a 24 hour average concentration in µg/m3 and
return the index. `AQIDesignationFromIndex` maps an index back to its category
name. All three return `goaqi.ErrBeyondTheScale` for inputs outside the defined
breakpoints, including negative values.

The PM10 concentration is truncated to a whole number before lookup, as the
document specifies. PM2.5 is not truncated; do that before calling if your
source requires it.

## Development

```sh
make        # wire up git hooks and install commitlint
make test   # go test -race -shuffle=on ./...
make lint   # golangci-lint and yamllint, the same checks CI runs
```

Commits follow [Conventional Commits](https://www.conventionalcommits.org/);
semantic-release cuts a GitHub release from them on every push to `main`.

## License

MIT
