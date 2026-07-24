-- Створюємо користувача для реплікації
CREATE ROLE replication_user WITH REPLICATION LOGIN ENCRYPTED PASSWORD 'replication_password_123';

-- Даємо права на читання (опціонально для логічної реплікації)
-- GRANT CONNECT ON DATABASE app TO replication_user;