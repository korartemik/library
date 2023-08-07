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
('Book11', 'Author12'),
('Book21', 'Author12'),
('Book31', 'Author22'),
('Book41', 'Author22'),
('Book5', 'Author3');