package multi-az-observability


// The options for calculating if a NAT Gateway is an outlier for packet loss.
// Experimental.
type PacketLossOutlierAlgorithm string

const (
	// This will take the availability threshold and calculate if one AZ is responsible for that percentage of packet loss.
	// Experimental.
	PacketLossOutlierAlgorithm_STATIC PacketLossOutlierAlgorithm = "STATIC"
)

