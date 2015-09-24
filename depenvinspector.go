package depenvinspector


import (
	"errors"
	"os"
)


var (
	errNotFound = errors.New("Not Found")
)


type Inspector interface {
	Register(string) Inspector
	Validate(string) error
	Inspect(string) (string, error)
}


type internalInspector struct {
	registry map[string]struct{}
}


func New() Inspector {
	registry := make(map[string]struct{})

	inspector := internalInspector{
		registry: registry,
	}

	return &inspector
}


func (inspector *internalInspector) Register(name string) Inspector {
	inspector.registry[name] = struct{}{}

	return inspector
}

func (inspector *internalInspector) Validate(deploymentEnvironmentName string) error {
	// Confirm that the 'deploymentEnvironment' is in the registry.
	// If it isn't, then return an error.
	if _, ok := inspector.registry[deploymentEnvironmentName]; !ok {
		return errNotFound
	}

	// Return.
	//
	// If we got this far in the code then everything is OK and
	// we do NOT return an error.
	return nil
}

func (inspector *internalInspector) Inspect(osEnvironmentVariableName string) (string, error) {

        fn := func() string {
                return os.Getenv(osEnvironmentVariableName)
        }

        return inspector.inspect(fn)
}

func (inspector *internalInspector) inspect(fn func()string) (string, error) {

	// Try to get the deployment environment name.
	deploymentEnvironmentName := fn()

	// Validate the deployment environment name.
	// I.e., it is in our registry.
	//
	// If it is not valid, then return an error.
	if err := inspector.Validate(deploymentEnvironmentName); nil != err {
		return "", err
	}

	// Return.
	return deploymentEnvironmentName, nil
}
