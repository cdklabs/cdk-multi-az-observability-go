package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Creates alarms and rules for an operation for both regional and zonal metrics.
// Experimental.
type IOperationAlarmsAndRules interface {
	// The aggregate, server-side and canary combined, zonal alarm indexed by Availability Zone name.
	// Experimental.
	AggregateZonalAlarms() *map[string]awscloudwatch.IAlarm
	// The canary regional alarms and rules.
	// Experimental.
	CanaryRegionalAlarmsAndRules() ICanaryOperationRegionalAlarmsAndRules
	// An alarm indicating availability or latency impact has been detected by the canary and the impact is regionally scoped, not zonal.
	// Experimental.
	CanaryRegionalImpactAlarm() awscloudwatch.IAlarm
	// The canary zonal alarms and rules, indexed by Availability Zone name.
	// Default: - This is an empty dictionary if canary metric details are not provided.
	//
	// Experimental.
	CanaryZonalAlarmsAndRules() *map[string]ICanaryOperationZonalAlarmsAndRules
	// The operation the alarms and rules are created for.
	// Experimental.
	Operation() IOperation
	// An alarm indicating availability or latency impact has been detected by the server-side  and/or canary (if present) and the impact is regionally scoped, not zonal.
	// Experimental.
	RegionalImpactAlarm() awscloudwatch.IAlarm
	// The server side regional alarms and rules.
	// Experimental.
	ServerSideRegionalAlarmsAndRules() IServerSideOperationRegionalAlarmsAndRules
	// An alarm indicating availability or latency impact has been detected by the server-side and the impact is regionally scoped, not zonal.
	// Experimental.
	ServerSideRegionalImpactAlarm() awscloudwatch.IAlarm
	// The server side zonal alarms and rules, indexed by Availability Zone name.
	// Experimental.
	ServerSideZonalAlarmsAndRules() *map[string]IServerSideOperationZonalAlarmsAndRules
}

// The jsii proxy for IOperationAlarmsAndRules
type jsiiProxy_IOperationAlarmsAndRules struct {
	_ byte // padding
}

func (j *jsiiProxy_IOperationAlarmsAndRules) AggregateZonalAlarms() *map[string]awscloudwatch.IAlarm {
	var returns *map[string]awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"aggregateZonalAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperationAlarmsAndRules) CanaryRegionalAlarmsAndRules() ICanaryOperationRegionalAlarmsAndRules {
	var returns ICanaryOperationRegionalAlarmsAndRules
	_jsii_.Get(
		j,
		"canaryRegionalAlarmsAndRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperationAlarmsAndRules) CanaryRegionalImpactAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"canaryRegionalImpactAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperationAlarmsAndRules) CanaryZonalAlarmsAndRules() *map[string]ICanaryOperationZonalAlarmsAndRules {
	var returns *map[string]ICanaryOperationZonalAlarmsAndRules
	_jsii_.Get(
		j,
		"canaryZonalAlarmsAndRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperationAlarmsAndRules) Operation() IOperation {
	var returns IOperation
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperationAlarmsAndRules) RegionalImpactAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"regionalImpactAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperationAlarmsAndRules) ServerSideRegionalAlarmsAndRules() IServerSideOperationRegionalAlarmsAndRules {
	var returns IServerSideOperationRegionalAlarmsAndRules
	_jsii_.Get(
		j,
		"serverSideRegionalAlarmsAndRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperationAlarmsAndRules) ServerSideRegionalImpactAlarm() awscloudwatch.IAlarm {
	var returns awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"serverSideRegionalImpactAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperationAlarmsAndRules) ServerSideZonalAlarmsAndRules() *map[string]IServerSideOperationZonalAlarmsAndRules {
	var returns *map[string]IServerSideOperationZonalAlarmsAndRules
	_jsii_.Get(
		j,
		"serverSideZonalAlarmsAndRules",
		&returns,
	)
	return returns
}

