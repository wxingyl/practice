# -*- coding: utf-8 -*-  
#用空格的缩进来表示隶属关系
i = 1
x = 1
if i > 0:
	x = x + 1
print x

i = 1
if i > 0:
	print 'positive i'
	i = i + 1
elif i == 0:
	print 'i is 0'
	i = i * 10
else:
	print 'negative i'
	i = i - 1
print 'new i:', i


i = 5
if i > 1:
	print 'i bigger than i'
	print 'good'
        if i > 2:
	    	print 'i bigger than 2'
	    	print 'even better'
	    	
i = 5
if i > 1:
	print 'i bigger than i'
	print 'good'
	if i > 2:
		print 'i bigger than 2'
		print 'even better'
		
idx = range(5)
print idx


for a in range(10):
	print a**2

i = 1
while i < 10:
	print i
	i = i + 1
	
for a in range(7):
	if a == 3:
		continue
	print a
	
for a in range(6):
	if a == 4:
		break
	print a	
	
def square_sum(a, b):
	c = a**2 + b**2
	return a, b, c

a = 4
b = 3
print square_sum(4,3)

a = 2
def change_int(x):
	x = x + 1
	return x
print change_int(x)
print a

b = [2, 4, 6]
def change_list(data):
	data[1] = 3
	return data
print change_list(b)
print b

class Bird(object):
     have_feather = True
     way_of_reproduction = 'egg'
     def move(self, dx, dy):
          position = [0, 0]
          position[0] = position[0] + dx
          position[1] = position[1] + dy
          return position
summer = Bird()
print 'after move:',summer.move(5, 8)

class Chicken(Bird):
     way_of_move = 'walk'
     possible_in_KFC = 'True'
summmer = Chicken()
print summer.have_feather
print summer.move(5, 8)

