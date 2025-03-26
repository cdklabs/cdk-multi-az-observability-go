package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Details for operation metrics in one perspective, such as server side latency.
// Experimental.
type IOperationLatencyMetricDetails interface {
	IOperationMetricDetails
	// The threshold for alarms associated with latency success metrics, for example if success latency exceeds 500 milliseconds.
	// Experimental.
	SuccessAlarmThreshold() awscdk.Duration
}

// The jsii proxy for IOperationLatencyMetricDetails
type jsiiProxy_IOperationLatencyMetricDetails struct {
	jsiiProxy_IOperationMetricDetails
}

func (j *jsiiProxy_IOperationLatencyMetricDetails) SuccessAlarmThreshold() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"successAlarmThreshold",
		&returns,
	)
	return returns
}

