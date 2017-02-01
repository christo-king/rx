import React from 'react'
import ReactDOM from 'react-dom'
import {Button, ButtonToolbar, Jumbotron, FormGroup, ControlLabel, FormControl, HelpBlock} from 'react-bootstrap'

export class AppHeader extends React.Component {
    render() {
        return (
            <Jumbotron>
                <h2>Standard Deviation testing page</h2>
                <br/>
            </Jumbotron>)
    }
}

export class ListStdDevs extends React.Component {
    componentDidMount() {
        fetch('standardDeviation').then((standardDeviations) => standardDeviations.json()).then((standardDeviations) => this.setState({
            standardDeviations: standardDeviations
        }));
    }

    render() {
        return (
            <div>Current List : {this.state}</div>
        )
    }

}
export class NewStdDevForm extends React.Component {
    constructor(props) {
        super(props)
        this.state = {value: ''}
    }

    getValidationState() {
        const length = this.state.value.length;
        if (length > 10) return 'success';
        else if (length > 5) return 'warning';
        else if (length > 0) return 'error';
    }

    handleChange(e) {
        this.setState({value: e.target.value});
    }

    render() {
        return (<div>
            <form>
                <FormGroup
                    controlId="formBasicText"
                    validationState={this.getValidationState()}
                >
                    <ControlLabel>Enter a list of numbers</ControlLabel>
                    <FormControl
                        type="text"
                        value={this.state.value}
                        placeholder="Example: 28 28 187 38 192 37"
                        onChange={this.handleChange}
                    />
                    <FormControl.Feedback />
                    <HelpBlock>Validation is based on string length.</HelpBlock>
                </FormGroup>
                <ButtonToolbar>
                    <Button>Add Standard Deviation</Button>
                </ButtonToolbar>
            </form>
        </div>)
    }
}

export default class App extends React.Component {
    render() {
        return (<div><AppHeader/><NewStdDevForm/><br/><ListStdDevs/></div>);
    }
}

ReactDOM.render(<App/>, document.getElementById("main-react-container"))
