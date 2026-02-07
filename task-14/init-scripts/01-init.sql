-- Создание пользователя приложения
CREATE USER aldpro_user WITH PASSWORD 'postgres';

-- Создание базы данных
CREATE DATABASE aldpro_db
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

-- Подключение к базе данных
\c aldpro_db;

-- Предоставление прав пользователю приложения
GRANT CONNECT ON DATABASE aldpro_db TO aldpro_user;

-- Создание схемы и предоставление прав
CREATE SCHEMA IF NOT EXISTS aldpro_schema;
GRANT USAGE ON SCHEMA aldpro_schema TO aldpro_user;
GRANT CREATE ON SCHEMA aldpro_schema TO aldpro_user;

-- Права на все таблицы в схеме
ALTER DEFAULT PRIVILEGES IN SCHEMA aldpro_schema
    GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO aldpro_user;

-- Права на последовательности
ALTER DEFAULT PRIVILEGES IN SCHEMA aldpro_schema
    GRANT USAGE, SELECT ON SEQUENCES TO aldpro_user;

-- Разрешение на выполнение функций
ALTER DEFAULT PRIVILEGES IN SCHEMA aldpro_schema
    GRANT EXECUTE ON FUNCTIONS TO aldpro_user;