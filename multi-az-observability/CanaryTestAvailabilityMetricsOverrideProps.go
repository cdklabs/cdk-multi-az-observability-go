package multi-az-observability

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for canary metrics in an operation.
// Experimental.
type CanaryTestAvailabilityMetricsOverrideProps struct {
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
	// The threshold for alarms associated with fault metrics, for example if measuring fault rate, the threshold may be 1, meaning you would want an alarm that triggers if the fault rate goes above 1%.
	// Default: - This property will use the default defined for the service.
	//
	// Experimental.
	FaultAlarmThreshold *float64 `field:"optional" json:"faultAlarmThreshold" yaml:"faultAlarmThreshold"`
	// The threshold for alarms associated with success metrics, for example if measuring success rate, the threshold may be 99, meaning you would want an alarm that triggers if success drops below 99%.
	// Default: - This property will use the default defined for the service.
	//
	// Experimental.
	SuccessAlarmThreshold *float64 `field:"optional" json:"successAlarmThreshold" yaml:"successAlarmThreshold"`
}

