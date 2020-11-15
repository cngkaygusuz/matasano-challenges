import math, base64

from challenge7 import ecb_enc_raw
from util import xor, randombytes


def little_endian(int):
    return int.to_bytes(8, byteorder='little')


def op_ctr(data: bytearray, key: bytearray, nonce: int):
    """
    Encrypt or decrypt under ctr mode.

    :return: bytearray
    """
    assert nonce <= 0xffffffffffffffff
    nonce_bytes = little_endian(nonce)
    counter = 0
    retval = bytearray()
    chunksize = math.ceil(len(data) / 16)

    for i in range(chunksize):
        dub = nonce_bytes + little_endian(counter)
        chunk = data[i*16:(i+1)*16]
        retval += op_ctr_raw(chunk, key, dub)
        counter += 1
    return retval


def op_ctr_raw(chunk, key, dub):
    keystream = ecb_enc_raw(dub, key)
    return xor(chunk, keystream[:len(chunk)])


def challenge():
    ciphertext = base64.b64decode('L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ==')
    plaintext = op_ctr(ciphertext, b'YELLOW SUBMARINE', 0)
    print(plaintext)


def main():
    key = randombytes(16)
    input = randombytes(16)

    assert op_ctr(op_ctr(input, key, 0), key, 0) == input

    challenge()


if __name__ == '__main__':
    main()
