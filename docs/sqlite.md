## SQLite

+ Create file `~/.sqliterc` with the below to show headers in column mode

  ```text
  .headers ON
  .mode column
  ```

+ Run sql script in sqlite

  ```zsh
  sqlite3 mockroblog.db < schema.sql
  ```

  OR

  ```zsh
  cat schema.sql | sqlite3 mockroblog.db
  ```

  OR

  ```zsh
  sqlite3
  .read mockroblog.db
  ```

  OR

  ```zsh
  sqlite3 mockroblog.db -init schema.sql
  ```

  `<` will exit the SQLite prompt immediately and return the error code to the shell, `.read` will leave the prompt up, and `-init` will always return 0. `<` is the best for scripting because it's cross-platform unlike `.read` which doesn't support Windows paths.

+ Open the newly created database in sqlite CLI

  ```zsh
  sqlite3 mockroblog.db
  ```

+ List tables

  ```sqlite
  .tables
  ```

+ Show schema of tables

  ```
  .schema user
  .schema post
  ```

+ Get data from tables (should be empty)

  ```sqlite
  select * from user;
  select * from post;
  ```
