import React from "react";

import {Table} from 'react-materialize';

export class StandardDeviationList extends React.Component {
    render() {
        return (
            <div>
                <Table>
                    <thead>
                    <th>ID</th>
                    <th>Standard Deviation</th>
                    <th>Data Point Input</th>
                    </thead>
                    <tbody>{this.props.standardDeviations.map(sd => {
                        return (
                            <tr className="show-grid" key={sd.id}>
                                <td>{sd.id}</td>
                                <tr>{sd.answer}</tr>
                                <tr>{sd.points.join(', ')}</tr>
                            </tr>)
                    })}
                    </tbody>
                </Table>
            </div>
        );
    }

}

module.exports = StandardDeviationList;