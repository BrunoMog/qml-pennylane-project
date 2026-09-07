package vqc

import "strings"

type GateType string

const (
	HGate    GateType = "h"
	XGate    GateType = "x"
	YGate    GateType = "y"
	ZGate    GateType = "z"
	RXGate   GateType = "rx"
	RYGate   GateType = "ry"
	RZGate   GateType = "rz"
	CNOTGate GateType = "cnot"
)

func (g GateType) Value() string {
	return string(g)
}

var singleQubitGates = []GateType{HGate, XGate, YGate, ZGate, RXGate, RYGate, RZGate}
var twoQubitGates = []GateType{CNOTGate}

func ParseGateType(gateTypeStr string) (GateType, error) {
	switch strings.ToLower(gateTypeStr) {
	case "h":
		return HGate, nil
	case "x":
		return XGate, nil
	case "y":
		return YGate, nil
	case "z":
		return ZGate, nil
	case "rx":
		return RXGate, nil
	case "ry":
		return RYGate, nil
	case "rz":
		return RZGate, nil
	case "cnot":
		return CNOTGate, nil
	default:
		return "", &InvalidParseGateTypeError{gateTypeStr}
	}
}
