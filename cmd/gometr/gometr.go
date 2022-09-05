package gometr

type GoMetClient struct {
	url          string
	numOfSeconds int
	//getHealth() HealthCheck
}

type HealthCheck struct {
	ServiceID string
}
