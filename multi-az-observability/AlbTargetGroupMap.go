package multi-az-observability

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// An object to map an ALB to its target groups.
// Experimental.
type AlbTargetGroupMap struct {
	// The application load balancer.
	// Experimental.
	ApplicationLoadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer `field:"required" json:"applicationLoadBalancer" yaml:"applicationLoadBalancer"`
	// The target groups associated with the ALB.
	// Default: No target groups are associated and will not display anomalous hosts or mitigated hosts on the dashboard.
	//
	// Experimental.
	TargetGroups *[]awselasticloadbalancingv2.ITargetGroup `field:"optional" json:"targetGroups" yaml:"targetGroups"`
}

