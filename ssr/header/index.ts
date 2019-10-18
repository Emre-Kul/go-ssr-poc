import { Renderer } from "./src/renderer";

export class Header {
    private renderer: Renderer;

    constructor() {
        this.renderer = new Renderer("HEADER-FRAGMENT");
    }

    public content(data: any) {
        return this.renderer.render(data);
    }

}