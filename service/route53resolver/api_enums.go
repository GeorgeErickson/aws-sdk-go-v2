// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

type IpAddressStatus string

// Enum values for IpAddressStatus
const (
	IpAddressStatusCreating               IpAddressStatus = "CREATING"
	IpAddressStatusFailedCreation         IpAddressStatus = "FAILED_CREATION"
	IpAddressStatusAttaching              IpAddressStatus = "ATTACHING"
	IpAddressStatusAttached               IpAddressStatus = "ATTACHED"
	IpAddressStatusRemapDetaching         IpAddressStatus = "REMAP_DETACHING"
	IpAddressStatusRemapAttaching         IpAddressStatus = "REMAP_ATTACHING"
	IpAddressStatusDetaching              IpAddressStatus = "DETACHING"
	IpAddressStatusFailedResourceGone     IpAddressStatus = "FAILED_RESOURCE_GONE"
	IpAddressStatusDeleting               IpAddressStatus = "DELETING"
	IpAddressStatusDeleteFailedFasExpired IpAddressStatus = "DELETE_FAILED_FAS_EXPIRED"
)

func (enum IpAddressStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IpAddressStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResolverEndpointDirection string

// Enum values for ResolverEndpointDirection
const (
	ResolverEndpointDirectionInbound  ResolverEndpointDirection = "INBOUND"
	ResolverEndpointDirectionOutbound ResolverEndpointDirection = "OUTBOUND"
)

func (enum ResolverEndpointDirection) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResolverEndpointDirection) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResolverEndpointStatus string

// Enum values for ResolverEndpointStatus
const (
	ResolverEndpointStatusCreating       ResolverEndpointStatus = "CREATING"
	ResolverEndpointStatusOperational    ResolverEndpointStatus = "OPERATIONAL"
	ResolverEndpointStatusUpdating       ResolverEndpointStatus = "UPDATING"
	ResolverEndpointStatusAutoRecovering ResolverEndpointStatus = "AUTO_RECOVERING"
	ResolverEndpointStatusActionNeeded   ResolverEndpointStatus = "ACTION_NEEDED"
	ResolverEndpointStatusDeleting       ResolverEndpointStatus = "DELETING"
)

func (enum ResolverEndpointStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResolverEndpointStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResolverRuleAssociationStatus string

// Enum values for ResolverRuleAssociationStatus
const (
	ResolverRuleAssociationStatusCreating   ResolverRuleAssociationStatus = "CREATING"
	ResolverRuleAssociationStatusComplete   ResolverRuleAssociationStatus = "COMPLETE"
	ResolverRuleAssociationStatusDeleting   ResolverRuleAssociationStatus = "DELETING"
	ResolverRuleAssociationStatusFailed     ResolverRuleAssociationStatus = "FAILED"
	ResolverRuleAssociationStatusOverridden ResolverRuleAssociationStatus = "OVERRIDDEN"
)

func (enum ResolverRuleAssociationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResolverRuleAssociationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResolverRuleStatus string

// Enum values for ResolverRuleStatus
const (
	ResolverRuleStatusComplete ResolverRuleStatus = "COMPLETE"
	ResolverRuleStatusDeleting ResolverRuleStatus = "DELETING"
	ResolverRuleStatusUpdating ResolverRuleStatus = "UPDATING"
	ResolverRuleStatusFailed   ResolverRuleStatus = "FAILED"
)

func (enum ResolverRuleStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResolverRuleStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RuleTypeOption string

// Enum values for RuleTypeOption
const (
	RuleTypeOptionForward   RuleTypeOption = "FORWARD"
	RuleTypeOptionSystem    RuleTypeOption = "SYSTEM"
	RuleTypeOptionRecursive RuleTypeOption = "RECURSIVE"
)

func (enum RuleTypeOption) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RuleTypeOption) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ShareStatus string

// Enum values for ShareStatus
const (
	ShareStatusNotShared    ShareStatus = "NOT_SHARED"
	ShareStatusSharedWithMe ShareStatus = "SHARED_WITH_ME"
	ShareStatusSharedByMe   ShareStatus = "SHARED_BY_ME"
)

func (enum ShareStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ShareStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
