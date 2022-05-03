n = int(input())

keys = [int(x) for x in input().split()]

k = int(input())

seq = (int(x) for x in input().split())

for key in seq:
	keys[key - 1] -= 1

for count in keys:
	if count <= 0:
		print("YES")
	else:
		print("NO")
	
