package nodeprobe

import "github.com/timob/javabind"

type LocatorEndpointSnitchInfoMBean struct {
	*javabind.Callable
}

// public String getRack(String host) throws UnknownHostException
func (jbobject *LocatorEndpointSnitchInfoMBean) GetRack(host string) (string, error) {
	conv_host := javabind.NewGoToJavaString()
	if err := conv_host.Convert(host); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallObj("getRack", "java.lang.String", javabind.CastObject(conv_host.Value(), "java.lang.String"))
	if err != nil {
		var zero string
		return zero, err
	}
	conv_host.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst, nil
}

// public String getDatacenter(String host) throws UnknownHostException
func (jbobject *LocatorEndpointSnitchInfoMBean) GetDatacenter(host string) (string, error) {
	conv_host := javabind.NewGoToJavaString()
	if err := conv_host.Convert(host); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallObj("getDatacenter", "java.lang.String", javabind.CastObject(conv_host.Value(), "java.lang.String"))
	if err != nil {
		var zero string
		return zero, err
	}
	conv_host.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(jret); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst, nil
}

// public String getSnitchName()
func (jbobject *LocatorEndpointSnitchInfoMBean) GetSnitchName() string {
	jret, err := jbobject.CallObj("getSnitchName", "java.lang.String")
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

