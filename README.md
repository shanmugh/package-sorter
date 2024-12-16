# package-sorter

## Instructions

Follow the below steps to build the CLI

```bash
make deps
make build
make run
```

## Sample Run

```bash
make deps
go mod tidy
go mod verify
all modules verified

make build
go build -o package-sorter main.go

make run
go build -o package-sorter main.go
./package-sorter -i ./input.json
{
  "standard": [
    {
      "name": "standard item",
      "width": 10,
      "height": 10,
      "length": 10,
      "mass": 10
    }
  ],
  "special": [
    {
      "name": "special bulky item by volume",
      "width": 100,
      "height": 100,
      "length": 100,
      "mass": 10
    },
    {
      "name": "special bulky item by length",
      "width": 10,
      "height": 10,
      "length": 150,
      "mass": 10
    },
    {
      "name": "special heavy item",
      "width": 10,
      "height": 10,
      "length": 10,
      "mass": 20
    }
  ],
  "rejected": [
    {
      "name": "rejected heavy and bulky item by volume",
      "width": 100,
      "height": 100,
      "length": 100,
      "mass": 20
    },
    {
      "name": "rejected heavy and bulky item by length",
      "width": 10,
      "height": 10,
      "length": 150,
      "mass": 20
    }
  ]
}
```

## Custom Run

To run with custom input

```bash
# Check help of the CLI
make help
# To use custom input, create a json input based on input.json
./package-sorter -i <custom.json>
```

## Library Use

The library function can be imported and used like below,

```golang
import "github.com/shanmugh/package-sorter/pkg/sorter"

func UseSorter() {
  s := sorter.NewSorter(150, 1_000_000_000, 20)
  stack := s.Sort(p.Width, p.Height, p.Length, p.Mass)
  ...
}
```

## Pre-requisites

```bash
make deps
```

## Run Application

```bash
make run
```

## Developing Application

### Building

```bash
make build
```

### Testing

```bash
make test
```

### Linting

```bash
make lint
```
