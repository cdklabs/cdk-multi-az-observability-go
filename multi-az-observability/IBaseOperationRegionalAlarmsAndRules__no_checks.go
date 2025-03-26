//go:build no_runtime_type_checking

package multi-az-observability

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IBaseOperationRegionalAlarmsAndRules) validateSetAvailabilityAlarmParameters(val awscloudwatch.IAlarm) error {
	return nil
}

func (j *jsiiProxy_IBaseOperationRegionalAlarmsAndRules) validateSetAvailabilityOrLatencyAlarmParameters(val awscloudwatch.IAlarm) error {
	return nil
}

func (j *jsiiProxy_IBaseOperationRegionalAlarmsAndRules) validateSetLatencyAlarmParameters(val awscloudwatch.IAlarm) error {
	return nil
}

