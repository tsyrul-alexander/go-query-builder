package parameter

type ArrayParameter struct {
	Parameters []QueryParameter
}

func (p *ArrayParameter)GetValueSql() string {
	return ""
}

func CreateParameters(values ... QueryParameter) QueryParameter {
	return &ArrayParameter{Parameters:values}
}
