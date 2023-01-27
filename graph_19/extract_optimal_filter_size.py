with open('crews_full_output.csv') as f:
    d = f.readlines()

data = [x.split(',') for x in d]

qualifies = lambda x: x[2]<(x[1]/10)

first_filter = [x for x in data if qualifies(x)]
