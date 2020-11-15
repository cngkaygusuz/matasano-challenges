import math


w = 32  # word size of generated things
n = 624
m = 397
r = 31
a = 0x9908b0df
u = 11
d = 0xffffffff
s = 7
b = 0x9d2c5680
t = 15
c = 0xefc60000
l = 18
f = 1812433253

_WMAX = 2**w - 1

MT = [0 for i in range(n)]
index = n+1
lower_mask = (1 << r) - 1
upper_mask = _WMAX & (~lower_mask)


def limit(val):
    """
    Limit expression into w-bit bits.
    :param val:
    :return:
    """
    return val & _WMAX


def seed(val):
    global index
    index = n
    MT[0] = val
    for i in range(1, len(MT)):
        MT[i] = f * (MT[i-1] ^ (MT[i-1] >> w-2)) + i
        MT[i] = limit(MT[i])


def twist():
    global index
    for i in range(n):
        x = MT[i] & upper_mask + (MT[(i+1) % n] & lower_mask)
        xA = x >> 1
        if x % 2 != 0:
            xA = xA ^ a

        MT[i] = MT[(i+m) % n] ^ xA

    index = 0


def extract():
    global index
    if index > n:
        raise Exception('seed pls')

    if index == n:
        twist()

    y = MT[index]
    printh(y)

    y = y ^ ((y >> u) & d)
    printh(y)

    y = y ^ ((y << s) & b)
    printh(y)

    y = y ^ ((y << t) & c)
    printh(y)

    y = y ^ (y >> l)
    printh(y)

    index += 1
    return limit(y)


def untamper(y):
    printh(y)

    y = untamper_shiftxor(y, l, 0xffffffff, 'right')  # no mask here, using all f for no op

    y = untamper_shiftxor(y, t, c, 'left')

    y = untamper_shiftxor(y, s, b, 'left')

    y = untamper_shiftxor(y, u, d, 'right')

    return y


def printh(num):
    print(f'{num:#011x}')


def printb(num):
    print(f'{num:#050b}')


def pmask(shiftcount, slot, direction):
    if direction == 'left':
        return (((1 << shiftcount)-1) << (shiftcount*slot)) & ((1 << 32)-1)
    elif direction == 'right':
        return (((1 << shiftcount)-1) << (32-(shiftcount))) >> (shiftcount*slot)

    raise Exception(f'direction?? {direction}')


def pt(val, shiftcount, slot, direction):
    return val & pmask(shiftcount, slot, direction)


def untamper_shift(val, shiftcount, direction):
    if direction == 'left':
        return val << shiftcount
    elif direction == 'right':
        return val >> shiftcount
    else:
        raise Exception(f'direction?? {direction}')


def untamper_shiftxor(val, shiftcount, andmask, direction):
    retval = val & pmask(shiftcount, 0, direction)

    for i in range(1, math.ceil(32/shiftcount)):
        decodedportion = pt(retval, shiftcount, i-1, direction)

        decodedportion = untamper_shift(decodedportion, shiftcount, direction)

        decodedportion = decodedportion & andmask

        decodedportion = decodedportion ^ pt(val, shiftcount, i, direction)

        retval = retval | decodedportion

    return retval


def main():
    seed(0)
    twist()

    initial_state = MT[:]
    reco_state = []

    for i in range(624):
        val = extract()
        reco_state.append(untamper(val))

    print(initial_state)
    print(reco_state)
    print(initial_state == reco_state)


if __name__ == '__main__':
    main()
