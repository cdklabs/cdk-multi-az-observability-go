![Build Workflow](https://github.com/cdklabs/cdk-multi-az-observability/actions/workflows/build.yml/badge.svg) ![Release Workflow](https://github.com/cdklabs/cdk-multi-az-observability/actions/workflows/release.yml/badge.svg) ![GitHub Release](https://img.shields.io/github/v/release/cdklabs/cdk-multi-az-observability?include_prereleases&sort=semver&logo=github&label=version)

# multi-az-observability

This is a CDK construct for multi-AZ observability to help detect single-AZ impairments. This is currently an `alpha` version, but is being used in the AWS [Advanced Multi-AZ Resilience Patterns](https://catalog.workshops.aws/multi-az-gray-failures/en-US) workshop.

There is a lot of available information to think through and combine to provide signals about single-AZ impact. To simplify the setup and use reasonable defaults, this construct (available in [TypeScript](https://www.npmjs.com/package/@cdklabs/multi-az-observability), [Go](https://github.com/cdklabs/cdk-multi-az-observability-go), [Python](https://pypi.org/project/cdklabs.multi-az-observability/), [.NET](https://www.nuget.org/packages/Cdklabs.MultiAZObservability), and [Java](https://central.sonatype.com/artifact/io.github.cdklabs/cdk-multi-az-observability)) sets up the necessary observability. To use the CDK construct, you first define your service like this:

```go
service := multi-az-observability.NewService(&ServiceProps{
	ServiceName: jsii.String("test"),
	AvailabilityZoneNames: vpc.AvailabilityZones,
	BaseUrl: jsii.String("http://www.example.com"),
	FaultCountThreshold: jsii.Number(25),
	Period: awscdk.Duration_Seconds(jsii.Number(60)),
	LoadBalancer: loadBalancer,
	TargetGroups: []iTargetGroup{
		targetGroup1,
		targetGroup2,
	},
	DefaultAvailabilityMetricDetails: multi-az-observability.NewServiceAvailabilityMetricDetails(&ServiceAvailabilityMetricDetailsProps{
		MetricNamespace: jsii.String("front-end/metrics"),
		SuccessMetricNames: []*string{
			jsii.String("Success"),
		},
		FaultMetricNames: []*string{
			jsii.String("Fault"),
			jsii.String("Error"),
		},
		AlarmStatistic: jsii.String("Sum"),
		Unit: awscdk.Unit_COUNT,
		Period: awscdk.Duration_*Seconds(jsii.Number(60)),
		EvaluationPeriods: jsii.Number(5),
		DatapointsToAlarm: jsii.Number(3),
		SuccessAlarmThreshold: jsii.Number(99.9),
		FaultAlarmThreshold: jsii.Number(0.1),
		GraphedFaultStatistics: []*string{
			jsii.String("Sum"),
		},
		GraphedSuccessStatistics: []*string{
			jsii.String("Sum"),
		},
	}),
	DefaultLatencyMetricDetails: multi-az-observability.NewServiceLatencyMetricDetails(&ServiceLatencyMetricDetailsProps{
		MetricNamespace: jsii.String("front-end/metrics"),
		SuccessMetricNames: []*string{
			jsii.String("SuccessLatency"),
		},
		FaultMetricNames: []*string{
			jsii.String("FaultLatency"),
		},
		AlarmStatistic: jsii.String("p99"),
		Unit: awscdk.Unit_MILLISECONDS,
		Period: awscdk.Duration_*Seconds(jsii.Number(60)),
		EvaluationPeriods: jsii.Number(5),
		DatapointsToAlarm: jsii.Number(3),
		SuccessAlarmThreshold: awscdk.Duration_Millis(jsii.Number(150)),
		GraphedFaultStatistics: []*string{
			jsii.String("p99"),
		},
		GraphedSuccessStatistics: []*string{
			jsii.String("p50"),
			jsii.String("p99"),
			jsii.String("tm99"),
		},
	}),
	DefaultContributorInsightRuleDetails: multi-az-observability.NewContributorInsightRuleDetails(&ContributorInsightRuleDetailsProps{
		SuccessLatencyMetricJsonPath: jsii.String("$.SuccessLatency"),
		FaultMetricJsonPath: jsii.String("$.Faults"),
		OperationNameJsonPath: jsii.String("$.Operation"),
		InstanceIdJsonPath: jsii.String("$.InstanceId"),
		AvailabilityZoneIdJsonPath: jsii.String("$.AZ-ID"),
		LogGroups: []iLogGroup{
			logGroup,
		},
	}),
	CanaryTestProps: &AddCanaryTestProps{
		RequestCount: jsii.Number(10),
		Schedule: jsii.String("rate(1 minute)"),
		LoadBalancer: loadBalancer,
		NetworkConfiguration: &NetworkConfigurationProps{
			Vpc: vpc,
			SubnetSelection: &SubnetSelection{
				SubnetType: awscdk.SubnetType_PRIVATE_ISOLATED,
			},
		},
	},
	MinimumUnhealthyTargets: &MinimumUnhealthyTargets{
		Percentage: jsii.Number(0.1),
	},
})

rideOperation := map[string]interface{}{
	"operationName": jsii.String("ride"),
	"service": service,
	"path": jsii.String("/ride"),
	"critical": jsii.Boolean(true),
	"httpMethods": []*string{
		jsii.String("GET"),
	},
	"serverSideContributorInsightRuleDetails": multi-az-observability.NewContributorInsightRuleDetails(&ContributorInsightRuleDetailsProps{
		"logGroups": []*iLogGroup{
			logGroup,
		},
		"successLatencyMetricJsonPath": jsii.String("$.SuccessLatency"),
		"faultMetricJsonPath": jsii.String("$.Faults"),
		"operationNameJsonPath": jsii.String("$.Operation"),
		"instanceIdJsonPath": jsii.String("$.InstanceId"),
		"availabilityZoneIdJsonPath": jsii.String("$.AZ-ID"),
	}),
	"serverSideAvailabilityMetricDetails": multi-az-observability.NewOperationAvailabilityMetricDetails(&OperationAvailabilityMetricDetailsProps{
		"operationName": jsii.String("ride"),
		"metricDimensions": multi-az-observability.NewMetricDimensions(map[string]*string{
			"Operation": jsii.String("ride"),
		}, jsii.String("AZ-ID"), jsii.String("Region")),
	}, service.defaultAvailabilityMetricDetails),
	"serverSideLatencyMetricDetails": multi-az-observability.NewOperationLatencyMetricDetails(&OperationLatencyMetricDetailsProps{
		"operationName": jsii.String("ride"),
		"metricDimensions": multi-az-observability.NewMetricDimensions(map[string]*string{
			"Operation": jsii.String("ride"),
		}, jsii.String("AZ-ID"), jsii.String("Region")),
	}, service.defaultLatencyMetricDetails),
}

payOperation := map[string]interface{}{
	"operationName": jsii.String("pay"),
	"service": service,
	"path": jsii.String("/pay"),
	"critical": jsii.Boolean(true),
	"httpMethods": []*string{
		jsii.String("GET"),
	},
	"serverSideContributorInsightRuleDetails": multi-az-observability.NewContributorInsightRuleDetails(&ContributorInsightRuleDetailsProps{
		"logGroups": []*iLogGroup{
			logGroup,
		},
		"successLatencyMetricJsonPath": jsii.String("$.SuccessLatency"),
		"faultMetricJsonPath": jsii.String("$.Faults"),
		"operationNameJsonPath": jsii.String("$.Operation"),
		"instanceIdJsonPath": jsii.String("$.InstanceId"),
		"availabilityZoneIdJsonPath": jsii.String("$.AZ-ID"),
	}),
	"serverSideAvailabilityMetricDetails": multi-az-observability.NewOperationAvailabilityMetricDetails(&OperationAvailabilityMetricDetailsProps{
		"operationName": jsii.String("pay"),
		"metricDimensions": multi-az-observability.NewMetricDimensions(map[string]*string{
			"Operation": jsii.String("ride"),
		}, jsii.String("AZ-ID"), jsii.String("Region")),
	}, service.defaultAvailabilityMetricDetails),
	"serverSideLatencyMetricDetails": multi-az-observability.NewOperationLatencyMetricDetails(&OperationLatencyMetricDetailsProps{
		"operationName": jsii.String("pay"),
		"metricDimensions": multi-az-observability.NewMetricDimensions(map[string]*string{
			"Operation": jsii.String("ride"),
		}, jsii.String("AZ-ID"), jsii.String("Region")),
	}, service.defaultLatencyMetricDetails),
}

service.AddOperation(rideOperation)
service.AddOperation(payOperation)
```

Then you provide that service definition to the CDK construct.

```go
multi-az-observability.NewInstrumentedServiceMultiAZObservability(stack, jsii.String("MAZObservability"), &InstrumentedServiceMultiAZObservabilityProps{
	CreateDashboards: jsii.Boolean(true),
	Service: service,
	Interval: awscdk.Duration_Minutes(jsii.Number(60)),
})
```

You define some characteristics of the service, default values for metrics and alarms, and then add operations as well as any overrides for default values that you need. The construct can also automatically create synthetic canaries that test each operation with a very simple HTTP check, or you can configure your own synthetics and just tell the construct about the metric details and optionally log files. This creates metrics, alarms, and dashboards that can be used to detect single-AZ impact. You can access these alarms from the `multiAvailabilityZoneObservability` object and use them in your CDK project to start automation, send SNS notifications, or incorporate in your own dashboards.

If you don't have service specific logs and custom metrics with per-AZ dimensions, you can still use the construct to evaluate ALB and/or NAT Gateway metrics to find single AZ impairments.

```go
multi-az-observability.NewBasicServiceMultiAZObservability(stack, jsii.String("MAZObservability"), &BasicServiceMultiAZObservabilityProps{
	ApplicationLoadBalancerProps: &ApplicationLoadBalancerDetectionProps{
		AlbTargetGroupMap: []albTargetGroupMap{
			&albTargetGroupMap{
				ApplicationLoadBalancer: awscdk.NewApplicationLoadBalancer(stack, jsii.String("alb"), &ApplicationLoadBalancerProps{
					Vpc: vpc,
					CrossZoneEnabled: jsii.Boolean(true),
				}),
				TargetGroups: []iTargetGroup{
					targetGroup1,
					targetGroup2,
				},
			},
		},
		FaultCountPercentThreshold: jsii.Number(1),
		LatencyStatistic: awscdk.Stats_Percentile(jsii.Number(99)),
		LatencyThreshold: awscdk.Duration_Millis(jsii.Number(200)),
		LatencyOutlierAlgorithm: multi-az-observability.ApplicationLoadBalancerLatencyOutlierAlgorithm_STATIC,
		LatencyOutlierThreshold: jsii.Number(45),
	},
	NatGatewayProps: &NatGatewayDetectionProps{
		NatGateways: map[string][]cfnNatGateway{
			"us-east-1a": []*cfnNatGateway{
				natGateway1,
			},
			"us-east-1b": []*cfnNatGateway{
				natGateway2,
			},
			"us-east-1c": []*cfnNatGateway{
				natGateway3,
			},
		},
		PacketLossPercentThreshold: jsii.Number(0.01),
	},
	ServiceName: jsii.String("test"),
	Period: awscdk.Duration_Seconds(jsii.Number(60)),
	CreateDashboard: jsii.Boolean(true),
	EvaluationPeriods: jsii.Number(5),
	DatapointsToAlarm: jsii.Number(3),
})
```

If you provide a load balancer, the construct assumes it is deployed in each AZ of the VPC the load balancer is associated with and will look for HTTP metrics using those AZs as dimensions.

Both options support running workloads on EC2, ECS, Lambda, and EKS.
