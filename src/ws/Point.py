
from itertools import permutations, combinations

class Point:
    pass

class Distance:
    pass

def get_distance(p1, p2, distances):
    for d in distances:
        if d.id1 == p1 and d.id2 == p2:
            return d.distance

    return 0

def get_route_length(route_pairs, distances):
    dst = 0
    
    for pair in route_pairs:
        dst += get_distance(pair[0], pair[1], distances)

    return dst

