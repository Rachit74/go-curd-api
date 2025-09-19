#### Basic Project structure

The project is split into layers, each with its own responsibility. This makes the code modular, testable, and maintainable as the project grows.

cmd/
Entry point of your application (main.go). Keeps the root clean.

config/
Handles app configuration (DB credentials, ports, etc.) so you don’t hardcode things.

db/
Contains database connection logic. Keeps DB setup separate from business logic.

models/
Defines the data structure (like Todo). Think of it as your schema.

repository/
Handles database queries (CRUD). This way, if you change the DB (say, from Postgres to MySQL), only this layer changes.

service/
Contains business logic. Example: mark todo as completed only if it has a title. Keeps rules separate from DB calls.

handlers/
API layer – translates HTTP requests into calls to services and sends responses back. Keeps routes clean.

routes/
Central place for all routes. Easy to see what endpoints your app has.