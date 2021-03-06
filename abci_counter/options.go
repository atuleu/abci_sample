package main

import "flag"

type options struct {
	Address  string
	ABCIType string
	Serial   bool
	Verbose  bool
}

var opts options

func ParseOptions() options {
	var opts options
	flag.StringVar(&opts.Address, "addr", "tcp://0.0.0.0:46659", "Listen address")
	flag.StringVar(&opts.ABCIType, "abci", "socket", "ABCI server: socket | grpc")
	flag.BoolVar(&opts.Serial, "serial", false, "Enforce incrementing (serial) txs")
	flag.BoolVar(&opts.Verbose, "verbose", false, "Set verbose output")
	flag.BoolVar(&opts.Verbose, "v", false, "Set verbose output")

	flag.Parse()
	return opts
}
