package dbsqlite

/*

#cgo LDFLAGS: -lsqlite3

#include "sqlexecute.h"
#include <stdio.h>
#include <stdbool.h>

int execStmt(char *sqlScript, bool isfile) {
  int res;
  if (isfile == true) {
	res = sqlExecuteFile(sqlScript);
  } else {
	res = sqlExecute(sqlScript);
  }
  if (res != EXIT_SUCCESS) {
    fprintf(stderr, "Error: executing sql statement from %s\n",
	  sqlScript);
    return EXIT_FAILURE;
  }
  return EXIT_SUCCESS;
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

// ExecuteSqliteStmtFile: This function is a wrapper for the C code above.
// It will execute the SQL statements from the  file path passed to the `.sql`
// file or a string with the SQL commands. An error is returned if the SQL
// statement wasn't executed properly.
func ExecuteSqliteFileOrStmt(fPathOrStmt string, isfile bool) error {
	cstr := C.CString(fPathOrStmt)
	defer C.free(unsafe.Pointer(cstr))
	res := C.execStmt(cstr, C.bool(isfile))
	if res != 0 {
		return fmt.Errorf("Error: executing SQL statement from %s\n", fPathOrStmt)
	}
	return nil
}
