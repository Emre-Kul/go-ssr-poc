import { Renderer } from "./src/renderer";

export class Footer {
    private renderer: Renderer;

    constructor() {
        this.renderer = new Renderer("FOOOTER-FRAGMENT");
    }

    public content(data: any) {
        return this.renderer.render(data);
    }

}
