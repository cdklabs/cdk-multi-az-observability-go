//go:build no_runtime_type_checking

package multi-az-observability

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules) validateSetAvailabilityAlarmParameters(val awscloudwatch.IAlarm) error {
	return nil
}

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules) validateSetAvailabilityZoneIsOutlierForFaultsParameters(val awscloudwatch.IAlarm) error {
	return nil
}

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules) validateSetAvailabilityZoneIsOutlierForLatencyParameters(val awscloudwatch.IAlarm) error {
	return nil
}

func (j *jsiiProxy_IBaseOperationZonalAlarmsAndRules) validateSetLatencyAlarmParameters(val awscloudwatch.IAlarm) error {
	return nil
}

