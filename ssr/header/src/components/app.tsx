import React from "react";

export class App extends React.Component {

    public render() {
        const name = ((this.props) as any).name;
        return <div>Hello From HEADER!</div>;
    }

}
