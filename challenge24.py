import math
import statistics
import challenge21 as mt
from util import xor, chunk


def mtcipherop(data, key):
    assert key < 2**16

    mt.seed(key)

    out = bytearray()

    for chk in chunk(data, 4):
        keystream = mt.extract().to_bytes(4, 'little')
        out += xor(chk, keystream[:len(chk)])

    return out


def breakmtcipher(ciphertext):
    lowestdev = 10000
    guessed_plaintext = b''

    for i in range(2**16):
        print(i)
        ptext = mtcipherop(ciphertext, i)
        sdev = statistics.pstdev(ptext)

        if sdev < lowestdev:
            lowestdev = sdev
            guessed_plaintext = ptext

    print(lowestdev)
    print(guessed_plaintext)


def main():
    ciphertext = mtcipherop(b'Brown fox jumped over a log or something', 10)
    breakmtcipher(ciphertext)


if __name__ == '__main__':
    main()



