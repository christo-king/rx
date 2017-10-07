import React from "react";


export class ListStdDevs extends React.Component {
    constructor(props) {
        super(props)
        this.state = {server: props.server, standardDeviations: props.standardDeviations}
    }

    render() {
        return (
            <div>
                <Panel collapsible defaultExpanded header="Existing Standard Deviations List">
                    <Grid>
                        <Row className="show-grid">
                            <Col xs={1} md={1}><h3>ID</h3></Col>
                            <Col xs={4} md={4}><h3>Standard Deviation</h3></Col>
                            <Col xs={6} md={6}><h3>Data Point Input</h3></Col>
                        </Row>
                        {this.state.standardDeviations.map(sd => {
                            return (
                                <Row className="show-grid" key={sd.id}>
                                    <Col xs={1} md={1}>{sd.id}</Col>
                                    <Col xs={4} md={4}>{sd.answer}</Col>
                                    <Col xs={6} md={6}>{sd.points.join(', ')}</Col>
                                </Row>
                            )
                        })}
                    </Grid>
                </Panel>
            </div>
        )
    }

}