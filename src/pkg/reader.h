#include <stdlib.h>
#include <stdio.h>

#ifndef READER_H_
#define READER_H_
#endif

// readFileMalloc: Writes the content of a whole file to a dynamic array.
// Care: this function `malloc`s but doesn't `free`. You need to `free` the
// variable where you return the `content`s.
char *readFileMalloc(char *filename);

/* FUNCTION DECLARATIONS */

char *readFileMalloc(char *filename) {
  FILE *fp;
  char *content;
  char c;
  int i;
  int length;

  fp = fopen(filename, "r");
  if (fp == NULL) {
    return NULL;
  }

  (void) fseek(fp, 0, SEEK_END); // Seek to EOF to get the length with ftell()
  length = ftell(fp);
  (void) fseek(fp, 0, SEEK_SET); // Reset seek to beginning of file
  content = (char *) malloc(sizeof(char) * (length + 1));

  i = 0;
  while ((c = fgetc(fp)) != EOF) {
    content[i] = c;
    i++;
  }

  (void) fclose(fp);

  return content;
}
