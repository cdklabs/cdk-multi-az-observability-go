package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Server side opertaion zonal alarms and rules.
// Experimental.
type IServerSideOperationZonalAlarmsAndRules interface {
	IBaseOperationZonalAlarmsAndRules
	// Insight rule that measures the number of instances contributing to faults in this AZ.
	// Experimental.
	InstanceContributorsToFaultsInThisAZ() awscloudwatch.CfnInsightRule
	// Experimental.
	SetInstanceContributorsToFaultsInThisAZ(i awscloudwatch.CfnInsightRule)
	// Insight rule that measures the number of instances contributing to high latency in this AZ.
	// Experimental.
	InstanceContributorsToHighLatencyInThisAZ() awscloudwatch.CfnInsightRule
	// Experimental.
	SetInstanceContributorsToHighLatencyInThisAZ(i awscloudwatch.CfnInsightRule)
	// Insight rule that is used to calculate the number of instances in this particular AZ that is used with metric math to calculate the percent of instances contributing to latency or faults.
	// Experimental.
	InstancesHandlingRequestsInThisAZ() awscloudwatch.CfnInsightRule
	// Experimental.
	SetInstancesHandlingRequestsInThisAZ(i awscloudwatch.CfnInsightRule)
	// Alarm that triggers if either latency or availability breach the specified threshold in this AZ and the AZ is an outlier for faults or latency.
	// Experimental.
	IsolatedImpactAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetIsolatedImpactAlarm(i awscloudwatch.IAlarm)
	// Alarm indicating that there are multiple instances producing faults in this AZ indicating the fault rate is not being caused by a single instance.
	// Experimental.
	MultipleInstancesProducingFaultsInThisAvailabilityZone() awscloudwatch.IAlarm
	// Experimental.
	SetMultipleInstancesProducingFaultsInThisAvailabilityZone(m awscloudwatch.IAlarm)
	// Alarm indicating that there are multiple instances producing high latency responses in this AZ indicating the latency is not being caused by a single instance.
	// Experimental.
	MultipleInstancesProducingHighLatencyInThisAZ() awscloudwatch.IAlarm
	// Experimental.
	SetMultipleInstancesProducingHighLatencyInThisAZ(m awscloudwatch.IAlarm)
}

// The jsii proxy for IServerSideOperationZonalAlarmsAndRules
type jsiiProxy_IServerSideOperationZonalAlarmsAndRules struct {
	jsiiProxy_IBaseOperationZonalAlarmsAndRules
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules) InstanceContributorsToFaultsInThisAZ() awscloudwatch.CfnInsightRule {
	var returns awscloudwatch.CfnInsightRule
	_jsii_.Get(
		j,
		"instanceContributorsToFaultsInThisAZ",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules)SetInstanceContributorsToFaultsInThisAZ(val awscloudwatch.CfnInsightRule) {
	_jsii_.Set(
		j,
		"instanceContributorsToFaultsInThisAZ",
		val,
	)
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules) InstanceContributorsToHighLatencyInThisAZ() awscloudwatch.CfnInsightRule {
	var returns awscloudwatch.CfnInsightRule
	_jsii_.Get(
		j,
		"instanceContributorsToHighLatencyInThisAZ",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules)SetInstanceContributorsToHighLatencyInThisAZ(val awscloudwatch.CfnInsightRule) {
	_jsii_.Set(
		j,
		"instanceContributorsToHighLatencyInThisAZ",
		val,
	)
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules) InstancesHandlingRequestsInThisAZ() awscloudwatch.CfnInsightRule {
	var returns awscloudwatch.CfnInsightRule
	_jsii_.Get(
		j,
		"instancesHandlingRequestsInThisAZ",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules)SetInstancesHandlingRequestsInThisAZ(val awscloudwatch.CfnInsightRule) {
	_jsii_.Set(
		j,
		"instancesHandlingRequestsInThisAZ",
		val,
	)
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules) IsolatedImpactAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"isolatedImpactAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules)SetIsolatedImpactAlarm(val awscloudwatch.IAlarm) {
	if err := j.validateSetIsolatedImpactAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isolatedImpactAlarm",
		val,
	)
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules) MultipleInstancesProducingFaultsInThisAvailabilityZone() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"multipleInstancesProducingFaultsInThisAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules)SetMultipleInstancesProducingFaultsInThisAvailabilityZone(val awscloudwatch.IAlarm) {
	_jsii_.Set(
		j,
		"multipleInstancesProducingFaultsInThisAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules) MultipleInstancesProducingHighLatencyInThisAZ() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"multipleInstancesProducingHighLatencyInThisAZ",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerSideOperationZonalAlarmsAndRules)SetMultipleInstancesProducingHighLatencyInThisAZ(val awscloudwatch.IAlarm) {
	_jsii_.Set(
		j,
		"multipleInstancesProducingHighLatencyInThisAZ",
		val,
	)
}

