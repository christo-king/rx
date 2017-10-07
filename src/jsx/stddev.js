import React from "react"


export class NewStdDevForm extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            server: props.server,
            value: '',
            validationState: '',
            valid: true,
            listener: props.newStandardDeviationListener
        }
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
            this.validation.valid = false
            console.log("invalid");
            return "error"
        }
        this.validation.valid = true
        return "success"
    }

    handleSubmit(e) {
        e.preventDefault()
        let self = this
        try {
            let points = self.getPointsList(this.state.value).filter(self.isNumeric).map(parseFloat)
            let postdata = JSON.stringify({points: points})
            var request = new Request('http://' + this.state.server + '/standardDeviation', {
                method: 'POST',
                mode: 'cors',
                body: postdata
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
        return (
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
                    <Button disabled={!this.validation.valid} onClick={(e) => this.handleSubmit(e)}>Add Standard
                        Deviation</Button>
                </ButtonToolbar>
            </form>
        )
    }
}
