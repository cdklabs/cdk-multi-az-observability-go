package multi-az-observability


// The options for calculating if an ALB is an outlier for availability.
// Experimental.
type ApplicationLoadBalancerAvailabilityOutlierAlgorithm string

const (
	// This will take the availability threshold and calculate if one AZ is responsible for that percentage of errors.
	// Experimental.
	ApplicationLoadBalancerAvailabilityOutlierAlgorithm_STATIC ApplicationLoadBalancerAvailabilityOutlierAlgorithm = "STATIC"
)

