/*
Package depenvinspector provides functionality for determining the deployment environment.

Example Usage

	// Create new deployment environment inspector.
	deploymentEnvironmentInspector := depenvinspector.New()
	
	// Register the valid deployment environments.
	//
	// Here our deployment environments are:
	//
	//	* DEV
	//	* STAGING
	//	* PROD
	deploymentEnvironmentInspector.Register("DEV")
	deploymentEnvironmentInspector.Register("STAGING")
	deploymentEnvironmentInspector.Register("PROD")
	
	// Figure out what our deployment environment is.
	//
	// Note that the variable 'deploymentEnvironmentFromCommandLinePtr' potentially
	// gets the deployment environment from the command line. (Code for that not
	// shown here.)
	//
	// We use Validate() to see if it matches on any of the valid deployment
	// environments that we already registered (with the Register() method).
	//
	// If we can't get it from there, then we to get the deployment environment from
	// the operating system environment variable named "MYAPP_ENV".
	// (YOU WOULD OF COURSE WANT TO CHANGE "MYAPP_ENV" TO SOMETHING MORE APPROPRIATE
	// FOR YOUR APPLICATION.)
	//
	//
	deploymentEnvironment := ""
	if err := deploymentEnvironmentInspector.Validate(*deploymentEnvironmentFromCommandLinePtr); nil == err {
		deploymentEnvironment = *deploymentEnvironmentFromCommandLinePtr
	} else {
		if name, err2 := deploymentEnvironmentInspector.Inspect("MYAPP_ENV"); nil == err2 {
			deploymentEnvironment = name
		} else {
			panic("Could not get deployment environment from the command line or the operation system environment.")
		}
	}

*/
package depenvinspector
