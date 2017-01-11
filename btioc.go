package btioc

import "fmt"

// A very simple inversion-of-control container
// modeled from the Laravel IoC container.

type Ioc struct {
	container map[string]interface{}
}

var i *Ioc

// Initialises a singleton IoC container.
func init() {
	i = New()
}

// Create a new IoC container.
func New() *Ioc {
	return &Ioc{
		container: make(map[string]interface{}),
	}
}

// Returns true if the specified component is registered.
func (ioc *Ioc) IsRegistered(name string) bool {
	if _, ok := ioc.container[name]; ok {
		return true
	}
	return false
}

// Registers a new object in the container.
func (ioc *Ioc) Register(name string, obj interface{}) error {
	// Check if already registered
	if ioc.IsRegistered(name) {
		return fmt.Errorf("'%s' is already registered in IoC container.", name)
	}

	ioc.container[name] = obj
	return nil
}

// Returns the specified instance if it's in the container, otherwise return an error.
func (ioc *Ioc) Retrieve(name string) (error, interface{}) {
	if !ioc.IsRegistered(name) {
		return fmt.Errorf("'%s' is not registered in IoC container.", name), nil
	}
	return nil, ioc.container[name]
}