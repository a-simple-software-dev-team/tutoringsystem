DROP DATABASE IF EXISTS tutoringsystem;
-- 创建数据库
CREATE DATABASE tutoringsystem;

-- 使用数据库
USE tutoringsystem;

-- 创建用户表
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    role ENUM('tutor', 'student') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建家教资料表
CREATE TABLE tutors (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    subjects VARCHAR(255),
    grades VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建学生需求表
CREATE TABLE students (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    subjects_needed VARCHAR(255),
    grades VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建日程管理表
CREATE TABLE schedules (
    id INT PRIMARY KEY AUTO_INCREMENT,
    tutor_id INT,
    student_id INT,
    date DATE,
    time TIME,
    FOREIGN KEY (tutor_id) REFERENCES users(id),
    FOREIGN KEY (student_id) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建家教资格认证表
CREATE TABLE tutor_certifications (
    id INT PRIMARY KEY AUTO_INCREMENT,
    tutor_id INT,
    degree VARCHAR(255),
    experience TEXT,
    certificates JSON,
    FOREIGN KEY (tutor_id) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建消息通知表
CREATE TABLE notifications (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    message TEXT,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_read BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 创建匹配结果表
CREATE TABLE matches (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    match_id INT,
    compatibility_score FLOAT,
    role ENUM('tutor', 'student'),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (match_id) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建可用时间表
CREATE TABLE available_times (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    day ENUM('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'),
    time_range VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users ADD COLUMN updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
ALTER TABLE users ADD COLUMN deleted_at DATETIME;

ALTER TABLE tutors ADD COLUMN updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
ALTER TABLE tutors ADD COLUMN deleted_at DATETIME;

ALTER TABLE students ADD COLUMN updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
ALTER TABLE students ADD COLUMN deleted_at DATETIME;

ALTER TABLE available_times ADD COLUMN updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
ALTER TABLE available_times ADD COLUMN deleted_at DATETIME;