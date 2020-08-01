package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	AccessTracing = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "access_log",
		},
		[]string{"trace_id", "span_id", "method"},
	)

	AccessErrorCode = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "access_errcode",
		},
		[]string{"trace_id", "span_id", "method", "code"},
	)

	RecoveryPanic = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "recovery_panic",
		},
		[]string{"trace_id", "span_id", "method"},
	)

	ErrorLogging = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "error_log",
		},
		[]string{"trace_id", "span_id", "level"},
	)

	SQLTracing = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "sql_tracing",
			Buckets: []float64{0.5, 1, 3, 5, 10, 30},
		},
		[]string{"trace_id", "span_id", "span_name", "instance", "user", "statement"},
	)

	HTTPTracing = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_tracing",
			Buckets: []float64{1, 3, 5, 10, 30},
		},
		[]string{"trace_id", "span_id", "span_name", "request_uri", "status_code"},
	)

	RPCTracing = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "rpc_tracing",
			Buckets: []float64{1, 3, 5, 10, 30},
		},
		[]string{"trace_id", "span_id", "span_name", "method"},
	)

	RedisTracing = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "redis_tracing",
			Buckets: []float64{0.1, 0.5, 1, 5, 10},
		},
		[]string{"trace_id", "span_id", "span_name", "command_name"},
	)
)

func init() {
	_ = prometheus.Register(AccessTracing)
	_ = prometheus.Register(AccessErrorCode)
	_ = prometheus.Register(RecoveryPanic)
	_ = prometheus.Register(ErrorLogging)
	_ = prometheus.Register(SQLTracing)
	_ = prometheus.Register(HTTPTracing)
	_ = prometheus.Register(RPCTracing)
	_ = prometheus.Register(RedisTracing)
}
