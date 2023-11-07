#include "reader.h"
#include "sqlite/sqlite3.h"
#include <limits.h>
#include <string.h>
#include <unistd.h>

#ifndef SQLEXECUTE_H_
#define SQLEXECUTE_H_
#endif

#define DATABASE_PATH "/data/example.db"

// sqlExecute: Accepts a string of SQL command and executes it for the database.
// Returns if the query was executed with `EXIT_SUCCESS`/ or `EXIT_FAILURE`.
// Prints the error to stderr.
int sqlExecute(char *scriptStmt) {
  char *err;
  char cwd[PATH_MAX];
  char dbPath[PATH_MAX];
  int resp = 0;
  sqlite3 *db;

  if (scriptStmt == NULL) {
    fprintf(stderr, "Error: No SQL commands passed.\n");
    return EXIT_FAILURE;
  }

  if (getcwd(cwd, sizeof(cwd)) == NULL) {
    fprintf(stderr, "Error: getting current working directory.\n");
    return EXIT_FAILURE;
  }

  (void) strncpy(dbPath, cwd, sizeof(dbPath));
  (void) strncat(dbPath, DATABASE_PATH, strlen(DATABASE_PATH));

  (void) sqlite3_open(dbPath, &db);
  resp = sqlite3_exec(db, scriptStmt, NULL, NULL, &err);

  (void) sqlite3_close(db);

  if (resp != SQLITE_OK) {
    fprintf(stderr, "Error: interfacing with database %s\n.", err);
    return EXIT_FAILURE;
  }

  return EXIT_SUCCESS;
}

// sqlExecuteFile: Accepts a file path to the sql query in a `.sql` file.
// Returns if the query was executed with `EXIT_SUCCESS`/ or `EXIT_FAILURE`.
// Prints the error to stderr.
int sqlExecuteFile(char *scriptPath) {
  char *err;
  char *sqlScriptContent;
  char cwd[PATH_MAX];
  char dbPath[PATH_MAX];
  char path[PATH_MAX];
  int resp = 0;
  sqlite3 *db;

  if (getcwd(cwd, sizeof(cwd)) == NULL) {
    fprintf(stderr, "Error: getting current working directory.\n");
    return EXIT_FAILURE;
  }

  (void) strncpy(dbPath, cwd, sizeof(dbPath));
  (void) strncat(dbPath, DATABASE_PATH, strlen(DATABASE_PATH));

  (void) strncpy(path, cwd, sizeof(path));
  (void) strncat(path, scriptPath, strlen(scriptPath));

  sqlScriptContent = readFileMalloc(path);
  if (sqlScriptContent == NULL) {
    fprintf(stderr, "Error: reading SQL script.\n");
    return EXIT_FAILURE;
  }

  (void) sqlite3_open(dbPath, &db);
  resp = sqlite3_exec(db, sqlScriptContent, NULL, NULL, &err);

  (void) sqlite3_close(db);
  free(sqlScriptContent); // Must free because readFileMalloc() calls malloc.

  if (resp != SQLITE_OK) {
    fprintf(stderr, "Error: interfacing with database %s\n.", err);
    return EXIT_FAILURE;
  }

  return EXIT_SUCCESS;
}
