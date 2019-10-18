import React from "react";

export class App extends React.Component {

    public render() {
        const text = ((this.props) as any).text;
        const type = ((this.props) as any).type;
        return <div> 
            <h1>NEWS</h1>
            <p>Text : {text} <b>{type}</b></p>
        </div>;
    }

}
