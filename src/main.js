/* global XMLHttpRequest, document */

var React = require('react');

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
        {this.props.sha1} {this.props.message}
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

document.addEventListener('DOMContentLoaded', function () {
  React.render(<Root />, document.body);
});
