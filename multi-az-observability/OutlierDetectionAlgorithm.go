package multi-az-observability


// Available algorithms for performing outlier detection.
// Experimental.
type OutlierDetectionAlgorithm string

const (
	// Defines using a static value to compare skew in faults or high latency responses.
	//
	// A good default threshold for this is .7 meaning one AZ
	// is responsible for 70% of the total errors or high latency responses.
	// Experimental.
	OutlierDetectionAlgorithm_STATIC OutlierDetectionAlgorithm = "STATIC"
	// Uses the chi squared statistic to determine if there is a statistically significant skew in fault rate or high latency distribution.
	//
	// A normal default threshold for this is 0.05, which means there is a 5% or
	// less chance of the skew in errors or high latency responses occuring.
	// Experimental.
	OutlierDetectionAlgorithm_CHI_SQUARED OutlierDetectionAlgorithm = "CHI_SQUARED"
	// Uses z-score to determine if the skew in faults or high latency respones exceeds a defined number of standard devations.
	//
	// A good default threshold value for this is 2, meaning the outlier value is outside
	// 95% of the normal distribution. Using 3 means the outlier is outside 99.7% of
	// the normal distribution.
	// Experimental.
	OutlierDetectionAlgorithm_Z_SCORE OutlierDetectionAlgorithm = "Z_SCORE"
	// Uses Interquartile Range Method to determine an outlier for faults or latency.
	//
	// No threshold is required for this method and will be ignored.
	// Experimental.
	OutlierDetectionAlgorithm_IQR OutlierDetectionAlgorithm = "IQR"
	// Median Absolute Deviation (MAD) to determine an outlier for faults or latency.
	//
	// A common default value threshold 3.
	// Experimental.
	OutlierDetectionAlgorithm_MAD OutlierDetectionAlgorithm = "MAD"
)

