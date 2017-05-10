package btioc

// Resets the global state.
func ResetGlobalState() {
	i = New()
}

func Container() map[string]interface{} {
	return i.Container()
}

// Returns the IoC container for testing.
func (ioc *Ioc) Container() map[string]interface{} {
	return ioc.container
}
