package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// ResourcesAdded tracks resources that were added
	ResourcesAdded = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "kubernetes_diff_resources_added_total",
		Help: "The total number of added Kubernetes resources",
	}, []string{"resource_type", "namespace"})

	// ResourcesDeleted tracks resources that were deleted
	ResourcesDeleted = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "kubernetes_diff_resources_deleted_total",
		Help: "The total number of deleted Kubernetes resources",
	}, []string{"resource_type", "namespace"})

	// ResourcesUpdated tracks resources that were updated
	ResourcesUpdated = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "kubernetes_diff_resources_updated_total",
		Help: "The total number of updated Kubernetes resources",
	}, []string{"resource_type", "namespace"})

	// DiffSize tracks the size of each diff
	DiffSize = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "kubernetes_diff_size_bytes",
		Help:    "Size of diffs in bytes",
		Buckets: prometheus.ExponentialBuckets(10, 10, 8), // 10, 100, 1000, ...
	}, []string{"resource_type", "namespace"})

	// ProcessingErrors tracks errors during processing
	ProcessingErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "kubernetes_diff_processing_errors_total",
		Help: "The total number of errors during diff processing",
	}, []string{"resource_type", "error_type"})

	// WatcherResyncs tracks number of full resyncs
	WatcherResyncs = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "kubernetes_diff_watcher_resyncs_total",
		Help: "The total number of watcher resyncs",
	}, []string{"resource_type"})
)
