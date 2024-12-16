package multi-az-observability


// Properties for an operation.
// Experimental.
type OperationProps struct {
	// Indicates this is a critical operation for the service and will be included in service level metrics and dashboards.
	// Experimental.
	Critical *bool `field:"required" json:"critical" yaml:"critical"`
	// The http methods supported by the operation.
	// Experimental.
	HttpMethods *[]*string `field:"required" json:"httpMethods" yaml:"httpMethods"`
	// The name of the operation.
	// Experimental.
	OperationName *string `field:"required" json:"operationName" yaml:"operationName"`
	// The HTTP path for the operation for canaries to run against, something like "/products/list".
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The server side availability metric details.
	// Experimental.
	ServerSideAvailabilityMetricDetails IOperationMetricDetails `field:"required" json:"serverSideAvailabilityMetricDetails" yaml:"serverSideAvailabilityMetricDetails"`
	// The server side latency metric details.
	// Experimental.
	ServerSideLatencyMetricDetails IOperationMetricDetails `field:"required" json:"serverSideLatencyMetricDetails" yaml:"serverSideLatencyMetricDetails"`
	// The service the operation is associated with.
	// Experimental.
	Service IService `field:"required" json:"service" yaml:"service"`
	// Optional metric details if the service has a canary.
	// Default: - No alarms, rules, or dashboards will be created
	// from canary metrics.
	//
	// Experimental.
	CanaryMetricDetails ICanaryMetrics `field:"optional" json:"canaryMetricDetails" yaml:"canaryMetricDetails"`
	// The override values for automatically created canary tests so you can use values other than the service defaults to define the thresholds for availability.
	// Default: - No availability metric details will be overridden and the
	// service defaults will be used for the automatically created canaries.
	//
	// Experimental.
	CanaryTestAvailabilityMetricsOverride ICanaryTestMetricsOverride `field:"optional" json:"canaryTestAvailabilityMetricsOverride" yaml:"canaryTestAvailabilityMetricsOverride"`
	// The override values for automatically created canary tests so you can use values other than the service defaults to define the thresholds for latency.
	// Default: - No latency metric details will be overridden and the
	// service defaults will be used for the automatically created canaries.
	//
	// Experimental.
	CanaryTestLatencyMetricsOverride ICanaryTestMetricsOverride `field:"optional" json:"canaryTestLatencyMetricsOverride" yaml:"canaryTestLatencyMetricsOverride"`
	// If you define this property, a synthetic canary will be provisioned to test the operation.
	// Default: - The default for the service will be used, if that
	// is undefined, then no canary will be provisioned for this operation.
	//
	// Experimental.
	CanaryTestProps *AddCanaryTestProps `field:"optional" json:"canaryTestProps" yaml:"canaryTestProps"`
	// Set to true if you have defined CanaryTestProps for your service, which applies to all operations, but you want to opt out of creating the canary test for this operation.
	// Default: - The operation is not opted out.
	//
	// Experimental.
	OptOutOfServiceCreatedCanary *bool `field:"optional" json:"optOutOfServiceCreatedCanary" yaml:"optOutOfServiceCreatedCanary"`
	// The server side details for contributor insights rules.
	// Default: - The default service contributor insight rule
	// details will be used. If those are not defined no Contributor Insight
	// rules will be created and the number of instances contributing to AZ
	// faults or high latency will not be considered, so a single bad instance
	// could make the AZ appear to look impaired.
	//
	// Experimental.
	ServerSideContributorInsightRuleDetails IContributorInsightRuleDetails `field:"optional" json:"serverSideContributorInsightRuleDetails" yaml:"serverSideContributorInsightRuleDetails"`
}

