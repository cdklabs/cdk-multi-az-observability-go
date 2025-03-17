package multi-az-observability

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
	// The service's name.
	// Experimental.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Properties for ALBs to detect single AZ impact.
	//
	// You must specify this
	// and/or natGatewayProps.
	// Default: "No ALBs will be used to calculate impact."
	//
	// Experimental.
	ApplicationLoadBalancerProps *ApplicationLoadBalancerDetectionProps `field:"optional" json:"applicationLoadBalancerProps" yaml:"applicationLoadBalancerProps"`
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
	// Properties for NAT Gateways to detect single AZ impact.
	//
	// You must specify
	// this and/or applicationLoadBalancerProps.
	// Default: "No NAT Gateways will be used to calculate impact."
	//
	// Experimental.
	NatGatewayProps *NatGatewayDetectionProps `field:"optional" json:"natGatewayProps" yaml:"natGatewayProps"`
	// The period to evaluate metrics.
	// Default: Duration.minutes(1)
	//
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
}

