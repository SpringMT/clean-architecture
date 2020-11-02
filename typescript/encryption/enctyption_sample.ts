import {Encryption} from "./domain/encryption";
import {ConsoleReader} from "./usecase/console_reader";
import {ConsoleWriter} from "./usecase/console_writer";

const enc = new Encryption(new ConsoleReader(), new ConsoleWriter())
enc.encryption()