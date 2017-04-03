package spec

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/v1"
)

// ResourceList is a set of (resource name, quantity) pairs.
type ResourceList v1.ResourceList

// Plan defines how resources could be managed and distributed
type Plan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PlanSpec `json:"spec"`
}

// PlanList is a list of ServicePlans
type PlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Plan `json:"items"`
}

// PlanSpec holds specification parameters of an Plan
type PlanSpec struct {
	// Compute Resources required by containers.
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
	// Hard is the set of desired hard limits for each named resource.
	Hard  ResourceList   `json:"hard,omitempty"`
	Roles []PlatformRole `json:"roles,omitempty"`
}

const (
	// ResourceNamespace , number
	ResourceNamespace api.ResourceName = "namespaces"
)

// PlatformRole is the name identifying various roles in a PlatformRoleList.
type PlatformRole string

const (
	// RoleExecAllow cluster role name
	RoleExecAllow PlatformRole = "exec-allow"
	// RolePortForwardAllow cluster role name
	RolePortForwardAllow PlatformRole = "portforward-allow"
	// RoleAutoScaleAllow cluster role name
	RoleAutoScaleAllow PlatformRole = "autoscale-allow"
	// RoleAttachAllow cluster role name
	RoleAttachAllow PlatformRole = "attach-allow"
	// RoleAddonManagement cluster role name
	RoleAddonManagement PlatformRole = "addon-management"
)

// ServicePlanPhase is the current lifecycle phase of the Service Plan.
type ServicePlanPhase string

const (
	// ServicePlanActive means the ServicePlan is available for use in the system
	ServicePlanActive ServicePlanPhase = "Active"
	// ServicePlanPending means the ServicePlan isn't associate with any global ServicePlan
	ServicePlanPending ServicePlanPhase = "Pending"
	// ServicePlanNotFound means the reference plan wasn't found
	ServicePlanNotFound ServicePlanPhase = "NotFound"
	// ServicePlanDisabled means the ServicePlan is disabled and cannot be associated with resources
	ServicePlanDisabled ServicePlanPhase = "Disabled"
)

// Addon defines integration with external resources
type Addon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddonSpec `json:"spec"`
}

// AddonList is a list of Addons.
type AddonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Addon `json:"items"`
}

// AddonSpec holds specification parameters of an addon
type AddonSpec struct {
	Type      string      `json:"type"`
	BaseImage string      `json:"baseImage"`
	Version   string      `json:"version"`
	Replicas  int32       `json:"replicas"`
	Port      int32       `json:"port"`
	Env       []v1.EnvVar `json:"env"`
	// More info: http://releases.k8s.io/HEAD/docs/user-guide/containers.md#containers-and-commands
	Args []string `json:"args,omitempty"`
}

// Release refers to compiled slug file versions
type Release struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReleaseSpec `json:"spec"`
}

// SourceType refers to the source of the build
type SourceType string

const (
	// GitHubSource means the build came from a webhook
	GitHubSource SourceType = "github"
	// GitLocalSource means the build came from the git local server
	GitLocalSource SourceType = "local"
)

// ReleaseSpec holds specification parameters of a release
type ReleaseSpec struct {
	// The URL of the git remote server to download the git revision tarball
	GitRemote     string     `json:"gitRemote"`
	GitRevision   string     `json:"gitRevision"`
	GitRepository string     `json:"gitRepository"`
	BuildRevision string     `json:"buildRevision"`
	AutoDeploy    bool       `json:"autoDeploy"`
	ExpireAfter   int32      `json:"expireAfter"`
	DeployName    string     `json:"deployName"`
	Build         bool       `json:"build"`
	AuthToken     string     `json:"authToken"` // expirable token
	Source        SourceType `json:"sourceType"`
}

// ReleaseList is a list of Release
type ReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Release `json:"items"`
}

// User identifies an user on the platform
type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Customer     string `json:"customer"`
	Organization string `json:"org"`
	// Groups are a set of strings which associate users with as set of commonly grouped users.
	// A group name is unique in the cluster and it's formed by it's namespace, customer or the organization name:
	// [org] - Matches all the namespaces of the broker
	// [customer]-[org] - Matches all namespaces from the customer broker
	// [name]-[customer]-[org] - Matches a specific namespace
	// http://kubernetes.io/docs/admin/authentication/
	Groups []string `json:"groups"`
}
