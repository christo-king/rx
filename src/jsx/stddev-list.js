import React from "react";

import {Table} from 'react-materialize';

export class StandardDeviationList extends React.Component {
   render() {
      return (
         <Table className="stddev-list-container">
            <thead>
            <tr>
               <th>Created</th>
               <th>Standard Deviation</th>
               <th>Data Point Input</th>
            </tr>
            </thead>
            <tbody>
            {this.props.standardDeviations.map(sd => {
               return (
                  <tr key={sd.id}>
                     <td>{(new Date(sd.created)).toDateString()}</td>
                     <td>{sd.answer}</td>
                     <td>{sd.points.join(', ')}</td>
                  </tr>)
            })}
            </tbody>
         </Table>
      );
   }
}

module.exports = StandardDeviationList;