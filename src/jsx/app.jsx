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
        this.state = {standardDeviations: props.standardDeviations}
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
        this.state = {value: '', validationState: '', valid: true, listener: props.newStandardDeviationListener}
        this.validation = {valid: true, message: ""}

    }

    handleChange(e) {
        this.setState({value: e.target.value});
    }

    isNumeric(n) {
        return !isNaN(parseFloat(n)) && isFinite(n);
    }

    getPointsList(value) {
        return value.replace(new RegExp(',', 'g'), '').split(" ")
    }

    handleValidation() {
        let invalids = this.getPointsList(this.state.value).filter(isNaN)
        if (invalids.length > 0) {
            return "error"
        }
        return "success"
    }

    handleSubmit(e) {
        e.preventDefault()
        let self = this
        try {
            let points = self.getPointsList(this.state.value).filter(self.isNumeric).map(parseFloat)
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
            }).then((newStandardDeviation) => {
                self.setState({value: ''})
                if (self.state.listener) {
                    self.state.listener(newStandardDeviation)
                }
            });
        } catch (e) {

        }
    }

    render() {
        return (<div className="std-dev-input-form">
            <form>
                <FormGroup controlId="formBasicText" validationState={this.handleValidation()}>
                    <ControlLabel>Enter a list of numbers</ControlLabel>
                    <FormControl
                        type="text"
                        value={this.state.value}
                        placeholder="Example: 28.232 28.442 187.644 38.1 192.0 37"
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
    constructor(props) {
        super(props)
        this.state = {standardDeviations: []}
    }

    componentDidMount() {
        this.refreshList();
    }

    refreshList() {
        let self = this;
        fetch('standardDeviation').then((standardDeviations) => standardDeviations.json()).then((standardDeviations) => {
            self.setState({
                standardDeviations: standardDeviations
            });
            self.standardDeviationList.setState({standardDeviations: standardDeviations})
        });
    }

    newStandardDeviation(standardDeviation) {
        this.standardDeviationList.setState((prevState, props) => {
            prevState.standardDeviations.unshift(standardDeviation)
            return {standardDeviations: prevState.standardDeviations};
        });
    }

    render() {
        return (<div>
            <AppHeader/>
            <NewStdDevForm newStandardDeviationListener={this.newStandardDeviation.bind(this)}/>
            <ListStdDevs standardDeviations={[]}
                         ref={(r) => this.standardDeviationList = r}/>
        </div>);
    }
}

ReactDOM.render(<App/>, document.getElementById("main-react-container"))
