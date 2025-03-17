package multi-az-observability

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The properties for performing zonal impact detection with NAT Gateway(s).
// Experimental.
type NatGatewayDetectionProps struct {
	// A list of NAT Gateways per Availability Zone (using the AZ name as the key).
	// Experimental.
	NatGateways *map[string]*[]awsec2.CfnNatGateway `field:"required" json:"natGateways" yaml:"natGateways"`
	// The algorithm to use to calculate an AZ as an outlier for packet loss.
	// Default: PacketLossOutlierAlgorithm.STATIC
	//
	// Experimental.
	PacketLossOutlierAlgorithm PacketLossOutlierAlgorithm `field:"optional" json:"packetLossOutlierAlgorithm" yaml:"packetLossOutlierAlgorithm"`
	// The threshold used with the outlier calculation.
	// Default: "This depends on the outlier algorithm. STATIC: 66. Z-SCORE: 3."
	//
	// Experimental.
	PacketLossOutlierThreshold *float64 `field:"optional" json:"packetLossOutlierThreshold" yaml:"packetLossOutlierThreshold"`
	// The percentage of packet loss at which you consider there to be impact.
	// Default: 0.01 (as in 0.01%)
	//
	// Experimental.
	PacketLossPercentThreshold *float64 `field:"optional" json:"packetLossPercentThreshold" yaml:"packetLossPercentThreshold"`
}

