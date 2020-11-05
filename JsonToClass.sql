CREATE TABLE `server` (
  `key` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` VARCHAR,
  `url` VARCHAR UNIQUE,
  `description` TEXT
);

CREATE TABLE `info` (
  `key` INTEGER PRIMARY KEY AUTOINCREMENT,
  `server_key` INTEGER REFERENCES server(key) ON DELETE CASCADE,
  `prefix` VARCHAR
);

CREATE TABLE `api` (
  `key` INTEGER PRIMARY KEY AUTOINCREMENT,
  `server_key` INTEGER REFERENCES server(key) ON DELETE CASCADE,
  `name` VARCHAR NOT NULL,
  `url` VARCHAR,
  `description` TEXT
);

CREATE TABLE `json_data` (
  `key` INTEGER PRIMARY KEY AUTOINCREMENT,
  `api_key` INTEGER REFERENCES api(key) ON DELETE CASCADE,
  `successed` BOOLEAN,
  `param_data` TEXT,
  `data` LONGTEXT,
  `description` TEXT,
  `created_date` DATETIME
);

CREATE TABLE `server_param` (
  `server_key` INTEGER REFERENCES server(key) ON DELETE CASCADE,
  `param_key` INTEGER REFERENCES param_data(key) ON DELETE CASCADE
);

CREATE TABLE `api_param` (
  `api_key` INTEGER REFERENCES api(key) ON DELETE CASCADE,
  `param_key` INTEGER REFERENCES param_data(key) ON DELETE CASCADE
);

CREATE TABLE `param_data` (
  `key` INTEGER PRIMARY KEY AUTOINCREMENT,
  `pkey` VARCHAR,
  `value` VARCHAR,
  `description` TEXT,
  `required` BOOLEAN DEFAULT true,
  `used` BOOLEAN DEFAULT true
);

CREATE TABLE `class_data` (
  `key` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` VARCHAR UNIQUE,
  `api_key` INTEGER REFERENCES api(key) ON DELETE CASCADE,
  `description` TEXT
);

CREATE TABLE `var_data` (
  `key` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` VARCHAR,
  `type_key` INTEGER REFERENCES var_type(key) ON DELETE CASCADE,
  `sub_type_key` INTEGER REFERENCES var_type(key) ON DELETE CASCADE,
  `description` TEXT
);

CREATE TABLE `class_var` (
  `class_key` INTEGER  REFERENCES class_data(key) ON DELETE CASCADE,
  `var_key` INTEGER REFERENCES var_type(key) ON DELETE CASCADE
);

CREATE TABLE `var_type` (
  `key` INTEGER PRIMARY KEY AUTOINCREMENT,
  `class_key` INTEGER
);