package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an operation in a service.
// Experimental.
type IOperation interface {
	// Optional metric details if the service has an existing canary.
	// Experimental.
	CanaryMetricDetails() ICanaryMetrics
	// The override values for automatically created canary tests so you can use values other than the service defaults to define the thresholds for availability.
	// Experimental.
	CanaryTestAvailabilityMetricsOverride() ICanaryTestAvailabilityMetricsOverride
	// The override values for automatically created canary tests so you can use values other than the service defaults to define the thresholds for latency.
	// Experimental.
	CanaryTestLatencyMetricsOverride() ICanaryTestLatencyMetricsOverride
	// If they have been added, the properties for creating new canary tests on this operation.
	// Experimental.
	CanaryTestProps() *AddCanaryTestProps
	// Indicates this is a critical operation for the service and will be included in service level metrics and dashboards.
	// Experimental.
	Critical() *bool
	// The http methods supported by the operation.
	// Experimental.
	HttpMethods() *[]*string
	// The name of the operation.
	// Experimental.
	OperationName() *string
	// Set to true if you have defined CanaryTestProps for your service, which applies to all operations, but you want to opt out of creating the canary test for this operation.
	// Default: - The operation is not opted out.
	//
	// Experimental.
	OptOutOfServiceCreatedCanary() *bool
	// The HTTP path for the operation for canaries to run against, something like "/products/list".
	// Experimental.
	Path() *string
	// The server side availability metric details.
	// Experimental.
	ServerSideAvailabilityMetricDetails() IOperationAvailabilityMetricDetails
	// The server side details for contributor insights rules.
	// Experimental.
	ServerSideContributorInsightRuleDetails() IContributorInsightRuleDetails
	// The server side latency metric details.
	// Experimental.
	ServerSideLatencyMetricDetails() IOperationLatencyMetricDetails
	// The service the operation is associated with.
	// Experimental.
	Service() IService
}

// The jsii proxy for IOperation
type jsiiProxy_IOperation struct {
	_ byte // padding
}

func (j *jsiiProxy_IOperation) CanaryMetricDetails() ICanaryMetrics {
	var returns ICanaryMetrics
	_jsii_.Get(
		j,
		"canaryMetricDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) CanaryTestAvailabilityMetricsOverride() ICanaryTestAvailabilityMetricsOverride {
	var returns ICanaryTestAvailabilityMetricsOverride
	_jsii_.Get(
		j,
		"canaryTestAvailabilityMetricsOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) CanaryTestLatencyMetricsOverride() ICanaryTestLatencyMetricsOverride {
	var returns ICanaryTestLatencyMetricsOverride
	_jsii_.Get(
		j,
		"canaryTestLatencyMetricsOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) CanaryTestProps() *AddCanaryTestProps {
	var returns *AddCanaryTestProps
	_jsii_.Get(
		j,
		"canaryTestProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) Critical() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"critical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) HttpMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) OptOutOfServiceCreatedCanary() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"optOutOfServiceCreatedCanary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) ServerSideAvailabilityMetricDetails() IOperationAvailabilityMetricDetails {
	var returns IOperationAvailabilityMetricDetails
	_jsii_.Get(
		j,
		"serverSideAvailabilityMetricDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) ServerSideContributorInsightRuleDetails() IContributorInsightRuleDetails {
	var returns IContributorInsightRuleDetails
	_jsii_.Get(
		j,
		"serverSideContributorInsightRuleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) ServerSideLatencyMetricDetails() IOperationLatencyMetricDetails {
	var returns IOperationLatencyMetricDetails
	_jsii_.Get(
		j,
		"serverSideLatencyMetricDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOperation) Service() IService {
	var returns IService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

