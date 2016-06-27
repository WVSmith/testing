package main

import (
	"fmt"
	as "github.com/aerospike/aerospike-client-go"
)

func main() {
	//client, err := as.NewClient("52.41.46.108", 3000)
	client, err := as.NewClient("52.38.52.159", 3000)
	// deal with the error here
	if nil != err {
		return
	}

	defer client.Close()

	spolicy := as.NewScanPolicy()
	spolicy.ScanPercent = 1
	spolicy.MultiPolicy.RecordQueueSize = 10
	spolicy.ConcurrentNodes = false
	spolicy.Priority = as.LOW
	spolicy.IncludeBinData = true

	recs, err := client.ScanAll(spolicy, "idmap", "usermap")
	// deal with the error here

	for res := range recs.Results() {
		if res.Err != nil {
			// handle error here
			// if you want to exit, cancel the recordset to release the resources
			recs.Close()
		} else {
			// process record here
			fmt.Println(res.Record.Bins["mfid"])
			recs.Close()
		}
	}
}
