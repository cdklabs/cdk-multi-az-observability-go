package multi-az-observability


// Properties for canary metrics in an operation.
// Experimental.
type CanaryMetricProps struct {
	// The canary availability metric details.
	// Experimental.
	CanaryAvailabilityMetricDetails IOperationMetricDetails `field:"required" json:"canaryAvailabilityMetricDetails" yaml:"canaryAvailabilityMetricDetails"`
	// The canary latency metric details.
	// Experimental.
	CanaryLatencyMetricDetails IOperationMetricDetails `field:"required" json:"canaryLatencyMetricDetails" yaml:"canaryLatencyMetricDetails"`
}

