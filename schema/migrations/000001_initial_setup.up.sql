-- Users 
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(255) UNIQUE NOT NULL,
    user_type VARCHAR(255) CHECK (user_type IN ('guest', 'manager')) NOT NULL
);
-- Hotels 
CREATE TABLE hotels (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location POINT NOT NULL,
    address VARCHAR(255) NOT NULL,
    manager_id UUID NOT NULL,
    FOREIGN KEY (manager_id) REFERENCES users(id) ON DELETE CASCADE
);
-- Room 
CREATE TABLE rooms (
    id UUID PRIMARY KEY,
    hotel_id UUID NOT NULL,
    floor INT NOT NULL,
    price INT NOT NULL CHECK (price >= 0),
    FOREIGN KEY (hotel_id) REFERENCES hotels(id) ON DELETE CASCADE
);
-- Bookings 
CREATE TABLE bookings (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    room_id UUID NOT NULL,
    hotel_id UUID NOT NULL,
    booking_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    duration INTERVAL NOT NULL CHECK (duration > '0 seconds'::INTERVAL),
    rating INT CHECK (
        rating BETWEEN 1 AND 5
    ),
    -- Optional user rating
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE,
    FOREIGN KEY (hotel_id) REFERENCES hotels(id) ON DELETE CASCADE
);
-- Indexes for quick lookups
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);
CREATE INDEX idx_hotels_location ON hotels(address);
CREATE INDEX idx_rooms_availability ON rooms(price);
CREATE INDEX idx_bookings_user_id ON bookings(user_id);
CREATE INDEX idx_bookings_room_id ON bookings(room_id);