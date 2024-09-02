
-- Insert data into the owners table
INSERT INTO owners (owner_id, first_name, last_name, email, phone, address) VALUES
(gen_random_uuid(), 'John', 'Doe', 'john.doe@example.com', '123-456-7890', '123 Elm Street'),
(gen_random_uuid(), 'Jane', 'Smith', 'jane.smith@example.com', '987-654-3210', '456 Oak Avenue'),
(gen_random_uuid(), 'Alice', 'Johnson', 'alice.johnson@example.com', '555-555-5555', '789 Maple Road');

-- Insert data into the pets table
INSERT INTO pets (pet_id, name, species, breed, age, owner_id) VALUES
(gen_random_uuid(), 'Max', 'Dog', 'Labrador', 5, (SELECT owner_id FROM owners WHERE first_name = 'John')),
(gen_random_uuid(), 'Bella', 'Cat', 'Siamese', 3, (SELECT owner_id FROM owners WHERE first_name = 'Jane')),
(gen_random_uuid(), 'Charlie', 'Dog', 'Beagle', 2, (SELECT owner_id FROM owners WHERE first_name = 'Alice'));

-- Insert data into the veterinarians table
INSERT INTO veterinarians (veterinarian_id, first_name, last_name, specialty, phone, email) VALUES
(gen_random_uuid(), 'Dr. Emily', 'Clark', 'General Practice', '111-222-3333', 'emily.clark@petclinic.com'),
(gen_random_uuid(), 'Dr. Michael', 'Brown', 'Surgery', '222-333-4444', 'michael.brown@petclinic.com');

-- Insert data into the appointments table
INSERT INTO appointments (appointment_id, appointment_date, pet_id, veterinarian_id, reason) VALUES
(gen_random_uuid(), '2024-09-05 10:00:00', (SELECT pet_id FROM pets WHERE name = 'Max'), (SELECT veterinarian_id FROM veterinarians WHERE first_name = 'Dr. Emily'), 'Annual Check-up'),
(gen_random_uuid(), '2024-09-06 11:30:00', (SELECT pet_id FROM pets WHERE name = 'Bella'), (SELECT veterinarian_id FROM veterinarians WHERE first_name = 'Dr. Michael'), 'Dental Cleaning'),
(gen_random_uuid(), '2024-09-07 14:00:00', (SELECT pet_id FROM pets WHERE name = 'Charlie'), (SELECT veterinarian_id FROM veterinarians WHERE first_name = 'Dr. Emily'), 'Vaccination');
