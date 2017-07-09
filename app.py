
import sys
from flask import Flask, render_template, request, make_response
from pprint import pprint
import json
import time

sys.path.append('./python/tsp')
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

    post_data = request.data.decode("utf-8")
    points = json.loads(post_data)
    points = [tsp.Point(p['id'], p['x'], p['y']) for p in points['points']]

    start_time = time.time()
    distances = tsp.calc_distance_between_points(points)
    route_info = tsp.calc_shortest_route(len(points), distances)
    route_info['duration'] = "%.4f" % (time.time() - start_time)

    res = make_response(json.dumps(route_info))
    res.headers['Access-Control-Allow-Origin'] = "*"

    return res
