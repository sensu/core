package v2

// automatically generated file, do not edit!

import (
	"testing"
)

func TestResolveAPIKey(t *testing.T) {
	var value interface{} = new(APIKey)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("APIKey"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("APIKey")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"APIKey" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveAdhocRequest(t *testing.T) {
	var value interface{} = new(AdhocRequest)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("AdhocRequest"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("AdhocRequest")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"AdhocRequest" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveAny(t *testing.T) {
	var value interface{} = new(Any)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Any"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Any")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Any" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveAsset(t *testing.T) {
	var value interface{} = new(Asset)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Asset"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Asset")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Asset" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveAssetBuild(t *testing.T) {
	var value interface{} = new(AssetBuild)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("AssetBuild"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("AssetBuild")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"AssetBuild" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveAssetList(t *testing.T) {
	var value interface{} = new(AssetList)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("AssetList"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("AssetList")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"AssetList" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveAuthProviderClaims(t *testing.T) {
	var value interface{} = new(AuthProviderClaims)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("AuthProviderClaims"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("AuthProviderClaims")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"AuthProviderClaims" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveCheck(t *testing.T) {
	var value interface{} = new(Check)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Check"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Check")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Check" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveCheckConfig(t *testing.T) {
	var value interface{} = new(CheckConfig)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("CheckConfig"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("CheckConfig")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"CheckConfig" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveCheckHistory(t *testing.T) {
	var value interface{} = new(CheckHistory)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("CheckHistory"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("CheckHistory")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"CheckHistory" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveCheckRequest(t *testing.T) {
	var value interface{} = new(CheckRequest)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("CheckRequest"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("CheckRequest")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"CheckRequest" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveClaims(t *testing.T) {
	var value interface{} = new(Claims)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Claims"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Claims")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Claims" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveClusterHealth(t *testing.T) {
	var value interface{} = new(ClusterHealth)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("ClusterHealth"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("ClusterHealth")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"ClusterHealth" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveClusterRole(t *testing.T) {
	var value interface{} = new(ClusterRole)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("ClusterRole"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("ClusterRole")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"ClusterRole" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveClusterRoleBinding(t *testing.T) {
	var value interface{} = new(ClusterRoleBinding)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("ClusterRoleBinding"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("ClusterRoleBinding")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"ClusterRoleBinding" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveDeregistration(t *testing.T) {
	var value interface{} = new(Deregistration)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Deregistration"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Deregistration")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Deregistration" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveEntity(t *testing.T) {
	var value interface{} = new(Entity)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Entity"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Entity")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Entity" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveEvent(t *testing.T) {
	var value interface{} = new(Event)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Event"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Event")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Event" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveEventFilter(t *testing.T) {
	var value interface{} = new(EventFilter)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("EventFilter"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("EventFilter")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"EventFilter" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveExtension(t *testing.T) {
	var value interface{} = new(Extension)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Extension"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Extension")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Extension" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveHandler(t *testing.T) {
	var value interface{} = new(Handler)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Handler"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Handler")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Handler" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveHandlerSocket(t *testing.T) {
	var value interface{} = new(HandlerSocket)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("HandlerSocket"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("HandlerSocket")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"HandlerSocket" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveHealthResponse(t *testing.T) {
	var value interface{} = new(HealthResponse)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("HealthResponse"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("HealthResponse")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"HealthResponse" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveHook(t *testing.T) {
	var value interface{} = new(Hook)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Hook"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Hook")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Hook" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveHookConfig(t *testing.T) {
	var value interface{} = new(HookConfig)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("HookConfig"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("HookConfig")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"HookConfig" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveHookList(t *testing.T) {
	var value interface{} = new(HookList)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("HookList"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("HookList")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"HookList" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveKeepaliveRecord(t *testing.T) {
	var value interface{} = new(KeepaliveRecord)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("KeepaliveRecord"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("KeepaliveRecord")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"KeepaliveRecord" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveMetricPoint(t *testing.T) {
	var value interface{} = new(MetricPoint)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("MetricPoint"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("MetricPoint")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"MetricPoint" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveMetricTag(t *testing.T) {
	var value interface{} = new(MetricTag)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("MetricTag"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("MetricTag")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"MetricTag" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveMetricThreshold(t *testing.T) {
	var value interface{} = new(MetricThreshold)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("MetricThreshold"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("MetricThreshold")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"MetricThreshold" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveMetricThresholdRule(t *testing.T) {
	var value interface{} = new(MetricThresholdRule)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("MetricThresholdRule"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("MetricThresholdRule")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"MetricThresholdRule" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveMetricThresholdTag(t *testing.T) {
	var value interface{} = new(MetricThresholdTag)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("MetricThresholdTag"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("MetricThresholdTag")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"MetricThresholdTag" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveMetrics(t *testing.T) {
	var value interface{} = new(Metrics)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Metrics"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Metrics")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Metrics" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveMutator(t *testing.T) {
	var value interface{} = new(Mutator)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Mutator"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Mutator")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Mutator" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveNamespace(t *testing.T) {
	var value interface{} = new(Namespace)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Namespace"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Namespace")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Namespace" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveNetwork(t *testing.T) {
	var value interface{} = new(Network)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Network"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Network")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Network" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveNetworkInterface(t *testing.T) {
	var value interface{} = new(NetworkInterface)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("NetworkInterface"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("NetworkInterface")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"NetworkInterface" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveObjectMeta(t *testing.T) {
	var value interface{} = new(ObjectMeta)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("ObjectMeta"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("ObjectMeta")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"ObjectMeta" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolvePipeline(t *testing.T) {
	var value interface{} = new(Pipeline)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Pipeline"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Pipeline")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Pipeline" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolvePipelineWorkflow(t *testing.T) {
	var value interface{} = new(PipelineWorkflow)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("PipelineWorkflow"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("PipelineWorkflow")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"PipelineWorkflow" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolvePostgresHealth(t *testing.T) {
	var value interface{} = new(PostgresHealth)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("PostgresHealth"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("PostgresHealth")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"PostgresHealth" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveProcess(t *testing.T) {
	var value interface{} = new(Process)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Process"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Process")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Process" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveProxyRequests(t *testing.T) {
	var value interface{} = new(ProxyRequests)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("ProxyRequests"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("ProxyRequests")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"ProxyRequests" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveResourceReference(t *testing.T) {
	var value interface{} = new(ResourceReference)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("ResourceReference"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("ResourceReference")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"ResourceReference" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveRole(t *testing.T) {
	var value interface{} = new(Role)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Role"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Role")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Role" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveRoleBinding(t *testing.T) {
	var value interface{} = new(RoleBinding)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("RoleBinding"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("RoleBinding")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"RoleBinding" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveRoleRef(t *testing.T) {
	var value interface{} = new(RoleRef)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("RoleRef"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("RoleRef")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"RoleRef" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveRule(t *testing.T) {
	var value interface{} = new(Rule)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Rule"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Rule")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Rule" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveSecret(t *testing.T) {
	var value interface{} = new(Secret)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Secret"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Secret")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Secret" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveSilenced(t *testing.T) {
	var value interface{} = new(Silenced)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Silenced"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Silenced")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Silenced" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveSubject(t *testing.T) {
	var value interface{} = new(Subject)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Subject"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Subject")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Subject" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveSystem(t *testing.T) {
	var value interface{} = new(System)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("System"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("System")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"System" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveTLSOptions(t *testing.T) {
	var value interface{} = new(TLSOptions)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("TLSOptions"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("TLSOptions")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"TLSOptions" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveTessenConfig(t *testing.T) {
	var value interface{} = new(TessenConfig)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("TessenConfig"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("TessenConfig")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"TessenConfig" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveTimeWindowDays(t *testing.T) {
	var value interface{} = new(TimeWindowDays)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("TimeWindowDays"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("TimeWindowDays")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"TimeWindowDays" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveTimeWindowRepeated(t *testing.T) {
	var value interface{} = new(TimeWindowRepeated)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("TimeWindowRepeated"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("TimeWindowRepeated")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"TimeWindowRepeated" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveTimeWindowTimeRange(t *testing.T) {
	var value interface{} = new(TimeWindowTimeRange)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("TimeWindowTimeRange"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("TimeWindowTimeRange")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"TimeWindowTimeRange" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveTimeWindowWhen(t *testing.T) {
	var value interface{} = new(TimeWindowWhen)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("TimeWindowWhen"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("TimeWindowWhen")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"TimeWindowWhen" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveTokens(t *testing.T) {
	var value interface{} = new(Tokens)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Tokens"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Tokens")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Tokens" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveTypeMeta(t *testing.T) {
	var value interface{} = new(TypeMeta)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("TypeMeta"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("TypeMeta")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"TypeMeta" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveUser(t *testing.T) {
	var value interface{} = new(User)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("User"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("User")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"User" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveVersion(t *testing.T) {
	var value interface{} = new(Version)
	if _, ok := value.(Resource); ok {
		if _, err := resolveResource("Version"); err != nil {
			t.Fatal(err)
		}
		return
	}
	_, err := resolveResource("Version")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Version" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveNotExists(t *testing.T) {
	_, err := resolveResource("!#$@$%@#$")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
}
