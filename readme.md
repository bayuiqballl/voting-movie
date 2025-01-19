# API Setup and Running Guide

This document will guide you through the steps to set up the environment, run migrations, and start the API.

## Prerequisites

Before you begin, make sure you have the following installed:

- **Database** (PostgreSQL)
- **Migrations** tool (GORM)

##  Setup Environment Variables

1. Create a `.env` file in the root directory of your project.
2. Add the following configuration variables based on your environment:

   ```ini
    DATABASE_USER=utsarrr
    DATABASE_PASSWORD=
    DATABASE_HOST=localhost
    DATABASE_NAME=postgres
    DATABASE_PORT=5432
    DATABASE_MIGRATION=true
    BASE_URL="http://localhost:5004" // host for get video
    JWT_SECRET=83xVjD9N/rUcZfqaT4KbtHYXc2HFY3u/ms0A132JY3c=
3. enabled migration DB
    ```ini
    DATABASE_MIGRATION=true
4. create folder uploads on root project
     ```ini
     uploads/
5. then change directory to /cmd then run main.go
    ```ini
    cd /cmd 
    go run main.go
6. test API on Postman
- API to login as an authenticated user and admin
    ```ini
        /user/login
        /admin/login
- API to register
    ```ini
    /admin
    /user
- API to create and upload movies. Required information related with a movies are at
least title, description, duration, artists, genres, watch URL (which points to the
uploaded movie file)
    ```ini
    /admin/upload-movie
    /uploads/:filename 
    /admin/insert-movie
- API to update movie
    ```ini
      /admin/update-movie
- API to track movie viewership
    ``` ini
        /viewership
- API to see most voted movie and most viewed genre, as an admin
    ``` ini
    /movie/most-data
- API to vote a movie as an authenticated user or to unvote a movie as an authenticated user
    ``` ini
    /vote
- API to list all of the user’s voted movie
    ``` ini
    /vote/list-user?movie_id=

