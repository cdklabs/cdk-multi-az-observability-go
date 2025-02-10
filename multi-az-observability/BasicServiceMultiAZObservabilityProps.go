package multi-az-observability

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Properties for creating basic multi-AZ observability.
// Experimental.
type BasicServiceMultiAZObservabilityProps struct {
	// The number of datapoints to alarm on for latency and availability alarms.
	// Experimental.
	DatapointsToAlarm *float64 `field:"required" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// The number of evaluation periods for latency and availabiltiy alarms.
	// Experimental.
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The percentage of faults for a single ALB to consider an AZ to be unhealthy, this should align with your availability goal.
	//
	// For example
	// 1% or 5%, specify as 1 or 5.
	// Experimental.
	FaultCountPercentageThreshold *float64 `field:"required" json:"faultCountPercentageThreshold" yaml:"faultCountPercentageThreshold"`
	// The statistic used to measure target response latency, like p99,  which can be specified using Stats.percentile(99) or "p99".
	// Experimental.
	LatencyStatistic *string `field:"required" json:"latencyStatistic" yaml:"latencyStatistic"`
	// The threshold in seconds for ALB targets whose responses are slower than this value at the specified percentile statistic.
	// Experimental.
	LatencyThreshold *float64 `field:"required" json:"latencyThreshold" yaml:"latencyThreshold"`
	// The service's name.
	// Experimental.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The application load balancers being used by the service.
	//
	// There will be an alarm created for
	// each AZ for each ALB. Then, there will be a composite alarm for AZ created from the input
	// of all ALBs. You must either specify an ALB or a NAT GW.
	// Default: "No alarms for ALBs will be created".
	//
	// Experimental.
	ApplicationLoadBalancers *[]awselasticloadbalancingv2.IApplicationLoadBalancer `field:"optional" json:"applicationLoadBalancers" yaml:"applicationLoadBalancers"`
	// If you are not using a static bucket to deploy assets, for example you are synthing this and it gets uploaded to a bucket whose name is unknown to you (maybe used as part of a central CI/CD system) and is provided as a parameter to your stack, specify that parameter name here.
	//
	// It will override the bucket location CDK provides by
	// default for bundled assets. The stack containing this contruct needs
	// to have a parameter defined that uses this name. The underlying
	// stacks in this construct that deploy assets will copy the parent stack's
	// value for this property.
	// Default: "The assets will be uploaded to the default defined asset location."
	//
	// Experimental.
	AssetsBucketParameterName *string `field:"optional" json:"assetsBucketParameterName" yaml:"assetsBucketParameterName"`
	// If you are not using a static bucket to deploy assets, for example you are synthing this and it gets uploaded to a bucket that uses a prefix that is unknown to you (maybe used as part of a central CI/CD system) and is provided as a parameter to your stack, specify that parameter name here.
	//
	// It will override the bucket prefix CDK provides by
	// default for bundled assets. This property only takes effect if you
	// defined the assetsBucketParameterName. The stack containing this contruct needs
	// to have a parameter defined that uses this name. The underlying
	// stacks in this construct that deploy assets will copy the parent stack's
	// value for this property.
	// Default: "No object prefix will be added to your custom assets location. However, if you have overridden something like the 'BucketPrefix' property in your stack synthesizer with a variable like '${AssetsBucketPrefix}', you will need to define this property so it doesn't cause a reference error even if the prefix value is blank."
	//
	// Experimental.
	AssetsBucketPrefixParameterName *string `field:"optional" json:"assetsBucketPrefixParameterName" yaml:"assetsBucketPrefixParameterName"`
	// Whether to create a dashboard displaying the metrics and alarms.
	// Default: false.
	//
	// Experimental.
	CreateDashboard *bool `field:"optional" json:"createDashboard" yaml:"createDashboard"`
	// Dashboard interval.
	// Default: Duration.hours(1)
	//
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// (Optional) A map of Availability Zone name to the NAT Gateways in that AZ.
	//
	// One alarm per NAT GW will be created. If multiple NAT GWs
	// are provided for a single AZ, those alarms will be aggregated into
	// a composite alarm for the AZ. You must either specify an ALB or a NAT GW.
	// Default: "No alarms for NAT Gateways will be created".
	//
	// Experimental.
	NatGateways *map[string]*[]awsec2.CfnNatGateway `field:"optional" json:"natGateways" yaml:"natGateways"`
	// The amount of packet loss in a NAT GW to determine if an AZ is actually impacted, recommendation is 0.01%.
	// Default: "0.01 (as in 0.01%)"
	//
	// Experimental.
	PacketLossImpactPercentageThreshold *float64 `field:"optional" json:"packetLossImpactPercentageThreshold" yaml:"packetLossImpactPercentageThreshold"`
	// The period to evaluate metrics.
	// Default: Duration.minutes(1)
	//
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
}

