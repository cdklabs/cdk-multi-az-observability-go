package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Provides overrides for the default metric settings used for the automatically created canary tests.
// Experimental.
type ICanaryTestLatencyMetricsOverride interface {
	ICanaryTestMetricsOverride
	// The threshold for alarms associated with latency success metrics, for example if success latency exceeds 500 milliseconds.
	// Experimental.
	SuccessAlarmThreshold() awscdk.Duration
}

// The jsii proxy for ICanaryTestLatencyMetricsOverride
type jsiiProxy_ICanaryTestLatencyMetricsOverride struct {
	jsiiProxy_ICanaryTestMetricsOverride
}

func (j *jsiiProxy_ICanaryTestLatencyMetricsOverride) SuccessAlarmThreshold() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"successAlarmThreshold",
		&returns,
	)
	return returns
}

