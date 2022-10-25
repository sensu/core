package v2

// automatically generated file, do not edit!

import (
	"testing"

	"github.com/sensu/sensu-api-tools"
)

func TestResolveAPIKey(t *testing.T) {
	var value interface{} = new(APIKey)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "APIKey"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*APIKey); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "APIKey")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveAdhocRequest(t *testing.T) {
	var value interface{} = new(AdhocRequest)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "AdhocRequest"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*AdhocRequest); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "AdhocRequest")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveAny(t *testing.T) {
	var value interface{} = new(Any)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Any"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Any); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Any")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveAsset(t *testing.T) {
	var value interface{} = new(Asset)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Asset"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Asset); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Asset")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveAssetBuild(t *testing.T) {
	var value interface{} = new(AssetBuild)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "AssetBuild"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*AssetBuild); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "AssetBuild")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveAssetList(t *testing.T) {
	var value interface{} = new(AssetList)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "AssetList"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*AssetList); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "AssetList")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveAuthProviderClaims(t *testing.T) {
	var value interface{} = new(AuthProviderClaims)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "AuthProviderClaims"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*AuthProviderClaims); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "AuthProviderClaims")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveCheck(t *testing.T) {
	var value interface{} = new(Check)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Check"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Check); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Check")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveCheckConfig(t *testing.T) {
	var value interface{} = new(CheckConfig)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "CheckConfig"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*CheckConfig); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "CheckConfig")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveCheckHistory(t *testing.T) {
	var value interface{} = new(CheckHistory)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "CheckHistory"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*CheckHistory); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "CheckHistory")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveCheckRequest(t *testing.T) {
	var value interface{} = new(CheckRequest)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "CheckRequest"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*CheckRequest); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "CheckRequest")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveClaims(t *testing.T) {
	var value interface{} = new(Claims)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Claims"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Claims); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Claims")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveClusterHealth(t *testing.T) {
	var value interface{} = new(ClusterHealth)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "ClusterHealth"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*ClusterHealth); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "ClusterHealth")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveClusterRole(t *testing.T) {
	var value interface{} = new(ClusterRole)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "ClusterRole"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*ClusterRole); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "ClusterRole")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveClusterRoleBinding(t *testing.T) {
	var value interface{} = new(ClusterRoleBinding)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "ClusterRoleBinding"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*ClusterRoleBinding); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "ClusterRoleBinding")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveDeregistration(t *testing.T) {
	var value interface{} = new(Deregistration)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Deregistration"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Deregistration); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Deregistration")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveEntity(t *testing.T) {
	var value interface{} = new(Entity)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Entity"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Entity); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Entity")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveEvent(t *testing.T) {
	var value interface{} = new(Event)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Event"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Event); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Event")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveEventFilter(t *testing.T) {
	var value interface{} = new(EventFilter)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "EventFilter"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*EventFilter); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "EventFilter")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveExtension(t *testing.T) {
	var value interface{} = new(Extension)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Extension"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Extension); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Extension")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveHandler(t *testing.T) {
	var value interface{} = new(Handler)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Handler"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Handler); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Handler")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveHandlerSocket(t *testing.T) {
	var value interface{} = new(HandlerSocket)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "HandlerSocket"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*HandlerSocket); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "HandlerSocket")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveHealthResponse(t *testing.T) {
	var value interface{} = new(HealthResponse)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "HealthResponse"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*HealthResponse); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "HealthResponse")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveHook(t *testing.T) {
	var value interface{} = new(Hook)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Hook"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Hook); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Hook")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveHookConfig(t *testing.T) {
	var value interface{} = new(HookConfig)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "HookConfig"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*HookConfig); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "HookConfig")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveHookList(t *testing.T) {
	var value interface{} = new(HookList)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "HookList"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*HookList); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "HookList")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveKeepaliveRecord(t *testing.T) {
	var value interface{} = new(KeepaliveRecord)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "KeepaliveRecord"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*KeepaliveRecord); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "KeepaliveRecord")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveMetricPoint(t *testing.T) {
	var value interface{} = new(MetricPoint)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "MetricPoint"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*MetricPoint); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "MetricPoint")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveMetricTag(t *testing.T) {
	var value interface{} = new(MetricTag)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "MetricTag"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*MetricTag); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "MetricTag")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveMetricThreshold(t *testing.T) {
	var value interface{} = new(MetricThreshold)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "MetricThreshold"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*MetricThreshold); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "MetricThreshold")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveMetricThresholdRule(t *testing.T) {
	var value interface{} = new(MetricThresholdRule)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "MetricThresholdRule"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*MetricThresholdRule); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "MetricThresholdRule")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveMetricThresholdTag(t *testing.T) {
	var value interface{} = new(MetricThresholdTag)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "MetricThresholdTag"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*MetricThresholdTag); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "MetricThresholdTag")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveMetrics(t *testing.T) {
	var value interface{} = new(Metrics)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Metrics"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Metrics); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Metrics")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveMutator(t *testing.T) {
	var value interface{} = new(Mutator)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Mutator"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Mutator); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Mutator")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveNamespace(t *testing.T) {
	var value interface{} = new(Namespace)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Namespace"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Namespace); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Namespace")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveNetwork(t *testing.T) {
	var value interface{} = new(Network)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Network"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Network); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Network")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveNetworkInterface(t *testing.T) {
	var value interface{} = new(NetworkInterface)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "NetworkInterface"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*NetworkInterface); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "NetworkInterface")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveObjectMeta(t *testing.T) {
	var value interface{} = new(ObjectMeta)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "ObjectMeta"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*ObjectMeta); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "ObjectMeta")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolvePipeline(t *testing.T) {
	var value interface{} = new(Pipeline)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Pipeline"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Pipeline); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Pipeline")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolvePipelineWorkflow(t *testing.T) {
	var value interface{} = new(PipelineWorkflow)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "PipelineWorkflow"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*PipelineWorkflow); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "PipelineWorkflow")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolvePostgresHealth(t *testing.T) {
	var value interface{} = new(PostgresHealth)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "PostgresHealth"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*PostgresHealth); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "PostgresHealth")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveProcess(t *testing.T) {
	var value interface{} = new(Process)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Process"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Process); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Process")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveProxyRequests(t *testing.T) {
	var value interface{} = new(ProxyRequests)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "ProxyRequests"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*ProxyRequests); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "ProxyRequests")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveResourceReference(t *testing.T) {
	var value interface{} = new(ResourceReference)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "ResourceReference"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*ResourceReference); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "ResourceReference")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveRole(t *testing.T) {
	var value interface{} = new(Role)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Role"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Role); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Role")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveRoleBinding(t *testing.T) {
	var value interface{} = new(RoleBinding)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "RoleBinding"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*RoleBinding); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "RoleBinding")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveRoleRef(t *testing.T) {
	var value interface{} = new(RoleRef)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "RoleRef"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*RoleRef); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "RoleRef")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveRule(t *testing.T) {
	var value interface{} = new(Rule)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Rule"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Rule); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Rule")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveSecret(t *testing.T) {
	var value interface{} = new(Secret)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Secret"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Secret); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Secret")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveSilenced(t *testing.T) {
	var value interface{} = new(Silenced)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Silenced"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Silenced); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Silenced")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveSubject(t *testing.T) {
	var value interface{} = new(Subject)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Subject"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Subject); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Subject")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveSystem(t *testing.T) {
	var value interface{} = new(System)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "System"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*System); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "System")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveTLSOptions(t *testing.T) {
	var value interface{} = new(TLSOptions)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "TLSOptions"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*TLSOptions); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "TLSOptions")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveTessenConfig(t *testing.T) {
	var value interface{} = new(TessenConfig)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "TessenConfig"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*TessenConfig); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "TessenConfig")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveTimeWindowDays(t *testing.T) {
	var value interface{} = new(TimeWindowDays)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "TimeWindowDays"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*TimeWindowDays); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "TimeWindowDays")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveTimeWindowRepeated(t *testing.T) {
	var value interface{} = new(TimeWindowRepeated)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "TimeWindowRepeated"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*TimeWindowRepeated); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "TimeWindowRepeated")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveTimeWindowTimeRange(t *testing.T) {
	var value interface{} = new(TimeWindowTimeRange)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "TimeWindowTimeRange"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*TimeWindowTimeRange); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "TimeWindowTimeRange")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveTimeWindowWhen(t *testing.T) {
	var value interface{} = new(TimeWindowWhen)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "TimeWindowWhen"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*TimeWindowWhen); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "TimeWindowWhen")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveTokens(t *testing.T) {
	var value interface{} = new(Tokens)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Tokens"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Tokens); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Tokens")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveTypeMeta(t *testing.T) {
	var value interface{} = new(TypeMeta)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "TypeMeta"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*TypeMeta); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "TypeMeta")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveUser(t *testing.T) {
	var value interface{} = new(User)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "User"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*User); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "User")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func TestResolveVersion(t *testing.T) {
	var value interface{} = new(Version)
	if _, ok := value.(Resource); ok {
		if actual, err := apitools.Resolve("core/v2", "Version"); err != nil {
			t.Fatal(err)
		} else if _, ok := actual.(*Version); !ok {
			t.Fatal("expected to resolve to type ")
		}
		return
	}
	_, err := apitools.Resolve("core/v2", "Version")
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}
