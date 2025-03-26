package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-multi-az-observability-go/multi-az-observability/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Provides overrides for the default metric settings used for the automatically created canary tests.
// Experimental.
type CanaryTestAvailabilityMetricsOverride interface {
	CanaryTestMetricsOverride
	// The statistic used for alarms, for availability metrics this should be "Sum", for latency metrics it could something like "p99" or "p99.9".
	// Experimental.
	AlarmStatistic() *string
	// The number of datapoints to alarm on for latency and availability alarms.
	// Experimental.
	DatapointsToAlarm() *float64
	// The number of evaluation periods for latency and availabiltiy alarms.
	// Experimental.
	EvaluationPeriods() *float64
	// The threshold for alarms associated with availability fault metrics, for example if measuring fault rate, the threshold may be 1, meaning you would want an alarm that triggers if the fault rate goes above 1%.
	// Experimental.
	FaultAlarmThreshold() *float64
	// The period for the metrics.
	// Experimental.
	Period() awscdk.Duration
	// The threshold for alarms associated with availability success metrics, for example if measuring success rate, the threshold may be 99, meaning you would want an alarm that triggers if success drops below 99%.
	// Experimental.
	SuccessAlarmThreshold() *float64
}

// The jsii proxy struct for CanaryTestAvailabilityMetricsOverride
type jsiiProxy_CanaryTestAvailabilityMetricsOverride struct {
	jsiiProxy_CanaryTestMetricsOverride
}

func (j *jsiiProxy_CanaryTestAvailabilityMetricsOverride) AlarmStatistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmStatistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanaryTestAvailabilityMetricsOverride) DatapointsToAlarm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanaryTestAvailabilityMetricsOverride) EvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanaryTestAvailabilityMetricsOverride) FaultAlarmThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"faultAlarmThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanaryTestAvailabilityMetricsOverride) Period() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanaryTestAvailabilityMetricsOverride) SuccessAlarmThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successAlarmThreshold",
		&returns,
	)
	return returns
}


// Experimental.
func NewCanaryTestAvailabilityMetricsOverride(props *CanaryTestAvailabilityMetricsOverrideProps) CanaryTestAvailabilityMetricsOverride {
	_init_.Initialize()

	if err := validateNewCanaryTestAvailabilityMetricsOverrideParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CanaryTestAvailabilityMetricsOverride{}

	_jsii_.Create(
		"@cdklabs/multi-az-observability.CanaryTestAvailabilityMetricsOverride",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCanaryTestAvailabilityMetricsOverride_Override(c CanaryTestAvailabilityMetricsOverride, props *CanaryTestAvailabilityMetricsOverrideProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/multi-az-observability.CanaryTestAvailabilityMetricsOverride",
		[]interface{}{props},
		c,
	)
}

