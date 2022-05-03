from collections import defaultdict

n, k = (int(x) for x in input().split())

s = input()

i, j = 0, 0
max_len = 0
max_index = 0
d = defaultdict(lambda : 0)

while j < n:
	while j < n and d[s[j]] <	 k:
		d[s[j]] += 1
		j += 1
		
	cur_len = j - i
	if cur_len > max_len:
		max_len = cur_len
		max_index = i

	d[s[i]] -= 1
	i += 1

max_index += 1

print(max_len, max_index)
