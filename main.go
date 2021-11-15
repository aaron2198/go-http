package main

import logger "github.com/aaron2198/vts_broker/logger"

// main function
func main() {
	//create environment
	env := getEnv()
	//Create the environment based on env.Conf
	env.boot()

	// Router entrypoint, a the core application routine
	routes(env)
}

// getEnv create a new environment and parse the hosts environment variables.
func getEnv() *Env {
	c := ReadConfig()
	env := &Env{
		Conf: c,
	}
	return env
}

// boot create a logger, connect to DB, migrate DB tables.
func (e *Env) boot() error {
	e.Logger = logger.CreateLogger(e.Conf.logLoc, e.Conf.logLevel)
	e.DBconnect()
	e.Migrate()
	return nil
}
