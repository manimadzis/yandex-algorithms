from collections import defaultdict

users = defaultdict(lambda: 0)

with open("inptut.txt") as f:
	line = f.readline().strip()
	
	
	while line:
		
		action, other = line.split(maxsplit=1)
		
		if action == "DEPOSIT":
			user, count = other.split()
			count = int(count)

			users[user] += count               
			
		elif action == "WITHDRAW":
			user, count = other.split()
			count = int(count)
			users[user] -= count
			
		elif action == "BALANCE":
			user = other
			
			if user in users:
				print(users[user])
			else:
				print("ERROR")
			
		elif action == "TRANSFER":
			user_from, user_to, count = other.split()
			count = int(count)
			
			users[user_from] -= count
			users[user_to] += count
		
		elif action == "INCOME":
			percents = int(other)
			
			for user in users:
				if users[user] > 0:
					users[user] = int(users[user] * (1.0 + percents / 100))
			
			
	
		line = f.readline().strip()