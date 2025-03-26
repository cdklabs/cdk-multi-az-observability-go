package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// The server side operation regional alarms and rules.
// Experimental.
type IServerSideOperationRegionalAlarmsAndRules interface {
	IBaseOperationRegionalAlarmsAndRules
	// A rule that shows which instances are contributing to faults.
	// Experimental.
	InstanceContributorsToRegionalFaults() awscloudwatch.CfnInsightRule
	// Experimental.
	SetInstanceContributorsToRegionalFaults(i awscloudwatch.CfnInsightRule)
	// A rule that shows which instances are contributing to high latency responses.
	// Experimental.
	InstanceContributorsToRegionalHighLatency() awscloudwatch.CfnInsightRule
	// Experimental.
	SetInstanceContributorsToRegionalHighLatency(i awscloudwatch.CfnInsightRule)
}

// The jsii proxy for IServerSideOperationRegionalAlarmsAndRules
type jsiiProxy_IServerSideOperationRegionalAlarmsAndRules struct {
	jsiiProxy_IBaseOperationRegionalAlarmsAndRules
}

func (j *jsiiProxy_IServerSideOperationRegionalAlarmsAndRules) InstanceContributorsToRegionalFaults() awscloudwatch.CfnInsightRule {
	var returns awscloudwatch.CfnInsightRule
	_jsii_.Get(
		j,
		"instanceContributorsToRegionalFaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerSideOperationRegionalAlarmsAndRules)SetInstanceContributorsToRegionalFaults(val awscloudwatch.CfnInsightRule) {
	_jsii_.Set(
		j,
		"instanceContributorsToRegionalFaults",
		val,
	)
}

func (j *jsiiProxy_IServerSideOperationRegionalAlarmsAndRules) InstanceContributorsToRegionalHighLatency() awscloudwatch.CfnInsightRule {
	var returns awscloudwatch.CfnInsightRule
	_jsii_.Get(
		j,
		"instanceContributorsToRegionalHighLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerSideOperationRegionalAlarmsAndRules)SetInstanceContributorsToRegionalHighLatency(val awscloudwatch.CfnInsightRule) {
	_jsii_.Set(
		j,
		"instanceContributorsToRegionalHighLatency",
		val,
	)
}

