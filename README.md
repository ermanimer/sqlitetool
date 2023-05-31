# sqlitetool

[![Build](https://github.com/ermanimer/sqlitetool/actions/workflows/build.yml/badge.svg)](https://github.com/ermanimer/sqlitetool/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ermanimer/sqlitetool)](https://goreportcard.com/report/github.com/ermanimer/sqlitetool)

**sqlitetool** provides basic commands for SQLite databases.

The [db-info](#db-info) command prints all table names of a database.

The [table-info](#table-info) command prints all column names and column types of a table. 

The [db-diff](#db-diff) command compares two databases by table names. 

The [table-diff](#table-diff) command compares two tables by column names and types.

The [table-to-struct](#table-to-struct) command converts the table to a Go struct.

# Installation

Run the following command to install **sqlitetool**.

```zsh
go install github.com/ermanimer/sqlitetool/cmd/sqlitetool@latest
```

You can also find the latest binaries in [Releases](https://github.com/ermanimer/sqlitetool/releases).

# Help

Type the following command for help.

```zsh
sqlitetool --help
```

# Commands

## db-info

**db-info** command prints all of the table names in a database.

**Usage:**

```
sqlitetool db-info <db>
```

**Sample Command:**

```zsh
sqlitetool db-info ./sample_dbs/1.db
```

**Sample Output:**

```
SQLite Tool - Database Info

Database: /home/ermanimer/sqlitetool/sample_dbs/1.db

Tables:
+---+---------+
| # |  NAME   |
+---+---------+
| 1 | table_1 |
| 2 | table_2 |
+---+---------+
```

## table-info

**table-info** command prints all of the column names and types in a table.

**Usage:**

```
sqlitetool table-info <db> <table>
```

**Sample Command:**

```zsh
sqlitetool table-info ./sample_dbs/1.db table_1
```

**Sample Output:**

```
SQLite Tool - Table Info

Database: /home/ermanimer/sqlitetool/sample_dbs/1.db
Table: table_1

Columns:
+---+----------+---------+
| # |   NAME   |  TYPE   |
+---+----------+---------+
| 1 | column_1 | INTEGER |
| 2 | column_2 | TEXT    |
| 3 | column_3 | BLOB    |
| 4 | column_4 | REAL    |
| 5 | column_5 | NUMERIC |
+---+----------+---------+
```

## db-diff

**db-diff** command prints table name differences of two databases.

**Usage:**

```
sqlitetool db-diff <dbA> <dbB>
```

**Sample Command:**

```zsh
sqlitetool db-diff ./sample_dbs/1.db ./test/data/2.db
```

**Sample Output:**

```
SQLite Tool - Database Diff

Database A: /home/ermanimer/sqlitetool/sample_dbs/1.db
Database B: /home/ermanimer/sqlitetool/sample_dbs/2.db

Tables:
+------+---------+
| DIFF |  NAME   |
+------+---------+
| -    | table_0 |
| ~    | table_1 |
| +    | table_2 |
+------+---------+
```

## table-diff

**table-diff** command prints column name and type differences of two tables.

**Usage:**

```
sqlitetool table-diff <dbA> <tableA> <dbB> <tableB>
```

**Sample Command:**

```zsh
sqlitetool table-diff ./sample_dbs/1.db table_1 ./test/data/2.db table_1
```

**Sample Output:**

```
SQLite Tool - Table Diff

Database A: /home/ermanimer/sqlitetool/sample_dbs/1.db
Table A: table_1
Database B: /home/ermanimer/sqlitetool/sample_dbs/2.db
Table B: table_1

Columns:
+------+----------+---------+
| DIFF |   NAME   |  TYPE   |
+------+----------+---------+
| -    | column_0 | NUMERIC |
| -    | column_3 | TEXT    |
| ~    | column_1 | INTEGER |
| ~    | column_2 | TEXT    |
| ~    | column_4 | REAL    |
| +    | column_3 | BLOB    |
| +    | column_5 | NUMERIC |
+------+----------+---------+
```

## table-to-struct

**table-to-struct** command converts the table to a Go struct.

**Usage:**

```
sqlitetool table-to-struct <db> <table> <struct>
```

**Sample Table:**

```
Database: /home/ermanimer/sqlitetool/sample_dbs/3.db
Table: table_1

Columns:
+---+----------+----------+
| # |   NAME   |   TYPE   |
+---+----------+----------+
| 1 | column_1 | INTEGER  |
| 2 | column_2 | TEXT     |
| 3 | column_3 | BLOB     |
| 4 | column_4 | REAL     |
| 5 | column_5 | DATE     |
| 6 | column_6 | DATETIME |
+---+----------+----------+ 
```

**Sample Command:**

```zsh
sqlitetool table-to-struct ./sample_dbs/1.db table_1 Table1
```

**Sample Output:**

```
SQLite Tool - Table To Struct

Database: /home/ermanimer/sqlitetool/sample_dbs/3.db
Table: table_1
Struct: Table1

type Table1 struct {
	Column1 int64
	Column2 string
	Column3 interface{}
	Column4 float64
	Column5 time.Time
	Column6 time.Time
}
```
