from collections import defaultdict

words = defaultdict(lambda: 0)

with open("input.txt") as f:
	string = f.readline().strip()
	
	while string:
		start = end = 0
		
		while end < len(string):
			if string[end] != " ":
				end += 1
			else:
				word = string[start: end]
				words[word] += 1

				while end < len(string) and string[end] == " ":
					end += 1

				start = end
		
		if end > start:
			word = string[start: end]
			words[word] += 1
			
				
		string = f.readline().strip()
	
max_count = max(words.values())
max_counted = []

for word, count in words.items():
	if count == max_count:
		max_counted.append(word)

max_counted.sort()

print(max_counted[0])
			
			