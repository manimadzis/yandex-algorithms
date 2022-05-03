n = int(input())
m = [int(x) for x in input().split()]
k = int(input())
p = [int(x) for x in input().split()]

i = j = 0
diff = 10_000_001
pair = i, j

while i < n and j < k and diff > 0:
	delta = abs(m[i] - p[j])

	if diff > delta:
		diff = delta
		pair = m[i], p[j]

	if m[i] >= p[j]:
		j += 1
	else:
		i += 1
print(pair[0], pair[1])