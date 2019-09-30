import { IPC } from "node-ipc";
import { Renderer } from "./src/renderer";

class App {
    private renderer: Renderer;
    private ipc: any;
    constructor() {
        this.renderer = new Renderer("HOME-FRAGMENT");
        this.ipc = new IPC();
        this.start();
    }

    public run() {
        this.ipc.server.start();
    }

    private start() {
        this.ipc.serve( () => {
            this.ipc.server.on("message", (data: any, socket: any) => {
                this.ipc.server.emit(
                    socket,
                    "message",
                    this.renderer.render(),
                );
            });
        });
    }

}
const app = new App();
app.run();
