import {CharReader} from "./char_reader";
import {CharWriter} from "./char_writer";

export class Encryption {
    reader: CharReader
    writer: CharWriter
    constructor(reader: CharReader, writer: CharWriter) {
        this.reader = reader
        this.writer = writer
    }

    encryption() {
        const input = this.reader.read()
        const encrypted = this.translate(input)
        this.writer.write(encrypted)
    }

    private translate(input: string): string {
        return input
    }
}
