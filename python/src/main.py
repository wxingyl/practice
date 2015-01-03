# -*- coding: utf-8 -*-  
# 用空格的缩进来表示隶属关系
from src.brid import Chicken, Bird


def square_sum(a, b):
    c = a ** 2 + b ** 2
    return a, b, c


def change_int(x):
    x += 1
    return x


def change_list(data):
    data[1] = 3
    return data


if __name__ == '__main__':
    summer = Chicken()
    print summer.have_feather
    print summer.move(5, 8)
    summer = Bird()
    print 'after move:', summer.move(5, 8)
    idx = range(5)
    print idx
    for a in range(10):
        print a ** 2
    i = 1
    while i < 10:
        print i
        i += 1
    for a in range(7):
        if a == 3:
            continue
        print a
    for a in range(6):
        if a == 4:
            break
        print a
    print square_sum(4, 3)