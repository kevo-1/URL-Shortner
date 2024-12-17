# URL-Shortner

## A simple URL shortener written in Go, which allows users to shorten URLs and redirect them using a shortened hash. 
### Note: this tool only runs locally until further updates.

## Features

- **Shorten URLs**: Users can generate short URLs by providing a URL.
- **Redirect to Original URL**: Shortened URLs will redirect users to the original URL.
- **Persistent Storage**: The URLs and their shortened hashes are stored in a database.

## Tech Stack

- **Go**: Programming language used for the server and logic.
- **MySQL**: Lightweight database used to store the original URLs and shortened hashes.
- **MD5**: Hashing algorithm used to generate unique shortened URL identifiers.

## Prerequisites

Before you begin, ensure that you have the following installed:

- Go (v1.16 or higher) - [Install Go](https://golang.org/doc/install)
- MySQL (if using MySQL as the database) - [Install MySQL](https://dev.mysql.com/downloads/installer/)

## Database

Run this script in your MySQL database to create the required table

```sql
CREATE TABLE urldata (
    shortened varchar(10) PRIMARY KEY,
    original varchar(2048) NOT NULL
);
