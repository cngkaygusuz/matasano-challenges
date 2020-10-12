from util import pkcs7


AES_BLOCK_SIZE = 16



if __name__ == '__main__':
    assert pkcs7(b'YELLOW SUBMARINE', 20) == b'YELLOW SUBMARINE\x04\x04\x04\x04'
    assert pkcs7(b'A' * 10) == b'A' * 10 + bytes([6]) * 6
