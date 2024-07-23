# mysql-ci

This tools is a simple cli of mysql for CI.

## How to install

### download binary

```sh
wget https://github.com/KamikazeZirou/mysql-ci/releases/download/v0.1.1/mysql-ci-v0.1.1-linux-amd64.tar.gz
tar xzf mysql-ci-v0.1.1-linux-amd64.tar.gz # `mysql-ci` is created.
```

### go install

```sh
# You need to add "$(go env GOPATH)/bin" to "PATH"
go install github.com/KamikazeZirou/mysql-ci@v0.1.1
```

## Usage

Write a file named 'init.sql' like below.

```sql
CREATE DATABASE IF NOT EXISTS testdb;
USE testdb;

CREATE TABLE `foobar` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    PRIMARY KEY(`id`)
);
```

Then, run a command like below.

```sh
mysql-ci -h 127.0.0.1 -p 3306 -pass password -f 'init.sql'
```

You can directly specify a query by using option '-q' instead of '-f'.
