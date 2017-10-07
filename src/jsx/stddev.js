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
        return (this.state.value &&
            this.state.value.length > 0 &&
            this.parse(this.state.value)) ? true : false;
    }

    save() {
        if (this.isValid()) {
            (this.props.listener || console.log)(this.parse(this.state.value));
        }
    }

    parse(numbers) {
        try {
            return numbers.split(/[, ]/)
                .filter(this.isNumeric)
                .map(parseFloat);
        } catch (e) {
            console.info(e);
        }
        return null;
    }

    render() {
        return (
            <div className="container">
                <Row>
                    <Input
                        type="text"
                        placeholder="Example: 28.232 28.442 187.644 38.1 192.0 37"
                        onChange={(e) => this.handleChange(e)}
                    />
                </Row>
                <Row>
                    <Button disabled={!this.isValid()} onClick={(e) => this.save(e)}>Add Standard
                        Deviation</Button>
                </Row>
            </div>
        )
    }
}

module.exports = StandardDeviation;