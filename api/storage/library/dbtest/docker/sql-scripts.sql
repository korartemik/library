CREATE TABLE library 
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
);

INSERT library(title, author) 
VALUES
('Book1', 'Author1'),
('Book2', 'Author1'),
('Book3', 'Author2'),
('Book4', 'Author2'),
('Book5', 'Author3');