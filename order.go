package main

import "fmt"

type order struct {
	ProductCode int
	Quantity    float64
	Status      orderStatus
}

func (o order) String() string {
	return fmt.Sprintf("ProductCode: %d, Quantity: %f, Status: %d", o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

func orderStatusToText(status orderStatus) string {
	switch status {
	case new:
		return "new"
	case received:
		return "received"
	case reserved:
		return "reserved"
	case filled:
		return "filled"
	default:
		return "none"
	}
}

type orderStatus int

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)
