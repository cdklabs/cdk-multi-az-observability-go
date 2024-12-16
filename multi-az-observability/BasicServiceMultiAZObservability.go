package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-multi-az-observability-go/multi-az-observability/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-multi-az-observability-go/multi-az-observability/internal"
)

// Basic observability for a service using metrics from ALBs and NAT Gateways.
// Experimental.
type BasicServiceMultiAZObservability interface {
	constructs.Construct
	IBasicServiceMultiAZObservability
	// The alarms indicating if an AZ has isolated impact from either ALB or NAT GW metrics.
	// Experimental.
	AggregateZonalIsolatedImpactAlarms() *map[string]awscloudwatch.IAlarm
	// Experimental.
	SetAggregateZonalIsolatedImpactAlarms(val *map[string]awscloudwatch.IAlarm)
	// The alarms indicating if an AZ is an outlier for ALB faults and has isolated impact.
	// Experimental.
	AlbZonalIsolatedImpactAlarms() *map[string]awscloudwatch.IAlarm
	// Experimental.
	SetAlbZonalIsolatedImpactAlarms(val *map[string]awscloudwatch.IAlarm)
	// The application load balancers being used by the service.
	// Experimental.
	ApplicationLoadBalancers() *[]awselasticloadbalancingv2.IApplicationLoadBalancer
	// Experimental.
	SetApplicationLoadBalancers(val *[]awselasticloadbalancingv2.IApplicationLoadBalancer)
	// The dashboard that is optionally created.
	// Experimental.
	Dashboard() awscloudwatch.Dashboard
	// Experimental.
	SetDashboard(val awscloudwatch.Dashboard)
	// The NAT Gateways being used in the service, each set of NAT Gateways are keyed by their Availability Zone Id.
	// Experimental.
	NatGateways() *map[string]*[]awsec2.CfnNatGateway
	// Experimental.
	SetNatGateways(val *map[string]*[]awsec2.CfnNatGateway)
	// The alarms indicating if an AZ is an outlier for NAT GW packet loss and has isolated impact.
	// Experimental.
	NatGWZonalIsolatedImpactAlarms() *map[string]awscloudwatch.IAlarm
	// Experimental.
	SetNatGWZonalIsolatedImpactAlarms(val *map[string]awscloudwatch.IAlarm)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The name of the service.
	// Experimental.
	ServiceName() *string
	// Experimental.
	SetServiceName(val *string)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BasicServiceMultiAZObservability
type jsiiProxy_BasicServiceMultiAZObservability struct {
	internal.Type__constructsConstruct
	jsiiProxy_IBasicServiceMultiAZObservability
}

func (j *jsiiProxy_BasicServiceMultiAZObservability) AggregateZonalIsolatedImpactAlarms() *map[string]awscloudwatch.IAlarm {
	var returns *map[string]awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"aggregateZonalIsolatedImpactAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicServiceMultiAZObservability) AlbZonalIsolatedImpactAlarms() *map[string]awscloudwatch.IAlarm {
	var returns *map[string]awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"albZonalIsolatedImpactAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicServiceMultiAZObservability) ApplicationLoadBalancers() *[]awselasticloadbalancingv2.IApplicationLoadBalancer {
	var returns *[]awselasticloadbalancingv2.IApplicationLoadBalancer
	_jsii_.Get(
		j,
		"applicationLoadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicServiceMultiAZObservability) Dashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard
	_jsii_.Get(
		j,
		"dashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicServiceMultiAZObservability) NatGateways() *map[string]*[]awsec2.CfnNatGateway {
	var returns *map[string]*[]awsec2.CfnNatGateway
	_jsii_.Get(
		j,
		"natGateways",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicServiceMultiAZObservability) NatGWZonalIsolatedImpactAlarms() *map[string]awscloudwatch.IAlarm {
	var returns *map[string]awscloudwatch.IAlarm
	_jsii_.Get(
		j,
		"natGWZonalIsolatedImpactAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicServiceMultiAZObservability) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicServiceMultiAZObservability) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}


// Experimental.
func NewBasicServiceMultiAZObservability(scope constructs.Construct, id *string, props *BasicServiceMultiAZObservabilityProps) BasicServiceMultiAZObservability {
	_init_.Initialize()

	if err := validateNewBasicServiceMultiAZObservabilityParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BasicServiceMultiAZObservability{}

	_jsii_.Create(
		"@cdklabs/multi-az-observability.BasicServiceMultiAZObservability",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBasicServiceMultiAZObservability_Override(b BasicServiceMultiAZObservability, scope constructs.Construct, id *string, props *BasicServiceMultiAZObservabilityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/multi-az-observability.BasicServiceMultiAZObservability",
		[]interface{}{scope, id, props},
		b,
	)
}

func (j *jsiiProxy_BasicServiceMultiAZObservability)SetAggregateZonalIsolatedImpactAlarms(val *map[string]awscloudwatch.IAlarm) {
	if err := j.validateSetAggregateZonalIsolatedImpactAlarmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregateZonalIsolatedImpactAlarms",
		val,
	)
}

func (j *jsiiProxy_BasicServiceMultiAZObservability)SetAlbZonalIsolatedImpactAlarms(val *map[string]awscloudwatch.IAlarm) {
	_jsii_.Set(
		j,
		"albZonalIsolatedImpactAlarms",
		val,
	)
}

func (j *jsiiProxy_BasicServiceMultiAZObservability)SetApplicationLoadBalancers(val *[]awselasticloadbalancingv2.IApplicationLoadBalancer) {
	_jsii_.Set(
		j,
		"applicationLoadBalancers",
		val,
	)
}

func (j *jsiiProxy_BasicServiceMultiAZObservability)SetDashboard(val awscloudwatch.Dashboard) {
	_jsii_.Set(
		j,
		"dashboard",
		val,
	)
}

func (j *jsiiProxy_BasicServiceMultiAZObservability)SetNatGateways(val *map[string]*[]awsec2.CfnNatGateway) {
	_jsii_.Set(
		j,
		"natGateways",
		val,
	)
}

func (j *jsiiProxy_BasicServiceMultiAZObservability)SetNatGWZonalIsolatedImpactAlarms(val *map[string]awscloudwatch.IAlarm) {
	_jsii_.Set(
		j,
		"natGWZonalIsolatedImpactAlarms",
		val,
	)
}

func (j *jsiiProxy_BasicServiceMultiAZObservability)SetServiceName(val *string) {
	if err := j.validateSetServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func BasicServiceMultiAZObservability_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBasicServiceMultiAZObservability_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/multi-az-observability.BasicServiceMultiAZObservability",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BasicServiceMultiAZObservability) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

