package multi-az-observability

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The properties for adding alarms and dashboards for an instrumented service.
// Experimental.
type InstrumentedServiceMultiAZObservabilityProps struct {
	// The algorithm to use for performing outlier detection.
	// Experimental.
	OutlierDetectionAlgorithm OutlierDetectionAlgorithm `field:"required" json:"outlierDetectionAlgorithm" yaml:"outlierDetectionAlgorithm"`
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
	// The outlier threshold for determining if an AZ is an outlier for latency or faults.
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
	// Z_SCORE: 2
	// IQR: 1.5
	// MAD: 3.
	// Default: - Depends on the outlier detection algorithm selected.
	//
	// Experimental.
	OutlierThreshold *float64 `field:"optional" json:"outlierThreshold" yaml:"outlierThreshold"`
}

