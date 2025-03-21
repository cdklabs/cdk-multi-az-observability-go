package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Service level alarms and rules using critical operations.
// Experimental.
type IServiceAlarmsAndRules interface {
	// An alarm indicating the canary has discovered an availability or latency impact on a critical operation while testing the regional endpoint.
	// Experimental.
	RegionalCanaryAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetRegionalCanaryAlarm(r awscloudwatch.IAlarm)
	// An alarm indicating there is availability or latency impact on a critical operation that is not scoped to a single availability zone as measured by the server-side and/or canary (if present).
	// Experimental.
	RegionalImpactAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetRegionalImpactAlarm(r awscloudwatch.IAlarm)
	// An alarm indicating there is availability or latency impact on a critical operation that is not scoped to a single availability zone as measured by the server-side.
	// Experimental.
	RegionalServerSideImpactAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetRegionalServerSideImpactAlarm(r awscloudwatch.IAlarm)
	// The service these alarms and rules are for.
	// Experimental.
	Service() IService
	// Experimental.
	SetService(s IService)
	// This is the top level alarm you should tie notifications/paging/alerting to.
	//
	// It triggers
	// on any impact to a critical operation either zonally scoped or regionally scoped.
	// Experimental.
	ServiceImpactAlarm() awscloudwatch.IAlarm
	// Experimental.
	SetServiceImpactAlarm(s awscloudwatch.IAlarm)
	// The zonal aggregate isolated impact alarms.
	//
	// There is 1 alarm per AZ that
	// triggers for availability or latency impact to any critical operation in that AZ
	// that indicates it has isolated impact as measured by canaries or server-side.
	// Experimental.
	ZonalAggregateIsolatedImpactAlarms() *map[string]awscloudwatch.IAlarm
	// Experimental.
	SetZonalAggregateIsolatedImpactAlarms(z *map[string]awscloudwatch.IAlarm)
	// The zonal server-side isolated impact alarms.
	//
	// There is 1 alarm per AZ that triggers
	// on availability or latency impact to any critical operation in that AZ. These are useful
	// for deployment monitoring to not inadvertently fail when a canary can't contact an AZ
	// during a deployment.
	// Experimental.
	ZonalServerSideIsolatedImpactAlarms() *map[string]awscloudwatch.IAlarm
	// Experimental.
	SetZonalServerSideIsolatedImpactAlarms(z *map[string]awscloudwatch.IAlarm)
}

// The jsii proxy for IServiceAlarmsAndRules
type jsiiProxy_IServiceAlarmsAndRules struct {
	_ byte // padding
}

func (j *jsiiProxy_IServiceAlarmsAndRules) RegionalCanaryAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"regionalCanaryAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetRegionalCanaryAlarm(val awscloudwatch.IAlarm) {
	_jsii_.Set(
		j,
		"regionalCanaryAlarm",
		val,
	)
}

func (j *jsiiProxy_IServiceAlarmsAndRules) RegionalImpactAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"regionalImpactAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetRegionalImpactAlarm(val awscloudwatch.IAlarm) {
	if err := j.validateSetRegionalImpactAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionalImpactAlarm",
		val,
	)
}

func (j *jsiiProxy_IServiceAlarmsAndRules) RegionalServerSideImpactAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"regionalServerSideImpactAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetRegionalServerSideImpactAlarm(val awscloudwatch.IAlarm) {
	if err := j.validateSetRegionalServerSideImpactAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionalServerSideImpactAlarm",
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

func (j *jsiiProxy_IServiceAlarmsAndRules) ServiceImpactAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"serviceImpactAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetServiceImpactAlarm(val awscloudwatch.IAlarm) {
	if err := j.validateSetServiceImpactAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceImpactAlarm",
		val,
	)
}

func (j *jsiiProxy_IServiceAlarmsAndRules) ZonalAggregateIsolatedImpactAlarms() *map[string]awscloudwatch.IAlarm {
	var returns *map[string]awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"zonalAggregateIsolatedImpactAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetZonalAggregateIsolatedImpactAlarms(val *map[string]awscloudwatch.IAlarm) {
	if err := j.validateSetZonalAggregateIsolatedImpactAlarmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zonalAggregateIsolatedImpactAlarms",
		val,
	)
}

func (j *jsiiProxy_IServiceAlarmsAndRules) ZonalServerSideIsolatedImpactAlarms() *map[string]awscloudwatch.IAlarm {
	var returns *map[string]awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"zonalServerSideIsolatedImpactAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceAlarmsAndRules)SetZonalServerSideIsolatedImpactAlarms(val *map[string]awscloudwatch.IAlarm) {
	if err := j.validateSetZonalServerSideIsolatedImpactAlarmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zonalServerSideIsolatedImpactAlarms",
		val,
	)
}

