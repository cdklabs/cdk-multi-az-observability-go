package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-multi-az-observability-go/multi-az-observability/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// The representation of a service composed of multiple operations.
// Experimental.
type Service interface {
	IService
	// A list of the Availability Zone names used by this application.
	// Experimental.
	AvailabilityZoneNames() *[]*string
	// The base endpoint for this service, like "https://www.example.com". Operation paths will be appended to this endpoint for canary testing the service.
	// Experimental.
	BaseUrl() *string
	// Define these settings if you want to automatically add canary tests to your operations.
	//
	// Operations can individually opt out
	// of canary test creation if you define this setting.
	// Default: - Automatic canary tests will not be created for
	// operations in this service.
	//
	// Experimental.
	CanaryTestProps() *AddCanaryTestProps
	// The default settings that are used for availability metrics for all operations unless specifically overridden in an operation definition.
	// Experimental.
	DefaultAvailabilityMetricDetails() IServiceAvailabilityMetricDetails
	// The default settings that are used for contributor insight rules.
	// Default: - No defaults are provided and must be specified per operation.
	//
	// Experimental.
	DefaultContributorInsightRuleDetails() IContributorInsightRuleDetails
	// The default settings that are used for availability metrics for all operations unless specifically overridden in an operation definition.
	// Experimental.
	DefaultLatencyMetricDetails() IServiceLatencyMetricDetails
	// The fault count threshold that indicates the service is unhealthy.
	//
	// This is an absolute value of faults
	// being produced by all critical operations in aggregate.
	// Experimental.
	FaultCountThreshold() *float64
	// The load balancer this service sits behind.
	// Default: - No load balancer metrics will be included in
	// dashboards and its ARN will not be added to top level AZ
	// alarm descriptions.
	//
	// Experimental.
	LoadBalancer() awselasticloadbalancingv2.ILoadBalancerV2
	// The minimum number of unhealthy targets to consider an AZ impaired.
	// Default: Count of 2.
	//
	// Experimental.
	MinimumUnhealthyTargets() *MinimumUnhealthyTargets
	// The operations that are part of this service.
	// Experimental.
	Operations() *[]IOperation
	// The period for which metrics for the service should be aggregated.
	// Experimental.
	Period() awscdk.Duration
	// The name of your service.
	// Experimental.
	ServiceName() *string
	// The target groups registered with the load balancer.
	// Default: Anomalous and mitigated host metrics will not be included on dashboards.
	//
	// Experimental.
	TargetGroups() *[]awselasticloadbalancingv2.ITargetGroup
	// Adds an operation to this service and sets the operation's service property.
	// Experimental.
	AddOperation(operation IOperation)
}

// The jsii proxy struct for Service
type jsiiProxy_Service struct {
	jsiiProxy_IService
}

func (j *jsiiProxy_Service) AvailabilityZoneNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZoneNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) CanaryTestProps() *AddCanaryTestProps {
	var returns *AddCanaryTestProps
	_jsii_.Get(
		j,
		"canaryTestProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) DefaultAvailabilityMetricDetails() IServiceAvailabilityMetricDetails {
	var returns IServiceAvailabilityMetricDetails
	_jsii_.Get(
		j,
		"defaultAvailabilityMetricDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) DefaultContributorInsightRuleDetails() IContributorInsightRuleDetails {
	var returns IContributorInsightRuleDetails
	_jsii_.Get(
		j,
		"defaultContributorInsightRuleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) DefaultLatencyMetricDetails() IServiceLatencyMetricDetails {
	var returns IServiceLatencyMetricDetails
	_jsii_.Get(
		j,
		"defaultLatencyMetricDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) FaultCountThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"faultCountThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) LoadBalancer() awselasticloadbalancingv2.ILoadBalancerV2 {
	var returns awselasticloadbalancingv2.ILoadBalancerV2
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) MinimumUnhealthyTargets() *MinimumUnhealthyTargets {
	var returns *MinimumUnhealthyTargets
	_jsii_.Get(
		j,
		"minimumUnhealthyTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Operations() *[]IOperation {
	var returns *[]IOperation
	_jsii_.Get(
		j,
		"operations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Period() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) TargetGroups() *[]awselasticloadbalancingv2.ITargetGroup {
	var returns *[]awselasticloadbalancingv2.ITargetGroup
	_jsii_.Get(
		j,
		"targetGroups",
		&returns,
	)
	return returns
}


// Experimental.
func NewService(props *ServiceProps) Service {
	_init_.Initialize()

	if err := validateNewServiceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Service{}

	_jsii_.Create(
		"@cdklabs/multi-az-observability.Service",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewService_Override(s Service, props *ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/multi-az-observability.Service",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_Service) AddOperation(operation IOperation) {
	if err := s.validateAddOperationParameters(operation); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOperation",
		[]interface{}{operation},
	)
}

