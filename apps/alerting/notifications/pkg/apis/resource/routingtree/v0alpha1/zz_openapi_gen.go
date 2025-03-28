//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by grafana-app-sdk. DO NOT EDIT.

package v0alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Matcher":             schema_apis_resource_routingtree_v0alpha1_Matcher(ref),
		"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.OperatorState":       schema_apis_resource_routingtree_v0alpha1_OperatorState(ref),
		"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Route":               schema_apis_resource_routingtree_v0alpha1_Route(ref),
		"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.RouteDefaults":       schema_apis_resource_routingtree_v0alpha1_RouteDefaults(ref),
		"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.RoutingTree":         schema_apis_resource_routingtree_v0alpha1_RoutingTree(ref),
		"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.RoutingTreeList":     schema_apis_resource_routingtree_v0alpha1_RoutingTreeList(ref),
		"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Spec":                schema_apis_resource_routingtree_v0alpha1_Spec(ref),
		"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Status":              schema_apis_resource_routingtree_v0alpha1_Status(ref),
		"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.StatusOperatorState": schema_apis_resource_routingtree_v0alpha1_StatusOperatorState(ref),
	}
}

func schema_apis_resource_routingtree_v0alpha1_Matcher(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Matcher defines model for Matcher.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"label": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"value": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"label", "type", "value"},
			},
		},
	}
}

func schema_apis_resource_routingtree_v0alpha1_OperatorState(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OperatorState defines model for OperatorState.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"descriptiveState": {
						SchemaProps: spec.SchemaProps{
							Description: "descriptiveState is an optional more descriptive state field which has no requirements on format",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"details": {
						SchemaProps: spec.SchemaProps{
							Description: "details contains any extra information that is operator-specific",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"object"},
										Format: "",
									},
								},
							},
						},
					},
					"lastEvaluation": {
						SchemaProps: spec.SchemaProps{
							Description: "lastEvaluation is the ResourceVersion last evaluated",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "state describes the state of the lastEvaluation. It is limited to three possible states for machine evaluation.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"lastEvaluation", "state"},
			},
		},
	}
}

func schema_apis_resource_routingtree_v0alpha1_Route(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Route defines model for Route.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"continue": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"group_by": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"group_interval": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"group_wait": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"matchers": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Matcher"),
									},
								},
							},
						},
					},
					"mute_time_intervals": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"receiver": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"repeat_interval": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"routes": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Route"),
									},
								},
							},
						},
					},
				},
				Required: []string{"continue"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Matcher", "github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Route"},
	}
}

func schema_apis_resource_routingtree_v0alpha1_RouteDefaults(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RouteDefaults defines model for RouteDefaults.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"group_by": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"group_interval": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"group_wait": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"receiver": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"repeat_interval": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"receiver"},
			},
		},
	}
}

func schema_apis_resource_routingtree_v0alpha1_RoutingTree(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Spec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Status"),
						},
					},
				},
				Required: []string{"metadata", "spec", "status"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Spec", "github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Status", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_apis_resource_routingtree_v0alpha1_RoutingTreeList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.RoutingTree"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.RoutingTree", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_apis_resource_routingtree_v0alpha1_Spec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Spec defines model for Spec.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"defaults": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.RouteDefaults"),
						},
					},
					"routes": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Route"),
									},
								},
							},
						},
					},
				},
				Required: []string{"defaults", "routes"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.Route", "github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.RouteDefaults"},
	}
}

func schema_apis_resource_routingtree_v0alpha1_Status(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Status defines model for Status.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"additionalFields": {
						SchemaProps: spec.SchemaProps{
							Description: "additionalFields is reserved for future use",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"object"},
										Format: "",
									},
								},
							},
						},
					},
					"operatorStates": {
						SchemaProps: spec.SchemaProps{
							Description: "operatorStates is a map of operator ID to operator state evaluations. Any operator which consumes this kind SHOULD add its state evaluation information to this field.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.StatusOperatorState"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/alerting/notifications/pkg/apis/resource/routingtree/v0alpha1.StatusOperatorState"},
	}
}

func schema_apis_resource_routingtree_v0alpha1_StatusOperatorState(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StatusOperatorState defines model for status.#OperatorState.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"descriptiveState": {
						SchemaProps: spec.SchemaProps{
							Description: "descriptiveState is an optional more descriptive state field which has no requirements on format",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"details": {
						SchemaProps: spec.SchemaProps{
							Description: "details contains any extra information that is operator-specific",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"object"},
										Format: "",
									},
								},
							},
						},
					},
					"lastEvaluation": {
						SchemaProps: spec.SchemaProps{
							Description: "lastEvaluation is the ResourceVersion last evaluated",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "state describes the state of the lastEvaluation. It is limited to three possible states for machine evaluation.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"lastEvaluation", "state"},
			},
		},
	}
}
