
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

typedef uint8_t u8;
typedef int32_t i32;

static inline u8* read(u8*);

// Entry Point
void main(int argc, u8 *argv[]) {
    u8* filename;
    u8* buffer;
       
    for (i32 i =0; i < argc; i++) {
        if (strstr(argv[i], ".bl")) {
            printf("File Name Found \n");
            filename = argv[i];
        }
    }

    buffer = read(filename);

    if (buffer == NULL) {
        return;
    }

    printf("Content of this file is \n");
    printf("%s", buffer);

    free(buffer);
    return;
}

static inline u8* read(u8* filename) {
    u8* buffer;
    FILE* ptr;
    long len;

    ptr = fopen(filename, "r");

    if (ptr == NULL) {
        printf("File cannot be opened \n");
        return NULL;
    }

    fseek(ptr, 0L, SEEK_END);
    len = ftell(ptr);

    // Reset Index Of The Stream
    fseek(ptr, 0L, SEEK_SET);

    buffer = (u8*)calloc(len, sizeof(u8));

    if (buffer == NULL) {
        printf("Buffer Memory failed to allocate \n");
        return NULL;
    }

    fread(buffer, sizeof(u8), len, ptr);
    fclose(ptr);

    return buffer;
}