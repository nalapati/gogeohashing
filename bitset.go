package main

import (
	"errors"
)

type Bitset interface {
	SetBit(int, bool) error
	GetBit(int) (bool, error)
	Size() int
}

type BitsetImpl struct {
	bits []byte
}

func (b BitsetImpl) SetBit(position int, val bool) error {
	if position>=b.Size() || position<0 {
		return errors.New("position must be less than size and atleast 0")
	}
	index := position/8
	bit := uint8(position%8)
	if val {
		b.bits[index] |= 1<<bit
	} else {
		b.bits[index] &= (0xFF ^ 1<<bit)
	}
	return nil
}

func (b BitsetImpl) GetBit(position int) (bool, error) {
	if position>=b.Size() || position<0 {
		return false, errors.New("position must be less than size and atleast 0")
	}
	index := position/8
	bit := uint8(position%8)
	return (b.bits[index] & (1<<bit)) != 0, nil
}

func (b BitsetImpl) Size() int {
	return len(b.bits)*8
}
