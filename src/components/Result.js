
import React, { Component } from 'react';

class Result extends Component {

  render() {

    let route = this.props.res.Route.join(", ");

    return(
      <div className="result">
        Shortest route: {route}<br/>
        Distance travelled: {this.props.res.Length}
      </div>
    );
  }
}

export default Result;

