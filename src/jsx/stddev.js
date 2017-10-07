import React from "react";
import {Button, Input, Row} from 'react-materialize';

module.exports = StandardDeviation;

export class StandardDeviation extends React.Component {

    constructor(props) {
        super(props)
        this.state = {value: ""}
    }

    handleChange(e) {
        this.setState({value: e.target.value});
    }

    isNumeric(n) {
        return !isNaN(parseFloat(n)) && isFinite(n);
    }

    isValid() {
        return this.parse(this.state.value) ? true : false;
    }

    save() {
        if (this.isValid()) {
            (this.props.listener || console.log)(this.parse(this.state.value));
        }
    }

    parse(numbers) {
        try {
            return self.getPointsList(this.state.value)
                .filter(self.isNumeric)
                .map(parseFloat);
        } catch (e) {
            console.debug(e);
        }
        return null;
    }

    render() {
        return (
            <Row>
                <Input
                    type="text"
                    placeholder="Example: 28.232 28.442 187.644 38.1 192.0 37"
                    onChange={(e) => this.handleChange(e)}
                />
                <Button disabled={!this.isValid()} onClick={(e) => this.save(e)}>Add Standard
                    Deviation</Button>
            </Row>
        )
    }
}
