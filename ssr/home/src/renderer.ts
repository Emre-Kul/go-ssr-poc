export class Renderer {
    private fragmentName: string;
    constructor(fragmentName: string) {
        this.fragmentName = fragmentName;
    }

    public render() {
        return `${this.fragmentName} will be rendered!`;
    }

}
