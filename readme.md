# Simple, scalable http server

This implementation documents a repository driven data access layer, an attempt to move GORM from the HTTP handlers. The goal was to add some advanced features that may not be ideal for this architecture in order to understand how to implement them.

## The pattern
At the core of this implementation, there are a series of repositories. These repositories are utilize our application model, but acqiusition could be from a database, a file system, an external datastore, depending on that models needs. All this repository should require is meeting the compliant handler interface.

`repository/community.go` implementes a database driven access layer that conforms to `interface/community.go`. Any repository can be substituted to `handler/community.go` as long as it implements `interface/community.go`.

## The surrounding bits

### Logging
This application has it's own logger.go, while this logger is simply logrus, it would be handy to set some expectations about logging verbosity and resolution at a global application scope. So, our logger contains a `logrus.Logger` and has neighboring functions based on application context to be used is standardized locations. For example, `RequestError` and `RequestDefaultError` are handy methods that append logrus fields to the logs by consuming the request data in a handler. `RequestDefaultError` adds the ability to provide a http.ResponseWriter, and allows the developer to respond to the user directly, or pass nil for a simple `httpInternalServerError` message.

### Configuration
This portion is just a simple way to parse some data from the applications environment, set defaults where helpful and provide the rest of the application with a valid configuration.

### Middleware
Simple, implemented as you'd expect, just a series of functions consuming http.HandlerFunc and returning one with some chunk of code prepended to it. There is currently profiling, http request logging, and CORS header middleware implemented. It would be good to add authentication to see how it fits this model. 