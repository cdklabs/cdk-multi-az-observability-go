//go:build no_runtime_type_checking

package multi-az-observability

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IServiceAlarmsAndRules) validateSetRegionalFaultCountServerSideAlarmParameters(val awscloudwatch.IAlarm) error {
	return nil
}

func (j *jsiiProxy_IServiceAlarmsAndRules) validateSetServiceParameters(val IService) error {
	return nil
}

func (j *jsiiProxy_IServiceAlarmsAndRules) validateSetZonalAggregateIsolatedImpactAlarmsParameters(val *map[string]awscloudwatch.IAlarm) error {
	return nil
}

func (j *jsiiProxy_IServiceAlarmsAndRules) validateSetZonalServerSideIsolatedImpactAlarmsParameters(val *map[string]awscloudwatch.IAlarm) error {
	return nil
}

