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

type ConcurrencyExample func()

var (
	RATE_LIMITING       ConcurrencyExample = rate_limiting.Execute
	FAN_OUT_FAN_IN      ConcurrencyExample = fan_out_fan_in.Execute
	PIPELINE            ConcurrencyExample = pipeline.Execute
	PUBLISH_SUBSCRIBE   ConcurrencyExample = publish_subscribe.Execute
	SELECT_WITH_TIMEOUT ConcurrencyExample = select_with_timeout.Execute
	SEMAPHORE           ConcurrencyExample = semaphore.Execute
	WORKER_POOL         ConcurrencyExample = worker_pool.Execute
)

func main() {
	example := FAN_OUT_FAN_IN

	example()
}
