n = int(input())

words = []

for _ in range(n):
	words.append(input().split())

word = input()

for word_1, word_2 in words:
	if word_1 == word:
		print(word_2)
		break
	elif word_2 == word:
		print(word_1)
		break
		