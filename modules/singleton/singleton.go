package singleton

import "../../types"

var pings []types.Pong

func GetPings() []types.Pong {
	return pings
}

func SetPings(m []types.Pong) {
	pings = m
}

