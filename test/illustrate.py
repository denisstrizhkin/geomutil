import matplotlib.pyplot as plt
import numpy as np

points = np.loadtxt("input.txt")
convex = np.loadtxt("conv_out.txt")
plt.plot(points[:, 0], points[:, 1], '.')
plt.plot(convex[:, 0], convex[:, 1])
plt.show()


