CREATE DATABASE IF NOT EXISTS goonstone;

CREATE TABLE IF NOT EXISTS goonstone.users (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY UNIQUE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    remember_token VARCHAR(100),
    email_verified_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS goonstone.photos (
    id VARCHAR(255) NOT NULL PRIMARY KEY UNIQUE,
    user_id INT NOT NULL,
    FOREIGN KEY fk_photos_user_id (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    filename VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS goonstone.comments (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY UNIQUE,
    photo_id VARCHAR(255) NOT NULL,
    FOREIGN KEY fk_comments_photo_id (photo_id) REFERENCES photos (id) ON DELETE CASCADE ON UPDATE CASCADE,
    user_id INT NOT NULL,
    FOREIGN KEY fk_comments_user_id (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS goonstone.likes (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY UNIQUE,
    photo_id VARCHAR(255) NOT NULL,
    FOREIGN KEY fk_likes_photo_id (photo_id) REFERENCES photos (id) ON DELETE CASCADE ON UPDATE CASCADE,
    user_id INT NOT NULL,
    FOREIGN KEY fk_likes_user_id (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

