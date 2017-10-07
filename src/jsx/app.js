import React from 'react'
import {Row} from 'react-materialize';

import service from './services';
import AppHeader from './header'
import ServerList from './server-list'
import StandardDeviation from './stddev'
// import ListStdDevs from './stddev-list'
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
        return (<div className="container">
                <Row>
                    <AppHeader/>
                </Row>
                <Row>
                    <ServerList listener={this.serverChanged.bind(this)}/>
                    <StandardDeviation listener={this.saveStandardDeviation.bind(this)}
                                       server={this.state.server} />
                </Row>
                {/*<ListStdDevs standardDeviations={[]} server={this.state.server}*/}
                {/*ref={(r) => this.standardDeviationList = r}/>*/}
            </div>
        );
    }
}

module.exports = App;
