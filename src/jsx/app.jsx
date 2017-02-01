import React from 'react'
import ReactDOM from 'react-dom'
import {Button, ButtonToolbar} from 'react-bootstrap'


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
    render() {
        return (<div>
            <form>
                <ButtonToolbar>
                    <Button>Submit New Standard Deviation</Button>
                </ButtonToolbar>
            </form>
        </div>)
    }
}

export default class App extends React.Component {
    render() {
        return (<div><NewStdDevForm/></div>);
    }
}

ReactDOM.render(<App/>, document.getElementById("main-react-container"))
