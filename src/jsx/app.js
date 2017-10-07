import React from 'react'
import AppHeader from './header'
// import ServerList from './server-list'
// import NewStdDevForm from './stddev'
// import ListStdDevs from './stddev-list'


class App extends React.Component {
    constructor(props) {
        super(props)
        this.state = {server: 'localhost:3000', standardDeviations: []}
        this.refetch()
    }

    refetch() {
        let self = this;
        fetch('http://' + this.state.server + '/standardDeviation').then((sds) => sds.json()).then((sds) => {
            self.setState({
                standardDeviations: sds || []
            });
        });
    }

    serverChanged(server) {
        this.setState({server: server})
        this.refetch()
    }

    render() {
        return (<div>
            <AppHeader/>
            {/*<div className="std-dev-input-form">*/}
                {/*<ServerList serverChangeListener={this.serverChanged.bind(this)}/>*/}
                {/*<NewStdDevForm newStandardDeviationListener={this.newStandardDeviation.bind(this)}*/}
                               {/*server={this.state.server}*/}
                               {/*ref={(r) => this.newStandardDeviationList = r}/>*/}
            {/*</div>*/}
            {/*<ListStdDevs standardDeviations={[]} server={this.state.server}*/}
                         {/*ref={(r) => this.standardDeviationList = r}/>*/}
        </div>);
    }
}

module.exports = App;
