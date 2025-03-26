package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Details for the defaults used in a service for metrics in one perspective, such as server side latency.
// Experimental.
type IServiceLatencyMetricDetails interface {
	IServiceMetricDetails
	// The threshold for alarms associated with latency success metrics, for example if success latency exceeds 500 milliseconds.
	// Experimental.
	SuccessAlarmThreshold() awscdk.Duration
}

// The jsii proxy for IServiceLatencyMetricDetails
type jsiiProxy_IServiceLatencyMetricDetails struct {
	jsiiProxy_IServiceMetricDetails
}

func (j *jsiiProxy_IServiceLatencyMetricDetails) SuccessAlarmThreshold() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"successAlarmThreshold",
		&returns,
	)
	return returns
}

