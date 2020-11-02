import {CharWriter} from "../domain/encryption";


export class ConsoleWriter implements CharWriter {
    write(str: string): void {
        console.log(str)
    }
}