CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    password TEXT NOT NULL,
    password_confirm TEXT NOT NULL,
    forgot_pas_code VARCHAR(10),
    image_urls TEXT
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    senders_name VARCHAR(255) NOT NULL,
    buyers_name VARCHAR(255) NOT NULL,
    from_where INTEGER NOT NULL,
    where_to INTEGER NOT NULL,
    type_of_service INTEGER NOT NULL,
    weight VARCHAR(50) NOT NULL,
    status INTEGER DEFAULT 0,
    users_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    seria_id VARCHAR(255),
    started_time VARCHAR(255),
    finished_time VARCHAR(255),
    total_steps INTEGER DEFAULT 0,
    current_step_number INTEGER DEFAULT 0
);

CREATE TABLE qrcodes (
    id SERIAL PRIMARY KEY,
    urls TEXT NOT NULL,
    data TEXT NOT NULL,
    orders_id INTEGER NOT NULL REFERENCES orders(id)
);

CREATE TABLE points (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE order_tracking_steps (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES orders(id),
    location INTEGER NOT NULL REFERENCES points(id),
    step_date TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE announcement (
    id SERIAL PRIMARY KEY,
    category TEXT NOT NULL,
    time TEXT NOT NULL,
    from_where INTEGER NOT NULL REFERENCES points(id),
    where_to INTEGER NOT NULL REFERENCES points(id),
    text TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    name TEXT NOT NULL
);
