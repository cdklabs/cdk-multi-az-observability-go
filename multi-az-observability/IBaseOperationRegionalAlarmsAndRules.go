package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Base regional alarms and rules.
// Experimental.
type IBaseOperationRegionalAlarmsAndRules interface {
	// Availability alarm for this operation.
	// Experimental.
	AvailabilityAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetAvailabilityAlarm(a awscloudwatch.IAlarm)
	// Composite alarm for either availabiltiy or latency impact to this operation.
	// Experimental.
	AvailabilityOrLatencyAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetAvailabilityOrLatencyAlarm(a awscloudwatch.IAlarm)
	// Latency alarm for this operation.
	// Experimental.
	LatencyAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetLatencyAlarm(l awscloudwatch.IAlarm)
}

// The jsii proxy for IBaseOperationRegionalAlarmsAndRules
type jsiiProxy_IBaseOperationRegionalAlarmsAndRules struct {
	_ byte // padding
}

func (j *jsiiProxy_IBaseOperationRegionalAlarmsAndRules) AvailabilityAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"availabilityAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBaseOperationRegionalAlarmsAndRules)SetAvailabilityAlarm(val awscloudwatch.IAlarm) {
	if err := j.validateSetAvailabilityAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityAlarm",
		val,
	)
}

func (j *jsiiProxy_IBaseOperationRegionalAlarmsAndRules) AvailabilityOrLatencyAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"availabilityOrLatencyAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBaseOperationRegionalAlarmsAndRules)SetAvailabilityOrLatencyAlarm(val awscloudwatch.IAlarm) {
	if err := j.validateSetAvailabilityOrLatencyAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityOrLatencyAlarm",
		val,
	)
}

func (j *jsiiProxy_IBaseOperationRegionalAlarmsAndRules) LatencyAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"latencyAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBaseOperationRegionalAlarmsAndRules)SetLatencyAlarm(val awscloudwatch.IAlarm) {
	if err := j.validateSetLatencyAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latencyAlarm",
		val,
	)
}

