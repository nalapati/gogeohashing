package main

import (
	"fmt"
	"os"
	"strconv"
)

// Geohash returns a geohash representation of the given latitude
// and longitude, This implementation returns a 60bit geohash..
func GeoHash(latitude float64, longitude float64) (Bitset, error) {
	// num-bits
	numBits := 60

	bits := make([]byte, 8)
	geohash := BitsetImpl{bits}

	// longitude min: -180, mid: 0 max: 180
	// latitude min: -180, mid: 0 max: 180
	// convention is min to mid is 0, and > mid to max is 1
	longMin := float64(-180.0)
	longMid := float64(0.0)
	longMax := float64(180.0)

	latMin := float64(-90.0)
	latMid := float64(0.0)
	latMax := float64(90.0)

	isLongitude := true

	for i:=0; i<numBits; i++ {
		if isLongitude {
			if longitude>=longMid {
				geohash.SetBit(i, true)
				longMin = longMid
			} else {
				longMax = longMid
			}
			longMid = (longMax + longMin) / 2
			isLongitude = false
		} else {
			if latitude>=latMid {
				geohash.SetBit(i, true)
				latMin = latMid
			} else {
				latMax = latMid
			}
			latMid = (latMax + latMin) / 2
			isLongitude = true
		}
	}
	return geohash, nil
}

func usage(val int) {
	fmt.Println("usage: ./main <latitude> <longitude>")
	os.Exit(val)
}

func main() {
	if len(os.Args) < 2 {
		usage(1)
	}

	latitude, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println(err)
		usage(1)
	}

	longitude, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println(err)
		usage(1)
	}

	geohash, err := GeoHash(latitude, longitude)

	geohashChars := "0123456789bcdefghjkmnpqrstuvwxyz"

	for i:=0; i<geohash.Size()/5; i++ {
		n:=0
		bit:=0
		for j:=0; j<5; j++ {
			res, _ := geohash.GetBit(i*5+j)
			if res {
				bit=1
			} else {
				bit=0
			}
			n=n*2+bit
		}
		fmt.Print(string(geohashChars[n]))
	}
	fmt.Println("")
}
