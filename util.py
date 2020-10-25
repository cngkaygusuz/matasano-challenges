import random
random.seed(0)


def chunk(input, size):
    for i in range(0, len(input), size):
        yield input[i:i+size]


def xor(a, b):
    assert len(a) == len(b)

    retval = b''

    for i in range(len(a)):
        retval += bytes([a[i] ^ b[i]])

    return retval


def randombytes(size):
    return bytearray([random.randint(0, 255) for _ in range(size)])


def pkcs7(input: bytes, block_size=16):
    mod = len(input) % block_size

    if mod == 0:
        pad = 16
    else:
        pad = block_size - mod

    return input + bytes([pad] * pad)


if __name__ == '__main__':
    assert len(list(chunk('foobar', 3))) == 2
    assert len(list(chunk('foo', 3))) == 1
    assert len(list(chunk('fo', 3))) == 1

    assert xor(b'\xf0\x0f', b'\x0f\x0f') == b'\xff\x00'

    print(randombytes(16))
