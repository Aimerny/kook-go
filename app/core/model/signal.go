package model

import (
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

// ws 连接相关
type Signal struct {
	SignalType   SignalType     `json:"s"`
	SerialNumber int            `json:"sn"`
	Data         map[string]any `json:"d"`
}
type SignalCode int
type SignalType int

const (
	Success          SignalCode = 0
	ConnMissParam    SignalCode = 40100
	InvalidToken     SignalCode = 40101
	ValidFailedToken SignalCode = 40102
	ExpiredToken     SignalCode = 40103

	ResumeMissParam SignalCode = 40106
	ExpiredSession  SignalCode = 40107
	InvalidSN       SignalCode = 40108
)

const (
	SIG_EVENT      SignalType = 0
	SIG_HELLO      SignalType = 1
	SIG_PING       SignalType = 2
	SIG_PONG       SignalType = 3
	SIG_RESUME     SignalType = 4
	SIG_RECONNECT  SignalType = 5
	SIG_RESUME_ACK SignalType = 6
)

func ParseSignal(data []byte) *Signal {
	res := &Signal{}
	err := jsoniter.Unmarshal(data, res)
	if err != nil {
		log.WithError(err).Errorf("parse signal failed")
		return nil
	}
	return res
}

func NewPing(sn int) *Signal {
	return &Signal{
		SignalType:   SIG_PING,
		SerialNumber: sn,
	}
}
