
import React, { Component } from 'react';

class Result extends Component {

  render() {

    let route = this.props.res.route.join(", ");

    return(
      <div className="result">
        Shortest route: {route}<br/>
        Distance travelled: {this.props.res.length}<br/>
				Calculation time: {this.props.res.Duration}
      </div>
    );
  }
}

export default Result;

