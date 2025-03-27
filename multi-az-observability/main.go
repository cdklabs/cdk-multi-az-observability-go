// A CDK construct for implementing multi-AZ observability to detect single AZ impairments
package multi-az-observability

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.AddCanaryTestProps",
		reflect.TypeOf((*AddCanaryTestProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/multi-az-observability.ApplicationLoadBalancerAvailabilityOutlierAlgorithm",
		reflect.TypeOf((*ApplicationLoadBalancerAvailabilityOutlierAlgorithm)(nil)).Elem(),
		map[string]interface{}{
			"STATIC": ApplicationLoadBalancerAvailabilityOutlierAlgorithm_STATIC,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.ApplicationLoadBalancerDetectionProps",
		reflect.TypeOf((*ApplicationLoadBalancerDetectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/multi-az-observability.ApplicationLoadBalancerLatencyOutlierAlgorithm",
		reflect.TypeOf((*ApplicationLoadBalancerLatencyOutlierAlgorithm)(nil)).Elem(),
		map[string]interface{}{
			"STATIC": ApplicationLoadBalancerLatencyOutlierAlgorithm_STATIC,
			"Z_SCORE": ApplicationLoadBalancerLatencyOutlierAlgorithm_Z_SCORE,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.AvailabilityZoneMapper",
		reflect.TypeOf((*AvailabilityZoneMapper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allAvailabilityZoneIdsAsArray", GoMethod: "AllAvailabilityZoneIdsAsArray"},
			_jsii_.MemberMethod{JsiiMethod: "allAvailabilityZoneIdsAsCommaDelimitedList", GoMethod: "AllAvailabilityZoneIdsAsCommaDelimitedList"},
			_jsii_.MemberMethod{JsiiMethod: "allAvailabilityZoneNamesAsCommaDelimitedList", GoMethod: "AllAvailabilityZoneNamesAsCommaDelimitedList"},
			_jsii_.MemberMethod{JsiiMethod: "availabilityZoneId", GoMethod: "AvailabilityZoneId"},
			_jsii_.MemberMethod{JsiiMethod: "availabilityZoneIdFromAvailabilityZoneLetter", GoMethod: "AvailabilityZoneIdFromAvailabilityZoneLetter"},
			_jsii_.MemberMethod{JsiiMethod: "availabilityZoneIdsAsArray", GoMethod: "AvailabilityZoneIdsAsArray"},
			_jsii_.MemberMethod{JsiiMethod: "availabilityZoneIdsAsCommaDelimitedList", GoMethod: "AvailabilityZoneIdsAsCommaDelimitedList"},
			_jsii_.MemberMethod{JsiiMethod: "availabilityZoneName", GoMethod: "AvailabilityZoneName"},
			_jsii_.MemberProperty{JsiiProperty: "function", GoGetter: "Function"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "mapper", GoGetter: "Mapper"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "regionPrefixForAvailabilityZoneIds", GoMethod: "RegionPrefixForAvailabilityZoneIds"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AvailabilityZoneMapper{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAvailabilityZoneMapper)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.AvailabilityZoneMapperProps",
		reflect.TypeOf((*AvailabilityZoneMapperProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.BasicServiceMultiAZObservability",
		reflect.TypeOf((*BasicServiceMultiAZObservability)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aggregateZonalIsolatedImpactAlarms", GoGetter: "AggregateZonalIsolatedImpactAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "albZonalIsolatedImpactAlarms", GoGetter: "AlbZonalIsolatedImpactAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "applicationLoadBalancers", GoGetter: "ApplicationLoadBalancers"},
			_jsii_.MemberProperty{JsiiProperty: "dashboard", GoGetter: "Dashboard"},
			_jsii_.MemberProperty{JsiiProperty: "natGateways", GoGetter: "NatGateways"},
			_jsii_.MemberProperty{JsiiProperty: "natGWZonalIsolatedImpactAlarms", GoGetter: "NatGWZonalIsolatedImpactAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BasicServiceMultiAZObservability{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBasicServiceMultiAZObservability)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.BasicServiceMultiAZObservabilityProps",
		reflect.TypeOf((*BasicServiceMultiAZObservabilityProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.CanaryMetricProps",
		reflect.TypeOf((*CanaryMetricProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.CanaryMetrics",
		reflect.TypeOf((*CanaryMetrics)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "canaryAvailabilityMetricDetails", GoGetter: "CanaryAvailabilityMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "canaryLatencyMetricDetails", GoGetter: "CanaryLatencyMetricDetails"},
		},
		func() interface{} {
			j := jsiiProxy_CanaryMetrics{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICanaryMetrics)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.CanaryTestAvailabilityMetricsOverride",
		reflect.TypeOf((*CanaryTestAvailabilityMetricsOverride)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultAlarmThreshold", GoGetter: "FaultAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
		},
		func() interface{} {
			j := jsiiProxy_CanaryTestAvailabilityMetricsOverride{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CanaryTestMetricsOverride)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.CanaryTestAvailabilityMetricsOverrideProps",
		reflect.TypeOf((*CanaryTestAvailabilityMetricsOverrideProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.CanaryTestLatencyMetricsOverride",
		reflect.TypeOf((*CanaryTestLatencyMetricsOverride)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
		},
		func() interface{} {
			j := jsiiProxy_CanaryTestLatencyMetricsOverride{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CanaryTestMetricsOverride)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.CanaryTestLatencyMetricsOverrideProps",
		reflect.TypeOf((*CanaryTestLatencyMetricsOverrideProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.CanaryTestMetricsOverride",
		reflect.TypeOf((*CanaryTestMetricsOverride)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
		},
		func() interface{} {
			j := jsiiProxy_CanaryTestMetricsOverride{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICanaryTestMetricsOverride)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.CanaryTestMetricsOverrideProps",
		reflect.TypeOf((*CanaryTestMetricsOverrideProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.ContributorInsightRuleDetails",
		reflect.TypeOf((*ContributorInsightRuleDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIdJsonPath", GoGetter: "AvailabilityZoneIdJsonPath"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricJsonPath", GoGetter: "FaultMetricJsonPath"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdJsonPath", GoGetter: "InstanceIdJsonPath"},
			_jsii_.MemberProperty{JsiiProperty: "logGroups", GoGetter: "LogGroups"},
			_jsii_.MemberProperty{JsiiProperty: "operationNameJsonPath", GoGetter: "OperationNameJsonPath"},
			_jsii_.MemberProperty{JsiiProperty: "successLatencyMetricJsonPath", GoGetter: "SuccessLatencyMetricJsonPath"},
		},
		func() interface{} {
			j := jsiiProxy_ContributorInsightRuleDetails{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IContributorInsightRuleDetails)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.ContributorInsightRuleDetailsProps",
		reflect.TypeOf((*ContributorInsightRuleDetailsProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IAvailabilityZoneMapper",
		reflect.TypeOf((*IAvailabilityZoneMapper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allAvailabilityZoneIdsAsArray", GoMethod: "AllAvailabilityZoneIdsAsArray"},
			_jsii_.MemberMethod{JsiiMethod: "allAvailabilityZoneIdsAsCommaDelimitedList", GoMethod: "AllAvailabilityZoneIdsAsCommaDelimitedList"},
			_jsii_.MemberMethod{JsiiMethod: "allAvailabilityZoneNamesAsCommaDelimitedList", GoMethod: "AllAvailabilityZoneNamesAsCommaDelimitedList"},
			_jsii_.MemberMethod{JsiiMethod: "availabilityZoneId", GoMethod: "AvailabilityZoneId"},
			_jsii_.MemberMethod{JsiiMethod: "availabilityZoneIdFromAvailabilityZoneLetter", GoMethod: "AvailabilityZoneIdFromAvailabilityZoneLetter"},
			_jsii_.MemberMethod{JsiiMethod: "availabilityZoneIdsAsArray", GoMethod: "AvailabilityZoneIdsAsArray"},
			_jsii_.MemberMethod{JsiiMethod: "availabilityZoneIdsAsCommaDelimitedList", GoMethod: "AvailabilityZoneIdsAsCommaDelimitedList"},
			_jsii_.MemberMethod{JsiiMethod: "availabilityZoneName", GoMethod: "AvailabilityZoneName"},
			_jsii_.MemberProperty{JsiiProperty: "function", GoGetter: "Function"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "mapper", GoGetter: "Mapper"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "regionPrefixForAvailabilityZoneIds", GoMethod: "RegionPrefixForAvailabilityZoneIds"},
		},
		func() interface{} {
			j := jsiiProxy_IAvailabilityZoneMapper{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IBaseOperationRegionalAlarmsAndRules",
		reflect.TypeOf((*IBaseOperationRegionalAlarmsAndRules)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityAlarm", GoGetter: "AvailabilityAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityOrLatencyAlarm", GoGetter: "AvailabilityOrLatencyAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "latencyAlarm", GoGetter: "LatencyAlarm"},
		},
		func() interface{} {
			return &jsiiProxy_IBaseOperationRegionalAlarmsAndRules{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IBaseOperationZonalAlarmsAndRules",
		reflect.TypeOf((*IBaseOperationZonalAlarmsAndRules)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityAlarm", GoGetter: "AvailabilityAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIsOutlierForFaults", GoGetter: "AvailabilityZoneIsOutlierForFaults"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIsOutlierForLatency", GoGetter: "AvailabilityZoneIsOutlierForLatency"},
			_jsii_.MemberProperty{JsiiProperty: "latencyAlarm", GoGetter: "LatencyAlarm"},
		},
		func() interface{} {
			return &jsiiProxy_IBaseOperationZonalAlarmsAndRules{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IBasicServiceMultiAZObservability",
		reflect.TypeOf((*IBasicServiceMultiAZObservability)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aggregateZonalIsolatedImpactAlarms", GoGetter: "AggregateZonalIsolatedImpactAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "albZonalIsolatedImpactAlarms", GoGetter: "AlbZonalIsolatedImpactAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "applicationLoadBalancers", GoGetter: "ApplicationLoadBalancers"},
			_jsii_.MemberProperty{JsiiProperty: "dashboard", GoGetter: "Dashboard"},
			_jsii_.MemberProperty{JsiiProperty: "natGateways", GoGetter: "NatGateways"},
			_jsii_.MemberProperty{JsiiProperty: "natGWZonalIsolatedImpactAlarms", GoGetter: "NatGWZonalIsolatedImpactAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
		},
		func() interface{} {
			j := jsiiProxy_IBasicServiceMultiAZObservability{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.ICanaryMetrics",
		reflect.TypeOf((*ICanaryMetrics)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "canaryAvailabilityMetricDetails", GoGetter: "CanaryAvailabilityMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "canaryLatencyMetricDetails", GoGetter: "CanaryLatencyMetricDetails"},
		},
		func() interface{} {
			return &jsiiProxy_ICanaryMetrics{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.ICanaryOperationRegionalAlarmsAndRules",
		reflect.TypeOf((*ICanaryOperationRegionalAlarmsAndRules)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityAlarm", GoGetter: "AvailabilityAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityOrLatencyAlarm", GoGetter: "AvailabilityOrLatencyAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "latencyAlarm", GoGetter: "LatencyAlarm"},
		},
		func() interface{} {
			j := jsiiProxy_ICanaryOperationRegionalAlarmsAndRules{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBaseOperationRegionalAlarmsAndRules)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.ICanaryOperationZonalAlarmsAndRules",
		reflect.TypeOf((*ICanaryOperationZonalAlarmsAndRules)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityAlarm", GoGetter: "AvailabilityAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIsOutlierForFaults", GoGetter: "AvailabilityZoneIsOutlierForFaults"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIsOutlierForLatency", GoGetter: "AvailabilityZoneIsOutlierForLatency"},
			_jsii_.MemberProperty{JsiiProperty: "isolatedImpactAlarm", GoGetter: "IsolatedImpactAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "latencyAlarm", GoGetter: "LatencyAlarm"},
		},
		func() interface{} {
			j := jsiiProxy_ICanaryOperationZonalAlarmsAndRules{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBaseOperationZonalAlarmsAndRules)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.ICanaryTestAvailabilityMetricsOverride",
		reflect.TypeOf((*ICanaryTestAvailabilityMetricsOverride)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultAlarmThreshold", GoGetter: "FaultAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
		},
		func() interface{} {
			j := jsiiProxy_ICanaryTestAvailabilityMetricsOverride{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICanaryTestMetricsOverride)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.ICanaryTestLatencyMetricsOverride",
		reflect.TypeOf((*ICanaryTestLatencyMetricsOverride)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
		},
		func() interface{} {
			j := jsiiProxy_ICanaryTestLatencyMetricsOverride{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICanaryTestMetricsOverride)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.ICanaryTestMetricsOverride",
		reflect.TypeOf((*ICanaryTestMetricsOverride)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
		},
		func() interface{} {
			return &jsiiProxy_ICanaryTestMetricsOverride{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IContributorInsightRuleDetails",
		reflect.TypeOf((*IContributorInsightRuleDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIdJsonPath", GoGetter: "AvailabilityZoneIdJsonPath"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricJsonPath", GoGetter: "FaultMetricJsonPath"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdJsonPath", GoGetter: "InstanceIdJsonPath"},
			_jsii_.MemberProperty{JsiiProperty: "logGroups", GoGetter: "LogGroups"},
			_jsii_.MemberProperty{JsiiProperty: "operationNameJsonPath", GoGetter: "OperationNameJsonPath"},
			_jsii_.MemberProperty{JsiiProperty: "successLatencyMetricJsonPath", GoGetter: "SuccessLatencyMetricJsonPath"},
		},
		func() interface{} {
			return &jsiiProxy_IContributorInsightRuleDetails{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IInstrumentedServiceMultiAZObservability",
		reflect.TypeOf((*IInstrumentedServiceMultiAZObservability)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "canaryLogGroup", GoGetter: "CanaryLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "operationDashboards", GoGetter: "OperationDashboards"},
			_jsii_.MemberProperty{JsiiProperty: "perOperationAlarmsAndRules", GoGetter: "PerOperationAlarmsAndRules"},
			_jsii_.MemberProperty{JsiiProperty: "perOperationZonalImpactAlarms", GoGetter: "PerOperationZonalImpactAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAlarms", GoGetter: "ServiceAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "serviceDashboard", GoGetter: "ServiceDashboard"},
		},
		func() interface{} {
			j := jsiiProxy_IInstrumentedServiceMultiAZObservability{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IOperation",
		reflect.TypeOf((*IOperation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "canaryMetricDetails", GoGetter: "CanaryMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "canaryTestAvailabilityMetricsOverride", GoGetter: "CanaryTestAvailabilityMetricsOverride"},
			_jsii_.MemberProperty{JsiiProperty: "canaryTestLatencyMetricsOverride", GoGetter: "CanaryTestLatencyMetricsOverride"},
			_jsii_.MemberProperty{JsiiProperty: "canaryTestProps", GoGetter: "CanaryTestProps"},
			_jsii_.MemberProperty{JsiiProperty: "critical", GoGetter: "Critical"},
			_jsii_.MemberProperty{JsiiProperty: "httpMethods", GoGetter: "HttpMethods"},
			_jsii_.MemberProperty{JsiiProperty: "operationName", GoGetter: "OperationName"},
			_jsii_.MemberProperty{JsiiProperty: "optOutOfServiceCreatedCanary", GoGetter: "OptOutOfServiceCreatedCanary"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideAvailabilityMetricDetails", GoGetter: "ServerSideAvailabilityMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideContributorInsightRuleDetails", GoGetter: "ServerSideContributorInsightRuleDetails"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideLatencyMetricDetails", GoGetter: "ServerSideLatencyMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
		},
		func() interface{} {
			return &jsiiProxy_IOperation{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IOperationAlarmsAndRules",
		reflect.TypeOf((*IOperationAlarmsAndRules)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aggregateZonalAlarms", GoGetter: "AggregateZonalAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "canaryRegionalAlarmsAndRules", GoGetter: "CanaryRegionalAlarmsAndRules"},
			_jsii_.MemberProperty{JsiiProperty: "canaryZonalAlarmsAndRules", GoGetter: "CanaryZonalAlarmsAndRules"},
			_jsii_.MemberProperty{JsiiProperty: "operation", GoGetter: "Operation"},
			_jsii_.MemberProperty{JsiiProperty: "regionalImpactAlarm", GoGetter: "RegionalImpactAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideRegionalAlarmsAndRules", GoGetter: "ServerSideRegionalAlarmsAndRules"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideZonalAlarmsAndRules", GoGetter: "ServerSideZonalAlarmsAndRules"},
		},
		func() interface{} {
			return &jsiiProxy_IOperationAlarmsAndRules{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IOperationAvailabilityMetricDetails",
		reflect.TypeOf((*IOperationAvailabilityMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultAlarmThreshold", GoGetter: "FaultAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricDimensions", GoGetter: "MetricDimensions"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "operationName", GoGetter: "OperationName"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			j := jsiiProxy_IOperationAvailabilityMetricDetails{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOperationMetricDetails)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IOperationLatencyMetricDetails",
		reflect.TypeOf((*IOperationLatencyMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricDimensions", GoGetter: "MetricDimensions"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "operationName", GoGetter: "OperationName"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			j := jsiiProxy_IOperationLatencyMetricDetails{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOperationMetricDetails)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IOperationMetricDetails",
		reflect.TypeOf((*IOperationMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricDimensions", GoGetter: "MetricDimensions"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "operationName", GoGetter: "OperationName"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			return &jsiiProxy_IOperationMetricDetails{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IServerSideOperationRegionalAlarmsAndRules",
		reflect.TypeOf((*IServerSideOperationRegionalAlarmsAndRules)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityAlarm", GoGetter: "AvailabilityAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityOrLatencyAlarm", GoGetter: "AvailabilityOrLatencyAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "instanceContributorsToRegionalFaults", GoGetter: "InstanceContributorsToRegionalFaults"},
			_jsii_.MemberProperty{JsiiProperty: "instanceContributorsToRegionalHighLatency", GoGetter: "InstanceContributorsToRegionalHighLatency"},
			_jsii_.MemberProperty{JsiiProperty: "latencyAlarm", GoGetter: "LatencyAlarm"},
		},
		func() interface{} {
			j := jsiiProxy_IServerSideOperationRegionalAlarmsAndRules{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBaseOperationRegionalAlarmsAndRules)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IServerSideOperationZonalAlarmsAndRules",
		reflect.TypeOf((*IServerSideOperationZonalAlarmsAndRules)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityAlarm", GoGetter: "AvailabilityAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIsOutlierForFaults", GoGetter: "AvailabilityZoneIsOutlierForFaults"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIsOutlierForLatency", GoGetter: "AvailabilityZoneIsOutlierForLatency"},
			_jsii_.MemberProperty{JsiiProperty: "instanceContributorsToFaultsInThisAZ", GoGetter: "InstanceContributorsToFaultsInThisAZ"},
			_jsii_.MemberProperty{JsiiProperty: "instanceContributorsToHighLatencyInThisAZ", GoGetter: "InstanceContributorsToHighLatencyInThisAZ"},
			_jsii_.MemberProperty{JsiiProperty: "instancesHandlingRequestsInThisAZ", GoGetter: "InstancesHandlingRequestsInThisAZ"},
			_jsii_.MemberProperty{JsiiProperty: "isolatedImpactAlarm", GoGetter: "IsolatedImpactAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "latencyAlarm", GoGetter: "LatencyAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "multipleInstancesProducingFaultsInThisAvailabilityZone", GoGetter: "MultipleInstancesProducingFaultsInThisAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "multipleInstancesProducingHighLatencyInThisAZ", GoGetter: "MultipleInstancesProducingHighLatencyInThisAZ"},
		},
		func() interface{} {
			j := jsiiProxy_IServerSideOperationZonalAlarmsAndRules{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBaseOperationZonalAlarmsAndRules)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IService",
		reflect.TypeOf((*IService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOperation", GoMethod: "AddOperation"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneNames", GoGetter: "AvailabilityZoneNames"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrl", GoGetter: "BaseUrl"},
			_jsii_.MemberProperty{JsiiProperty: "canaryTestProps", GoGetter: "CanaryTestProps"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAvailabilityMetricDetails", GoGetter: "DefaultAvailabilityMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "defaultContributorInsightRuleDetails", GoGetter: "DefaultContributorInsightRuleDetails"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLatencyMetricDetails", GoGetter: "DefaultLatencyMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "faultCountThreshold", GoGetter: "FaultCountThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancer", GoGetter: "LoadBalancer"},
			_jsii_.MemberProperty{JsiiProperty: "operations", GoGetter: "Operations"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
		},
		func() interface{} {
			return &jsiiProxy_IService{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IServiceAlarmsAndRules",
		reflect.TypeOf((*IServiceAlarmsAndRules)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "regionalCanaryAlarm", GoGetter: "RegionalCanaryAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "regionalImpactAlarm", GoGetter: "RegionalImpactAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "regionalServerSideImpactAlarm", GoGetter: "RegionalServerSideImpactAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "serviceImpactAlarm", GoGetter: "ServiceImpactAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "zonalAggregateIsolatedImpactAlarms", GoGetter: "ZonalAggregateIsolatedImpactAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "zonalServerSideIsolatedImpactAlarms", GoGetter: "ZonalServerSideIsolatedImpactAlarms"},
		},
		func() interface{} {
			return &jsiiProxy_IServiceAlarmsAndRules{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IServiceAvailabilityMetricDetails",
		reflect.TypeOf((*IServiceAvailabilityMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultAlarmThreshold", GoGetter: "FaultAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			j := jsiiProxy_IServiceAvailabilityMetricDetails{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServiceMetricDetails)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IServiceLatencyMetricDetails",
		reflect.TypeOf((*IServiceLatencyMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			j := jsiiProxy_IServiceLatencyMetricDetails{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServiceMetricDetails)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/multi-az-observability.IServiceMetricDetails",
		reflect.TypeOf((*IServiceMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			return &jsiiProxy_IServiceMetricDetails{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.InstrumentedServiceMultiAZObservability",
		reflect.TypeOf((*InstrumentedServiceMultiAZObservability)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "canaryLogGroup", GoGetter: "CanaryLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "operationDashboards", GoGetter: "OperationDashboards"},
			_jsii_.MemberProperty{JsiiProperty: "perOperationAlarmsAndRules", GoGetter: "PerOperationAlarmsAndRules"},
			_jsii_.MemberProperty{JsiiProperty: "perOperationZonalImpactAlarms", GoGetter: "PerOperationZonalImpactAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAlarms", GoGetter: "ServiceAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "serviceDashboard", GoGetter: "ServiceDashboard"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_InstrumentedServiceMultiAZObservability{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInstrumentedServiceMultiAZObservability)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.InstrumentedServiceMultiAZObservabilityProps",
		reflect.TypeOf((*InstrumentedServiceMultiAZObservabilityProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/multi-az-observability.LatencyOutlierMetricAggregation",
		reflect.TypeOf((*LatencyOutlierMetricAggregation)(nil)).Elem(),
		map[string]interface{}{
			"COUNT": LatencyOutlierMetricAggregation_COUNT,
			"VALUE": LatencyOutlierMetricAggregation_VALUE,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.MetricDimensions",
		reflect.TypeOf((*MetricDimensions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIdKey", GoGetter: "AvailabilityZoneIdKey"},
			_jsii_.MemberMethod{JsiiMethod: "regionalDimensions", GoMethod: "RegionalDimensions"},
			_jsii_.MemberProperty{JsiiProperty: "regionKey", GoGetter: "RegionKey"},
			_jsii_.MemberProperty{JsiiProperty: "staticDimensions", GoGetter: "StaticDimensions"},
			_jsii_.MemberMethod{JsiiMethod: "zonalDimensions", GoMethod: "ZonalDimensions"},
		},
		func() interface{} {
			return &jsiiProxy_MetricDimensions{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.NatGatewayDetectionProps",
		reflect.TypeOf((*NatGatewayDetectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.NetworkConfigurationProps",
		reflect.TypeOf((*NetworkConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.Operation",
		reflect.TypeOf((*Operation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "canaryMetricDetails", GoGetter: "CanaryMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "canaryTestAvailabilityMetricsOverride", GoGetter: "CanaryTestAvailabilityMetricsOverride"},
			_jsii_.MemberProperty{JsiiProperty: "canaryTestLatencyMetricsOverride", GoGetter: "CanaryTestLatencyMetricsOverride"},
			_jsii_.MemberProperty{JsiiProperty: "canaryTestProps", GoGetter: "CanaryTestProps"},
			_jsii_.MemberProperty{JsiiProperty: "critical", GoGetter: "Critical"},
			_jsii_.MemberProperty{JsiiProperty: "httpMethods", GoGetter: "HttpMethods"},
			_jsii_.MemberProperty{JsiiProperty: "operationName", GoGetter: "OperationName"},
			_jsii_.MemberProperty{JsiiProperty: "optOutOfServiceCreatedCanary", GoGetter: "OptOutOfServiceCreatedCanary"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideAvailabilityMetricDetails", GoGetter: "ServerSideAvailabilityMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideContributorInsightRuleDetails", GoGetter: "ServerSideContributorInsightRuleDetails"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideLatencyMetricDetails", GoGetter: "ServerSideLatencyMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
		},
		func() interface{} {
			j := jsiiProxy_Operation{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOperation)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.OperationAvailabilityMetricDetails",
		reflect.TypeOf((*OperationAvailabilityMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultAlarmThreshold", GoGetter: "FaultAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricDimensions", GoGetter: "MetricDimensions"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "operationName", GoGetter: "OperationName"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			j := jsiiProxy_OperationAvailabilityMetricDetails{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_OperationMetricDetails)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOperationAvailabilityMetricDetails)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.OperationAvailabilityMetricDetailsProps",
		reflect.TypeOf((*OperationAvailabilityMetricDetailsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.OperationLatencyMetricDetails",
		reflect.TypeOf((*OperationLatencyMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricDimensions", GoGetter: "MetricDimensions"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "operationName", GoGetter: "OperationName"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			j := jsiiProxy_OperationLatencyMetricDetails{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_OperationMetricDetails)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOperationLatencyMetricDetails)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.OperationLatencyMetricDetailsProps",
		reflect.TypeOf((*OperationLatencyMetricDetailsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.OperationMetricDetails",
		reflect.TypeOf((*OperationMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricDimensions", GoGetter: "MetricDimensions"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "operationName", GoGetter: "OperationName"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			j := jsiiProxy_OperationMetricDetails{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOperationMetricDetails)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.OperationMetricDetailsProps",
		reflect.TypeOf((*OperationMetricDetailsProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.OperationProps",
		reflect.TypeOf((*OperationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/multi-az-observability.OutlierDetectionAlgorithm",
		reflect.TypeOf((*OutlierDetectionAlgorithm)(nil)).Elem(),
		map[string]interface{}{
			"STATIC": OutlierDetectionAlgorithm_STATIC,
			"CHI_SQUARED": OutlierDetectionAlgorithm_CHI_SQUARED,
			"Z_SCORE": OutlierDetectionAlgorithm_Z_SCORE,
			"IQR": OutlierDetectionAlgorithm_IQR,
			"MAD": OutlierDetectionAlgorithm_MAD,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/multi-az-observability.PacketLossOutlierAlgorithm",
		reflect.TypeOf((*PacketLossOutlierAlgorithm)(nil)).Elem(),
		map[string]interface{}{
			"STATIC": PacketLossOutlierAlgorithm_STATIC,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.Service",
		reflect.TypeOf((*Service)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOperation", GoMethod: "AddOperation"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneNames", GoGetter: "AvailabilityZoneNames"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrl", GoGetter: "BaseUrl"},
			_jsii_.MemberProperty{JsiiProperty: "canaryTestProps", GoGetter: "CanaryTestProps"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAvailabilityMetricDetails", GoGetter: "DefaultAvailabilityMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "defaultContributorInsightRuleDetails", GoGetter: "DefaultContributorInsightRuleDetails"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLatencyMetricDetails", GoGetter: "DefaultLatencyMetricDetails"},
			_jsii_.MemberProperty{JsiiProperty: "faultCountThreshold", GoGetter: "FaultCountThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancer", GoGetter: "LoadBalancer"},
			_jsii_.MemberProperty{JsiiProperty: "operations", GoGetter: "Operations"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
		},
		func() interface{} {
			j := jsiiProxy_Service{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IService)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.ServiceAvailabilityMetricDetails",
		reflect.TypeOf((*ServiceAvailabilityMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultAlarmThreshold", GoGetter: "FaultAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceAvailabilityMetricDetails{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceMetricDetails)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServiceAvailabilityMetricDetails)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.ServiceAvailabilityMetricDetailsProps",
		reflect.TypeOf((*ServiceAvailabilityMetricDetailsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.ServiceLatencyMetricDetails",
		reflect.TypeOf((*ServiceLatencyMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successAlarmThreshold", GoGetter: "SuccessAlarmThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceLatencyMetricDetails{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceMetricDetails)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServiceLatencyMetricDetails)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.ServiceLatencyMetricDetailsProps",
		reflect.TypeOf((*ServiceLatencyMetricDetailsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/multi-az-observability.ServiceMetricDetails",
		reflect.TypeOf((*ServiceMetricDetails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmStatistic", GoGetter: "AlarmStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "faultMetricNames", GoGetter: "FaultMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "graphedFaultStatistics", GoGetter: "GraphedFaultStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "graphedSuccessStatistics", GoGetter: "GraphedSuccessStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "metricNamespace", GoGetter: "MetricNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "successMetricNames", GoGetter: "SuccessMetricNames"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceMetricDetails{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServiceMetricDetails)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.ServiceMetricDetailsProps",
		reflect.TypeOf((*ServiceMetricDetailsProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/multi-az-observability.ServiceProps",
		reflect.TypeOf((*ServiceProps)(nil)).Elem(),
	)
}
