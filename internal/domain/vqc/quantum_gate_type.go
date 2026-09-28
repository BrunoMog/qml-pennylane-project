package vqc

import (
	"slices"
	"strings"
)

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

const maxGateTypeLength = 10

func (g GateType) String() string {
	return string(g)
}

var singleQubitGates = []GateType{HGate, XGate, YGate, ZGate, RXGate, RYGate, RZGate}
var twoQubitGates = []GateType{CNOTGate}

func (g GateType) isValid() bool {
	switch g {
	case HGate, XGate, YGate, ZGate, RXGate, RYGate, RZGate, CNOTGate:
		return true
	default:
		return false
	}
}

func (g GateType) isSingleQubitGate() bool {
	return slices.Contains(singleQubitGates, g)
}

func (g GateType) isTwoQubitGate() bool {
	return slices.Contains(twoQubitGates, g)
}

func ParseGateType(gateTypeStr string) (GateType, error) {
	if len(gateTypeStr) > maxGateTypeLength {
		return "", ErrParseGateType
	}
	switch strings.ToLower(strings.TrimSpace(gateTypeStr)) {
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
		return "", ErrParseGateType
	}
}
