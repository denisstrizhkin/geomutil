import matplotlib.pyplot as plt
import numpy as np

points = np.loadtxt("input.txt")
print(points)
point_x = points[:, 0]
point_y = points[:, 1]
plt.plot(point_x, point_y, '.')
plt.show()


