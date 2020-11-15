import time
import random
from challenge21 import seed, extract

orig_seed = None


def routine():
    global orig_seed
    time.sleep(random.randint(0, 60))
    orig_seed = int(time.time())
    seed(orig_seed)
    time.sleep(random.randint(0, 60))
    return extract()


def main():
    val = routine()

    start = int(time.time()) - 180

    for i in range(180):
        sd = start+i
        seed(sd)
        test = extract()

        if val == test:
            print(sd, orig_seed, orig_seed == sd)
            return
    print('???')


if __name__ == '__main__':
    main()
