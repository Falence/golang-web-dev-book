# Basic CRUD web app with Go HTML Templates
Building a note taker web app using Go html/templates that performs Create, Read, Update and Delete operations.

## Routes
- GET: **/**  - Gets all notes page 
- GET: **/notes/add**  - Create note page
- GET: **/notes/edit/{id}** - update note page

- POST: **/notes/save**  - Create note route
- PUT: **/notes/update/{id}** - update note route
- DELETE: **/notes/delete/{id}** - delete note route