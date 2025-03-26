package multi-az-observability

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// The properties for default service metric details.
// Experimental.
type ServiceMetricDetailsProps struct {
	// The statistic used for alarms, for availability metrics this should be "Sum", for latency metrics it could something like "p99" or "p99.9".
	// Experimental.
	AlarmStatistic *string `field:"required" json:"alarmStatistic" yaml:"alarmStatistic"`
	// The number of datapoints to alarm on for latency and availability alarms.
	// Experimental.
	DatapointsToAlarm *float64 `field:"required" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// The number of evaluation periods for latency and availabiltiy alarms.
	// Experimental.
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The names of fault indicating metrics.
	// Experimental.
	FaultMetricNames *[]*string `field:"required" json:"faultMetricNames" yaml:"faultMetricNames"`
	// The CloudWatch metric namespace for these metrics.
	// Experimental.
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The period for the metrics.
	// Experimental.
	Period awscdk.Duration `field:"required" json:"period" yaml:"period"`
	// The names of success indicating metrics.
	// Experimental.
	SuccessMetricNames *[]*string `field:"required" json:"successMetricNames" yaml:"successMetricNames"`
	// The unit used for these metrics.
	// Experimental.
	Unit awscloudwatch.Unit `field:"required" json:"unit" yaml:"unit"`
	// The statistics for faults you want to appear on dashboards, for example, with latency metrics, you might want p50, p99, and tm99.
	//
	// For availability
	// metrics this will typically just be "Sum".
	// Default: - For availability metrics, this will be "Sum", for latency metrics it will be just "p99".
	//
	// Experimental.
	GraphedFaultStatistics *[]*string `field:"optional" json:"graphedFaultStatistics" yaml:"graphedFaultStatistics"`
	// The statistics for successes you want to appear on dashboards, for example, with latency metrics, you might want p50, p99, and tm99.
	//
	// For availability
	// metrics this will typically just be "Sum".
	// Default: - For availability metrics, this will be "Sum", for latency metrics it will be just "p99".
	//
	// Experimental.
	GraphedSuccessStatistics *[]*string `field:"optional" json:"graphedSuccessStatistics" yaml:"graphedSuccessStatistics"`
}

