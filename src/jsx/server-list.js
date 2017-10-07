
import React from "react";

export class ServerList extends React.Component {

    constructor(props) {
        super(props)
        this.state = {server: 'localhost:3000', listener: props.serverChangeListener}
    }

    handleServerChange(e) {
        console.log(e.target.value);
        this.setState({server: e.target.value})
        if (this.state.listener) {
            this.state.listener(e.target.value);
        }
    }

    render() {
        return (
            <FormGroup controlId="formControlsSelectMultiple">
                <ControlLabel>Select Server</ControlLabel>
                <FormControl componentClass="select" onChange={(e) => this.handleServerChange(e)}>
                    <option value="localhost:3000">Go</option>
                    <option value="localhost:3002">Ruby</option>
                </FormControl>
            </FormGroup>
        )
    }
}

module.exports = ServerList