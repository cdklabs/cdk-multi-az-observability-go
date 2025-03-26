package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The metric definitions for metric produced by the canary.
// Experimental.
type ICanaryMetrics interface {
	// The canary availability metric details.
	// Experimental.
	CanaryAvailabilityMetricDetails() IOperationAvailabilityMetricDetails
	// The canary latency metric details.
	// Experimental.
	CanaryLatencyMetricDetails() IOperationLatencyMetricDetails
}

// The jsii proxy for ICanaryMetrics
type jsiiProxy_ICanaryMetrics struct {
	_ byte // padding
}

func (j *jsiiProxy_ICanaryMetrics) CanaryAvailabilityMetricDetails() IOperationAvailabilityMetricDetails {
	var returns IOperationAvailabilityMetricDetails
	_jsii_.Get(
		j,
		"canaryAvailabilityMetricDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICanaryMetrics) CanaryLatencyMetricDetails() IOperationLatencyMetricDetails {
	var returns IOperationLatencyMetricDetails
	_jsii_.Get(
		j,
		"canaryLatencyMetricDetails",
		&returns,
	)
	return returns
}

