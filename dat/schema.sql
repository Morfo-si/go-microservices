CREATE TABLE owners (
    owner_id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address VARCHAR(255)
);

CREATE TABLE pets (
    pet_id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    species VARCHAR(50) NOT NULL,
    breed VARCHAR(50),
    age INTEGER,
    owner_id UUID NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES owners(owner_id)
);

CREATE TABLE veterinarians (
    veterinarian_id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    specialty VARCHAR(100),
    phone VARCHAR(20),
    email VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE appointments (
    appointment_id UUID PRIMARY KEY,
    appointment_date TIMESTAMP NOT NULL,
    pet_id UUID NOT NULL,
    veterinarian_id UUID NOT NULL,
    reason TEXT,
    FOREIGN KEY (pet_id) REFERENCES pets(pet_id),
    FOREIGN KEY (veterinarian_id) REFERENCES veterinarians(veterinarian_id)
);