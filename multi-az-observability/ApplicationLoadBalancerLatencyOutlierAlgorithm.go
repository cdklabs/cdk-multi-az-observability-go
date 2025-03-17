package multi-az-observability


// The options for calculating if an AZ is an outlier for latency for ALBs.
// Experimental.
type ApplicationLoadBalancerLatencyOutlierAlgorithm string

const (
	// This will take the latency threshold and count the number of requests per AZ  that exceed this threshold and then calculate the percentage of requests exceeding this threshold belong to each AZ.
	//
	// This provides a static comparison
	// of the number of high latency requests in one AZ versus the others.
	// Experimental.
	ApplicationLoadBalancerLatencyOutlierAlgorithm_STATIC ApplicationLoadBalancerLatencyOutlierAlgorithm = "STATIC"
	// This calculates the z score of latency in one AZ against the other AZs.
	//
	// It uses
	// the target response time of all requests to calculate the standard deviation and
	// average for all AZs. This is the default.
	// Experimental.
	ApplicationLoadBalancerLatencyOutlierAlgorithm_Z_SCORE ApplicationLoadBalancerLatencyOutlierAlgorithm = "Z_SCORE"
)

