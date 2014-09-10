package nodeprobe

import "github.com/timob/javabind"

type ServiceStorageProxyMBean struct {
	*javabind.Callable
}

// public abstract long getReadOperations()
func (jbobject *ServiceStorageProxyMBean) GetReadOperations() int64 {
	jret, err := jbobject.CallLong("getReadOperations")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getTotalReadLatencyMicros()
func (jbobject *ServiceStorageProxyMBean) GetTotalReadLatencyMicros() int64 {
	jret, err := jbobject.CallLong("getTotalReadLatencyMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getRecentReadLatencyMicros()
func (jbobject *ServiceStorageProxyMBean) GetRecentReadLatencyMicros() float64 {
	jret, err := jbobject.CallDouble("getRecentReadLatencyMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getTotalReadLatencyHistogramMicros()
func (jbobject *ServiceStorageProxyMBean) GetTotalReadLatencyHistogramMicros() []int64 {
	jret, err := jbobject.CallLongArray("getTotalReadLatencyHistogramMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getRecentReadLatencyHistogramMicros()
func (jbobject *ServiceStorageProxyMBean) GetRecentReadLatencyHistogramMicros() []int64 {
	jret, err := jbobject.CallLongArray("getRecentReadLatencyHistogramMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getRangeOperations()
func (jbobject *ServiceStorageProxyMBean) GetRangeOperations() int64 {
	jret, err := jbobject.CallLong("getRangeOperations")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getTotalRangeLatencyMicros()
func (jbobject *ServiceStorageProxyMBean) GetTotalRangeLatencyMicros() int64 {
	jret, err := jbobject.CallLong("getTotalRangeLatencyMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getRecentRangeLatencyMicros()
func (jbobject *ServiceStorageProxyMBean) GetRecentRangeLatencyMicros() float64 {
	jret, err := jbobject.CallDouble("getRecentRangeLatencyMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getTotalRangeLatencyHistogramMicros()
func (jbobject *ServiceStorageProxyMBean) GetTotalRangeLatencyHistogramMicros() []int64 {
	jret, err := jbobject.CallLongArray("getTotalRangeLatencyHistogramMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getRecentRangeLatencyHistogramMicros()
func (jbobject *ServiceStorageProxyMBean) GetRecentRangeLatencyHistogramMicros() []int64 {
	jret, err := jbobject.CallLongArray("getRecentRangeLatencyHistogramMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getWriteOperations()
func (jbobject *ServiceStorageProxyMBean) GetWriteOperations() int64 {
	jret, err := jbobject.CallLong("getWriteOperations")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getTotalWriteLatencyMicros()
func (jbobject *ServiceStorageProxyMBean) GetTotalWriteLatencyMicros() int64 {
	jret, err := jbobject.CallLong("getTotalWriteLatencyMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getRecentWriteLatencyMicros()
func (jbobject *ServiceStorageProxyMBean) GetRecentWriteLatencyMicros() float64 {
	jret, err := jbobject.CallDouble("getRecentWriteLatencyMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getTotalWriteLatencyHistogramMicros()
func (jbobject *ServiceStorageProxyMBean) GetTotalWriteLatencyHistogramMicros() []int64 {
	jret, err := jbobject.CallLongArray("getTotalWriteLatencyHistogramMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long[] getRecentWriteLatencyHistogramMicros()
func (jbobject *ServiceStorageProxyMBean) GetRecentWriteLatencyHistogramMicros() []int64 {
	jret, err := jbobject.CallLongArray("getRecentWriteLatencyHistogramMicros")
	if err != nil {
		panic(err)
	}
	return jret
}

// public long getTotalHints()
func (jbobject *ServiceStorageProxyMBean) GetTotalHints() int64 {
	jret, err := jbobject.CallLong("getTotalHints")
	if err != nil {
		panic(err)
	}
	return jret
}

// public boolean getHintedHandoffEnabled()
func (jbobject *ServiceStorageProxyMBean) GetHintedHandoffEnabled() bool {
	jret, err := jbobject.CallBoolean("getHintedHandoffEnabled")
	if err != nil {
		panic(err)
	}
	return jret
}

// public Set<String> getHintedHandoffEnabledByDC()
func (jbobject *ServiceStorageProxyMBean) GetHintedHandoffEnabledByDC() []string {
	jret, err := jbobject.CallObj("getHintedHandoffEnabledByDC", "java.util.Set")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoSet(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setHintedHandoffEnabled(boolean b)
func (jbobject *ServiceStorageProxyMBean) SetHintedHandoffEnabled(b bool)  {
	err := jbobject.CallVoid("setHintedHandoffEnabled", b)
	if err != nil {
		panic(err)
	}

}

// public void setHintedHandoffEnabledByDCList(String dcs)
func (jbobject *ServiceStorageProxyMBean) SetHintedHandoffEnabledByDCList(dcs string)  {
	conv_dcs := javabind.NewGoToJavaString()
	if err := conv_dcs.Convert(dcs); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("setHintedHandoffEnabledByDCList", javabind.CastObject(conv_dcs.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_dcs.CleanUp()

}

// public int getMaxHintWindow()
func (jbobject *ServiceStorageProxyMBean) GetMaxHintWindow() int {
	jret, err := jbobject.CallInt("getMaxHintWindow")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setMaxHintWindow(int ms)
func (jbobject *ServiceStorageProxyMBean) SetMaxHintWindow(ms int)  {
	err := jbobject.CallVoid("setMaxHintWindow", ms)
	if err != nil {
		panic(err)
	}

}

// public int getMaxHintsInProgress()
func (jbobject *ServiceStorageProxyMBean) GetMaxHintsInProgress() int {
	jret, err := jbobject.CallInt("getMaxHintsInProgress")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setMaxHintsInProgress(int qs)
func (jbobject *ServiceStorageProxyMBean) SetMaxHintsInProgress(qs int)  {
	err := jbobject.CallVoid("setMaxHintsInProgress", qs)
	if err != nil {
		panic(err)
	}

}

// public int getHintsInProgress()
func (jbobject *ServiceStorageProxyMBean) GetHintsInProgress() int {
	jret, err := jbobject.CallInt("getHintsInProgress")
	if err != nil {
		panic(err)
	}
	return jret
}

// public Long getRpcTimeout()
func (jbobject *ServiceStorageProxyMBean) GetRpcTimeout() int64 {
	jret, err := jbobject.CallObj("getRpcTimeout", "java.lang.Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setRpcTimeout(Long timeoutInMillis)
func (jbobject *ServiceStorageProxyMBean) SetRpcTimeout(timeoutInMillis int64)  {
	conv_timeoutInMillis := javabind.NewGoToJavaLong()
	if err := conv_timeoutInMillis.Convert(timeoutInMillis); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("setRpcTimeout", javabind.CastObject(conv_timeoutInMillis.Value(), "java.lang.Long"))
	if err != nil {
		panic(err)
	}
	conv_timeoutInMillis.CleanUp()

}

// public Long getReadRpcTimeout()
func (jbobject *ServiceStorageProxyMBean) GetReadRpcTimeout() int64 {
	jret, err := jbobject.CallObj("getReadRpcTimeout", "java.lang.Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setReadRpcTimeout(Long timeoutInMillis)
func (jbobject *ServiceStorageProxyMBean) SetReadRpcTimeout(timeoutInMillis int64)  {
	conv_timeoutInMillis := javabind.NewGoToJavaLong()
	if err := conv_timeoutInMillis.Convert(timeoutInMillis); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("setReadRpcTimeout", javabind.CastObject(conv_timeoutInMillis.Value(), "java.lang.Long"))
	if err != nil {
		panic(err)
	}
	conv_timeoutInMillis.CleanUp()

}

// public Long getWriteRpcTimeout()
func (jbobject *ServiceStorageProxyMBean) GetWriteRpcTimeout() int64 {
	jret, err := jbobject.CallObj("getWriteRpcTimeout", "java.lang.Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setWriteRpcTimeout(Long timeoutInMillis)
func (jbobject *ServiceStorageProxyMBean) SetWriteRpcTimeout(timeoutInMillis int64)  {
	conv_timeoutInMillis := javabind.NewGoToJavaLong()
	if err := conv_timeoutInMillis.Convert(timeoutInMillis); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("setWriteRpcTimeout", javabind.CastObject(conv_timeoutInMillis.Value(), "java.lang.Long"))
	if err != nil {
		panic(err)
	}
	conv_timeoutInMillis.CleanUp()

}

// public Long getCasContentionTimeout()
func (jbobject *ServiceStorageProxyMBean) GetCasContentionTimeout() int64 {
	jret, err := jbobject.CallObj("getCasContentionTimeout", "java.lang.Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setCasContentionTimeout(Long timeoutInMillis)
func (jbobject *ServiceStorageProxyMBean) SetCasContentionTimeout(timeoutInMillis int64)  {
	conv_timeoutInMillis := javabind.NewGoToJavaLong()
	if err := conv_timeoutInMillis.Convert(timeoutInMillis); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("setCasContentionTimeout", javabind.CastObject(conv_timeoutInMillis.Value(), "java.lang.Long"))
	if err != nil {
		panic(err)
	}
	conv_timeoutInMillis.CleanUp()

}

// public Long getRangeRpcTimeout()
func (jbobject *ServiceStorageProxyMBean) GetRangeRpcTimeout() int64 {
	jret, err := jbobject.CallObj("getRangeRpcTimeout", "java.lang.Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setRangeRpcTimeout(Long timeoutInMillis)
func (jbobject *ServiceStorageProxyMBean) SetRangeRpcTimeout(timeoutInMillis int64)  {
	conv_timeoutInMillis := javabind.NewGoToJavaLong()
	if err := conv_timeoutInMillis.Convert(timeoutInMillis); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("setRangeRpcTimeout", javabind.CastObject(conv_timeoutInMillis.Value(), "java.lang.Long"))
	if err != nil {
		panic(err)
	}
	conv_timeoutInMillis.CleanUp()

}

// public Long getTruncateRpcTimeout()
func (jbobject *ServiceStorageProxyMBean) GetTruncateRpcTimeout() int64 {
	jret, err := jbobject.CallObj("getTruncateRpcTimeout", "java.lang.Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTruncateRpcTimeout(Long timeoutInMillis)
func (jbobject *ServiceStorageProxyMBean) SetTruncateRpcTimeout(timeoutInMillis int64)  {
	conv_timeoutInMillis := javabind.NewGoToJavaLong()
	if err := conv_timeoutInMillis.Convert(timeoutInMillis); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("setTruncateRpcTimeout", javabind.CastObject(conv_timeoutInMillis.Value(), "java.lang.Long"))
	if err != nil {
		panic(err)
	}
	conv_timeoutInMillis.CleanUp()

}

// public void reloadTriggerClasses()
func (jbobject *ServiceStorageProxyMBean) ReloadTriggerClasses()  {
	err := jbobject.CallVoid("reloadTriggerClasses")
	if err != nil {
		panic(err)
	}

}

// public long getReadRepairAttempted()
func (jbobject *ServiceStorageProxyMBean) GetReadRepairAttempted() int64 {
	jret, err := jbobject.CallLong("getReadRepairAttempted")
	if err != nil {
		panic(err)
	}
	return jret
}

// public long getReadRepairRepairedBlocking()
func (jbobject *ServiceStorageProxyMBean) GetReadRepairRepairedBlocking() int64 {
	jret, err := jbobject.CallLong("getReadRepairRepairedBlocking")
	if err != nil {
		panic(err)
	}
	return jret
}

// public long getReadRepairRepairedBackground()
func (jbobject *ServiceStorageProxyMBean) GetReadRepairRepairedBackground() int64 {
	jret, err := jbobject.CallLong("getReadRepairRepairedBackground")
	if err != nil {
		panic(err)
	}
	return jret
}

// public Map<String, List<String>> getSchemaVersions()
func (jbobject *ServiceStorageProxyMBean) GetSchemaVersions() map[string][]string {
	jret, err := jbobject.CallObj("getSchemaVersions", "java.util.Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoList(javabind.NewJavaToGoString()))
	dst := new(map[string][]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

