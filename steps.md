# Steps

## Poject setup
### Enabling modules
Let Go know that we want to use modules functionality, to help manage any
third-party packages that our project imports.

    What module path for out project should be?

The module path is essentially just the canonical name or identifier for your
project. Needs to be globally unique and unlikely to be used by anyone else.
I might namesapce my module paths by basing them on a URL that I own.

    danicos.dev/snippetbox

## Web App Essentials

### Handler
They’re responsible for executing your application logic and for writing HTTP
response headers and bodies.
### Router
The second component is a router (or servemux in Go terminology). This stores
a mapping between the URL patterns for your application and the corresponding
handlers. Usually you have one servemux for your application containing all your
routes.
### Web Server
The last thing we need is a web server. One of the great things about Go is that
you can establish a web server and listen for incoming requests as part of your
application itself. You don’t need an external third-party server like Nginx or
Apache.

## Project structure
The cmd directory will contain the **application-specific** (application logic) code for the executable applications in the project. For now we’ll have just one executable application — the web application — which will live under the cmd/web directory.

The pkg directory will contain the ancillary non-application-specific code used in the project. We’ll use it to hold potentially reusable code like validation helpers and the SQL database models for the project.

The ui directory will contain the user-interface assets used by the web application. Specifically, the ui/html directory will contain HTML templates, and the ui/static directory will contain static files (like CSS and images).

## Configuration and error handling
- Set configuration settings for the application at runtime in an easy and idiomatic way using command-line flags. flag package.
- Improve the application log messages to include more information, and manage them differently depending on the type (or level) of log message.
- Make dependencies available to the handlers in a way that’s extensible, type-safe, and doesn’t get in the way when it comes to writing tests.
- Centralize error handling so that there is no need to repeat oneself when writing code.

### Dependency Injection
Most web applications will have multiple dependencies that their handlers need to access, such as a database connection pool, centralized error handlers, and template caches. What we really want to answer is: how can we make any dependency available to our handlers?

But in general, it is good practice to inject dependencies into your handlers. It makes your code more explicit, less error-prone and easier to unit test than if you use global variables.

For applications where all your handlers are in the same package, like ours, a neat way to inject dependencies is to put them into a custom application struct, and then define your handler functions as methods against application.

#### Closures for Dependency Injection

The pattern that we’re using to inject dependencies won’t work if your handlers are spread across multiple packages. In that case, an alternative approach is to create a config package exporting an Application struct and have your handler functions close over this to form a closure. Very roughly:

```go
func main() {
    app := &config.Application{
        ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
    }

    mux.Handle("/", handlers.Home(app))
}
```

```go
func Home(app *config.Application) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ...
        ts, err := template.ParseFiles(files...)
        if err != nil {
            app.ErrorLog.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
            return
        }
        ...
    }
}
```
### Centralize Error Handling

## Isolating the Application Routes
The responsibilities of main() function are limited to:

    Parsing the runtime configuration settings for the application;
    Establishing the dependencies for the handlers; and
    Running the HTTP server.

## Database-Driven Responses
For our Snippetbox web application to become truly useful we need somewhere to store (or persist) the data entered by users, and the ability to query this data store dynamically at runtime.

There are many different data stores we could use for our application — each with different pros and cons — but we’ll opt for the popular relational database MySQL.

In this section you’ll learn how to:

- Connect to MySQL from your web application (specifically, you’ll learn how to establish a pool of reusable connections).
- Create a standalone models package, so that your database logic is reusable and decoupled from your web application.
- Use the appropriate functions in Go’s database/sql package to execute different types of SQL statements, and how to avoid common errors that can lead to your server running out of resources.
- Prevent SQL injection attacks by correctly using placeholder parameters.
    Use transactions, so that you can execute multiple SQL statements in one atomic action.

### Installing a database driver.

