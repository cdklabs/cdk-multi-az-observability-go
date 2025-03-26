package multi-az-observability

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for canary metrics in an operation.
// Experimental.
type CanaryTestLatencyMetricsOverrideProps struct {
	// The statistic used for alarms, for availability metrics this should be "Sum", for latency metrics it could something like "p99" or "p99.9".
	// Default: - This property will use the default defined for the service.
	//
	// Experimental.
	AlarmStatistic *string `field:"optional" json:"alarmStatistic" yaml:"alarmStatistic"`
	// The number of datapoints to alarm on for latency and availability alarms.
	// Default: - This property will use the default defined for the service.
	//
	// Experimental.
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// The number of evaluation periods for latency and availabiltiy alarms.
	// Default: - This property will use the default defined for the service.
	//
	// Experimental.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The period for the metrics.
	// Default: - This property will use the default defined for the service.
	//
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// The threshold for alarms associated with latency success metrics, for example if success latency exceeds 500 milliseconds.
	// Default: - This property will use the default defined for the service.
	//
	// Experimental.
	SuccessAlarmThreshold awscdk.Duration `field:"optional" json:"successAlarmThreshold" yaml:"successAlarmThreshold"`
}

