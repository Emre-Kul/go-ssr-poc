import { Renderer } from "./src/renderer";

export class News {
    private renderer: Renderer;
    
    constructor() {
        this.renderer = new Renderer("HOME-FRAGMENT");
    }

    public content(data: any) {
        return this.renderer.render(data);
    }
}

