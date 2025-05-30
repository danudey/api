// Copyright (c) 2019-2024 Tigera, Inc. All rights reserved.

package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupName is the group name use in this package
const GroupName = "projectcalico.org"

// SchemeGroupVersion is group version used to register these objects
var (
	SchemeGroupVersion         = schema.GroupVersion{Group: GroupName, Version: "v3"}
	SchemeGroupVersionInternal = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
)

var (
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
	AllKnownTypes      = []runtime.Object{
		&NetworkPolicy{},
		&NetworkPolicyList{},
		&GlobalNetworkPolicy{},
		&GlobalNetworkPolicyList{},
		&GlobalNetworkSet{},
		&GlobalNetworkSetList{},
		&HostEndpoint{},
		&HostEndpointList{},
		&IPPool{},
		&IPPoolList{},
		&IPReservation{},
		&IPReservationList{},
		&BGPConfiguration{},
		&BGPConfigurationList{},
		&BGPFilter{},
		&BGPFilterList{},
		&BGPPeer{},
		&BGPPeerList{},
		&Profile{},
		&ProfileList{},
		&FelixConfiguration{},
		&FelixConfigurationList{},
		&KubeControllersConfiguration{},
		&KubeControllersConfigurationList{},
		&ClusterInformation{},
		&ClusterInformationList{},
		&NetworkSet{},
		&NetworkSetList{},
		&CalicoNodeStatus{},
		&CalicoNodeStatusList{},
		&IPAMConfiguration{},
		&IPAMConfigurationList{},
		&BlockAffinity{},
		&BlockAffinityList{},
		&Tier{},
		&TierList{},

		// Enterprise-only types.
		&AlertException{},
		&AlertExceptionList{},
		&AuthenticationReview{},
		&AuthenticationReviewList{},
		&AuthorizationReview{},
		&AuthorizationReviewList{},
		&DeepPacketInspection{},
		&DeepPacketInspectionList{},
		&EgressGatewayPolicy{},
		&EgressGatewayPolicyList{},
		&GlobalAlert{},
		&GlobalAlertList{},
		&GlobalAlertTemplate{},
		&GlobalAlertTemplateList{},
		&GlobalReport{},
		&GlobalReportList{},
		&GlobalReportType{},
		&GlobalReportTypeList{},
		&GlobalThreatFeed{},
		&GlobalThreatFeedList{},
		&LicenseKey{},
		&LicenseKeyList{},
		&ManagedCluster{},
		&ManagedClusterList{},
		&PacketCapture{},
		&PacketCaptureList{},
		&PolicyRecommendationScope{},
		&PolicyRecommendationScopeList{},
		&RemoteClusterConfiguration{},
		&RemoteClusterConfigurationList{},
		&StagedGlobalNetworkPolicy{},
		&StagedGlobalNetworkPolicyList{},
		&StagedKubernetesNetworkPolicy{},
		&StagedKubernetesNetworkPolicyList{},
		&StagedNetworkPolicy{},
		&StagedNetworkPolicyList{},
		&UISettingsGroup{},
		&UISettingsGroupList{},
		&UISettings{},
		&UISettingsList{},
		&ExternalNetwork{},
		&ExternalNetworkList{},
		&SecurityEventWebhook{},
		&SecurityEventWebhookList{},
		&BFDConfiguration{},
		&BFDConfigurationList{},
	}
)

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes, addConversionFuncs)
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, AllKnownTypes...)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
