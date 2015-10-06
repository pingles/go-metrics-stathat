package metricsstathat

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/stathat/stathatgo"
	"time"
)

// Output the stats in `metrics.Registry` r every `interval` seconds using `key` API-key
func StatHat(r metrics.Registry, interval int, key string) {
	for {
		r.Each(func(name string, i interface{}) {
			n := func(p string) string {
				return fmt.Sprintf("%s.%s", name, p)
			}
			switch m := i.(type) {
			case metrics.Counter:
				stathat.PostEZValue(n("count"), key, float64(m.Count()))
			case metrics.Gauge:
				stathat.PostEZValue(n("value"), key, float64(m.Value()))
			case metrics.Healthcheck:
				val := 0.0
				if m.Error() == nil {
					val = 1.0
				}
				stathat.PostEZValue(n("healthy"), key, val)
			case metrics.Histogram:
				ps := m.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
				stathat.PostEZValue(n("count"), key, float64(m.Count()))
				stathat.PostEZValue(n("min"), key, float64(m.Min()))
				stathat.PostEZValue(n("max"), key, float64(m.Max()))
				stathat.PostEZValue(n("mean"), key, m.Mean())
				stathat.PostEZValue(n("stddev"), key, m.StdDev())
				stathat.PostEZValue(n("median"), key, ps[0])
				stathat.PostEZValue(n("75"), key, ps[1])
				stathat.PostEZValue(n("95"), key, ps[2])
				stathat.PostEZValue(n("99"), key, ps[3])
				stathat.PostEZValue(n("999"), key, ps[4])
			case metrics.Meter:
				stathat.PostEZValue(n("count"), key, float64(m.Count()))
				stathat.PostEZValue(n("1minRate"), key, m.Rate1())
				stathat.PostEZValue(n("5minRate"), key, m.Rate5())
				stathat.PostEZValue(n("15minRate"), key, m.Rate15())
				stathat.PostEZValue(n("meanRate"), key, m.RateMean())
			case metrics.Timer:
				ps := m.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
				stathat.PostEZValue(n("count"), key, float64(m.Count()))
				stathat.PostEZValue(n("min"), key, float64(m.Min()))
				stathat.PostEZValue(n("max"), key, float64(m.Max()))
				stathat.PostEZValue(n("mean"), key, m.Mean())
				stathat.PostEZValue(n("stddev"), key, m.StdDev())
				stathat.PostEZValue(n("median"), key, ps[0])
				stathat.PostEZValue(n("75"), key, ps[1])
				stathat.PostEZValue(n("95"), key, ps[2])
				stathat.PostEZValue(n("99"), key, ps[3])
				stathat.PostEZValue(n("999"), key, ps[4])
				stathat.PostEZValue(n("1minRate"), key, m.Rate1())
				stathat.PostEZValue(n("5minRate"), key, m.Rate5())
				stathat.PostEZValue(n("15minRate"), key, m.Rate15())
				stathat.PostEZValue(n("meanRate"), key, m.RateMean())
			}
		})
		time.Sleep(time.Duration(time.Second * time.Duration(interval)))
	}
}
