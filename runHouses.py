import pandas as pd
import time

N = 100
start_time = time.time()

with open('housesOutputPy.txt', 'wt') as outfile:
    for i in range(N):
        houses = pd.read_csv("housesInput.csv")
        outfile.write(houses.describe().to_string(header=True, index=True))
        outfile.write("\n")

end_time = time.time()
print("Processing time: {} seconds".format(end_time - start_time))





