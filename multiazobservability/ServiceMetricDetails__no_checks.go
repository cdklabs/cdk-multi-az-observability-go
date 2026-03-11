//go:build no_runtime_type_checking

package multiazobservability

// Building without runtime type checking enabled, so all the below just return nil

func validateNewServiceMetricDetailsParameters(props *ServiceMetricDetailsProps) error {
	return nil
}

