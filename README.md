# number-converter

Number system converter for Golang

### Installation

    go get -u github.com/kenticny/numconvert

### Usage

#### Convert(number int64, radix int64) (string, error)

> notice: radix maximum is 62

    import (
      ncv "github.com/kenticny/numconvert"
    )

    if result, err := ncv.Convert(123456, 20); err != nil {
      panic(err)
    } else {
      fmt.Println(result) // f8cg
    }
