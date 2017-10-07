import React from 'react'
import {Row} from 'react-materialize';

import service from './services';
import AppHeader from './header'
import ServerList from './server-list'
import StandardDeviation from './stddev'
import StandardDeviationList from './stddev-list'
import './style/main.scss'

class App extends React.Component {
    constructor(props) {
        super(props)
        this.state = {server: 'localhost:3000', standardDeviations: []}
        this.refetch()
    }

    refetch() {
        service.listStandardDeviations().then(sds => {
            this.setState({
                standardDeviations: sds || []
            });
        });
    }

    serverChanged(server) {
        this.setState({server: server})
        this.refetch()
    }

    saveStandardDeviation(numberList) {
        service.saveStandardDeviation(numberList).then(sd => {
            console.log(sd);
            this.refetch();
        }).catch(console.error)
    }

    render() {
        return (<div className="container center content-container">
                <Row>
                    <AppHeader/>
                </Row>
                <Row>
                    <ServerList listener={this.serverChanged.bind(this)}/>
                </Row>
                <Row>
                    <StandardDeviation listener={this.saveStandardDeviation.bind(this)}
                                       server={this.state.server}/>
                </Row>
                <StandardDeviationList standardDeviations={this.state.standardDeviations}/>
            </div>
        );
    }
}

module.exports = App;
