package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// The base operation zonal alarms and rules.
// Experimental.
type IBaseOperationZonalAlarmsAndRules interface {
	// Availability alarm for this operation.
	// Experimental.
	AvailabilityAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetAvailabilityAlarm(a awscloudwatch.IAlarm)
	// Alarm that indicates that this AZ is an outlier for fault rate.
	// Experimental.
	AvailabilityZoneIsOutlierForFaults() awscloudwatch.IAlarm
	// Experimental.
	SetAvailabilityZoneIsOutlierForFaults(a awscloudwatch.IAlarm)
	// Alarm that indicates this AZ is an outlier for high latency.
	// Experimental.
	AvailabilityZoneIsOutlierForLatency() awscloudwatch.IAlarm
	// Experimental.
	SetAvailabilityZoneIsOutlierForLatency(a awscloudwatch.IAlarm)
	// Latency alarm for this operation.
	// Experimental.
	LatencyAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetLatencyAlarm(l awscloudwatch.IAlarm)
}

// The jsii proxy for IBaseOperationZonalAlarmsAndRules
type jsiiProxy_IBaseOperationZonalAlarmsAndRules struct {
	_ byte // padding
}

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules) AvailabilityAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"availabilityAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules)SetAvailabilityAlarm(val awscloudwatch.IAlarm) {
	if err := j.validateSetAvailabilityAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityAlarm",
		val,
	)
}

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules) AvailabilityZoneIsOutlierForFaults() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"availabilityZoneIsOutlierForFaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules)SetAvailabilityZoneIsOutlierForFaults(val awscloudwatch.IAlarm) {
	if err := j.validateSetAvailabilityZoneIsOutlierForFaultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneIsOutlierForFaults",
		val,
	)
}

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules) AvailabilityZoneIsOutlierForLatency() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"availabilityZoneIsOutlierForLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules)SetAvailabilityZoneIsOutlierForLatency(val awscloudwatch.IAlarm) {
	if err := j.validateSetAvailabilityZoneIsOutlierForLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneIsOutlierForLatency",
		val,
	)
}

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules) LatencyAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"latencyAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules)SetLatencyAlarm(val awscloudwatch.IAlarm) {
	if err := j.validateSetLatencyAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latencyAlarm",
		val,
	)
}

