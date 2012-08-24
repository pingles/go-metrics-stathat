go-metrics-stathat
==========

A [stathat](http://www.stathat.com/) sender for [go-metrics](http://github.com/rcrowley/go-metrics)

Usage
-----

Create and update metrics on stathat:

```go
import (
	"github.com/rcrowley/go-metrics"           // to get the "metrics" namespace
	"github.com/samuraisam/go-metrics-stathat" // to get the "metricsstathat" namespace
)

// use this registry as you would normally using go-metrics
reg := metrics.NewRegistry()

// every 60 seconds, all metrics that have been added to the provied registry to stathat
go metricsstathat.StatHat(reg, 60, "MYEZKEY")
```

Installation
------------

```sh
go install github.com/samuraisam/go-metrics-stathat
```
