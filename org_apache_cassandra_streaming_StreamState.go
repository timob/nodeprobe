package nodeprobe

import "github.com/timob/javabind"

type StreamingStreamState struct {
	*javabind.Callable
}

// public boolean hasFailedSession()
func (jbobject *StreamingStreamState) HasFailedSession() bool {
	jret, err := jbobject.CallBoolean("hasFailedSession")
	if err != nil {
		panic(err)
	}
	return jret
}

