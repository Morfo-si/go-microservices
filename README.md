# Go Microservices

A sample [Golang](https://go.dev) application to illustrate how a fully functional `microservice` can be written. The source code is provided as the final solution to the online course ["Build a Microservice with Go" by Frank P Moley III](https://www.linkedin.com/learning/build-a-microservice-with-go).

## Pre-requisites

The following should be installed on your local system:

* The [Go](https://go.dev) programming language;
* [Podman](https://podman.io/) or [Docker](https://www.docker.com) to run a containerized [PostgreSQL](https://www.postgresql.org/) database.
  
> [!TIP]
> **(OPTIONAL)** If you prefer to use a locally installed instance of [PostgreSQL](https://www.postgresql.org/), skip the next step to create the database 

## Create a Postgresql database

Execute the following command to start a `postgresql` database container:

```shell
make start-db
```

> [!NOTE] The database container will be created with the following configuration:
>
> ```text
>    Container name: local-pg
>    Database name: postgres
>    Postgresql user: postgres
>    Postgresql password: postgres
>    Postgresql port: 5432
> ```

## Create the database schema

Execute the following command to create the `wisdom` database schema:

```shell
make apply-db-schema
```

### Database Schema Description

The schema `wisdom` contains the following tables:

#### Table: `services`

This table stores information about different services.

* **service_id**: `UUID` - Primary key, a unique identifier for each service.
* **name**: `VARCHAR` - A unique name for the service.
* **price**: `NUMERIC(12,2)` - The price of the service.

#### Table: `customers`

This table stores information about customers.

* **customer_id**: `UUID` - Primary key, a unique identifier for each customer.
* **first_name**: `VARCHAR` - The first name of the customer.
* **last_name**: `VARCHAR` - The last name of the customer.
* **email**: `VARCHAR` - The email address of the customer.
* **phone**: `VARCHAR` - The phone number of the customer.
* **address**: `VARCHAR` - The physical address of the customer.

#### Table: `vendors`

This table stores information about vendors.

* **vendor_id**: `UUID` - Primary key, a unique identifier for each vendor.
* **name**: `VARCHAR` - The name of the vendor. This field is required.
* **contact**: `VARCHAR` - The contact person's name at the vendor.
* **phone**: `VARCHAR` - The phone number of the vendor.
* **email**: `VARCHAR` - The email address of the vendor.
* **address**: `VARCHAR` - The physical address of the vendor.

#### Table: `products`

This table stores information about products.

* **product_id**: `UUID` - Primary key, a unique identifier for each product.
* **name**: `VARCHAR` - A unique name for the product.
* **price**: `NUMERIC(12,2)` - The price of the product.
* **vendor_id**: `UUID` - Foreign key, references the `vendor_id` in the `vendors` table.

## Populate the database

Once the `wisdom` database schema is created, execute the following command to populate the `postgresql` database with sample data:

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
* `http://localhost:8080/customers`;
* `http://localhost:8080/products`;
* `http://localhost:8080/services`;
* `http://localhost:8080/vendors`;

## Cleaning up the database

You can always start with a fresh `postgresql` database by running the following command:

```shell
make clean-db
```
