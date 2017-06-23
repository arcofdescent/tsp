import React, { Component } from 'react';
import Map from './components/Map';
import Result from './components/Result';
import './App.css';
import request from 'request';

class App extends Component {

  constructor() {
    super(...arguments);

    this.state = {
      points: [],
      result: {
        Route: [],
        Length: 0,
      },
    };

    this.addPoint = this.addPoint.bind(this);
    this.clearMap = this.clearMap.bind(this);
  }

  render() {

    let h_style = {
      marginBottom: 5,
      marginTop: 5,
    };

    return (
      <div className="App">
        <h3 style={h_style}>Traveling Salesperson problem solved using brute force</h3>
        <h4 style={h_style}>Click on the area below to set the cities (max 8)</h4>
        <Map points={this.state.points} addPoint={this.addPoint} result={this.state.result} />
        <br />
        <button onClick={this.calcRoute.bind(this)}>Calculate shortest route</button>
        <button onClick={this.clearMap}>Clear</button>
        <Result res={this.state.result} />
      </div>
    );
  }

  clearMap() {
    this.setState({points: [], result: {Route: [], Length: 0}});
  }

  addPoint(p) {
    //console.log('adding point');

    let newPoints = this.state.points.concat([p]);
    this.setState({points: newPoints});
  }

  async calcRoute() {
    let data = await this.getRoute();
    let json = JSON.parse(data);
    //console.log({json});
    this.setState({result: json});
  }

  getRoute() {
    return new Promise((resolve, reject) => {
      request.post({ 
          url: 'http://localhost:8081/', 
          form: JSON.stringify(this.state.points),
        },
        function(err, response, body) {
          if (err) {
            reject(err); return;
          }
          resolve(body);
        });
    });
  }
}

export default App;
