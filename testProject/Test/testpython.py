import time
j = 10000
l = 0
print time.ctime()
for d in xrange(0,j):
    for i in xrange(0,j):
        l = l + 1
print time.ctime()
print l