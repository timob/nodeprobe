package nodeprobe

import "github.com/timob/javabind"

type DbColumnFamilyStoreMBean struct {
	*javabind.Callable
}

// public String getColumnFamilyName()
func (jbobject *DbColumnFamilyStoreMBean) GetColumnFamilyName() string {
	jret, err := jbobject.CallObj("getColumnFamilyName", "java.lang.String")
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

// public abstract long getMemtableDataSize()
func (jbobject *DbColumnFamilyStoreMBean) GetMemtableDataSize() int64 {
	jret, err := jbobject.CallLong("getMemtableDataSize")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getMemtableColumnsCount()
func (jbobject *DbColumnFamilyStoreMBean) GetMemtableColumnsCount() int64 {
	jret, err := jbobject.CallLong("getMemtableColumnsCount")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract int getMemtableSwitchCount()
func (jbobject *DbColumnFamilyStoreMBean) GetMemtableSwitchCount() int {
	jret, err := jbobject.CallInt("getMemtableSwitchCount")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getRecentSSTablesPerReadHistogram()
func (jbobject *DbColumnFamilyStoreMBean) GetRecentSSTablesPerReadHistogram() []int64 {
	jret, err := jbobject.CallLongArray("getRecentSSTablesPerReadHistogram")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getSSTablesPerReadHistogram()
func (jbobject *DbColumnFamilyStoreMBean) GetSSTablesPerReadHistogram() []int64 {
	jret, err := jbobject.CallLongArray("getSSTablesPerReadHistogram")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getReadCount()
func (jbobject *DbColumnFamilyStoreMBean) GetReadCount() int64 {
	jret, err := jbobject.CallLong("getReadCount")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getTotalReadLatencyMicros()
func (jbobject *DbColumnFamilyStoreMBean) GetTotalReadLatencyMicros() int64 {
	jret, err := jbobject.CallLong("getTotalReadLatencyMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getLifetimeReadLatencyHistogramMicros()
func (jbobject *DbColumnFamilyStoreMBean) GetLifetimeReadLatencyHistogramMicros() []int64 {
	jret, err := jbobject.CallLongArray("getLifetimeReadLatencyHistogramMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getRecentReadLatencyHistogramMicros()
func (jbobject *DbColumnFamilyStoreMBean) GetRecentReadLatencyHistogramMicros() []int64 {
	jret, err := jbobject.CallLongArray("getRecentReadLatencyHistogramMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getRecentReadLatencyMicros()
func (jbobject *DbColumnFamilyStoreMBean) GetRecentReadLatencyMicros() float64 {
	jret, err := jbobject.CallDouble("getRecentReadLatencyMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getWriteCount()
func (jbobject *DbColumnFamilyStoreMBean) GetWriteCount() int64 {
	jret, err := jbobject.CallLong("getWriteCount")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getTotalWriteLatencyMicros()
func (jbobject *DbColumnFamilyStoreMBean) GetTotalWriteLatencyMicros() int64 {
	jret, err := jbobject.CallLong("getTotalWriteLatencyMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getLifetimeWriteLatencyHistogramMicros()
func (jbobject *DbColumnFamilyStoreMBean) GetLifetimeWriteLatencyHistogramMicros() []int64 {
	jret, err := jbobject.CallLongArray("getLifetimeWriteLatencyHistogramMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getRecentWriteLatencyHistogramMicros()
func (jbobject *DbColumnFamilyStoreMBean) GetRecentWriteLatencyHistogramMicros() []int64 {
	jret, err := jbobject.CallLongArray("getRecentWriteLatencyHistogramMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getRecentWriteLatencyMicros()
func (jbobject *DbColumnFamilyStoreMBean) GetRecentWriteLatencyMicros() float64 {
	jret, err := jbobject.CallDouble("getRecentWriteLatencyMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract int getPendingTasks()
func (jbobject *DbColumnFamilyStoreMBean) GetPendingTasks() int {
	jret, err := jbobject.CallInt("getPendingTasks")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract int getLiveSSTableCount()
func (jbobject *DbColumnFamilyStoreMBean) GetLiveSSTableCount() int {
	jret, err := jbobject.CallInt("getLiveSSTableCount")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getLiveDiskSpaceUsed()
func (jbobject *DbColumnFamilyStoreMBean) GetLiveDiskSpaceUsed() int64 {
	jret, err := jbobject.CallLong("getLiveDiskSpaceUsed")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getTotalDiskSpaceUsed()
func (jbobject *DbColumnFamilyStoreMBean) GetTotalDiskSpaceUsed() int64 {
	jret, err := jbobject.CallLong("getTotalDiskSpaceUsed")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void forceMajorCompaction() throws ExecutionException, InterruptedException
func (jbobject *DbColumnFamilyStoreMBean) ForceMajorCompaction() error {
	err := jbobject.CallVoid("forceMajorCompaction")
	if err != nil {
		return err
	}
	return nil
}

// public abstract long getMinRowSize()
func (jbobject *DbColumnFamilyStoreMBean) GetMinRowSize() int64 {
	jret, err := jbobject.CallLong("getMinRowSize")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getMaxRowSize()
func (jbobject *DbColumnFamilyStoreMBean) GetMaxRowSize() int64 {
	jret, err := jbobject.CallLong("getMaxRowSize")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getMeanRowSize()
func (jbobject *DbColumnFamilyStoreMBean) GetMeanRowSize() int64 {
	jret, err := jbobject.CallLong("getMeanRowSize")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getBloomFilterFalsePositives()
func (jbobject *DbColumnFamilyStoreMBean) GetBloomFilterFalsePositives() int64 {
	jret, err := jbobject.CallLong("getBloomFilterFalsePositives")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getRecentBloomFilterFalsePositives()
func (jbobject *DbColumnFamilyStoreMBean) GetRecentBloomFilterFalsePositives() int64 {
	jret, err := jbobject.CallLong("getRecentBloomFilterFalsePositives")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getBloomFilterFalseRatio()
func (jbobject *DbColumnFamilyStoreMBean) GetBloomFilterFalseRatio() float64 {
	jret, err := jbobject.CallDouble("getBloomFilterFalseRatio")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getRecentBloomFilterFalseRatio()
func (jbobject *DbColumnFamilyStoreMBean) GetRecentBloomFilterFalseRatio() float64 {
	jret, err := jbobject.CallDouble("getRecentBloomFilterFalseRatio")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getBloomFilterDiskSpaceUsed()
func (jbobject *DbColumnFamilyStoreMBean) GetBloomFilterDiskSpaceUsed() int64 {
	jret, err := jbobject.CallLong("getBloomFilterDiskSpaceUsed")
	if err != nil {
		panic(err)
	}
	return jret
}

// public int getMinimumCompactionThreshold()
func (jbobject *DbColumnFamilyStoreMBean) GetMinimumCompactionThreshold() int {
	jret, err := jbobject.CallInt("getMinimumCompactionThreshold")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setMinimumCompactionThreshold(int threshold)
func (jbobject *DbColumnFamilyStoreMBean) SetMinimumCompactionThreshold(threshold int)  {
	err := jbobject.CallVoid("setMinimumCompactionThreshold", threshold)
	if err != nil {
		panic(err)
	}

}

// public int getMaximumCompactionThreshold()
func (jbobject *DbColumnFamilyStoreMBean) GetMaximumCompactionThreshold() int {
	jret, err := jbobject.CallInt("getMaximumCompactionThreshold")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setCompactionThresholds(int minThreshold, int maxThreshold)
func (jbobject *DbColumnFamilyStoreMBean) SetCompactionThresholds(minThreshold int, maxThreshold int)  {
	err := jbobject.CallVoid("setCompactionThresholds", minThreshold, maxThreshold)
	if err != nil {
		panic(err)
	}

}

// public void setMaximumCompactionThreshold(int threshold)
func (jbobject *DbColumnFamilyStoreMBean) SetMaximumCompactionThreshold(threshold int)  {
	err := jbobject.CallVoid("setMaximumCompactionThreshold", threshold)
	if err != nil {
		panic(err)
	}

}

// public void setCompactionStrategyClass(String className)
func (jbobject *DbColumnFamilyStoreMBean) SetCompactionStrategyClass(className string)  {
	conv_className := javabind.NewGoToJavaString()
	if err := conv_className.Convert(className); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("setCompactionStrategyClass", javabind.CastObject(conv_className.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_className.CleanUp()

}

// public String getCompactionStrategyClass()
func (jbobject *DbColumnFamilyStoreMBean) GetCompactionStrategyClass() string {
	jret, err := jbobject.CallObj("getCompactionStrategyClass", "java.lang.String")
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

// public Map<String,String> getCompressionParameters()
func (jbobject *DbColumnFamilyStoreMBean) GetCompressionParameters() map[string]string {
	jret, err := jbobject.CallObj("getCompressionParameters", "java.util.Map")
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

// public abstract void setCompressionParameters(java.util.Map<java.lang.String, java.lang.String>)
func (jbobject *DbColumnFamilyStoreMBean) SetCompressionParameters(a map[string]string)  {
	conv_a := javabind.NewGoToJavaMap(javabind.NewGoToJavaString(), javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("setCompressionParameters", javabind.CastObject(conv_a.Value(), "java.util.Map"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setCrcCheckChance(double crcCheckChance)
func (jbobject *DbColumnFamilyStoreMBean) SetCrcCheckChance(crcCheckChance float64)  {
	err := jbobject.CallVoid("setCrcCheckChance", crcCheckChance)
	if err != nil {
		panic(err)
	}

}

// public boolean isAutoCompactionDisabled()
func (jbobject *DbColumnFamilyStoreMBean) IsAutoCompactionDisabled() bool {
	jret, err := jbobject.CallBoolean("isAutoCompactionDisabled")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getTombstonesPerSlice()
func (jbobject *DbColumnFamilyStoreMBean) GetTombstonesPerSlice() float64 {
	jret, err := jbobject.CallDouble("getTombstonesPerSlice")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getLiveCellsPerSlice()
func (jbobject *DbColumnFamilyStoreMBean) GetLiveCellsPerSlice() float64 {
	jret, err := jbobject.CallDouble("getLiveCellsPerSlice")
	if err != nil {
		panic(err)
	}
	return jret
}

// public long estimateKeys()
func (jbobject *DbColumnFamilyStoreMBean) EstimateKeys() int64 {
	jret, err := jbobject.CallLong("estimateKeys")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getEstimatedRowSizeHistogram()
func (jbobject *DbColumnFamilyStoreMBean) GetEstimatedRowSizeHistogram() []int64 {
	jret, err := jbobject.CallLongArray("getEstimatedRowSizeHistogram")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getEstimatedColumnCountHistogram()
func (jbobject *DbColumnFamilyStoreMBean) GetEstimatedColumnCountHistogram() []int64 {
	jret, err := jbobject.CallLongArray("getEstimatedColumnCountHistogram")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getCompressionRatio()
func (jbobject *DbColumnFamilyStoreMBean) GetCompressionRatio() float64 {
	jret, err := jbobject.CallDouble("getCompressionRatio")
	if err != nil {
		panic(err)
	}
	return jret
}

// public List<String> getBuiltIndexes()
func (jbobject *DbColumnFamilyStoreMBean) GetBuiltIndexes() []string {
	jret, err := jbobject.CallObj("getBuiltIndexes", "java.util.List")
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

// public List<String> getSSTablesForKey(String key)
func (jbobject *DbColumnFamilyStoreMBean) GetSSTablesForKey(key string) []string {
	conv_key := javabind.NewGoToJavaString()
	if err := conv_key.Convert(key); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallObj("getSSTablesForKey", "java.util.List", javabind.CastObject(conv_key.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
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

// public void loadNewSSTables()
func (jbobject *DbColumnFamilyStoreMBean) LoadNewSSTables()  {
	err := jbobject.CallVoid("loadNewSSTables")
	if err != nil {
		panic(err)
	}

}

// public int getUnleveledSSTables()
func (jbobject *DbColumnFamilyStoreMBean) GetUnleveledSSTables() int {
	jret, err := jbobject.CallInt("getUnleveledSSTables")
	if err != nil {
		panic(err)
	}
	return jret
}

// public int[] getSSTableCountPerLevel()
func (jbobject *DbColumnFamilyStoreMBean) GetSSTableCountPerLevel() []int {
	jret, err := jbobject.CallIntArray("getSSTableCountPerLevel")
	if err != nil {
		panic(err)
	}
	return jret
}

// public double getDroppableTombstoneRatio()
func (jbobject *DbColumnFamilyStoreMBean) GetDroppableTombstoneRatio() float64 {
	jret, err := jbobject.CallDouble("getDroppableTombstoneRatio")
	if err != nil {
		panic(err)
	}
	return jret
}

