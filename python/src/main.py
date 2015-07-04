# -*- coding: utf-8 -*-  
    
class bird(object):
    feather = True
    
class chicken(bird):
    fly = False
    def __init__(self, age):
        self.age = age
    def _getattr_(self, name):
        if name == 'adult':
            if self.age > 1.0:
                return True
            else :
                return False
        else :
            raise AttributeError(name)
summer = chicken(2)

print (summer.adult)
summer.age = 0.5

def square_sum(a, b):
    c = a ** 2 + b ** 2
    return a, b, c

print (summer.male) 

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
