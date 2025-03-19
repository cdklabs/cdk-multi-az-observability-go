package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Service level alarms and rules using critical operations.
// Experimental.
type IServiceAlarmsAndRules interface {
	// An alarm for regional availability impact of any critical operation as measured by the canary.
	// Experimental.
	RegionalAvailabilityCanaryAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetRegionalAvailabilityCanaryAlarm(r awscloudwatch.IAlarm)
	// An alarm for regional availability or latency impact of any critical operation as measured by the canary.
	// Experimental.
	RegionalAvailabilityOrLatencyCanaryAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetRegionalAvailabilityOrLatencyCanaryAlarm(r awscloudwatch.IAlarm)
	// An alarm for fault count exceeding a regional threshold for all critical operations.
	// Experimental.
	RegionalFaultCountServerSideAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetRegionalFaultCountServerSideAlarm(r awscloudwatch.IAlarm)
	// The service these alarms and rules are for.
	// Experimental.
	Service() IService
	// Experimental.
	SetService(s IService)
	// The zonal aggregate isolated impact alarms.
	//
	// There is 1 alarm per AZ that
	// triggers for availability or latency impact to any critical operation in that AZ
	// that indicates it has isolated impact as measured by canaries or server-side.
	// Experimental.
	ZonalAggregateIsolatedImpactAlarms() *[]awscloudwatch.IAlarm
	// Experimental.
	SetZonalAggregateIsolatedImpactAlarms(z *[]awscloudwatch.IAlarm)
	// The zonal server-side isolated impact alarms.
	//
	// There is 1 alarm per AZ that triggers
	// on availability or atency impact to any critical operation in that AZ. These are useful
	// for deployment monitoring to not inadvertently fail when a canary can't contact an AZ
	// during a deployment.
	// Experimental.
	ZonalServerSideIsolatedImpactAlarms() *[]awscloudwatch.IAlarm
	// Experimental.
	SetZonalServerSideIsolatedImpactAlarms(z *[]awscloudwatch.IAlarm)
}

// The jsii proxy for IServiceAlarmsAndRules
type jsiiProxy_IServiceAlarmsAndRules struct {
	_ byte // padding
}

func (j *jsiiProxy_IServiceAlarmsAndRules) RegionalAvailabilityCanaryAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"regionalAvailabilityCanaryAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetRegionalAvailabilityCanaryAlarm(val awscloudwatch.IAlarm) {
	_jsii_.Set(
		j,
		"regionalAvailabilityCanaryAlarm",
		val,
	)
}

func (j *jsiiProxy_IServiceAlarmsAndRules) RegionalAvailabilityOrLatencyCanaryAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"regionalAvailabilityOrLatencyCanaryAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetRegionalAvailabilityOrLatencyCanaryAlarm(val awscloudwatch.IAlarm) {
	_jsii_.Set(
		j,
		"regionalAvailabilityOrLatencyCanaryAlarm",
		val,
	)
}

func (j *jsiiProxy_IServiceAlarmsAndRules) RegionalFaultCountServerSideAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"regionalFaultCountServerSideAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetRegionalFaultCountServerSideAlarm(val awscloudwatch.IAlarm) {
	if err := j.validateSetRegionalFaultCountServerSideAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionalFaultCountServerSideAlarm",
		val,
	)
}

func (j *jsiiProxy_IServiceAlarmsAndRules) Service() IService {
	var returns IService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetService(val IService) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_IServiceAlarmsAndRules) ZonalAggregateIsolatedImpactAlarms() *[]awscloudwatch.IAlarm {
	var returns *[]awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"zonalAggregateIsolatedImpactAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetZonalAggregateIsolatedImpactAlarms(val *[]awscloudwatch.IAlarm) {
	if err := j.validateSetZonalAggregateIsolatedImpactAlarmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zonalAggregateIsolatedImpactAlarms",
		val,
	)
}

func (j *jsiiProxy_IServiceAlarmsAndRules) ZonalServerSideIsolatedImpactAlarms() *[]awscloudwatch.IAlarm {
	var returns *[]awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"zonalServerSideIsolatedImpactAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetZonalServerSideIsolatedImpactAlarms(val *[]awscloudwatch.IAlarm) {
	if err := j.validateSetZonalServerSideIsolatedImpactAlarmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zonalServerSideIsolatedImpactAlarms",
		val,
	)
}

