package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-multi-az-observability-go/multi-az-observability/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Default latency metric details for a service.
// Experimental.
type ServiceLatencyMetricDetails interface {
	ServiceMetricDetails
	IServiceLatencyMetricDetails
	// The statistic used for alarms, for availability metrics this should be "Sum", for latency metrics it could something like "p99" or "p99.9".
	// Experimental.
	AlarmStatistic() *string
	// The number of datapoints to alarm on for latency and availability alarms.
	// Experimental.
	DatapointsToAlarm() *float64
	// The number of evaluation periods for latency and availabiltiy alarms.
	// Experimental.
	EvaluationPeriods() *float64
	// The names of fault indicating metrics.
	// Experimental.
	FaultMetricNames() *[]*string
	// The statistics for faults you want to appear on dashboards, for example, with latency metrics, you might want p50, p99, and tm99.
	//
	// For availability
	// metrics this will typically just be "Sum".
	// Default: - For availability metrics, this will be "Sum", for latency metrics it will be just "p99".
	//
	// Experimental.
	GraphedFaultStatistics() *[]*string
	// The statistics for successes you want to appear on dashboards, for example, with latency metrics, you might want p50, p99, and tm99.
	//
	// For availability
	// metrics this will typically just be "Sum".
	// Default: - For availability metrics, this will be "Sum", for latency metrics it will be just "p99".
	//
	// Experimental.
	GraphedSuccessStatistics() *[]*string
	// The CloudWatch metric namespace for these metrics.
	// Experimental.
	MetricNamespace() *string
	// The period for the metrics.
	// Experimental.
	Period() awscdk.Duration
	// The threshold for alarms associated with latency success metrics, for example if success latency exceeds 500 milliseconds.
	// Experimental.
	SuccessAlarmThreshold() awscdk.Duration
	// The names of success indicating metrics.
	// Experimental.
	SuccessMetricNames() *[]*string
	// The unit used for these metrics.
	// Experimental.
	Unit() awscloudwatch.Unit
}

// The jsii proxy struct for ServiceLatencyMetricDetails
type jsiiProxy_ServiceLatencyMetricDetails struct {
	jsiiProxy_ServiceMetricDetails
	jsiiProxy_IServiceLatencyMetricDetails
}

func (j *jsiiProxy_ServiceLatencyMetricDetails) AlarmStatistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmStatistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLatencyMetricDetails) DatapointsToAlarm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLatencyMetricDetails) EvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLatencyMetricDetails) FaultMetricNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"faultMetricNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLatencyMetricDetails) GraphedFaultStatistics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"graphedFaultStatistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLatencyMetricDetails) GraphedSuccessStatistics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"graphedSuccessStatistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLatencyMetricDetails) MetricNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLatencyMetricDetails) Period() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLatencyMetricDetails) SuccessAlarmThreshold() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"successAlarmThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLatencyMetricDetails) SuccessMetricNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"successMetricNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLatencyMetricDetails) Unit() awscloudwatch.Unit {
	var returns awscloudwatch.Unit
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}


// Experimental.
func NewServiceLatencyMetricDetails(props *ServiceLatencyMetricDetailsProps) ServiceLatencyMetricDetails {
	_init_.Initialize()

	if err := validateNewServiceLatencyMetricDetailsParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceLatencyMetricDetails{}

	_jsii_.Create(
		"@cdklabs/multi-az-observability.ServiceLatencyMetricDetails",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewServiceLatencyMetricDetails_Override(s ServiceLatencyMetricDetails, props *ServiceLatencyMetricDetailsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/multi-az-observability.ServiceLatencyMetricDetails",
		[]interface{}{props},
		s,
	)
}

