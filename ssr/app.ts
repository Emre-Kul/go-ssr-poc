import { IPC } from "node-ipc";
import { Footer } from "./footer/index";
import { Header } from "./header/index";
import { News } from "./news/index";

export class App {
    private ipc: any;
    constructor() {
        this.ipc = new IPC();
    }

    public run() {
        this.start();
        this.ipc.config.logger = () => {};
        this.ipc.server.start();
    }

    private registerFragments(server: any) {
        const news = new News();
        const footer = new Footer();
        const header = new Header();

        server.on("NEWS-CONTENT", (data: any, socket: any) => {
            server.emit(socket, "NEWS-CONTENT", news.content(data));
        });

        server.on("FOOTER-CONTENT", (data: any, socket: any) => {
            server.emit(socket, "FOOTER-CONTENT", footer.content(data));
        });

        server.on("HEADER-CONTENT", (data: any, socket: any) => {
            server.emit(socket, "HEADER-CONTENT", header.content(data));
        });
    }

    private start() {
        this.ipc.serveNet( 8000, () => {
            this.registerFragments(this.ipc.server);
        });
    }
}
