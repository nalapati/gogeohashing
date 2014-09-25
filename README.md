#### Geohashing utility in Go
##### Setup:
```
$ go get github.com/nalapati/gogeohashing
```
##### Usage:
```
$ go $GOPATH/bin/gogeohashing <latitude> <longitude>

All parameters are required:
<latitude> -90.0,90.0 (50N -> 50.0, 60S -> -60.0)
<longitude> -180.0,180.0 (120E -> 120.0, 80W -> -80.0)
```
##### Example:
```
$ go $GOPATH/bin/gogeohashing 47.609700 -122.333100
```
