g, s = (int(x) for x in input().split())
w = input()
string = input()

dict_w = dict()
dict_string = dict()

count = 0
occurances = 0

for c in w:
	if c not in dict_w:
		dict_w[c] = 0
	
	dict_w[c] += 1

for c in string[:g]:
	if c in dict_w:
		if c not in dict_string:
			dict_string[c] = 0 
		dict_string[c] += 1

for c in dict_w:
	if c in dict_string and dict_w[c] == dict_string[c]:
		count += 1

if count == len(dict_w):
	occurances += 1

for i in range(1, s - g + 1):
	del_c = string[i - 1]
	new_c = string[i + g - 1]
	
	if del_c in dict_w:
		if del_c not in dict_string:
			dict_string[del_c] = 0

		if dict_string[del_c] == dict_w[del_c]:
			count -= 1
		elif dict_string[del_c] - 1 == dict_w[del_c]:
			count += 1

		dict_string[del_c] -= 1

	if new_c in dict_w:
		if new_c not in dict_string:
			dict_string[new_c] = 0

		if dict_string[new_c] == dict_w[new_c]:
			count -= 1
		elif dict_string[new_c] + 1 == dict_w[new_c]:
			count += 1

		dict_string[new_c] += 1	
	
	if count == len(dict_w):
		occurances += 1
		
print(occurances)
