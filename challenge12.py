import base64

from challenge7 import ecb_enc

KEY = b'\xcd\x0e1\x80\xc9\xbb\xe1\xfe\xb5j\x80\x96/\x04\xb6<'

ENC_SUFFIX = base64.b64decode(
    'Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIG'
    'Rvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGll'
    'cyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQ'
    'pEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK'
)


def enc(value):
    return ecb_enc(value + ENC_SUFFIX, KEY)


def bytes_to_next_block():
    val = 1
    base_length = len(enc(b''))

    while base_length == len(enc(b'A' * val)):
        val += 1

    return val


if __name__ == '__main__':
    bnb = bytes_to_next_block()
    print(bnb)
