package multi-az-observability

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// The props for requesting a canary be made for an operation.
// Experimental.
type AddCanaryTestProps struct {
	// The load balancer that will be tested against.
	// Experimental.
	LoadBalancer awselasticloadbalancingv2.ILoadBalancerV2 `field:"required" json:"loadBalancer" yaml:"loadBalancer"`
	// The number of requests to send on each test.
	// Experimental.
	RequestCount *float64 `field:"required" json:"requestCount" yaml:"requestCount"`
	// A schedule expression.
	// Experimental.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// Any headers to include.
	// Default: - No additional headers are added to the requests.
	//
	// Experimental.
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Defining this will override the methods defined in the operation and will use these instead.
	// Default: - The operation's defined HTTP methods will be used to
	// conduct the canary tests.
	//
	// Experimental.
	HttpMethods *[]*string `field:"optional" json:"httpMethods" yaml:"httpMethods"`
	// Whether to ignore TLS validation errors.
	// Default: - false.
	//
	// Experimental.
	IgnoreTlsErrors *bool `field:"optional" json:"ignoreTlsErrors" yaml:"ignoreTlsErrors"`
	// The VPC network configuration.
	// Default: - The Lambda function is not run in a VPC.
	//
	// Experimental.
	NetworkConfiguration *NetworkConfigurationProps `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Data to supply in a POST, PUT, or PATCH operation.
	// Default: - No data is sent in a POST, PUT, or PATCH request.
	//
	// Experimental.
	PostData *string `field:"optional" json:"postData" yaml:"postData"`
	// Specifies a separate number of request to send to the regional endpoint.
	// Default: - The same number of requests specified by the requestCount property is used.
	//
	// Experimental.
	RegionalRequestCount *float64 `field:"optional" json:"regionalRequestCount" yaml:"regionalRequestCount"`
	// The timeout for each individual HTTP request.
	// Default: - Defaults to 2 seconds.
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

