package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Details for the defaults used in a service for metrics in one perspective, such as server side latency.
// Experimental.
type IServiceAvailabilityMetricDetails interface {
	IServiceMetricDetails
	// The threshold for alarms associated with availability fault metrics, for example if measuring fault rate, the threshold may be 1, meaning you would want an alarm that triggers if the fault rate goes above 1%.
	// Experimental.
	FaultAlarmThreshold() *float64
	// The threshold for alarms associated with availability success metrics, for example if measuring success rate, the threshold may be 99, meaning you would want an alarm that triggers if success drops below 99%.
	// Experimental.
	SuccessAlarmThreshold() *float64
}

// The jsii proxy for IServiceAvailabilityMetricDetails
type jsiiProxy_IServiceAvailabilityMetricDetails struct {
	jsiiProxy_IServiceMetricDetails
}

func (j *jsiiProxy_IServiceAvailabilityMetricDetails) FaultAlarmThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"faultAlarmThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAvailabilityMetricDetails) SuccessAlarmThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successAlarmThreshold",
		&returns,
	)
	return returns
}

