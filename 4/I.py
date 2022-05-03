from collections import defaultdict

def find_stress(string):
	for i, x in enumerate(string):
		if x.isupper():
			return i


n = int(input())
dictionary = defaultdict(lambda: set())

for _ in range(n):
	word = input()
	
	stress = find_stress(word)
	
	dictionary[word.tolower()].add(stress)
	
text = input()
count = 0

for word in text.split():
	if word.islower():
		count += 1
	elif len([x for x in word if x.isupper()]) ==  1:
		lower_word = word.tolower()
		
		if lower_word in dictionary and find_stress(word) not in dictionary[lower_word]:
			count += 1
	else:
		count += 1
	
	
print(count)
	
	