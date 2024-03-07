package githubcollected

import (
	"github.com/Legit-Labs/legitify/internal/clients/github/types"
	"github.com/Legit-Labs/legitify/internal/common/namespace"
	"github.com/Legit-Labs/legitify/internal/common/permissions"
)

type Enterprise struct {
	MembersCanChangeRepositoryVisibilitySetting   string `json:"members_can_change_repository_visibility"`
	RepositoriesForkingPolicy                     string `json:"repositories_forking_policy"`
	ExternalCollaboratorsInvitePolicy             string `json:"external_collaborators_invite_policy"`
	TwoFactorRequiredSetting                      string `json:"two_factor_required_setting"`
	SamlEnabled                                   bool   `json:"saml_enabled"`
	EnterpriseName                                string `json:"name"`
	Url                                           string `json:"url"`
	Id                                            int64  `json:"id"`
	UserRole                                      string
	MembersCanCreatePublicRepositoriesSetting     bool                               `json:"members_can_create_public_repositories"`
	DefaultRepositoryPermissionSetting            string                             `json:"default_repository_permission_settings"`
	MembersCanDeleteRepositoriesSetting           string                             `json:"member_can_delete_repository"`
	NotificationDeliveryRestrictionEnabledSetting string                             `json:"notification_delivery_restriction_enabled"`
	CodeAndSecurityPolicySettings                 *types.AnalysisAndSecurityPolicies `json:"code_analysis_and_security_policies"`
}

func NewEnterprise(membersCanChangeRepositoryVisibilitySetting string, name string, Url string, Id int64, isAdmin bool, repositoriesForkingPolicy string,
	externalCollaboratorsInvitePolicy string, membersCanCreatePublicRepositoriesSetting bool, twoFactorRequiredSetting string, defaultRepositoryPermissionSetting string, membersCanDeleteRepositoriesSetting string, notificationDeliveryRestrictionEnabledSetting string, samlEnabled bool, codeAndSecurityPolicySettings *types.AnalysisAndSecurityPolicies) Enterprise {
	UserRole := permissions.EnterpriseNonAdminRole
	if isAdmin {
		UserRole = permissions.EnterpriseAdminRole
	}
	return Enterprise{
		MembersCanChangeRepositoryVisibilitySetting: membersCanChangeRepositoryVisibilitySetting,
		RepositoriesForkingPolicy:                   repositoriesForkingPolicy,
		TwoFactorRequiredSetting:                    twoFactorRequiredSetting,
		ExternalCollaboratorsInvitePolicy:           externalCollaboratorsInvitePolicy,
		EnterpriseName:                              name,
		SamlEnabled:                                 samlEnabled,
		Url:                                         Url,
		Id:                                          Id,
		UserRole:                                    UserRole,
		MembersCanCreatePublicRepositoriesSetting:     membersCanCreatePublicRepositoriesSetting,
		DefaultRepositoryPermissionSetting:            defaultRepositoryPermissionSetting,
		MembersCanDeleteRepositoriesSetting:           membersCanDeleteRepositoriesSetting,
		NotificationDeliveryRestrictionEnabledSetting: notificationDeliveryRestrictionEnabledSetting,
		CodeAndSecurityPolicySettings:                 codeAndSecurityPolicySettings,
	}
}

func (o Enterprise) ViolationEntityType() string {
	return namespace.Enterprise
}

func (o Enterprise) CanonicalLink() string {
	return o.Url
}

func (o Enterprise) Name() string {
	return o.EnterpriseName
}

func (o Enterprise) ID() int64 {
	return o.Id
}
