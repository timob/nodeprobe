package nodeprobe

import "github.com/timob/javabind"

type ConcurrentJMXEnabledThreadPoolExecutorMBean struct {
	*javabind.Callable
}

// public int getTotalBlockedTasks()
func (jbobject *ConcurrentJMXEnabledThreadPoolExecutorMBean) GetTotalBlockedTasks() int {
	jret, err := jbobject.CallInt("getTotalBlockedTasks")
	if err != nil {
		panic(err)
	}
	return jret
}

// public int getCurrentlyBlockedTasks()
func (jbobject *ConcurrentJMXEnabledThreadPoolExecutorMBean) GetCurrentlyBlockedTasks() int {
	jret, err := jbobject.CallInt("getCurrentlyBlockedTasks")
	if err != nil {
		panic(err)
	}
	return jret
}

// public int getCoreThreads()
func (jbobject *ConcurrentJMXEnabledThreadPoolExecutorMBean) GetCoreThreads() int {
	jret, err := jbobject.CallInt("getCoreThreads")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setCoreThreads(int number)
func (jbobject *ConcurrentJMXEnabledThreadPoolExecutorMBean) SetCoreThreads(number int)  {
	err := jbobject.CallVoid("setCoreThreads", number)
	if err != nil {
		panic(err)
	}

}

// public int getMaximumThreads()
func (jbobject *ConcurrentJMXEnabledThreadPoolExecutorMBean) GetMaximumThreads() int {
	jret, err := jbobject.CallInt("getMaximumThreads")
	if err != nil {
		panic(err)
	}
	return jret
}

// public void setMaximumThreads(int number)
func (jbobject *ConcurrentJMXEnabledThreadPoolExecutorMBean) SetMaximumThreads(number int)  {
	err := jbobject.CallVoid("setMaximumThreads", number)
	if err != nil {
		panic(err)
	}

}

