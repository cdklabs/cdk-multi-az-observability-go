package multi-az-observability

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The properties for adding alarms and dashboards for an instrumented service.
// Experimental.
type InstrumentedServiceMultiAZObservabilityProps struct {
	// The service that the alarms and dashboards are being crated for.
	// Experimental.
	Service IService `field:"required" json:"service" yaml:"service"`
	// If you are not using a static bucket to deploy assets, for example you are synthing this and it gets uploaded to a bucket whose name is unknown to you (maybe used as part of a central CI/CD system) and is provided as a parameter to your stack, specify that parameter name here.
	//
	// It will override the bucket location CDK provides by
	// default for bundled assets. The stack containing this contruct needs
	// to have a parameter defined that uses this name. The underlying
	// stacks in this construct that deploy assets will copy the parent stack's
	// value for this property.
	// Default: - The assets will be uploaded to the default defined
	// asset location.
	//
	// Experimental.
	AssetsBucketParameterName *string `field:"optional" json:"assetsBucketParameterName" yaml:"assetsBucketParameterName"`
	// If you are not using a static bucket to deploy assets, for example you are synthing this and it gets uploaded to a bucket that uses a prefix that is unknown to you (maybe used as part of a central CI/CD system) and is provided as a parameter to your stack, specify that parameter name here.
	//
	// It will override the bucket prefix CDK provides by
	// default for bundled assets. This property only takes effect if you
	// defined the assetsBucketParameterName. The stack containing this contruct needs
	// to have a parameter defined that uses this name. The underlying
	// stacks in this construct that deploy assets will copy the parent stack's
	// value for this property.
	// Default: - No object prefix will be added to your custom assets location.
	// However, if you have overridden something like the 'BucketPrefix' property
	// in your stack synthesizer with a variable like "${AssetsBucketPrefix",
	// you will need to define this property so it doesn't cause a reference error
	// even if the prefix value is blank.
	//
	// Experimental.
	AssetsBucketPrefixParameterName *string `field:"optional" json:"assetsBucketPrefixParameterName" yaml:"assetsBucketPrefixParameterName"`
	// The algorithm to use for performing outlier detection for availability metrics.
	//
	// ** Currently only STATIC is supported **.
	// Default: OutlierDetectionAlgorithm.STATIC
	//
	// Experimental.
	AvailabilityOutlierDetectionAlgorithm OutlierDetectionAlgorithm `field:"optional" json:"availabilityOutlierDetectionAlgorithm" yaml:"availabilityOutlierDetectionAlgorithm"`
	// The outlier threshold for determining if an AZ is an outlier for faults.
	//
	// This number is interpreted
	// differently for different outlier algorithms. When used with
	// STATIC, the number should be between 0 and 1 to represent the
	// percentage of errors (like .7) that an AZ must be responsible
	// for to be considered an outlier. When used with CHI_SQUARED, it
	// represents the p value that indicates statistical significance, like
	// 0.05 which means the skew has less than or equal to a 5% chance of
	// occuring. When used with Z_SCORE it indicates how many standard
	// deviations to evaluate for an AZ being an outlier, typically 3 is
	// standard for Z_SCORE.
	//
	// Standard defaults based on the outlier detection algorithm:
	// STATIC: 0.7
	// CHI_SQUARED: 0.05
	// Z_SCORE: 3
	// IQR: 1.5
	// MAD: 3.
	// Default: - Depends on the outlier detection algorithm selected.
	//
	// Experimental.
	AvailabilityOutlierThreshold *float64 `field:"optional" json:"availabilityOutlierThreshold" yaml:"availabilityOutlierThreshold"`
	// Indicates whether to create per operation and overall service dashboards.
	// Default: - No dashboards are created.
	//
	// Experimental.
	CreateDashboards *bool `field:"optional" json:"createDashboards" yaml:"createDashboards"`
	// The interval used in the dashboard, defaults to 60 minutes.
	// Default: - 60 minutes.
	//
	// Experimental.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The algorithm to use for performing outlier detection for latency metrics.
	//
	// ** Currently only STATIC is supported **.
	// Default: OutlierDetectionAlgorithm.STATIC
	//
	// Experimental.
	LatencyOutlierDetectionAlgorithm OutlierDetectionAlgorithm `field:"optional" json:"latencyOutlierDetectionAlgorithm" yaml:"latencyOutlierDetectionAlgorithm"`
	// The metric for latency to use in outlier detection, which means whether the algorithm uses a count of requests exceeding your latency threshold or whether it uses the actual latency values at your latency alarm threshold statistic.
	// Default: LatencyOutlierMetric.COUNT
	//
	// Experimental.
	LatencyOutlierMetricAggregation LatencyOutlierMetricAggregation `field:"optional" json:"latencyOutlierMetricAggregation" yaml:"latencyOutlierMetricAggregation"`
	// The outlier threshold for determining if an AZ is an outlier for latency.
	//
	// This number is interpreted
	// differently for different outlier algorithms. When used with
	// STATIC, the number should be between 0 and 1 to represent the
	// percentage of errors (like .7) that an AZ must be responsible
	// for to be considered an outlier. When used with CHI_SQUARED, it
	// represents the p value that indicates statistical significance, like
	// 0.05 which means the skew has less than or equal to a 5% chance of
	// occuring. When used with Z_SCORE it indicates how many standard
	// deviations to evaluate for an AZ being an outlier, typically 3 is
	// standard for Z_SCORE.
	//
	// Standard defaults based on the outlier detection algorithm:
	// STATIC: 0.7
	// CHI_SQUARED: 0.05
	// Z_SCORE: 3
	// IQR: 1.5
	// MAD: 3.
	// Default: - Depends on the outlier detection algorithm selected.
	//
	// Experimental.
	LatencyOutlierThreshold *float64 `field:"optional" json:"latencyOutlierThreshold" yaml:"latencyOutlierThreshold"`
}

