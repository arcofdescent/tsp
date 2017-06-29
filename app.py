
import sys
from flask import Flask, render_template, request, make_response
from pprint import pprint
import json

sys.path.append('./python/tsp')
pprint(sys.path)
import tsp

app = Flask(__name__)

@app.route('/')
def index():
    return render_template("index.html")

@app.route('/dev')
def index_dev():
    return render_template("index_dev.html")

@app.route('/get_shortest_route', methods=['POST'])
def get_shortest_route():

    pprint(request.data)
    post_data = request.data.decode("utf-8")
    pprint(post_data)
    points = json.loads(post_data)
    pprint(points)
    
    points = [tsp.Point(p['id'], p['x'], p['y']) for p in points['points']]
    distances = tsp.calc_distance_between_points(points)
    print("distances: %s" % distances)

    shortest_route_length = tsp.calc_shortest_route(len(points), distances)

    ret = { 'length': shortest_route_length }
    ret_json = json.dumps(ret)
    pprint(ret_json)
        
    res = make_response("sdf")
    res.headers['Access-Control-Allow-Origin'] = "*"

    return res
