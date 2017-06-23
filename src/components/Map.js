
import React, { Component } from 'react';
import Point from './Point';

class Map extends Component {

  constructor() {
    super(...arguments);
    this.drawPoint = this.drawPoint.bind(this);
  }

  drawPoint(event) {
    //console.log('drawPoint');
    let numPoints = this.props.points.length;

    if (numPoints === 8) {
      alert("Maximum 8 :)");
      return;
    }

    let rect = event.target.getBoundingClientRect();
    let x = event.clientX - rect.left;
    let y = event.clientY - rect.top;

    let newID = 'p' + (numPoints + 1).toString();

    this.props.addPoint({id: newID, x: x, y: y});
  }

  render() {

    let points = [];
    let point_info = {};

    for (var i = 0; i < this.props.points.length; i++) {

      // save points to object 
      point_info[this.props.points[i].id] = this.props.points[i];

      points.push(
        <Point key={i} id={this.props.points[i].id} 
          x={this.props.points[i].x} y={this.props.points[i].y} />
      );
    }

    let lines = [];
    let line_style = {
      stroke: "#000", 
    };

    for (var j = 0; j < this.props.result.Route.length; j++) {
      if (j === this.props.result.Route.length - 1) {
        break;
      }

      let x1 =point_info[this.props.result.Route[j]].x;
      let y1 =point_info[this.props.result.Route[j]].y;
      let x2 =point_info[this.props.result.Route[j+1]].x;
      let y2 =point_info[this.props.result.Route[j+1]].y;

      lines.push(
        <line key={j} x1={x1} y1={y1} x2={x2} y2={y2} style={line_style} />
      );
    }

    return (
      <svg id="map" width="600" height="400" onClick={this.drawPoint}>
        {points}
        {lines}
      </svg>
    );
  }
}

export default Map;

