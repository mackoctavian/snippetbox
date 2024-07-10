#SNIPPETBOX
To run this application Go version 1.19 or high must be installed
In the root of the project directory run "go run ./cmd/web" without qoutes
Open your browser to localhost:4000

# Code Structure
**cmd/web**
This is the main routing and business logic functionality of the app

**main.go**
Is the main entry point and contains func main()

**handlers.go**
Contains handlers which perform the required logic etc for a particular route

**routes.go**
Contains the routes. Handle() or HandleFunc() are used to map the route url to a particular handler function

**middleware.go**
Contains functions that return an Handler type. Internally they execute some additional logic before or after the route handler is executed. Middleware functions can be chained because they return the same type as a handler.

**helpers.go**
Contains various little functions which are used in the handlers e.g. ServerError, ClientError, template caching, decode posted form data etc

**templates.go**
Contains some functions specifically related to templates including newTemplateCache() and human readable date function

**cmd/internal/models**
The SQL model structures and methods on those structures perform DB CRUD operations

**errors.go**
this is a set of functions for capturing errors generated when performing CRUD operations

**cmd/internal/validator**
Contains validator.go which has validation functions for request form fields and other

**ui/html**
HTML template for base.html

**ui/html/pages**
HTML templates for we pages

**ui/html/partials**
HTML partials e.g. nav.html
