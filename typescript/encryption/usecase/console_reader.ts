import {CharReader} from "../domain/encryption";

export class ConsoleReader implements CharReader{
    read(): string {
        return "dummy string"
    }
}