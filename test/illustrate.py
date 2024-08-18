import matplotlib.pyplot as plt
import numpy as np

points = np.loadtxt("a.txt")
convex = np.loadtxt("conv_out.txt")
plt.plot(points[:, 0], points[:, 1], '.')
plt.plot(np.append(convex[:, 0], convex[0, 0]), np.append(convex[:, 1], convex[0, 1]))
plt.show()


