#!/usr/bin/env python3

import tsp

points = []
points.append(tsp.Point('p1', 40, 30))
points.append(tsp.Point('p2', 140, 80))
points.append(tsp.Point('p3', 20, 35))
points.append(tsp.Point('p4', 200, 135))

distances = tsp.calc_distance_between_points(points)
print(distances)
print()

tsp.calc_shortest_route(len(points), distances)



