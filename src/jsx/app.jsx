import React from 'react'
import ReactDOM from 'react-dom'
import {
    Button,
    ButtonToolbar,
    Jumbotron,
    FormGroup,
    ControlLabel,
    FormControl,
    HelpBlock,
    Panel,
    Grid, Row, Col
} from 'react-bootstrap'

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
    constructor(props) {
        super(props)
        this.state = {standardDeviations: []}
    }

    componentDidMount() {
        let self = this;
        fetch('standardDeviation').then((standardDeviations) => standardDeviations.json()).then((standardDeviations) => this.setState({
            standardDeviations: standardDeviations
        }));
    }

    render() {
        return (
            <div>
                <Panel collapsible defaultExpanded header="Existing Standard Deviations List">
                    <Grid>
                        <Row className="show-grid">
                            <Col xs={3} md={3}><h3>ID</h3></Col>
                            <Col xs={3} md={3}><h3>Standard Deviation</h3></Col>
                            <Col xs={6} md={6}><h3>Data Point Input</h3></Col>
                        </Row>
                        {this.state.standardDeviations.map((sd) => {
                            return (
                                <Row className="show-grid" key={sd.id}>
                                    <Col xs={3} md={3}>{sd.id}</Col>
                                    <Col xs={3} md={3}>{sd.answer}</Col>
                                    <Col xs={6} md={6}>{sd.points.join(', ')}</Col>
                                </Row>
                            )
                        })}
                    </Grid>
                </Panel>
            </div>
        )
    }

}
export class NewStdDevForm extends React.Component {
    constructor(props) {
        super(props)
        this.state = {value: '', validationState: ''}
    }

    handleChange(e) {
        this.setState({value: e.target.value});
    }

    handleSubmit(e) {
        e.preventDefault()
        console.log("Submitting " + this.state.value);
        let points = this.state.value.split(" ").map(parseFloat)
        let postdata = {points: points}

        var request = new Request('/standardDeviation', {
            method: 'POST',
            mode: 'same-origin',
            headers: new Headers({
                'Content-Type': 'application/json'
            }),
            body: JSON.stringify(postdata)
        });
        fetch(request).then((response) => {
            return response.json();
        }).then(() => {
            this.state.value = ''
        });
    }

    render() {
        return (<div>
            <form>
                <FormGroup
                    controlId="formBasicText"
                >
                    <ControlLabel>Enter a list of numbers</ControlLabel>
                    <FormControl
                        type="text"
                        value={this.state.value}
                        placeholder="Example: 28 28 187 38 192 37"
                        onChange={(e) => this.handleChange(e)}
                    />
                    <FormControl.Feedback />
                </FormGroup>
                <ButtonToolbar>
                    <Button onClick={(e) => this.handleSubmit(e)}>Add Standard Deviation</Button>
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
