package nodeprobe

import "github.com/timob/javabind"

type ToolsNodeProbe struct {
	*javabind.Callable
}

// public NodeProbe(String host, int port, String username, String password) throws IOException
func NewToolsNodeProbe(host string, port int, username string, password string) (*ToolsNodeProbe, error) {
	conv_host := javabind.NewGoToJavaString()
	conv_username := javabind.NewGoToJavaString()
	conv_password := javabind.NewGoToJavaString()
	if err := conv_host.Convert(host); err != nil {
		panic(err)
	}
	if err := conv_username.Convert(username); err != nil {
		panic(err)
	}
	if err := conv_password.Convert(password); err != nil {
		panic(err)
	}

	obj, err := javabind.Env.NewInstanceStr("org.apache.cassandra.tools.NodeProbe", javabind.CastObject(conv_host.Value(), "java.lang.String"), port, javabind.CastObject(conv_username.Value(), "java.lang.String"), javabind.CastObject(conv_password.Value(), "java.lang.String"))
	if err != nil {
		return nil, err
	}
	conv_host.CleanUp()
	conv_username.CleanUp()
	conv_password.CleanUp()
	return &ToolsNodeProbe{&javabind.Callable{obj, javabind.Env}}, nil
}

// public NodeProbe(String host, int port) throws IOException
func NewToolsNodeProbe2(host string, port int) (*ToolsNodeProbe, error) {
	conv_host := javabind.NewGoToJavaString()
	if err := conv_host.Convert(host); err != nil {
		panic(err)
	}

	obj, err := javabind.Env.NewInstanceStr("org.apache.cassandra.tools.NodeProbe", javabind.CastObject(conv_host.Value(), "java.lang.String"), port)
	if err != nil {
		return nil, err
	}
	conv_host.CleanUp()
	return &ToolsNodeProbe{&javabind.Callable{obj, javabind.Env}}, nil
}

// public NodeProbe(String host) throws IOException
func NewToolsNodeProbe3(host string) (*ToolsNodeProbe, error) {
	conv_host := javabind.NewGoToJavaString()
	if err := conv_host.Convert(host); err != nil {
		panic(err)
	}

	obj, err := javabind.Env.NewInstanceStr("org.apache.cassandra.tools.NodeProbe", javabind.CastObject(conv_host.Value(), "java.lang.String"))
	if err != nil {
		return nil, err
	}
	conv_host.CleanUp()
	return &ToolsNodeProbe{&javabind.Callable{obj, javabind.Env}}, nil
}

// public void close() throws IOException
func (jbobject *ToolsNodeProbe) Close() error {
	err := jbobject.CallVoid("close")
	if err != nil {
		return err
	}
	return nil
}

// public void forceKeyspaceCleanup(String keyspaceName, String... columnFamilies) throws IOException, ExecutionException, InterruptedException
func (jbobject *ToolsNodeProbe) ForceKeyspaceCleanup(keyspaceName string, columnFamilies ...string) error {
	conv_keyspaceName := javabind.NewGoToJavaString()
	conv_columnFamilies := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_keyspaceName.Convert(keyspaceName); err != nil {
		panic(err)
	}
	if err := conv_columnFamilies.Convert(columnFamilies); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("forceKeyspaceCleanup", javabind.CastObject(conv_keyspaceName.Value(), "java.lang.String"), javabind.ObjectArray(conv_columnFamilies.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_keyspaceName.CleanUp()
	conv_columnFamilies.CleanUp()
	return nil
}

// public void scrub(boolean disableSnapshot, boolean skipCorrupted, String keyspaceName, String... columnFamilies) throws IOException, ExecutionException, InterruptedException
func (jbobject *ToolsNodeProbe) Scrub(disableSnapshot bool, skipCorrupted bool, keyspaceName string, columnFamilies ...string) error {
	conv_keyspaceName := javabind.NewGoToJavaString()
	conv_columnFamilies := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_keyspaceName.Convert(keyspaceName); err != nil {
		panic(err)
	}
	if err := conv_columnFamilies.Convert(columnFamilies); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("scrub", disableSnapshot, skipCorrupted, javabind.CastObject(conv_keyspaceName.Value(), "java.lang.String"), javabind.ObjectArray(conv_columnFamilies.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_keyspaceName.CleanUp()
	conv_columnFamilies.CleanUp()
	return nil
}

// public void upgradeSSTables(String keyspaceName, boolean excludeCurrentVersion, String... columnFamilies) throws IOException, ExecutionException, InterruptedException
func (jbobject *ToolsNodeProbe) UpgradeSSTables(keyspaceName string, excludeCurrentVersion bool, columnFamilies ...string) error {
	conv_keyspaceName := javabind.NewGoToJavaString()
	conv_columnFamilies := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_keyspaceName.Convert(keyspaceName); err != nil {
		panic(err)
	}
	if err := conv_columnFamilies.Convert(columnFamilies); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("upgradeSSTables", javabind.CastObject(conv_keyspaceName.Value(), "java.lang.String"), excludeCurrentVersion, javabind.ObjectArray(conv_columnFamilies.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_keyspaceName.CleanUp()
	conv_columnFamilies.CleanUp()
	return nil
}

// public void forceKeyspaceCompaction(String keyspaceName, String... columnFamilies) throws IOException, ExecutionException, InterruptedException
func (jbobject *ToolsNodeProbe) ForceKeyspaceCompaction(keyspaceName string, columnFamilies ...string) error {
	conv_keyspaceName := javabind.NewGoToJavaString()
	conv_columnFamilies := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_keyspaceName.Convert(keyspaceName); err != nil {
		panic(err)
	}
	if err := conv_columnFamilies.Convert(columnFamilies); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("forceKeyspaceCompaction", javabind.CastObject(conv_keyspaceName.Value(), "java.lang.String"), javabind.ObjectArray(conv_columnFamilies.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_keyspaceName.CleanUp()
	conv_columnFamilies.CleanUp()
	return nil
}

// public void forceKeyspaceFlush(String keyspaceName, String... columnFamilies) throws IOException, ExecutionException, InterruptedException
func (jbobject *ToolsNodeProbe) ForceKeyspaceFlush(keyspaceName string, columnFamilies ...string) error {
	conv_keyspaceName := javabind.NewGoToJavaString()
	conv_columnFamilies := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_keyspaceName.Convert(keyspaceName); err != nil {
		panic(err)
	}
	if err := conv_columnFamilies.Convert(columnFamilies); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("forceKeyspaceFlush", javabind.CastObject(conv_keyspaceName.Value(), "java.lang.String"), javabind.ObjectArray(conv_columnFamilies.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_keyspaceName.CleanUp()
	conv_columnFamilies.CleanUp()
	return nil
}

// public void forceKeyspaceRepair(String keyspaceName, boolean isSequential, boolean isLocal, String... columnFamilies) throws IOException
func (jbobject *ToolsNodeProbe) ForceKeyspaceRepair(keyspaceName string, isSequential bool, isLocal bool, columnFamilies ...string) error {
	conv_keyspaceName := javabind.NewGoToJavaString()
	conv_columnFamilies := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_keyspaceName.Convert(keyspaceName); err != nil {
		panic(err)
	}
	if err := conv_columnFamilies.Convert(columnFamilies); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("forceKeyspaceRepair", javabind.CastObject(conv_keyspaceName.Value(), "java.lang.String"), isSequential, isLocal, javabind.ObjectArray(conv_columnFamilies.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_keyspaceName.CleanUp()
	conv_columnFamilies.CleanUp()
	return nil
}

// public void forceKeyspaceRepairPrimaryRange(String keyspaceName, boolean isSequential, boolean isLocal, String... columnFamilies) throws IOException
func (jbobject *ToolsNodeProbe) ForceKeyspaceRepairPrimaryRange(keyspaceName string, isSequential bool, isLocal bool, columnFamilies ...string) error {
	conv_keyspaceName := javabind.NewGoToJavaString()
	conv_columnFamilies := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_keyspaceName.Convert(keyspaceName); err != nil {
		panic(err)
	}
	if err := conv_columnFamilies.Convert(columnFamilies); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("forceKeyspaceRepairPrimaryRange", javabind.CastObject(conv_keyspaceName.Value(), "java.lang.String"), isSequential, isLocal, javabind.ObjectArray(conv_columnFamilies.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_keyspaceName.CleanUp()
	conv_columnFamilies.CleanUp()
	return nil
}

// public void forceKeyspaceRepairRange(String beginToken, String endToken, String keyspaceName, boolean isSequential, boolean isLocal, String... columnFamilies) throws IOException
func (jbobject *ToolsNodeProbe) ForceKeyspaceRepairRange(beginToken string, endToken string, keyspaceName string, isSequential bool, isLocal bool, columnFamilies ...string) error {
	conv_beginToken := javabind.NewGoToJavaString()
	conv_endToken := javabind.NewGoToJavaString()
	conv_keyspaceName := javabind.NewGoToJavaString()
	conv_columnFamilies := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_beginToken.Convert(beginToken); err != nil {
		panic(err)
	}
	if err := conv_endToken.Convert(endToken); err != nil {
		panic(err)
	}
	if err := conv_keyspaceName.Convert(keyspaceName); err != nil {
		panic(err)
	}
	if err := conv_columnFamilies.Convert(columnFamilies); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("forceKeyspaceRepairRange", javabind.CastObject(conv_beginToken.Value(), "java.lang.String"), javabind.CastObject(conv_endToken.Value(), "java.lang.String"), javabind.CastObject(conv_keyspaceName.Value(), "java.lang.String"), isSequential, isLocal, javabind.ObjectArray(conv_columnFamilies.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_beginToken.CleanUp()
	conv_endToken.CleanUp()
	conv_keyspaceName.CleanUp()
	conv_columnFamilies.CleanUp()
	return nil
}

// public void invalidateKeyCache()
func (jbobject *ToolsNodeProbe) InvalidateKeyCache()  {
	err := jbobject.CallVoid("invalidateKeyCache")
	if err != nil {
		panic(err)
	}

}

// public void invalidateRowCache()
func (jbobject *ToolsNodeProbe) InvalidateRowCache()  {
	err := jbobject.CallVoid("invalidateRowCache")
	if err != nil {
		panic(err)
	}

}

// public void drain() throws IOException, InterruptedException, ExecutionException
func (jbobject *ToolsNodeProbe) Drain() error {
	err := jbobject.CallVoid("drain")
	if err != nil {
		return err
	}
	return nil
}

// public Map<String, String> getTokenToEndpointMap()
func (jbobject *ToolsNodeProbe) GetTokenToEndpointMap() map[string]string {
	jret, err := jbobject.CallObj("getTokenToEndpointMap", "java.util.Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoString())
	dst := new(map[string]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public List<String> getLiveNodes()
func (jbobject *ToolsNodeProbe) GetLiveNodes() []string {
	jret, err := jbobject.CallObj("getLiveNodes", "java.util.List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public List<String> getJoiningNodes()
func (jbobject *ToolsNodeProbe) GetJoiningNodes() []string {
	jret, err := jbobject.CallObj("getJoiningNodes", "java.util.List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public List<String> getLeavingNodes()
func (jbobject *ToolsNodeProbe) GetLeavingNodes() []string {
	jret, err := jbobject.CallObj("getLeavingNodes", "java.util.List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public List<String> getMovingNodes()
func (jbobject *ToolsNodeProbe) GetMovingNodes() []string {
	jret, err := jbobject.CallObj("getMovingNodes", "java.util.List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public List<String> getUnreachableNodes()
func (jbobject *ToolsNodeProbe) GetUnreachableNodes() []string {
	jret, err := jbobject.CallObj("getUnreachableNodes", "java.util.List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public Map<String, String> getLoadMap()
func (jbobject *ToolsNodeProbe) GetLoadMap() map[string]string {
	jret, err := jbobject.CallObj("getLoadMap", "java.util.Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoString())
	dst := new(map[string]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public Map<InetAddress, Float> getOwnership()
func (jbobject *ToolsNodeProbe) GetOwnership() map[string]float32 {
	jret, err := jbobject.CallObj("getOwnership", "java.util.Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoInetAddress(), javabind.NewJavaToGoFloat())
	dst := new(map[string]float32)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public Map<InetAddress, Float> effectiveOwnership(String keyspace) throws IllegalStateException
func (jbobject *ToolsNodeProbe) EffectiveOwnership(keyspace string) (map[string]float32, error) {
	conv_keyspace := javabind.NewGoToJavaString()
	if err := conv_keyspace.Convert(keyspace); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallObj("effectiveOwnership", "java.util.Map", javabind.CastObject(conv_keyspace.Value(), "java.lang.String"))
	if err != nil {
		var zero map[string]float32
		return zero, err
	}
	conv_keyspace.CleanUp()
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoInetAddress(), javabind.NewJavaToGoFloat())
	dst := new(map[string]float32)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst, nil
}

// public CacheServiceMBean getCacheServiceMBean()
func (jbobject *ToolsNodeProbe) GetCacheServiceMBean() *ServiceCacheServiceMBean {
	jret, err := jbobject.CallObj("getCacheServiceMBean", "org.apache.cassandra.service.CacheServiceMBean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return &ServiceCacheServiceMBean{dst}
}

// public Iterator<Map.Entry<String, ColumnFamilyStoreMBean>> getColumnFamilyStoreMBeanProxies()
func (jbobject *ToolsNodeProbe) GetColumnFamilyStoreMBeanProxies() []struct{Key string; Value *DbColumnFamilyStoreMBean} {
	jret, err := jbobject.CallObj("getColumnFamilyStoreMBeanProxies", "java.util.Iterator")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoIterator(javabind.NewJavaToGoMap_Entry(javabind.NewJavaToGoString(), javabind.NewJavaToGoCallable()))
	dst := new([]struct{Key string; Value *DbColumnFamilyStoreMBean})
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public CompactionManagerMBean getCompactionManagerProxy()
func (jbobject *ToolsNodeProbe) GetCompactionManagerProxy() *DbCompactionCompactionManagerMBean {
	jret, err := jbobject.CallObj("getCompactionManagerProxy", "org.apache.cassandra.db.compaction.CompactionManagerMBean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return &DbCompactionCompactionManagerMBean{dst}
}

// public List<String> getTokens()
func (jbobject *ToolsNodeProbe) GetTokens() []string {
	jret, err := jbobject.CallObj("getTokens", "java.util.List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public List<String> getTokens(String endpoint)
func (jbobject *ToolsNodeProbe) GetTokens2(endpoint string) []string {
	conv_endpoint := javabind.NewGoToJavaString()
	if err := conv_endpoint.Convert(endpoint); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallObj("getTokens", "java.util.List", javabind.CastObject(conv_endpoint.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_endpoint.CleanUp()
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public String getLocalHostId()
func (jbobject *ToolsNodeProbe) GetLocalHostId() string {
	jret, err := jbobject.CallObj("getLocalHostId", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public Map<String, String> getHostIdMap()
func (jbobject *ToolsNodeProbe) GetHostIdMap() map[string]string {
	jret, err := jbobject.CallObj("getHostIdMap", "java.util.Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoString())
	dst := new(map[string]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public String getLoadString()
func (jbobject *ToolsNodeProbe) GetLoadString() string {
	jret, err := jbobject.CallObj("getLoadString", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public String getReleaseVersion()
func (jbobject *ToolsNodeProbe) GetReleaseVersion() string {
	jret, err := jbobject.CallObj("getReleaseVersion", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getCurrentGenerationNumber()
func (jbobject *ToolsNodeProbe) GetCurrentGenerationNumber() int {
	jret, err := jbobject.CallInt("getCurrentGenerationNumber")
	if err != nil {
		panic(err)
	}
	return jret
}

// public long getUptime()
func (jbobject *ToolsNodeProbe) GetUptime() int64 {
	jret, err := jbobject.CallLong("getUptime")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void takeSnapshot(String snapshotName, String columnFamily, String... keyspaces) throws IOException
func (jbobject *ToolsNodeProbe) TakeSnapshot(snapshotName string, columnFamily string, keyspaces ...string) error {
	conv_snapshotName := javabind.NewGoToJavaString()
	conv_columnFamily := javabind.NewGoToJavaString()
	conv_keyspaces := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_snapshotName.Convert(snapshotName); err != nil {
		panic(err)
	}
	if err := conv_columnFamily.Convert(columnFamily); err != nil {
		panic(err)
	}
	if err := conv_keyspaces.Convert(keyspaces); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("takeSnapshot", javabind.CastObject(conv_snapshotName.Value(), "java.lang.String"), javabind.CastObject(conv_columnFamily.Value(), "java.lang.String"), javabind.ObjectArray(conv_keyspaces.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_snapshotName.CleanUp()
	conv_columnFamily.CleanUp()
	conv_keyspaces.CleanUp()
	return nil
}

// public void clearSnapshot(String tag, String... keyspaces) throws IOException
func (jbobject *ToolsNodeProbe) ClearSnapshot(tag string, keyspaces ...string) error {
	conv_tag := javabind.NewGoToJavaString()
	conv_keyspaces := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_tag.Convert(tag); err != nil {
		panic(err)
	}
	if err := conv_keyspaces.Convert(keyspaces); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("clearSnapshot", javabind.CastObject(conv_tag.Value(), "java.lang.String"), javabind.ObjectArray(conv_keyspaces.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_tag.CleanUp()
	conv_keyspaces.CleanUp()
	return nil
}

// public boolean isJoined()
func (jbobject *ToolsNodeProbe) IsJoined() bool {
	jret, err := jbobject.CallBoolean("isJoined")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void joinRing() throws IOException
func (jbobject *ToolsNodeProbe) JoinRing() error {
	err := jbobject.CallVoid("joinRing")
	if err != nil {
		return err
	}
	return nil
}

// public void decommission() throws InterruptedException
func (jbobject *ToolsNodeProbe) Decommission() error {
	err := jbobject.CallVoid("decommission")
	if err != nil {
		return err
	}
	return nil
}

// public void move(String newToken) throws IOException
func (jbobject *ToolsNodeProbe) Move(newToken string) error {
	conv_newToken := javabind.NewGoToJavaString()
	if err := conv_newToken.Convert(newToken); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("move", javabind.CastObject(conv_newToken.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_newToken.CleanUp()
	return nil
}

// public void takeTokens(String[] tokens) throws IOException
func (jbobject *ToolsNodeProbe) TakeTokens(tokens []string) error {
	conv_tokens := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_tokens.Convert(tokens); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("takeTokens", javabind.ObjectArray(conv_tokens.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_tokens.CleanUp()
	return nil
}

// public void removeNode(String token)
func (jbobject *ToolsNodeProbe) RemoveNode(token string)  {
	conv_token := javabind.NewGoToJavaString()
	if err := conv_token.Convert(token); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("removeNode", javabind.CastObject(conv_token.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_token.CleanUp()

}

// public String getRemovalStatus()
func (jbobject *ToolsNodeProbe) GetRemovalStatus() string {
	jret, err := jbobject.CallObj("getRemovalStatus", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void forceRemoveCompletion()
func (jbobject *ToolsNodeProbe) ForceRemoveCompletion()  {
	err := jbobject.CallVoid("forceRemoveCompletion")
	if err != nil {
		panic(err)
	}

}

// public Iterator<Map.Entry<String, JMXEnabledThreadPoolExecutorMBean>> getThreadPoolMBeanProxies()
func (jbobject *ToolsNodeProbe) GetThreadPoolMBeanProxies() []struct{Key string; Value *ConcurrentJMXEnabledThreadPoolExecutorMBean} {
	jret, err := jbobject.CallObj("getThreadPoolMBeanProxies", "java.util.Iterator")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoIterator(javabind.NewJavaToGoMap_Entry(javabind.NewJavaToGoString(), javabind.NewJavaToGoCallable()))
	dst := new([]struct{Key string; Value *ConcurrentJMXEnabledThreadPoolExecutorMBean})
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setCompactionThreshold(String ks, String cf, int minimumCompactionThreshold, int maximumCompactionThreshold)
func (jbobject *ToolsNodeProbe) SetCompactionThreshold(ks string, cf string, minimumCompactionThreshold int, maximumCompactionThreshold int)  {
	conv_ks := javabind.NewGoToJavaString()
	conv_cf := javabind.NewGoToJavaString()
	if err := conv_ks.Convert(ks); err != nil {
		panic(err)
	}
	if err := conv_cf.Convert(cf); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("setCompactionThreshold", javabind.CastObject(conv_ks.Value(), "java.lang.String"), javabind.CastObject(conv_cf.Value(), "java.lang.String"), minimumCompactionThreshold, maximumCompactionThreshold)
	if err != nil {
		panic(err)
	}
	conv_ks.CleanUp()
	conv_cf.CleanUp()

}

// public void disableAutoCompaction(java.lang.String, java.lang.String...) throws java.io.IOException
func (jbobject *ToolsNodeProbe) DisableAutoCompaction(a string, b ...string) error {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("disableAutoCompaction", javabind.CastObject(conv_a.Value(), "java.lang.String"), javabind.ObjectArray(conv_b.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return nil
}

// public void enableAutoCompaction(java.lang.String, java.lang.String...) throws java.io.IOException
func (jbobject *ToolsNodeProbe) EnableAutoCompaction(a string, b ...string) error {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("enableAutoCompaction", javabind.CastObject(conv_a.Value(), "java.lang.String"), javabind.ObjectArray(conv_b.Value(), "java.lang.String"))
	if err != nil {
		return err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return nil
}

// public void setIncrementalBackupsEnabled(boolean enabled)
func (jbobject *ToolsNodeProbe) SetIncrementalBackupsEnabled(enabled bool)  {
	err := jbobject.CallVoid("setIncrementalBackupsEnabled", enabled)
	if err != nil {
		panic(err)
	}

}

// public void setCacheCapacities(int keyCacheCapacity, int rowCacheCapacity)
func (jbobject *ToolsNodeProbe) SetCacheCapacities(keyCacheCapacity int, rowCacheCapacity int)  {
	err := jbobject.CallVoid("setCacheCapacities", keyCacheCapacity, rowCacheCapacity)
	if err != nil {
		panic(err)
	}

}

// public void setCacheKeysToSave(int keyCacheKeysToSave, int rowCacheKeysToSave)
func (jbobject *ToolsNodeProbe) SetCacheKeysToSave(keyCacheKeysToSave int, rowCacheKeysToSave int)  {
	err := jbobject.CallVoid("setCacheKeysToSave", keyCacheKeysToSave, rowCacheKeysToSave)
	if err != nil {
		panic(err)
	}

}

// public List<InetAddress> getEndpoints(String keyspace, String cf, String key)
func (jbobject *ToolsNodeProbe) GetEndpoints(keyspace string, cf string, key string) []string {
	conv_keyspace := javabind.NewGoToJavaString()
	conv_cf := javabind.NewGoToJavaString()
	conv_key := javabind.NewGoToJavaString()
	if err := conv_keyspace.Convert(keyspace); err != nil {
		panic(err)
	}
	if err := conv_cf.Convert(cf); err != nil {
		panic(err)
	}
	if err := conv_key.Convert(key); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallObj("getEndpoints", "java.util.List", javabind.CastObject(conv_keyspace.Value(), "java.lang.String"), javabind.CastObject(conv_cf.Value(), "java.lang.String"), javabind.CastObject(conv_key.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_keyspace.CleanUp()
	conv_cf.CleanUp()
	conv_key.CleanUp()
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoInetAddress())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public List<String> getSSTables(String keyspace, String cf, String key)
func (jbobject *ToolsNodeProbe) GetSSTables(keyspace string, cf string, key string) []string {
	conv_keyspace := javabind.NewGoToJavaString()
	conv_cf := javabind.NewGoToJavaString()
	conv_key := javabind.NewGoToJavaString()
	if err := conv_keyspace.Convert(keyspace); err != nil {
		panic(err)
	}
	if err := conv_cf.Convert(cf); err != nil {
		panic(err)
	}
	if err := conv_key.Convert(key); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallObj("getSSTables", "java.util.List", javabind.CastObject(conv_keyspace.Value(), "java.lang.String"), javabind.CastObject(conv_cf.Value(), "java.lang.String"), javabind.CastObject(conv_key.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_keyspace.CleanUp()
	conv_cf.CleanUp()
	conv_key.CleanUp()
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public Set<StreamState> getStreamStatus()
func (jbobject *ToolsNodeProbe) GetStreamStatus() []*StreamingStreamState {
	jret, err := jbobject.CallObj("getStreamStatus", "java.util.Set")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoSet(javabind.NewJavaToGoCallable())
	dst := new([]*StreamingStreamState)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public String getOperationMode()
func (jbobject *ToolsNodeProbe) GetOperationMode() string {
	jret, err := jbobject.CallObj("getOperationMode", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void truncate(String keyspaceName, String cfName)
func (jbobject *ToolsNodeProbe) Truncate(keyspaceName string, cfName string)  {
	conv_keyspaceName := javabind.NewGoToJavaString()
	conv_cfName := javabind.NewGoToJavaString()
	if err := conv_keyspaceName.Convert(keyspaceName); err != nil {
		panic(err)
	}
	if err := conv_cfName.Convert(cfName); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("truncate", javabind.CastObject(conv_keyspaceName.Value(), "java.lang.String"), javabind.CastObject(conv_cfName.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_keyspaceName.CleanUp()
	conv_cfName.CleanUp()

}

// public EndpointSnitchInfoMBean getEndpointSnitchInfoProxy()
func (jbobject *ToolsNodeProbe) GetEndpointSnitchInfoProxy() *LocatorEndpointSnitchInfoMBean {
	jret, err := jbobject.CallObj("getEndpointSnitchInfoProxy", "org.apache.cassandra.locator.EndpointSnitchInfoMBean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return &LocatorEndpointSnitchInfoMBean{dst}
}

// public ColumnFamilyStoreMBean getCfsProxy(String ks, String cf)
func (jbobject *ToolsNodeProbe) GetCfsProxy(ks string, cf string) *DbColumnFamilyStoreMBean {
	conv_ks := javabind.NewGoToJavaString()
	conv_cf := javabind.NewGoToJavaString()
	if err := conv_ks.Convert(ks); err != nil {
		panic(err)
	}
	if err := conv_cf.Convert(cf); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallObj("getCfsProxy", "org.apache.cassandra.db.ColumnFamilyStoreMBean", javabind.CastObject(conv_ks.Value(), "java.lang.String"), javabind.CastObject(conv_cf.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_ks.CleanUp()
	conv_cf.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return &DbColumnFamilyStoreMBean{dst}
}

// public StorageProxyMBean getSpProxy()
func (jbobject *ToolsNodeProbe) GetSpProxy() *ServiceStorageProxyMBean {
	jret, err := jbobject.CallObj("getSpProxy", "org.apache.cassandra.service.StorageProxyMBean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return &ServiceStorageProxyMBean{dst}
}

// public String getEndpoint()
func (jbobject *ToolsNodeProbe) GetEndpoint() string {
	jret, err := jbobject.CallObj("getEndpoint", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public String getDataCenter()
func (jbobject *ToolsNodeProbe) GetDataCenter() string {
	jret, err := jbobject.CallObj("getDataCenter", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public String getRack()
func (jbobject *ToolsNodeProbe) GetRack() string {
	jret, err := jbobject.CallObj("getRack", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public List<String> getKeyspaces()
func (jbobject *ToolsNodeProbe) GetKeyspaces() []string {
	jret, err := jbobject.CallObj("getKeyspaces", "java.util.List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public String getClusterName()
func (jbobject *ToolsNodeProbe) GetClusterName() string {
	jret, err := jbobject.CallObj("getClusterName", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public String getPartitioner()
func (jbobject *ToolsNodeProbe) GetPartitioner() string {
	jret, err := jbobject.CallObj("getPartitioner", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void disableHintedHandoff()
func (jbobject *ToolsNodeProbe) DisableHintedHandoff()  {
	err := jbobject.CallVoid("disableHintedHandoff")
	if err != nil {
		panic(err)
	}

}

// public void enableHintedHandoff()
func (jbobject *ToolsNodeProbe) EnableHintedHandoff()  {
	err := jbobject.CallVoid("enableHintedHandoff")
	if err != nil {
		panic(err)
	}

}

// public void enableHintedHandoff(String dcNames)
func (jbobject *ToolsNodeProbe) EnableHintedHandoff2(dcNames string)  {
	conv_dcNames := javabind.NewGoToJavaString()
	if err := conv_dcNames.Convert(dcNames); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("enableHintedHandoff", javabind.CastObject(conv_dcNames.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_dcNames.CleanUp()

}

// public void pauseHintsDelivery()
func (jbobject *ToolsNodeProbe) PauseHintsDelivery()  {
	err := jbobject.CallVoid("pauseHintsDelivery")
	if err != nil {
		panic(err)
	}

}

// public void resumeHintsDelivery()
func (jbobject *ToolsNodeProbe) ResumeHintsDelivery()  {
	err := jbobject.CallVoid("resumeHintsDelivery")
	if err != nil {
		panic(err)
	}

}

// public void truncateHints(final String host)
func (jbobject *ToolsNodeProbe) TruncateHints(host string)  {
	conv_host := javabind.NewGoToJavaString()
	if err := conv_host.Convert(host); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("truncateHints", javabind.CastObject(conv_host.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_host.CleanUp()

}

// public void truncateHints()
func (jbobject *ToolsNodeProbe) TruncateHints2()  {
	err := jbobject.CallVoid("truncateHints")
	if err != nil {
		panic(err)
	}

}

// public void stopNativeTransport()
func (jbobject *ToolsNodeProbe) StopNativeTransport()  {
	err := jbobject.CallVoid("stopNativeTransport")
	if err != nil {
		panic(err)
	}

}

// public void startNativeTransport()
func (jbobject *ToolsNodeProbe) StartNativeTransport()  {
	err := jbobject.CallVoid("startNativeTransport")
	if err != nil {
		panic(err)
	}

}

// public boolean isNativeTransportRunning()
func (jbobject *ToolsNodeProbe) IsNativeTransportRunning() bool {
	jret, err := jbobject.CallBoolean("isNativeTransportRunning")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void stopGossiping()
func (jbobject *ToolsNodeProbe) StopGossiping()  {
	err := jbobject.CallVoid("stopGossiping")
	if err != nil {
		panic(err)
	}

}

// public void startGossiping()
func (jbobject *ToolsNodeProbe) StartGossiping()  {
	err := jbobject.CallVoid("startGossiping")
	if err != nil {
		panic(err)
	}

}

// public void stopThriftServer()
func (jbobject *ToolsNodeProbe) StopThriftServer()  {
	err := jbobject.CallVoid("stopThriftServer")
	if err != nil {
		panic(err)
	}

}

// public void startThriftServer()
func (jbobject *ToolsNodeProbe) StartThriftServer()  {
	err := jbobject.CallVoid("startThriftServer")
	if err != nil {
		panic(err)
	}

}

// public boolean isThriftServerRunning()
func (jbobject *ToolsNodeProbe) IsThriftServerRunning() bool {
	jret, err := jbobject.CallBoolean("isThriftServerRunning")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void stopCassandraDaemon()
func (jbobject *ToolsNodeProbe) StopCassandraDaemon()  {
	err := jbobject.CallVoid("stopCassandraDaemon")
	if err != nil {
		panic(err)
	}

}

// public boolean isInitialized()
func (jbobject *ToolsNodeProbe) IsInitialized() bool {
	jret, err := jbobject.CallBoolean("isInitialized")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setCompactionThroughput(int value)
func (jbobject *ToolsNodeProbe) SetCompactionThroughput(value int)  {
	err := jbobject.CallVoid("setCompactionThroughput", value)
	if err != nil {
		panic(err)
	}

}

// public int getCompactionThroughput()
func (jbobject *ToolsNodeProbe) GetCompactionThroughput() int {
	jret, err := jbobject.CallInt("getCompactionThroughput")
	if err != nil {
		panic(err)
	}
	return jret
}

// public int getStreamThroughput()
func (jbobject *ToolsNodeProbe) GetStreamThroughput() int {
	jret, err := jbobject.CallInt("getStreamThroughput")
	if err != nil {
		panic(err)
	}
	return jret
}

// public int getExceptionCount()
func (jbobject *ToolsNodeProbe) GetExceptionCount() int {
	jret, err := jbobject.CallInt("getExceptionCount")
	if err != nil {
		panic(err)
	}
	return jret
}

// public Map<String, Integer> getDroppedMessages()
func (jbobject *ToolsNodeProbe) GetDroppedMessages() map[string]int {
	jret, err := jbobject.CallObj("getDroppedMessages", "java.util.Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoInteger())
	dst := new(map[string]int)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void loadNewSSTables(String ksName, String cfName)
func (jbobject *ToolsNodeProbe) LoadNewSSTables(ksName string, cfName string)  {
	conv_ksName := javabind.NewGoToJavaString()
	conv_cfName := javabind.NewGoToJavaString()
	if err := conv_ksName.Convert(ksName); err != nil {
		panic(err)
	}
	if err := conv_cfName.Convert(cfName); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("loadNewSSTables", javabind.CastObject(conv_ksName.Value(), "java.lang.String"), javabind.CastObject(conv_cfName.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_ksName.CleanUp()
	conv_cfName.CleanUp()

}

// public void rebuildIndex(String ksName, String cfName, String... idxNames)
func (jbobject *ToolsNodeProbe) RebuildIndex(ksName string, cfName string, idxNames ...string)  {
	conv_ksName := javabind.NewGoToJavaString()
	conv_cfName := javabind.NewGoToJavaString()
	conv_idxNames := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString())
	if err := conv_ksName.Convert(ksName); err != nil {
		panic(err)
	}
	if err := conv_cfName.Convert(cfName); err != nil {
		panic(err)
	}
	if err := conv_idxNames.Convert(idxNames); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("rebuildIndex", javabind.CastObject(conv_ksName.Value(), "java.lang.String"), javabind.CastObject(conv_cfName.Value(), "java.lang.String"), javabind.ObjectArray(conv_idxNames.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_ksName.CleanUp()
	conv_cfName.CleanUp()
	conv_idxNames.CleanUp()

}

// public String getGossipInfo()
func (jbobject *ToolsNodeProbe) GetGossipInfo() string {
	jret, err := jbobject.CallObj("getGossipInfo", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void stop(String string)
func (jbobject *ToolsNodeProbe) Stop(string string)  {
	conv_string := javabind.NewGoToJavaString()
	if err := conv_string.Convert(string); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("stop", javabind.CastObject(conv_string.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_string.CleanUp()

}

// public void setStreamThroughput(int value)
func (jbobject *ToolsNodeProbe) SetStreamThroughput(value int)  {
	err := jbobject.CallVoid("setStreamThroughput", value)
	if err != nil {
		panic(err)
	}

}

// public void setTraceProbability(double value)
func (jbobject *ToolsNodeProbe) SetTraceProbability(value float64)  {
	err := jbobject.CallVoid("setTraceProbability", value)
	if err != nil {
		panic(err)
	}

}

// public String getSchemaVersion()
func (jbobject *ToolsNodeProbe) GetSchemaVersion() string {
	jret, err := jbobject.CallObj("getSchemaVersion", "java.lang.String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public List<String> describeRing(String keyspaceName) throws IOException
func (jbobject *ToolsNodeProbe) DescribeRing(keyspaceName string) ([]string, error) {
	conv_keyspaceName := javabind.NewGoToJavaString()
	if err := conv_keyspaceName.Convert(keyspaceName); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallObj("describeRing", "java.util.List", javabind.CastObject(conv_keyspaceName.Value(), "java.lang.String"))
	if err != nil {
		var zero []string
		return zero, err
	}
	conv_keyspaceName.CleanUp()
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst, nil
}

// public void rebuild(String sourceDc)
func (jbobject *ToolsNodeProbe) Rebuild(sourceDc string)  {
	conv_sourceDc := javabind.NewGoToJavaString()
	if err := conv_sourceDc.Convert(sourceDc); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("rebuild", javabind.CastObject(conv_sourceDc.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_sourceDc.CleanUp()

}

// public List<String> sampleKeyRange()
func (jbobject *ToolsNodeProbe) SampleKeyRange() []string {
	jret, err := jbobject.CallObj("sampleKeyRange", "java.util.List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void resetLocalSchema() throws IOException
func (jbobject *ToolsNodeProbe) ResetLocalSchema() error {
	err := jbobject.CallVoid("resetLocalSchema")
	if err != nil {
		return err
	}
	return nil
}

// public boolean isFailed()
func (jbobject *ToolsNodeProbe) IsFailed() bool {
	jret, err := jbobject.CallBoolean("isFailed")
	if err != nil {
		panic(err)
	}
	return jret
}

// public long getReadRepairAttempted()
func (jbobject *ToolsNodeProbe) GetReadRepairAttempted() int64 {
	jret, err := jbobject.CallLong("getReadRepairAttempted")
	if err != nil {
		panic(err)
	}
	return jret
}

// public long getReadRepairRepairedBlocking()
func (jbobject *ToolsNodeProbe) GetReadRepairRepairedBlocking() int64 {
	jret, err := jbobject.CallLong("getReadRepairRepairedBlocking")
	if err != nil {
		panic(err)
	}
	return jret
}

// public long getReadRepairRepairedBackground()
func (jbobject *ToolsNodeProbe) GetReadRepairRepairedBackground() int64 {
	jret, err := jbobject.CallLong("getReadRepairRepairedBackground")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void reloadTriggers()
func (jbobject *ToolsNodeProbe) ReloadTriggers()  {
	err := jbobject.CallVoid("reloadTriggers")
	if err != nil {
		panic(err)
	}

}

// public void setLoggingLevel(String classQualifier, String level)
func (jbobject *ToolsNodeProbe) SetLoggingLevel(classQualifier string, level string)  {
	conv_classQualifier := javabind.NewGoToJavaString()
	conv_level := javabind.NewGoToJavaString()
	if err := conv_classQualifier.Convert(classQualifier); err != nil {
		panic(err)
	}
	if err := conv_level.Convert(level); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("setLoggingLevel", javabind.CastObject(conv_classQualifier.Value(), "java.lang.String"), javabind.CastObject(conv_level.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_classQualifier.CleanUp()
	conv_level.CleanUp()

}

// public Map<String, String> getLoggingLevels()
func (jbobject *ToolsNodeProbe) GetLoggingLevels() map[string]string {
	jret, err := jbobject.CallObj("getLoggingLevels", "java.util.Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoString())
	dst := new(map[string]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

