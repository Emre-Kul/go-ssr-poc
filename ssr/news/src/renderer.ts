
import React from "react";
import ReactDOM from "react-dom/server";
import { App } from "./components/app";

export class Renderer {
    private fragmentName: string;
    constructor(fragmentName: string) {
        this.fragmentName = fragmentName;
    }

    public render(data: any) {
        const props = { text: data.text, type: data.type };
        const output = ReactDOM.renderToString(React.createElement(App, props as any));
        return output;
    }

}
