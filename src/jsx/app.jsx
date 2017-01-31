import React from 'react'
import ReactDOM from 'react-dom'
import {Button, ButtonToolbar} from 'react-bootstrap'

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
