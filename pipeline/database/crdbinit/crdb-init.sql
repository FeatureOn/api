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

INSERT INTO featureon.users (ID, Name, UserName, Password) VALUES ('40e6215d-b5c6-4896-987c-f30f3678f608', 'First User', 'firstu', b'\x8d\xc2\x80\xa0\x49\x00\x2b\x44\xdc\xa2\xe2\x8b\x72\x21\x3a\x5a\xaf\x0e\x30\x22');

CREATE TABLE products
(
    ID UUID PRIMARY KEY,
    Name VARCHAR(100) NOT NULL
);

INSERT INTO featureon.products (ID, Name) VALUES ('4f0e9aea-e4bd-11ec-8fea-0242ac120002', 'Product1');

CREATE TABLE environments
(
    ID UUID PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    ProductID UUID,
    CONSTRAINT fk_product
        FOREIGN KEY(ProductID)
        REFERENCES products(ID)
);

INSERT INTO featureon.environments (ID, Name, ProductID) VALUES ('4f0e9c8e-e4bd-11ec-8fea-0242ac120002', 'Dev', '4f0e9aea-e4bd-11ec-8fea-0242ac120002');
INSERT INTO featureon.environments (ID, Name, ProductID) VALUES ('4f0e9d92-e4bd-11ec-8fea-0242ac120002', 'Prod', '4f0e9aea-e4bd-11ec-8fea-0242ac120002');

CREATE TABLE features
(
    Key         VARCHAR(100) PRIMARY KEY,
    Name        VARCHAR(250) NOT NULL,
    ProductID UUID NOT NULL,
    Description  VARCHAR(1000) NOT NULL,
    DefaultState BOOLEAN NOT NULL,
    Active       BOOLEAN NOT NULL,
    CONSTRAINT fk_product
        FOREIGN KEY(ProductID)
            REFERENCES products(ID)
);

INSERT INTO featureon.features (Key, Name, ProductID, Description, DefaultState, Active) VALUES ('NewLogin', 'New Login Feature', '4f0e9aea-e4bd-11ec-8fea-0242ac120002', 'New loging feature with added captcha', true, true);
INSERT INTO featureon.features (Key, Name, ProductID, Description, DefaultState, Active) VALUES ('ProfPic', 'Profile picture feature', '4f0e9aea-e4bd-11ec-8fea-0242ac120002', 'Adding profile picture feature to the users', false, true);
INSERT INTO featureon.features (Key, Name, ProductID, Description, DefaultState, Active) VALUES ('TwitterInt', 'Twitter integration', '4f0e9aea-e4bd-11ec-8fea-0242ac120002', 'Twitter integration for users', false, true);

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

INSERT INTO featureon.flags (FeatureKey, EnvironmentID, Value) VALUES ('NewLogin', '4f0e9c8e-e4bd-11ec-8fea-0242ac120002', true);
INSERT INTO featureon.flags (FeatureKey, EnvironmentID, Value) VALUES ('NewLogin', '4f0e9d92-e4bd-11ec-8fea-0242ac120002', true);
INSERT INTO featureon.flags (FeatureKey, EnvironmentID, Value) VALUES ('ProfPic', '4f0e9c8e-e4bd-11ec-8fea-0242ac120002', true);
INSERT INTO featureon.flags (FeatureKey, EnvironmentID, Value) VALUES ('ProfPic', '4f0e9d92-e4bd-11ec-8fea-0242ac120002', false);
INSERT INTO featureon.flags (FeatureKey, EnvironmentID, Value) VALUES ('TwitterInt', '4f0e9c8e-e4bd-11ec-8fea-0242ac120002', false);
INSERT INTO featureon.flags (FeatureKey, EnvironmentID, Value) VALUES ('TwitterInt', '4f0e9d92-e4bd-11ec-8fea-0242ac120002', false);
