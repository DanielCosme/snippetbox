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
