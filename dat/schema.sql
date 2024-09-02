CREATE SCHEMA pet_clinic;

CREATE TABLE pet_clinic.owners (
    owner_id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address VARCHAR(255)
);

CREATE TABLE pet_clinic.pets (
    pet_id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    species VARCHAR(50) NOT NULL,
    breed VARCHAR(50),
    age INTEGER,
    owner_id UUID NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES pet_clinic.owners(owner_id)
);

CREATE TABLE pet_clinic.veterinarians (
    veterinarian_id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    specialty VARCHAR(100),
    phone VARCHAR(20),
    email VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE pet_clinic.appointments (
    appointment_id UUID PRIMARY KEY,
    appointment_date TIMESTAMP NOT NULL,
    pet_id UUID NOT NULL,
    veterinarian_id UUID NOT NULL,
    reason TEXT,
    FOREIGN KEY (pet_id) REFERENCES pet_clinic.pets(pet_id),
    FOREIGN KEY (veterinarian_id) REFERENCES pet_clinic.veterinarians(veterinarian_id)
);