package model

import (
	oidc "istio.io/api/oidc/v1alpha1"
)

func GetOidcPolicy(store IstioConfigStore, service *Service, port *Port) *oidc.Policy {
	config := store.OidcPolicyByDestination(service, port)
	if config != nil {
		policy := config.Spec.(*oidc.Policy)
		log.Infof("sso - " +
			" returning policy: %+v", policy)
		return policy
	}

	return nil
}
