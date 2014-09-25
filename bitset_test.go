package main

import (
	"testing"
)

func TestSetGetBitTrueSuccess(t *testing.T) {
	bits := make([]byte, 8)
	bitset := BitsetImpl{bits}
	if bitset.Size() != 64 {
		t.Error("size was not 64")
	}
	bitset.SetBit(5, true)
	isSet, err := bitset.GetBit(5)
	if !isSet {
		t.Error("serialized ack does not match the expected data block bytes")
	}
	if err != nil {
		t.Error("error was not null")
	}
}

func TestSetGetBitFalseSuccess(t *testing.T) {
	bits := make([]byte, 8)
	bitset := BitsetImpl{bits}
	if bitset.Size() != 64 {
		t.Error("size was not 64")
	}
	bitset.SetBit(5, true)
	bitset.SetBit(5, false)
	isSet, err := bitset.GetBit(5)
	if isSet {
		t.Error("serialized ack does not match the expected data block bytes")
	}
	if err != nil {
		t.Error("error was not null")
	}
}

func TestSetBitOutOfBoundsFailure(t *testing.T) {
	bits := make([]byte, 8)
	bitset := BitsetImpl{bits}
	err := bitset.SetBit(65, true)
	if err == nil {
		t.Error("error was nil, should have returned an error when setting bit out of bounds")
	}
}

func TestGetBitOutOfBoundsFailure(t *testing.T) {
	bits := make([]byte, 8)
	bitset := BitsetImpl{bits}
	_, err := bitset.GetBit(65)
	if err == nil {
		t.Error("error was nil, should have returned an error when setting bit out of bounds")
	}
}

func TestSetGetAllBitsTrueFalseSuccess(t *testing.T) {
	bits := make([]byte, 20)
	bitset := BitsetImpl{bits}
	for i:=0; i<bitset.Size(); i++ {
		bitset.SetBit(i, true)
	}

	for i:=0; i<bitset.Size(); i++ {
		isSet, err := bitset.GetBit(5)
		if !isSet {
			t.Error("serialized ack does not match the expected data block bytes")
		}
		if err != nil {
			t.Error("error was not null")
		}
	}

	for i:=0; i<bitset.Size(); i++ {
		bitset.SetBit(i, false)
	}

	for i:=0; i<bitset.Size(); i++ {
		isSet, err := bitset.GetBit(5)
		if isSet {
			t.Error("serialized ack does not match the expected data block bytes")
		}
		if err != nil {
			t.Error("error was not null")
		}
	}
}
