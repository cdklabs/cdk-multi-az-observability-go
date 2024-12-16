package multi-az-observability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-multi-az-observability-go/multi-az-observability/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-multi-az-observability-go/multi-az-observability/internal"
)

// A construct that allows you to map AZ names to ids and back.
// Experimental.
type AvailabilityZoneMapper interface {
	constructs.Construct
	IAvailabilityZoneMapper
	// The function that does the mapping.
	// Experimental.
	Function() awslambda.IFunction
	// Experimental.
	SetFunction(val awslambda.IFunction)
	// The log group for the function's logs.
	// Experimental.
	LogGroup() awslogs.ILogGroup
	// Experimental.
	SetLogGroup(val awslogs.ILogGroup)
	// The custom resource that can be referenced to use Fn::GetAtt functions on to retrieve availability zone names and ids.
	// Experimental.
	Mapper() awscdk.CustomResource
	// Experimental.
	SetMapper(val awscdk.CustomResource)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a reference that can be cast to a string array with all of the Availability Zone Ids.
	// Experimental.
	AllAvailabilityZoneIdsAsArray() awscdk.Reference
	// Returns a comma delimited list of Availability Zone Ids for the supplied Availability Zone names.
	//
	// You can use this string with Fn.Select(x, Fn.Split(",", azs)) to
	// get a specific Availability Zone Id.
	// Experimental.
	AllAvailabilityZoneIdsAsCommaDelimitedList() *string
	// Gets all of the Availability Zone names in this Region as a comma delimited list.
	//
	// You can use this string with Fn.Select(x, Fn.Split(",", azs)) to
	// get a specific Availability Zone Name.
	// Experimental.
	AllAvailabilityZoneNamesAsCommaDelimitedList() *string
	// Gets the Availability Zone Id for the given Availability Zone Name in this account.
	// Experimental.
	AvailabilityZoneId(availabilityZoneName *string) *string
	// Given a letter like "f" or "a", returns the Availability Zone Id for that Availability Zone name in this account.
	// Experimental.
	AvailabilityZoneIdFromAvailabilityZoneLetter(letter *string) *string
	// Returns an array for Availability Zone Ids for the supplied Availability Zone names, they are returned in the same order the names were provided.
	// Experimental.
	AvailabilityZoneIdsAsArray(availabilityZoneNames *[]*string) *[]*string
	// Returns a comma delimited list of Availability Zone Ids for the supplied Availability Zone names.
	//
	// You can use this string with Fn.Select(x, Fn.Split(",", azs)) to
	// get a specific Availability Zone Id.
	// Experimental.
	AvailabilityZoneIdsAsCommaDelimitedList(availabilityZoneNames *[]*string) *string
	// Gets the Availability Zone Name for the given Availability Zone Id in this account.
	// Experimental.
	AvailabilityZoneName(availabilityZoneId *string) *string
	// Gets the prefix for the region used with Availability Zone Ids, for example in us-east-1, this returns "use1".
	// Experimental.
	RegionPrefixForAvailabilityZoneIds() *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AvailabilityZoneMapper
type jsiiProxy_AvailabilityZoneMapper struct {
	internal.Type__constructsConstruct
	jsiiProxy_IAvailabilityZoneMapper
}

func (j *jsiiProxy_AvailabilityZoneMapper) Function() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AvailabilityZoneMapper) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AvailabilityZoneMapper) Mapper() awscdk.CustomResource {
	var returns awscdk.CustomResource
	_jsii_.Get(
		j,
		"mapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AvailabilityZoneMapper) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewAvailabilityZoneMapper(scope constructs.Construct, id *string, props *AvailabilityZoneMapperProps) AvailabilityZoneMapper {
	_init_.Initialize()

	if err := validateNewAvailabilityZoneMapperParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AvailabilityZoneMapper{}

	_jsii_.Create(
		"@cdklabs/multi-az-observability.AvailabilityZoneMapper",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAvailabilityZoneMapper_Override(a AvailabilityZoneMapper, scope constructs.Construct, id *string, props *AvailabilityZoneMapperProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/multi-az-observability.AvailabilityZoneMapper",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AvailabilityZoneMapper)SetFunction(val awslambda.IFunction) {
	if err := j.validateSetFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"function",
		val,
	)
}

func (j *jsiiProxy_AvailabilityZoneMapper)SetLogGroup(val awslogs.ILogGroup) {
	if err := j.validateSetLogGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroup",
		val,
	)
}

func (j *jsiiProxy_AvailabilityZoneMapper)SetMapper(val awscdk.CustomResource) {
	if err := j.validateSetMapperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapper",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AvailabilityZoneMapper_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAvailabilityZoneMapper_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/multi-az-observability.AvailabilityZoneMapper",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AvailabilityZoneMapper) AllAvailabilityZoneIdsAsArray() awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		a,
		"allAvailabilityZoneIdsAsArray",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AvailabilityZoneMapper) AllAvailabilityZoneIdsAsCommaDelimitedList() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"allAvailabilityZoneIdsAsCommaDelimitedList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AvailabilityZoneMapper) AllAvailabilityZoneNamesAsCommaDelimitedList() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"allAvailabilityZoneNamesAsCommaDelimitedList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AvailabilityZoneMapper) AvailabilityZoneId(availabilityZoneName *string) *string {
	if err := a.validateAvailabilityZoneIdParameters(availabilityZoneName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"availabilityZoneId",
		[]interface{}{availabilityZoneName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AvailabilityZoneMapper) AvailabilityZoneIdFromAvailabilityZoneLetter(letter *string) *string {
	if err := a.validateAvailabilityZoneIdFromAvailabilityZoneLetterParameters(letter); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"availabilityZoneIdFromAvailabilityZoneLetter",
		[]interface{}{letter},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AvailabilityZoneMapper) AvailabilityZoneIdsAsArray(availabilityZoneNames *[]*string) *[]*string {
	if err := a.validateAvailabilityZoneIdsAsArrayParameters(availabilityZoneNames); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"availabilityZoneIdsAsArray",
		[]interface{}{availabilityZoneNames},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AvailabilityZoneMapper) AvailabilityZoneIdsAsCommaDelimitedList(availabilityZoneNames *[]*string) *string {
	if err := a.validateAvailabilityZoneIdsAsCommaDelimitedListParameters(availabilityZoneNames); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"availabilityZoneIdsAsCommaDelimitedList",
		[]interface{}{availabilityZoneNames},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AvailabilityZoneMapper) AvailabilityZoneName(availabilityZoneId *string) *string {
	if err := a.validateAvailabilityZoneNameParameters(availabilityZoneId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"availabilityZoneName",
		[]interface{}{availabilityZoneId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AvailabilityZoneMapper) RegionPrefixForAvailabilityZoneIds() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"regionPrefixForAvailabilityZoneIds",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AvailabilityZoneMapper) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

