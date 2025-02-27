
#ifndef EXAMPLE_H
#define EXAMPLE_H
#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct ByteArray {
    unsigned char *data;
    int length;
} ByteArray;

void generateKey(const char *secretKeyPath, const char *keyDirPath);
void *createCrytoLabBySeceryKeyAndKeyDir(const char *secretKeyPath,
                                         const char *keyDirPath);
void *createCrytoLabByKeyDir(const char *keyDirPath);
ByteArray *encrypt_(void *cryptoLabPtr, uint64_t plainText,
                    ByteArray *byteArray); // gtest exist encrypt function.
uint64_t decrypt(void *cryptoLabPtr, ByteArray byteArray);
void freeByteArray(ByteArray *byteArray);
ByteArray *add(void *cryptoLabPtr, ByteArray byteArray1, ByteArray byteArray2);
ByteArray *addScalar(void *cryptoLabPtr, ByteArray byteArray,
                     uint64_t plainText);
ByteArray *sub(void *cryptoLabPtr, ByteArray byteArray1, ByteArray byteArray2);
ByteArray *subScalar(void *cryptoLabPtr, ByteArray byteArray,
                     uint64_t plainText);

#ifdef __cplusplus
}
#endif

#endif /* EXAMPLE_H */
