package multi-az-observability


// The minimum unhealthy targets for an AZ to be considered impaired instead of individual targets in the zone.
//
// You must specify either count or percentage,
// if you specify both, only count will be used.
// Experimental.
type MinimumUnhealthyTargets struct {
	// The minimum number of targets that must be unhealthy.
	//
	// If the number of
	// unhealthy targets is equal to this, or greater, the impact ot a single AZ
	// is not considered ot be from a "single" target. You must specify either count or percentage,
	// if you specify both, only count will be used.
	// Default: This value is not used.
	//
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The minimum percentage of targets that must be unhealthy.
	//
	// If the percentage
	// of unhealthy targets is equal to this, or greater, the impact to a single AZ
	// is not considered to be from a "single" target. You must specify either count or percentage,
	// if you specify both, only count will be used.
	// Default: This value is not used.
	//
	// Experimental.
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

