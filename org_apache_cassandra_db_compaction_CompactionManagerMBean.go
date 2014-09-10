package nodeprobe

import "github.com/timob/javabind"

type DbCompactionCompactionManagerMBean struct {
	*javabind.Callable
}

// public List<Map<String, String>> getCompactions()
func (jbobject *DbCompactionCompactionManagerMBean) GetCompactions() []map[string]string {
	jret, err := jbobject.CallObj("getCompactions", "java.util.List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoString()))
	dst := new([]map[string]string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public List<String> getCompactionSummary()
func (jbobject *DbCompactionCompactionManagerMBean) GetCompactionSummary() []string {
	jret, err := jbobject.CallObj("getCompactionSummary", "java.util.List")
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

// public abstract int getPendingTasks()
func (jbobject *DbCompactionCompactionManagerMBean) GetPendingTasks() int {
	jret, err := jbobject.CallInt("getPendingTasks")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getCompletedTasks()
func (jbobject *DbCompactionCompactionManagerMBean) GetCompletedTasks() int64 {
	jret, err := jbobject.CallLong("getCompletedTasks")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getTotalBytesCompacted()
func (jbobject *DbCompactionCompactionManagerMBean) GetTotalBytesCompacted() int64 {
	jret, err := jbobject.CallLong("getTotalBytesCompacted")
	if err != nil {
		panic(err)
	}
	return jret
}

// public abstract long getTotalCompactionsCompleted()
func (jbobject *DbCompactionCompactionManagerMBean) GetTotalCompactionsCompleted() int64 {
	jret, err := jbobject.CallLong("getTotalCompactionsCompleted")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void forceUserDefinedCompaction(String dataFiles)
func (jbobject *DbCompactionCompactionManagerMBean) ForceUserDefinedCompaction(dataFiles string)  {
	conv_dataFiles := javabind.NewGoToJavaString()
	if err := conv_dataFiles.Convert(dataFiles); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("forceUserDefinedCompaction", javabind.CastObject(conv_dataFiles.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_dataFiles.CleanUp()

}

// public void stopCompaction(String type)
func (jbobject *DbCompactionCompactionManagerMBean) StopCompaction(type_gen string)  {
	conv_type := javabind.NewGoToJavaString()
	if err := conv_type.Convert(type_gen); err != nil {
		panic(err)
	}
	err := jbobject.CallVoid("stopCompaction", javabind.CastObject(conv_type.Value(), "java.lang.String"))
	if err != nil {
		panic(err)
	}
	conv_type.CleanUp()

}

// public int getCoreCompactorThreads()
func (jbobject *DbCompactionCompactionManagerMBean) GetCoreCompactorThreads() int {
	jret, err := jbobject.CallInt("getCoreCompactorThreads")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setCoreCompactorThreads(int number)
func (jbobject *DbCompactionCompactionManagerMBean) SetCoreCompactorThreads(number int)  {
	err := jbobject.CallVoid("setCoreCompactorThreads", number)
	if err != nil {
		panic(err)
	}

}

// public int getMaximumCompactorThreads()
func (jbobject *DbCompactionCompactionManagerMBean) GetMaximumCompactorThreads() int {
	jret, err := jbobject.CallInt("getMaximumCompactorThreads")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setMaximumCompactorThreads(int number)
func (jbobject *DbCompactionCompactionManagerMBean) SetMaximumCompactorThreads(number int)  {
	err := jbobject.CallVoid("setMaximumCompactorThreads", number)
	if err != nil {
		panic(err)
	}

}

// public int getCoreValidationThreads()
func (jbobject *DbCompactionCompactionManagerMBean) GetCoreValidationThreads() int {
	jret, err := jbobject.CallInt("getCoreValidationThreads")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setCoreValidationThreads(int number)
func (jbobject *DbCompactionCompactionManagerMBean) SetCoreValidationThreads(number int)  {
	err := jbobject.CallVoid("setCoreValidationThreads", number)
	if err != nil {
		panic(err)
	}

}

// public int getMaximumValidatorThreads()
func (jbobject *DbCompactionCompactionManagerMBean) GetMaximumValidatorThreads() int {
	jret, err := jbobject.CallInt("getMaximumValidatorThreads")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setMaximumValidatorThreads(int number)
func (jbobject *DbCompactionCompactionManagerMBean) SetMaximumValidatorThreads(number int)  {
	err := jbobject.CallVoid("setMaximumValidatorThreads", number)
	if err != nil {
		panic(err)
	}

}

