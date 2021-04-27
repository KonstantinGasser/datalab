export class DataKraken {

    private APP_TOKEN: string

    constructor(app_token: string) {
        this.APP_TOKEN = app_token
        document.addEventListener("mouseover", this.onClick)
    }

    // functions for events
    // events:
    //      - element clicked
    //      - URL change
    //      - hover time over element
    //      - time per URL
    //      - referrer page
    //      - device info
    //      - 

    private onClick(event: any) {
        console.log("Mouse-Over", event)
        event.target.addEventListener("click", function() {
            console.log("Clicked: ", event)
        })
    }
}