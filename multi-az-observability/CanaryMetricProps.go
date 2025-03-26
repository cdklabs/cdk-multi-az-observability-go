package multi-az-observability


// Properties for canary metrics in an operation.
// Experimental.
type CanaryMetricProps struct {
	// The canary availability metric details.
	// Experimental.
	CanaryAvailabilityMetricDetails IOperationAvailabilityMetricDetails `field:"required" json:"canaryAvailabilityMetricDetails" yaml:"canaryAvailabilityMetricDetails"`
	// The canary latency metric details.
	// Experimental.
	CanaryLatencyMetricDetails IOperationLatencyMetricDetails `field:"required" json:"canaryLatencyMetricDetails" yaml:"canaryLatencyMetricDetails"`
}

