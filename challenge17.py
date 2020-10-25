import base64

from challenge10 import cbc_dec, cbc_enc
from util import randombytes, xor

KEY = b"\xd5TF\xdd\xde22\x82\x9e}_'8\x9e\x0e\x99"


def get_plaintexts():
    retval = """MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=
    MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=
    MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==
    MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==
    MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl
    MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==
    MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==
    MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=
    MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=
    MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93
    """.split('\n')

    retval = list(map(lambda el: base64.b64decode(el), retval))

    return retval


PLAINTEXTS = get_plaintexts()


def is_valid_padding(plaintext):
    assert plaintext and len(plaintext) % 16 == 0

    last_byte = plaintext[-1]

    #print(plaintext[-16:])

    # last byte cannot be 0; if no padding needed than last block is all 0x10's
    if last_byte == 0x00:
        return False

    last_block = plaintext[-int(last_byte):]
    expected = bytes([last_byte] * int(last_byte))

    #print(last_block, last_block == expected)
    return last_block == expected


def server(ciphertext, iv):
    """
    emulate ciphertext that is sent to server; return true if padding is correct, false otherwise
    """
    plaintext = cbc_dec(ciphertext, KEY, iv)
    return is_valid_padding(plaintext)


def guess(index, ciphertext, iv, guessed_plaintext):
    """
    Crack the byte at given index of ciphertext

    :param index:
    :param ciphertext:
    :param iv:
    :param guessed_plaintext: Guesses plaintext so for
    """
    assert index < len(ciphertext)
    ccopy = ciphertext[:]
    block_no = index // 16

    cipher_all = iv + ccopy[:(block_no+1) * 16]

    pbyten = index % 16
    if pbyten == 15:
        pass
    else:
        pbytes = guessed_plaintext[:(15-pbyten)]
        probe_block = cipher_all[-32:-16]

        fixed = probe_block[:pbyten+1]
        fxor = xor(pbytes, [16-pbyten]*len(pbytes))
        fxor = xor(probe_block[pbyten+1:], fxor)

        fixed += fxor

        cipher_all = cipher_all[:-32] + fixed + cipher_all[-16:]

    # index of probe byte is 16 before the target
    # though no calculation needed; we are prepending 16 bytes to ciphertext

    original_probe = cipher_all[index]
    succeeding_byte = -1

    for i in range(256):
        if not guessed_plaintext and i == original_probe:
            # first character we're guessing will be the padding size
            # there will be 2 succeesses in this case, and we're not interested in the case of unmodified state
            continue

        cipher_all[index] = i
        if server(cipher_all[16:], cipher_all[:16]):
            succeeding_byte = i
            break

    cracked_byte = succeeding_byte ^ original_probe ^ (16-pbyten)
    return cracked_byte

    # 0 -> 16
    # 1 -> 15
    # 14 -> 2
    # c_-1 ^ d_0 = p_0
    # (c_-1 ^ p_0 ^ 0x01) ^ d_0 = 0x01


def get_padding_size(ciphertext: bytearray, iv: bytearray):
    cipher_all = iv + ciphertext

    probe_original = cipher_all[-17]
    non_original_success = -1

    for i in range(0, 256):
        if i == probe_original:
            continue

        cipher_all[-17] = i
        if server(cipher_all[16:], cipher_all[:16]):
            non_original_success = i

    if non_original_success != -1:
        retval = non_original_success ^ probe_original ^ 0x01
        assert 1 <= retval <= 16
        return retval
    else:
        return 0x01


def game():
    iv = randombytes(16)
    ciphertext = cbc_enc(PLAINTEXTS[3], KEY, iv)
    guessed_plaintext = bytearray()

    for i in range(len(ciphertext)):
        guessing_index = len(ciphertext) - i - 1
        guessed_byte = guess(guessing_index, ciphertext, iv, guessed_plaintext)
        guessed_plaintext.insert(0, guessed_byte)

    print(guessed_plaintext)


def _game():
    ptext = b'012345678901234'

    iv = randombytes(16)
    ctext = cbc_enc(ptext, KEY, iv)
    guessed = bytearray(b'4\x01')

    cb = guess(13, ctext, iv, guessed)
    guessed.insert(0, cb)

    cb = guess(12, ctext, iv, guessed)
    guessed.insert(0, cb)

    return


def main():
    game()
    return
    test_iv = randombytes(16)

    assert is_valid_padding(b'a' * 16 + b'\x10' * 16)
    assert is_valid_padding(b'a' * 31 + b'\x01' * 1)
    assert is_valid_padding(b'a' * 30 + b'\x02' * 2)
    assert is_valid_padding(b'a' * 32) is False

    assert get_padding_size(cbc_enc(b'a' * 16, KEY, test_iv), test_iv) == 16
    assert get_padding_size(cbc_enc(b'a' * 15, KEY, test_iv), test_iv) == 1
    assert get_padding_size(cbc_enc(b'a' * 14, KEY, test_iv), test_iv) == 2
    assert get_padding_size(cbc_enc(b'a' * 10, KEY, test_iv), test_iv) == 6

    assert guess(31, cbc_enc(b'0'*16, KEY, test_iv), test_iv, b'') == 0x10
    assert guess(30, cbc_enc(b'0'*16, KEY, test_iv), test_iv, b'\x10') == 0x10
    assert guess(16, cbc_enc(b'0'*16, KEY, test_iv), test_iv, b'\x10'*15) == 0x10
    assert guess(15, cbc_enc(b'0'*16, KEY, test_iv), test_iv, b'\x10'*16) == ord('0')
    return


if __name__ == '__main__':
    main()


