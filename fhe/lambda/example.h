#ifndef EXAMPLE_H
#define EXAMPLE_H
#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    unsigned char* data;
    int length;
} ByteArray;
uint64_t myUnsigned64BitInteger;

void generateKey(char* secretKeyPath,char* keyDirPath);
void* makeContext();
void* makePack(void* contextPtr);
void* createCrytoLab(char* secretKeyPath,char* keyDirPath);
ByteArray* encrypt(void *cryptoLabPtr,uint64_t plainText, ByteArray* byteArray);
uint64_t decrypt(void *cryptoLabPtr, ByteArray byteArray);
void initPack();
void freeByteArray(ByteArray* byteArray);
ByteArray* add(void* cryptoLabPtr, ByteArray byteArray1, ByteArray byteArray2);
ByteArray* addScalar(void* cryptoLabPtr, ByteArray byteArray, uint64_t plainText);
ByteArray* sub(void* cryptoLabPtr, ByteArray byteArray1, ByteArray byteArray2);
ByteArray* subScalar(void* cryptoLabPtr, ByteArray byteArray, uint64_t plainText);



#ifdef __cplusplus
}
#endif

#endif /* EXAMPLE_H */
