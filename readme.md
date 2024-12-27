# How to run server
1. Install Docker
2. Run `docker-compose up` in the root directory
3. navigate into the postgres running container and run `psql -U postgres` to access the database and create the instashop database `CREATE DATABASE instashop;`
4. drop the running containers and start it again with `docker-compose up` to apply the changes.

# Routes and Endpoints
You can find the documented endpoints (necessary endpoints) over at [Postman](https://documenter.getpostman.com/view/28281208/2sAYJ6BenF)
