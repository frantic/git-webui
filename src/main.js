/**
 * @flow
 */

'use strict';

var React = require('react');
var ReactStyle = require('react-style');

function loadJSON(path, cb) {
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
    if (xhr.readyState === 4) {
      var log = JSON.parse(xhr.responseText);
      cb(log);
    }
  };
  xhr.open('GET', path);
  xhr.send();
}

var Root = React.createClass({
  getInitialState: function() {
    return { log: [] };
  },

  componentDidMount: function() {
    loadJSON('/log', this.handleLog);
  },

  handleLog: function(log) {
    this.setState({log: log});
  },

  render: function() {
    var items = this.state.log.map(
      (item) => <LogEntry key={item.sha1} sha1={item.sha1} message={item.message} />
    );
    return <pre>{items}</pre>;
  }
});

var LogEntry = React.createClass({
  render: function() {
    return (
      <div onClick={this.handleClick}>
        <span styles={[styles]}>{this.props.sha1.substr(0, 8)}</span> {this.props.message}
        {this.state.diff}
      </div>
    );
  },

  getInitialState: function() {
    return { diff: null };
  },

  handleClick: function() {
    loadJSON('/diff/' + this.props.sha1, this.handleDiff);
  },

  handleDiff: function(diff) {
    this.setState({diff: diff});
  }
});

var styles = ReactStyle({
  color: '#119911',
  fontSize: 8,
});

document.addEventListener('DOMContentLoaded', function () {
  ReactStyle.inject();
  React.render(<Root />, document.body);
});
