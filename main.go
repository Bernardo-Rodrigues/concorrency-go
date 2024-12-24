package main

import (
	fan_out_fan_in "concorrency/fan-out-fan-in"
	"concorrency/pipeline"
	publish_subscribe "concorrency/publish-subscribe"
	rate_limiting "concorrency/rate-limiting"
	select_with_timeout "concorrency/select-with-timeout"
	"concorrency/semaphore"
	worker_pool "concorrency/worker-pool"
)

type ConcurrencyExample string

const (
	RATE_LIMITING       ConcurrencyExample = "rate-limiting"
	FAN_OUT_FAN_IN      ConcurrencyExample = "fan-out-fan-in"
	PIPELINE            ConcurrencyExample = "pipeline"
	PUBLISH_SUBSCRIBE   ConcurrencyExample = "publish-subscribe"
	SELECT_WITH_TIMEOUT ConcurrencyExample = "select-with-timeout"
	SEMAPHORE           ConcurrencyExample = "semaphore"
	WORKER_POOL         ConcurrencyExample = "worker-pool"
)

func main() {
	examples := map[ConcurrencyExample]func(){
		RATE_LIMITING:       rate_limiting.Execute,
		FAN_OUT_FAN_IN:      fan_out_fan_in.Execute,
		PIPELINE:            pipeline.Execute,
		PUBLISH_SUBSCRIBE:   publish_subscribe.Execute,
		SELECT_WITH_TIMEOUT: select_with_timeout.Execute,
		SEMAPHORE:           semaphore.Execute,
		WORKER_POOL:         worker_pool.Execute,
	}

	example := FAN_OUT_FAN_IN

	if fn, ok := examples[example]; ok {
		fn()
	} else {
		panic("example not found")
	}
}
