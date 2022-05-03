from collections import defaultdict

customers = defaultdict(lambda: defaultdict(lambda: 0))


with open("input.txt") as f:
	line = f.readline().strip()
	
	while line:
		customer, goods, count = line.split()
		count = int(count)
		
		customers[customer][goods] += count
		
		line = f.readline().strip()
		
sorted_customers = sorted(customers.keys())

for customer in sorted_customers:

	sorted_goods = sorted(customers[customer].keys())
	
	print(f"{customer}:")
	
	for goods in sorted_goods:
		print(f"{goods} {customers[customer][goods]}")
		