import csv
with open('baby-names.csv', 'rb') as names:
    name_reader = csv.reader(names, delimiter = ',')
    headers = name_reader[0];
    print(headers)
