# Go Microservices

A sample [Golang](https://go.dev) application to illustrate how a fully functional `microservice` can be written. The source code is a modificated version for the example provided in the online course ["Build a Microservice with Go" by Frank P Moley III](https://www.linkedin.com/learning/build-a-microservice-with-go).

## Pre-requisites

The following should be installed on your local system:

* The [Go](https://go.dev) programming language;
* [Podman](https://podman.io/) or [Docker](https://www.docker.com) to run a containerized [PostgreSQL](https://www.postgresql.org/) database.
  
> [!TIP]
> **(OPTIONAL)** If you prefer to use a locally installed instance of [PostgreSQL](https://www.postgresql.org/), skip the next step to create the database.

## Create a Postgresql database

Execute the following command to start a `postgresql` database container:

```shell
make start-db
```

> [!NOTE] The database container will be created with the following configuration:
>
> ```text
>    Container name: petclinic-pg
>    Database name: postgres
>    Postgresql user: postgres
>    Postgresql password: postgres
>    Postgresql port: 5432
> ```

## Create the database schema

Execute the following command to create the `petclinic` database schema:

```shell
make apply-db-schema
```

### Database Schema Description

The schema `petclinic` contains the following tables:

#### Table: `owners`

This table stores information about the owners.

* **owner_id**: `UUID` - Primary key, a unique identifier for each owner.
* **first_name**: `VARCHAR(50)` - The first name of the owner.
* **last_name**: `VARCHAR(50)` - The last name of the owner.
* **email**: `VARCHAR(100)` - The email address of the owner, must be unique.
* **phone**: `VARCHAR(20)` - The phone number of the owner.
* **address**: `VARCHAR(255)` - The physical address of the owner.

#### Table: `pets`

This table stores information about the pets.

* **pet_id**: `UUID` - Primary key, a unique identifier for each pet.
* **name**: `VARCHAR(50)` - The name of the pet.
* **species**: `VARCHAR(50)` - The species of the pet (e.g., dog, cat).
* **breed**: `VARCHAR(50)` - The breed of the pet.
* **age**: `INTEGER` - The age of the pet.
* **owner_id**: `UUID` - Foreign key, references the owner_id in the owners table.

#### Table: `veterinarians`

This table stores information about the veterinarians.

* **veterinarian_id**: `UUID - Primary key, a unique identifier for each veterinarian.
* **first_name**: `VARCHAR(50) - The first name of the veterinarian.
* **last_name**: `VARCHAR(50) - The last name of the veterinarian.
* **specialty**: `VARCHAR(100) - The specialty of the veterinarian (e.g., surgery, general practice).
* **phone**: `VARCHAR(20) - The phone number of the veterinarian.
* **email**: `VARCHAR(100) - The email address of the veterinarian, must be unique.

#### Table: `appointments`

This table stores information about the appointments.

* **appointment_id**: `UUID` - Primary key, a unique identifier for each appointment.
* **appointment_date**: `TIMESTAMP` - The date and time of the appointment.
* **pet_id**: `UUID` - Foreign key, references the pet_id in the pets table.
* **veterinarian_id**: `UUID` - Foreign key, references the veterinarian_id in the veterinarians table.
* **reason**: `TEXT` - The reason for the appointment (e.g., check-up, vaccination).

## Populate the database

Once the `petclinic` database schema is created, execute the following command to populate the `postgresql` database with sample data:

```shell
make populate-db
```

## Run the Go microservice API service

Now you can start the Go microservice API service with the following command:

```shell
make run
```

The Go microservice API service can be reached at following endpoints:

* `http://localhost:8080/liveness`;
* `http://localhost:8080/readiness`;
* `http://localhost:8080/owners`;
* `http://localhost:8080/appointments`;
* `http://localhost:8080/pets`;
* `http://localhost:8080/veterinarians`;

## Cleaning up the database

You can always start with a fresh `postgresql` database by running the following command:

```shell
make clean-db
```
