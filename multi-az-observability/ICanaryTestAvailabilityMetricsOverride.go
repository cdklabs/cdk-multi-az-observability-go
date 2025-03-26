package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Provides overrides for the default metric settings used for the automatically created canary tests.
// Experimental.
type ICanaryTestAvailabilityMetricsOverride interface {
	ICanaryTestMetricsOverride
	// The threshold for alarms associated with fault metrics, for example if measuring fault rate, the threshold may be 1, meaning you would want an alarm that triggers if the fault rate goes above 1%.
	// Default: - This property will use the default defined for the service.
	//
	// Experimental.
	FaultAlarmThreshold() *float64
	// The threshold for alarms associated with success metrics, for example if measuring success rate, the threshold may be 99, meaning you would want an alarm that triggers if success drops below 99%.
	// Default: - This property will use the default defined for the service.
	//
	// Experimental.
	SuccessAlarmThreshold() *float64
}

// The jsii proxy for ICanaryTestAvailabilityMetricsOverride
type jsiiProxy_ICanaryTestAvailabilityMetricsOverride struct {
	jsiiProxy_ICanaryTestMetricsOverride
}

func (j *jsiiProxy_ICanaryTestAvailabilityMetricsOverride) FaultAlarmThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"faultAlarmThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICanaryTestAvailabilityMetricsOverride) SuccessAlarmThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successAlarmThreshold",
		&returns,
	)
	return returns
}

