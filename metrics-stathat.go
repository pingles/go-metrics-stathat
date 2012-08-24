package metricsstathat

import (
	// "log"
	"fmt"
	"time"
	"github.com/rcrowley/go-metrics"
	"github.com/stathat/stathatgo"
)

// Output each metric in the given registry periodically using the given
// logger.  The interval is to be given in seconds.
func StatHat(r metrics.Registry, interval int, key string) {
	for {
		r.Each(func(name string, i interface{}) {
			n := func(p string) string {
				return fmt.Sprintf("%s.%s", name, p)
			}
			switch m := i.(type) {
			case metrics.Counter:
				// l.Printf("counter %s\n", name)
				// l.Printf("  count:       %9d\n", m.Count())
				stathat.PostEZCount(n("count"), key, int(m.Count()))
			case metrics.Gauge:
				// l.Printf("gauge %s\n", name)
				// l.Printf("  value:       %9d\n", m.Value())
				stathat.PostEZValue(n("value"), key, float64(m.Value()))
			case metrics.Healthcheck:
				// m.Check()
				// l.Printf("healthcheck %s\n", name)
				// l.Printf("  error:       %v\n", m.Error())
				val := 0.0
				if m.Error() == nil {
					val = 1.0 
				}
				stathat.PostEZValue(n("healthy"), key, val)
			case metrics.Histogram:
				ps := m.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
				// l.Printf("histogram %s\n", name)
				// l.Printf("  count:       %9d\n", m.Count())
				// l.Printf("  min:         %9d\n", m.Min())
				// l.Printf("  max:         %9d\n", m.Max())
				// l.Printf("  mean:        %12.2f\n", m.Mean())
				// l.Printf("  stddev:      %12.2f\n", m.StdDev())
				// l.Printf("  median:      %12.2f\n", ps[0])
				// l.Printf("  75%%:         %12.2f\n", ps[1])
				// l.Printf("  95%%:         %12.2f\n", ps[2])
				// l.Printf("  99%%:         %12.2f\n", ps[3])
				// l.Printf("  99.9%%:       %12.2f\n", ps[4])
				stathat.PostEZCount(n("count"), key, int(m.Count()))
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
				// l.Printf("meter %s\n", name)
				// l.Printf("  count:       %9d\n", m.Count())
				// l.Printf("  1-min rate:  %12.2f\n", m.Rate1())
				// l.Printf("  5-min rate:  %12.2f\n", m.Rate5())
				// l.Printf("  15-min rate: %12.2f\n", m.Rate15())
				// l.Printf("  mean rate:   %12.2f\n", m.RateMean())
				stathat.PostEZCount(n("count"), key, int(m.Count()))
				stathat.PostEZValue(n("1minRate"), key, m.Rate1())
				stathat.PostEZValue(n("5minRate"), key, m.Rate5())
				stathat.PostEZValue(n("15minRate"), key, m.Rate15())
				stathat.PostEZValue(n("meanRate"), key, m.RateMean())
			case metrics.Timer:
				ps := m.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
				stathat.PostEZCount(n("count"), key, int(m.Count()))
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
				
				// l.Printf("timer %s\n", name)
				// l.Printf("  count:       %9d\n", m.Count())
				// l.Printf("  min:         %9d\n", m.Min())
				// l.Printf("  max:         %9d\n", m.Max())
				// l.Printf("  mean:        %12.2f\n", m.Mean())
				// l.Printf("  stddev:      %12.2f\n", m.StdDev())
				// l.Printf("  median:      %12.2f\n", ps[0])
				// l.Printf("  75%%:         %12.2f\n", ps[1])
				// l.Printf("  95%%:         %12.2f\n", ps[2])
				// l.Printf("  99%%:         %12.2f\n", ps[3])
				// l.Printf("  99.9%%:       %12.2f\n", ps[4])
				// l.Printf("  1-min rate:  %12.2f\n", m.Rate1())
				// l.Printf("  5-min rate:  %12.2f\n", m.Rate5())
				// l.Printf("  15-min rate: %12.2f\n", m.Rate15())
				// l.Printf("  mean rate:   %12.2f\n", m.RateMean())
			}
		})
		time.Sleep(time.Duration(time.Second * time.Duration(interval)))
	}
}
