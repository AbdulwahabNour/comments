GO REST API 
==========
## Technologies Used:

* **Postgres**
* **Docker + Docker-Compose**
* **Go** (This one was fairly obvious no?)
* **Postman** - for testing our service manually with HTTP request.

## Frameworks + Libraries Used

* **sqlx** - for simplifying our interactions with the database
* **golang-migrate** - for running our migrations on app startup
* **dgrijalva/jwt-go** - for working with JWTs in our transport layer.
* **satori/go.uuid** - for generating and working with UUIDs.
* **sirupsen/logrus** - nicer logging.
* **stretchr/testify** - For easier testing!