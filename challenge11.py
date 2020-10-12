import random

from util import randombytes, chunk

from challenge7 import ecb_enc
from challenge10 import cbc_enc

ENC_BASE = b'A' * 100


def random_enc():
    enc_data = randombytes(random.randint(5, 10)) + ENC_BASE + randombytes(random.randint(5, 10))
    key = randombytes(16)
    enc_method = random.choice(['ecb', 'cbc'])

    if enc_method == 'ecb':
        ciphertext = ecb_enc(enc_data, key)
    elif enc_method == 'cbc':
        ciphertext = cbc_enc(enc_data, key, randombytes(16))
    else:
        raise Exception(f'eh?: {enc_method}')

    return ciphertext, enc_method


def guess_enc_method(ciphertext):
    blocks = {}

    for chk in chunk(ciphertext, 16):
        if blocks.get(chk):
            return 'ecb'
        else:
            blocks[chk] = True

    return 'cbc'


def detect_enc():
    ROUNDS = 100

    correctly_guessed = 0

    for i in range(ROUNDS):
        ciphertext, enc_method = random_enc()

        if enc_method == guess_enc_method(ciphertext):
            correctly_guessed += 1

    print(ROUNDS, correctly_guessed)


if __name__ == '__main__':
    assert guess_enc_method(b'A'*100) == 'ecb'
    assert guess_enc_method(randombytes(100)) == 'cbc'

    detect_enc()
