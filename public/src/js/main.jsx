const React = require('react');
const ReactDOM = require('react-dom');
const request = require('superagent');
const ActivateButton = React.createClass({
	handleClick: function(e) {
		request.post('/api/command/activate').end();
	},
	render: function() {
		return (
			<button onClick={this.handleClick}>Activate</button>
		)
	}
});
const ImageUploader = React.createClass({
	getInitialState: function() {
		return {
			imgurl: "#"
		}
	},
	handleChange: function(e) {
		var reader = new FileReader();
		reader.onload = (e) => {
			this.setState({ imgurl: e.target.result });
		}
		reader.readAsDataURL(e.target.files[0]);
	},
	render: function() {
		return (
			<div>
				<input type="file" onChange={this.handleChange}/>
				<img src={this.state.imgurl} width="300" />
			</div>
		)
	}
});
const CommandMonitor = React.createClass({
	getInitialState: function() {
		return {
			log: [], ws: null
		}
	},
	componentDidMount: function() {
		var ws = new WebSocket("ws://localhost:9000/api/command/activity");
		this.setState({ ws: ws });
		ws.onmessage = (e) => {
			ws.send('0');
			console.log(e.data);
			this.setState(function(previousState) {
				var log = previousState.log.concat(e.data);
				return { log: log };
			})
		};
	},
	componentWillUnmount: function() {
		this.state.ws.close();
		console.log("umount");
	},
	render: function() {
		return (
			<div>
				CommandMonitor
				{this.state.log.map(function(line) {
					return <p>{line}</p>;
				})}
			</div>
		);
	}
});
const Hello = React.createClass({
	render: function() {
		return (
			<div>
				<div>Hello</div>
				<ActivateButton />
				<CommandMonitor />
				<ImageUploader />
				<ImageUploader />
			</div>
		);
	}
});


ReactDOM.render(
	<Hello />,
	document.getElementById("container")
);

window.onbeforeunload = function() {
	ReactDOM.unmountComponentAtNode(
		document.getElementById("container")
	);
};