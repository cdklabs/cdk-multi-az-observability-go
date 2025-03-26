package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-multi-az-observability-go/multi-az-observability/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Provides overrides for the default metric settings used for the automatically created canary tests.
// Experimental.
type CanaryTestMetricsOverride interface {
	ICanaryTestMetricsOverride
	// The statistic used for alarms, for availability metrics this should be "Sum", for latency metrics it could something like "p99" or "p99.9".
	// Experimental.
	AlarmStatistic() *string
	// The number of datapoints to alarm on for latency and availability alarms.
	// Experimental.
	DatapointsToAlarm() *float64
	// The number of evaluation periods for latency and availabiltiy alarms.
	// Experimental.
	EvaluationPeriods() *float64
	// The period for the metrics.
	// Experimental.
	Period() awscdk.Duration
}

// The jsii proxy struct for CanaryTestMetricsOverride
type jsiiProxy_CanaryTestMetricsOverride struct {
	jsiiProxy_ICanaryTestMetricsOverride
}

func (j *jsiiProxy_CanaryTestMetricsOverride) AlarmStatistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmStatistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanaryTestMetricsOverride) DatapointsToAlarm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanaryTestMetricsOverride) EvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanaryTestMetricsOverride) Period() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}


// Experimental.
func NewCanaryTestMetricsOverride_Override(c CanaryTestMetricsOverride, props *CanaryTestMetricsOverrideProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/multi-az-observability.CanaryTestMetricsOverride",
		[]interface{}{props},
		c,
	)
}

