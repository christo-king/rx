import React from "react";
import {Button, Input} from 'react-materialize';

import './style/_stddev.scss'

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
            <div className="stddev-container center">
                <div className="controls-container center">
                    <Input className="center" label="Data Points" type="text"
                           placeholder="Example: 28.232 28.442 187.644"
                           onChange={(e) => this.handleChange(e)}/>
                    <Button disabled={!this.isValid()} onClick={e => this.save(e)}>Add</Button>
                </div>
            </div>);
    }
}

module.exports = StandardDeviation;