package nodeprobe

import "github.com/timob/javabind"

type ServiceCacheServiceMBean struct {
	*javabind.Callable
}

// public int getRowCacheSavePeriodInSeconds()
func (jbobject *ServiceCacheServiceMBean) GetRowCacheSavePeriodInSeconds() int {
	jret, err := jbobject.CallInt("getRowCacheSavePeriodInSeconds")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setRowCacheSavePeriodInSeconds(int rcspis)
func (jbobject *ServiceCacheServiceMBean) SetRowCacheSavePeriodInSeconds(rcspis int)  {
	err := jbobject.CallVoid("setRowCacheSavePeriodInSeconds", rcspis)
	if err != nil {
		panic(err)
	}

}

// public int getKeyCacheSavePeriodInSeconds()
func (jbobject *ServiceCacheServiceMBean) GetKeyCacheSavePeriodInSeconds() int {
	jret, err := jbobject.CallInt("getKeyCacheSavePeriodInSeconds")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setKeyCacheSavePeriodInSeconds(int kcspis)
func (jbobject *ServiceCacheServiceMBean) SetKeyCacheSavePeriodInSeconds(kcspis int)  {
	err := jbobject.CallVoid("setKeyCacheSavePeriodInSeconds", kcspis)
	if err != nil {
		panic(err)
	}

}

// public int getRowCacheKeysToSave()
func (jbobject *ServiceCacheServiceMBean) GetRowCacheKeysToSave() int {
	jret, err := jbobject.CallInt("getRowCacheKeysToSave")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setRowCacheKeysToSave(int rckts)
func (jbobject *ServiceCacheServiceMBean) SetRowCacheKeysToSave(rckts int)  {
	err := jbobject.CallVoid("setRowCacheKeysToSave", rckts)
	if err != nil {
		panic(err)
	}

}

// public int getKeyCacheKeysToSave()
func (jbobject *ServiceCacheServiceMBean) GetKeyCacheKeysToSave() int {
	jret, err := jbobject.CallInt("getKeyCacheKeysToSave")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setKeyCacheKeysToSave(int kckts)
func (jbobject *ServiceCacheServiceMBean) SetKeyCacheKeysToSave(kckts int)  {
	err := jbobject.CallVoid("setKeyCacheKeysToSave", kckts)
	if err != nil {
		panic(err)
	}

}

// public void invalidateKeyCache()
func (jbobject *ServiceCacheServiceMBean) InvalidateKeyCache()  {
	err := jbobject.CallVoid("invalidateKeyCache")
	if err != nil {
		panic(err)
	}

}

// public void invalidateRowCache()
func (jbobject *ServiceCacheServiceMBean) InvalidateRowCache()  {
	err := jbobject.CallVoid("invalidateRowCache")
	if err != nil {
		panic(err)
	}

}

// public void setRowCacheCapacityInMB(long capacity)
func (jbobject *ServiceCacheServiceMBean) SetRowCacheCapacityInMB(capacity int64)  {
	err := jbobject.CallVoid("setRowCacheCapacityInMB", capacity)
	if err != nil {
		panic(err)
	}

}

// public void setKeyCacheCapacityInMB(long capacity)
func (jbobject *ServiceCacheServiceMBean) SetKeyCacheCapacityInMB(capacity int64)  {
	err := jbobject.CallVoid("setKeyCacheCapacityInMB", capacity)
	if err != nil {
		panic(err)
	}

}

// public void saveCaches() throws ExecutionException, InterruptedException
func (jbobject *ServiceCacheServiceMBean) SaveCaches() error {
	err := jbobject.CallVoid("saveCaches")
	if err != nil {
		return err
	}
	return nil
}

// public abstract long getKeyCacheHits()
func (jbobject *ServiceCacheServiceMBean) GetKeyCacheHits() int64 {
	jret, err := jbobject.CallLong("getKeyCacheHits")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getRowCacheHits()
func (jbobject *ServiceCacheServiceMBean) GetRowCacheHits() int64 {
	jret, err := jbobject.CallLong("getRowCacheHits")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getKeyCacheRequests()
func (jbobject *ServiceCacheServiceMBean) GetKeyCacheRequests() int64 {
	jret, err := jbobject.CallLong("getKeyCacheRequests")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getRowCacheRequests()
func (jbobject *ServiceCacheServiceMBean) GetRowCacheRequests() int64 {
	jret, err := jbobject.CallLong("getRowCacheRequests")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getKeyCacheRecentHitRate()
func (jbobject *ServiceCacheServiceMBean) GetKeyCacheRecentHitRate() float64 {
	jret, err := jbobject.CallDouble("getKeyCacheRecentHitRate")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract double getRowCacheRecentHitRate()
func (jbobject *ServiceCacheServiceMBean) GetRowCacheRecentHitRate() float64 {
	jret, err := jbobject.CallDouble("getRowCacheRecentHitRate")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getRowCacheCapacityInMB()
func (jbobject *ServiceCacheServiceMBean) GetRowCacheCapacityInMB() int64 {
	jret, err := jbobject.CallLong("getRowCacheCapacityInMB")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getRowCacheCapacityInBytes()
func (jbobject *ServiceCacheServiceMBean) GetRowCacheCapacityInBytes() int64 {
	jret, err := jbobject.CallLong("getRowCacheCapacityInBytes")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getKeyCacheCapacityInMB()
func (jbobject *ServiceCacheServiceMBean) GetKeyCacheCapacityInMB() int64 {
	jret, err := jbobject.CallLong("getKeyCacheCapacityInMB")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getKeyCacheCapacityInBytes()
func (jbobject *ServiceCacheServiceMBean) GetKeyCacheCapacityInBytes() int64 {
	jret, err := jbobject.CallLong("getKeyCacheCapacityInBytes")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getRowCacheSize()
func (jbobject *ServiceCacheServiceMBean) GetRowCacheSize() int64 {
	jret, err := jbobject.CallLong("getRowCacheSize")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getRowCacheEntries()
func (jbobject *ServiceCacheServiceMBean) GetRowCacheEntries() int64 {
	jret, err := jbobject.CallLong("getRowCacheEntries")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getKeyCacheSize()
func (jbobject *ServiceCacheServiceMBean) GetKeyCacheSize() int64 {
	jret, err := jbobject.CallLong("getKeyCacheSize")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getKeyCacheEntries()
func (jbobject *ServiceCacheServiceMBean) GetKeyCacheEntries() int64 {
	jret, err := jbobject.CallLong("getKeyCacheEntries")
	if err != nil {
		panic(err)
	}
	return jret
}

