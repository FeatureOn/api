DROP DATABASE IF EXISTS featureon;
CREATE DATABASE featureon;
USE featureon;
CREATE TABLE users
(
    ID UUID PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    UserName VARCHAR(100) NOT NULL,
    Password BYTES NOT NULL
);

CREATE TABLE products
(
    ID UUID PRIMARY KEY,
    Name VARCHAR(100) NOT NULL
);

CREATE TABLE environments
(
    ID UUID PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    ProductID UUID,
    CONSTRAINT fk_product
        FOREIGN KEY(ProductID)
        REFERENCES products(ID)
);

CREATE TABLE features
(
    Key         VARCHAR(100) PRIMARY KEY,
    Name        VARCHAR(250) NOT NULL,
    Description  VARCHAR(1000) NOT NULL,
    DefaultState BOOLEAN NOT NULL,
    Active       BOOLEAN NOT NULL
);

CREATE TABLE flags
(
    FeatureKey VARCHAR(100) NOT NULL,
    EnvironmentID UUID NOT NULL,
    Value BOOLEAN NOT  NULL,
    CONSTRAINT fk_environment
        FOREIGN KEY(EnvironmentID)
        REFERENCES environments(ID),
    CONSTRAINT fk_feature
        FOREIGN KEY(FeatureKey)
        REFERENCES features(Key)
);