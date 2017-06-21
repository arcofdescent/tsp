
import tsp

points = []
points.append(tsp.Point('p1', 40, 30))
points.append(tsp.Point('p2', 140, 80))
points.append(tsp.Point('p3', 20, 35))

distances = tsp.calc_distance_between_points(points)
print(distances)


