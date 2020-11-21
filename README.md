# Dinah

[![PkgGoDev](https://pkg.go.dev/badge/github.com/fischersean/dinah)](https://pkg.go.dev/github.com/fischersean/dinah) [![Go Report Card](https://goreportcard.com/badge/github.com/fischersean/dinah)](https://goreportcard.com/badge/github.com/fischersean/dinah)

 The Dinah package provides data structures and fetch functionality for MLB's Statcast and Chadwick's Register datasets.

## Installation

`go get -u github.com/fischersean/dinah`

## Example

Consider we wanted to download pitch by pitch Statcast data for the month of September 2020. 

```go
package main

import (
    "fmt"
    "time"
    "github.com/fischersean/dianh/statcast"
)

func main(){
  
    d0 := time.Date(2020, time.September, 1, 0, 0, 0, 0, time.UTC)
    d1 := time.Date(2020, time.September, 30, 0, 0, 0, 0, time.UTC)

    ds, err := statcast.FromHttp(d0, d1)
  
    if err != nil {
        panic(err)
    }
  
    fmt.Println("%#v\n", ds)

}
```

If we wanted to patch up the MLBAM id's for the players within the Statcast data, we could pull the mapping from Chadwick's people register. It's recommended to use a local file for this. The full dataset is 45+MB and my dramatically slow execution time.

```go
package main

import (
    "fmt"
    "time"
    "github.com/fischersean/dianh/chadwick"
)

func main(){
  
    people, err := chadwick.PeopleFromCsv("cwickpeople.csv")

    if err != nil {
        panic(err)
    }
  
    fmt.Println("%#v\n", people)
  
}
```



## Acknowledgments

[Chadwick people register](https://github.com/chadwickbureau/register)

[Baseball Savant](https://baseballsavant.mlb.com)

