package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-multi-az-observability-go/multi-az-observability/internal"
)

// Observability for an instrumented service.
// Experimental.
type IInstrumentedServiceMultiAZObservability interface {
	constructs.IConstruct
	// If the service is configured to have canary tests created, this will be the log group where the canary's logs are stored.
	// Default: - No log group is created if the canary is not requested.
	//
	// Experimental.
	CanaryLogGroup() awslogs.ILogGroup
	// The dashboards for each operation.
	// Experimental.
	OperationDashboards() *[]awscloudwatch.Dashboard
	// Key represents the operation name and the value is the set of zonal alarms and rules for that operation.
	//
	// You can get the
	// granular alarms that compose the higher level aggregate alarms
	// for each operation.
	// Experimental.
	PerOperationAlarmsAndRules() *map[string]IOperationAlarmsAndRules
	// Index into the dictionary by operation name, then by Availability Zone Name to get the alarms that indicate an AZ shows isolated impact from availability or latency as seen by either the server-side or canary.
	//
	// This is a shortcut to
	// access the same alarms from the perOperationAlarmsAndRules property.
	// Experimental.
	PerOperationZonalImpactAlarms() *map[string]*map[string]awscloudwatch.IAlarm
	// The alarms and rules for the overall service.
	// Experimental.
	ServiceAlarms() IServiceAlarmsAndRules
	// The service level dashboard.
	// Experimental.
	ServiceDashboard() awscloudwatch.Dashboard
}

// The jsii proxy for IInstrumentedServiceMultiAZObservability
type jsiiProxy_IInstrumentedServiceMultiAZObservability struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInstrumentedServiceMultiAZObservability) CanaryLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"canaryLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstrumentedServiceMultiAZObservability) OperationDashboards() *[]awscloudwatch.Dashboard {
	var returns *[]awscloudwatch.Dashboard
	_jsii_.Get(
		j,
		"operationDashboards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstrumentedServiceMultiAZObservability) PerOperationAlarmsAndRules() *map[string]IOperationAlarmsAndRules {
	var returns *map[string]IOperationAlarmsAndRules
	_jsii_.Get(
		j,
		"perOperationAlarmsAndRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstrumentedServiceMultiAZObservability) PerOperationZonalImpactAlarms() *map[string]*map[string]awscloudwatch.IAlarm {
	var returns *map[string]*map[string]awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"perOperationZonalImpactAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstrumentedServiceMultiAZObservability) ServiceAlarms() IServiceAlarmsAndRules {
	var returns IServiceAlarmsAndRules
	_jsii_.Get(
		j,
		"serviceAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstrumentedServiceMultiAZObservability) ServiceDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard
	_jsii_.Get(
		j,
		"serviceDashboard",
		&returns,
	)
	return returns
}

