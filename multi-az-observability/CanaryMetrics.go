package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-multi-az-observability-go/multi-az-observability/jsii"
)

// Represents metrics for a canary testing a service.
// Experimental.
type CanaryMetrics interface {
	ICanaryMetrics
	// The canary availability metric details.
	// Experimental.
	CanaryAvailabilityMetricDetails() IOperationAvailabilityMetricDetails
	// The canary latency metric details.
	// Experimental.
	CanaryLatencyMetricDetails() IOperationLatencyMetricDetails
}

// The jsii proxy struct for CanaryMetrics
type jsiiProxy_CanaryMetrics struct {
	jsiiProxy_ICanaryMetrics
}

func (j *jsiiProxy_CanaryMetrics) CanaryAvailabilityMetricDetails() IOperationAvailabilityMetricDetails {
	var returns IOperationAvailabilityMetricDetails
	_jsii_.Get(
		j,
		"canaryAvailabilityMetricDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanaryMetrics) CanaryLatencyMetricDetails() IOperationLatencyMetricDetails {
	var returns IOperationLatencyMetricDetails
	_jsii_.Get(
		j,
		"canaryLatencyMetricDetails",
		&returns,
	)
	return returns
}


// Experimental.
func NewCanaryMetrics(props *CanaryMetricProps) CanaryMetrics {
	_init_.Initialize()

	if err := validateNewCanaryMetricsParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CanaryMetrics{}

	_jsii_.Create(
		"@cdklabs/multi-az-observability.CanaryMetrics",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCanaryMetrics_Override(c CanaryMetrics, props *CanaryMetricProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/multi-az-observability.CanaryMetrics",
		[]interface{}{props},
		c,
	)
}

