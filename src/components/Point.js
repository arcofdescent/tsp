import React, { Component } from 'react';

class Point extends Component {

  constructor() {
    super(...arguments);
    this.radius = 10;
  }

  render() {

    let style = {
      stroke: '#000',
    };

    return (
      <g>
        <text x={this.props.x - this.radius} y={this.props.y - this.radius - 2}>{this.props.id}</text>
        <circle id={this.props.id} cx={this.props.x} cy={this.props.y} r={this.radius} style={style} />
      </g>
    );
  }

}

export default Point;

