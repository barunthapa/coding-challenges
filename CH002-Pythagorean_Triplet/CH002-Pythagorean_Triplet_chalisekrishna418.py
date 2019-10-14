#!/bin/python
total = 1000

for first_num in xrange(1, total/3):
	for second_num in xrange(first_num + 1,total/2):
		third_num = total - first_num - second_num
		if ( (first_num * first_num) + (second_num * second_num) == (third_num * third_num) ):
			print("first_num = %d" % first_num)
			print("second_num = %d" % second_num)
			print("third_num = %d" % third_num)