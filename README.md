# API Documentation

This repository provides an API with the following endpoints:

1. **Endpoint:** `/api/players/goals_and_penalties/match/:id`
   - **Method:** GET
   - **Description:** Retrieves the goals and penalties information for a specific match.
   - **Parameters:**
     - `:id` (path parameter) - The ID of the match.
   - **Example:** `GET /api/players/goals_and_penalties/match/12345`

2. **Endpoint:** `/api/team/:teamId/match/:matchId`
   - **Method:** GET
   - **Description:** Retrieves the team coach information for a specific match.
   - **Parameters:**
     - `:teamId` (path parameter) - The ID of the team.
     - `:matchId` (path parameter) - The ID of the match.
   - **Example:** `GET /api/team/98765/match/12345`

3. **Endpoint:** `/api/players/:playeId/stats/season/:startYear/:endYear`
   - **Method:** GET
   - **Description:** Retrieves the player's statistics for a specific season.
   - **Parameters:**
     - `:playeId` (path parameter) - The ID of the player.
     - `:startYear` (path parameter) - The start year of the season.
     - `:endYear` (path parameter) - The end year of the season.
   - **Example:** `GET /api/players/54321/stats/season/2022/2023`

4. **Endpoint:** `/api/teams/:teamId/stats/season/:startYear/:endYear`
   - **Method:** GET
   - **Description:** Retrieves the team's statistics for a specific season.
   - **Parameters:**
     - `:teamId` (path parameter) - The ID of the team.
     - `:startYear` (path parameter) - The start year of the season.
     - `:endYear` (path parameter) - The end year of the season.
   - **Example:** `GET /api/teams/98765/stats/season/2022/2023`

## Running the API

To run the API, follow these steps:

1. Make sure you have Go version 1.2 or higher installed.
2. Install the dependencies located in the `go.mod` file by running the following command:
   ```shell
   go mod download
   ```
3. Run a postgres server v15 locally or type
   ```shell
   docker-compose up -d
   ```
   to run the docker-compose.yml file which will create the 
   postgres server and also expose pgadmin on port :5050. In order to connect to 
   the created database via pgAdmin, email: admin@admin.com & password: root
   and the host that the database is running is db
4. Inside the postgres server create the database that you want to use and apply
   the migrations located in /db/schema.sql file
5. Create the views using the /db/views.sql file
5. Populate the database using the file /db/data.sql file
6. Run the application:
   Linux/MacOS
   ```shell
   DB_USERNAME=username DB_PASSWORD=password DB_NAME=name go run main.go
   ```
   Windows:
   ```shell
   set DB_USERNAME=username 
   set DB_PASSWORD=password 
   set DB_NAME=name
   go run main.go
   ```
   Where:
   DB_USERNAME username of database created (use guest if you run docker-compose file)
   DB_PASSWORD password of database created (use guest if you run docker-compose file)
   DB_NAME name of database created