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
    y = y ^ ((y >> u) & d)
    y = y ^ ((y << s) & b)
    y = y ^ ((y << t) & c)
    y = y ^ (y >> l)

    index += 1
    return limit(y)


def main():
    seed(1)
    for i in range(10):
        randint = extract()
        print(randint)


if __name__ == '__main__':
    main()
