import numpy
from scipy.optimize import curve_fit
import pandas

# Read file, split x and y columns
data = pandas.read_csv('1exp.out', header=0, names=['x','y'])
np_data = data.to_numpy()
x = np_data[:,0]
y = np_data[:,1]
print(x)

# Fit curve
[a, b], res1 = curve_fit(lambda x1,a,b: a * numpy.exp(b*x1), x, y)
print(f"a:{a} b:{b}")

# Calculate 
print(f"")

