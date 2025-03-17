package multi-az-observability

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// The properties for performing zonal impact detection with ALB(s).
// Experimental.
type ApplicationLoadBalancerDetectionProps struct {
	// The application load balancers to collect metrics from.
	// Experimental.
	ApplicationLoadBalancers *[]awselasticloadbalancingv2.IApplicationLoadBalancer `field:"required" json:"applicationLoadBalancers" yaml:"applicationLoadBalancers"`
	// The percentage of faults for a single ALB to consider an AZ to be unhealthy, this should align with your availability goal.
	//
	// For example
	// 1% or 5%, 0.01 or 0.05.
	// Experimental.
	FaultCountPercentThreshold *float64 `field:"required" json:"faultCountPercentThreshold" yaml:"faultCountPercentThreshold"`
	// The statistic used to measure target response latency, like p99,  which can be specified using Stats.percentile(99) or "p99".
	// Experimental.
	LatencyStatistic *string `field:"required" json:"latencyStatistic" yaml:"latencyStatistic"`
	// The threshold in milliseconds for ALB targets whose responses are slower than this value at the specified percentile statistic.
	// Experimental.
	LatencyThreshold *float64 `field:"required" json:"latencyThreshold" yaml:"latencyThreshold"`
	// The method used to determine if an AZ is an outlier for availability for Application Load Balancer metrics.
	// Default: STATIC.
	//
	// Experimental.
	AvailabilityOutlierAlgorithm ApplicationLoadBalancerAvailabilityOutlierAlgorithm `field:"optional" json:"availabilityOutlierAlgorithm" yaml:"availabilityOutlierAlgorithm"`
	// The threshold for the outlier detection algorithm.
	// Default: "This depends on the algorithm used. STATIC: 0.66."
	//
	// Experimental.
	AvailabilityOutlierThreshold *float64 `field:"optional" json:"availabilityOutlierThreshold" yaml:"availabilityOutlierThreshold"`
	// The method used to determine if an AZ is an outlier for latency for Application Load Balancer metrics.
	// Default: Z_SCORE.
	//
	// Experimental.
	LatencyOutlierAlgorithm ApplicationLoadBalancerLatencyOutlierAlgorithm `field:"optional" json:"latencyOutlierAlgorithm" yaml:"latencyOutlierAlgorithm"`
	// The threshold for the outlier detection algorithm.
	// Default: "This depends on the algorithm used. STATIC: 0.66. Z_SCORE: 3."
	//
	// Experimental.
	LatencyOutlierThreshold *float64 `field:"optional" json:"latencyOutlierThreshold" yaml:"latencyOutlierThreshold"`
}

