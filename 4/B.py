from collections import defaultdict


words = defaultdict(lambda: 0)

with open("input.txt") as f:

	string = f.readline().strip()
	while string:
		start = end = 0

		while end < len(string):
			if string[end] == " ":
				word = string[start: end]
				
				print(words[word], end=" ")
				words[word] += 1
				
				while end < len(string) and string[end] == " ":
					end += 1
				
				start = end
			else:
				end += 1

		if end > start:
			word = string[start: end]
			print(words[word], end=" ")

		string = f.readline().strip()
