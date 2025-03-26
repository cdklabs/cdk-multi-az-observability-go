package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Alarms and rules for canary metrics.
// Experimental.
type ICanaryOperationZonalAlarmsAndRules interface {
	IBaseOperationZonalAlarmsAndRules
	// Alarm that triggers if either latency or availability breach the specified threshold in this AZ and the AZ is an outlier for faults or latency.
	// Experimental.
	IsolatedImpactAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetIsolatedImpactAlarm(i awscloudwatch.IAlarm)
}

// The jsii proxy for ICanaryOperationZonalAlarmsAndRules
type jsiiProxy_ICanaryOperationZonalAlarmsAndRules struct {
	jsiiProxy_IBaseOperationZonalAlarmsAndRules
}

func (j *jsiiProxy_ICanaryOperationZonalAlarmsAndRules) IsolatedImpactAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"isolatedImpactAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICanaryOperationZonalAlarmsAndRules)SetIsolatedImpactAlarm(val awscloudwatch.IAlarm) {
	if err := j.validateSetIsolatedImpactAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isolatedImpactAlarm",
		val,
	)
}

