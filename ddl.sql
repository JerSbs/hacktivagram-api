-- TABLE

-- USERS TABLE
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    age INT NOT NULL
);

-- CATEGORIES TABLE
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- ARTICLES TABLE (POSTS)
CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    image_url TEXT NOT NULL,
    author_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category_id INT REFERENCES categories(id) ON DELETE SET NULL
);

-- LIKES TABLE
CREATE TABLE likes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    article_id INT NOT NULL REFERENCES articles(id) ON DELETE CASCADE,
    UNIQUE (user_id, article_id)
);

-- COMMENTS TABLE
CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    article_id INT NOT NULL REFERENCES articles(id) ON DELETE CASCADE,
    UNIQUE (user_id, article_id)
);

-- USER ACTIVITY LOG TABLE
CREATE TABLE user_activity_logs (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ====

-- Querry Seeding

-- Insert users
INSERT INTO users (full_name, username, email, password, address, age) VALUES
('Test1 User1', 'testuser1', 'test1@example.com', 'password1', 'Jakarta', 25),
('Test2 User2', 'testuser2', 'test2@example.com', 'password2', 'Bandung', 30),
('Test3 User3', 'testuser3', 'test3@example.com', 'password3', 'Surabaya', 28);

-- Insert categories
INSERT INTO categories (name) VALUES
('Food'),
('Travel'),
('Technology');

-- Insert articles (posts)
INSERT INTO articles (title, content, image_url, author_id, category_id) VALUES
('Post 1', 'Content1', 'https://example.com/image1', 1, 1),
('Post 2', 'Content2', 'https://example.com/image2', 2, 2),
('Post 3', 'Content3', 'https://example.com/image3', 3, 3);

-- Insert likes
INSERT INTO likes (user_id, article_id) VALUES
(1, 2),
(2, 3),
(3, 1);

-- Insert comments
INSERT INTO comments (content, user_id, article_id) VALUES
('Good!', 2, 1),
('Amazing!', 3, 2),
('Fantastic!', 1, 3);

-- Insert user activity logs
INSERT INTO user_activity_logs (user_id, description) VALUES
(1, 'user create new POST with ID 1'),
(2, 'user create new POST with ID 2'),
(3, 'user create new POST with ID 3'),
(1, 'user COMMENT on POST with ID 2'),
(2, 'user COMMENT on POST with ID 1');