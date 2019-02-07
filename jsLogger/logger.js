Object.defineProperty(exports, "__esModule", { value: true });
const stream = require("stream");
const moment = require("moment");
const fs = require("fs");
const encodings = ['utf8', 'json'];
const levels = ['trace', 'debug', 'info', 'warn', 'error'];
class Logger {
    constructor(encoding = 'utf8', level = 'trace') {
        // check level exists, then set it
        if (levels.indexOf(level) > -1) {
            this.level = level;
        }
        else {
            console.log('invalid level, set to default (trace)');
            this.level = 'trace';
        }
        // check encoding exists, then set it
        if (encodings.indexOf(encoding) > -1) {
            this.encoding = encoding;
        }
        else {
            console.log('invalid encoding, set to default (utf8)');
            this.encoding = 'utf8';
        }
        // set log level
        this.SetLogStream(level);
    }
    // Sets logger output stream to a file
    SetLogFile(fpath, level = this.level) {
        this.level = level;
        var logFile;
        try {
            fs.accessSync(fpath, fs.constants.W_OK);
            logFile = fs.createWriteStream(fpath, { flags: 'a' });
            // set logs based on level
            switch (this.level) {
                case 'trace': {
                    this.Trace = new Log(false, 'trace', logFile, this.encoding);
                    this.Debug = new Log(false, 'debug', logFile, this.encoding);
                    this.Info = new Log(false, 'info', logFile, this.encoding);
                    this.Warn = new Log(false, 'warn', logFile, this.encoding);
                    this.Error = new Log(false, 'error', logFile, this.encoding);
                    break;
                }
                case 'debug': {
                    this.Trace = new Log(true, 'trace', logFile, this.encoding);
                    this.Debug = new Log(false, 'debug', logFile, this.encoding);
                    this.Info = new Log(false, 'info', logFile, this.encoding);
                    this.Warn = new Log(false, 'warn', logFile, this.encoding);
                    this.Error = new Log(false, 'error', logFile, this.encoding);
                    break;
                }
                case 'info': {
                    this.Trace = new Log(true, 'trace', logFile, this.encoding);
                    this.Debug = new Log(true, 'debug', logFile, this.encoding);
                    this.Info = new Log(false, 'info', logFile, this.encoding);
                    this.Warn = new Log(false, 'warn', logFile, this.encoding);
                    this.Error = new Log(false, 'error', logFile, this.encoding);
                    break;
                }
                case 'warn': {
                    this.Trace = new Log(true, 'trace', logFile, this.encoding);
                    this.Debug = new Log(true, 'debug', logFile, this.encoding);
                    this.Info = new Log(true, 'info', logFile, this.encoding);
                    this.Warn = new Log(false, 'warn', logFile, this.encoding);
                    this.Error = new Log(false, 'error', logFile, this.encoding);
                    break;
                }
                case 'error': {
                    this.Trace = new Log(true, 'trace', logFile, this.encoding);
                    this.Debug = new Log(true, 'debug', logFile, this.encoding);
                    this.Info = new Log(true, 'info', logFile, this.encoding);
                    this.Warn = new Log(true, 'warn', logFile, this.encoding);
                    this.Error = new Log(false, 'error', logFile, this.encoding);
                }
                default: {
                    console.log('invalid log level');
                }
            }
        }
        catch (err) {
            console.error('no access!');
        }
    }
    // Set logger output to the IO stream
    SetLogStream(level) {
        // set logs base on level
        switch (this.level) {
            case 'trace': {
                this.Trace = new Log(false, 'trace', process.stdout, this.encoding);
                this.Debug = new Log(false, 'debug', process.stdout, this.encoding);
                this.Info = new Log(false, 'info', process.stdout, this.encoding);
                this.Warn = new Log(false, 'warn', process.stdout, this.encoding);
                this.Error = new Log(false, 'error', process.stdout, this.encoding);
                break;
            }
            case 'debug': {
                this.Trace = new Log(true, 'trace', process.stdout, this.encoding);
                this.Debug = new Log(false, 'debug', process.stdout, this.encoding);
                this.Info = new Log(false, 'info', process.stdout, this.encoding);
                this.Warn = new Log(false, 'warn', process.stdout, this.encoding);
                this.Error = new Log(false, 'error', process.stdout, this.encoding);
                break;
            }
            case 'info': {
                this.Trace = new Log(true, 'trace', process.stdout, this.encoding);
                this.Debug = new Log(true, 'debug', process.stdout, this.encoding);
                this.Info = new Log(false, 'info', process.stdout, this.encoding);
                this.Warn = new Log(false, 'warn', process.stdout, this.encoding);
                this.Error = new Log(false, 'error', process.stdout, this.encoding);
                break;
            }
            case 'warn': {
                this.Trace = new Log(true, 'trace', process.stdout, this.encoding);
                this.Debug = new Log(true, 'debug', process.stdout, this.encoding);
                this.Info = new Log(true, 'info', process.stdout, this.encoding);
                this.Warn = new Log(false, 'warn', process.stdout, this.encoding);
                this.Error = new Log(false, 'error', process.stdout, this.encoding);
                break;
            }
            case 'error': {
                this.Trace = new Log(true, 'trace', process.stdout, this.encoding);
                this.Debug = new Log(true, 'debug', process.stdout, this.encoding);
                this.Info = new Log(true, 'info', process.stdout, this.encoding);
                this.Warn = new Log(true, 'warn', process.stdout, this.encoding);
                this.Error = new Log(false, 'error', process.stdout, this.encoding);
                break;
            }
            default: {
                console.log("invalid log level");
            }
        }
    }
    // Get encoding
    GetEncoding() {
        return this.encoding;
    }
    // Set encoding
    SetEncoding(type) {
        if (encodings.indexOf(type) > -1) {
            this.encoding = type;
        }
        else {
            console.log('invalid encoding');
        }
    }
    // Get log level
    GetLogLevel() {
        return this.level;
    }
    // Set log level to input
    SetLogLevel(level) {
        if (levels.indexOf(level) > -1) {
            this.level = level;
            this.SetLogStream(level);
        }
        else {
            console.log('invalid level');
        }
    }
    Log(msg) {
        let encoding = this.encoding;
        this.Trace.Output(msg, encoding);
        this.Debug.Output(msg, encoding);
        this.Info.Output(msg, encoding);
        this.Warn.Output(msg, encoding);
        this.Error.Output(msg, encoding);
    }
}
exports.default = Logger;
class Log {
    constructor(discard, level, stream, encoding) {
        this.level = level;
        this.encoding = encoding;
        if (discard) {
            // don't write to file
            this.out = new DiscardStream();
        }
        else {
            // write to file
            this.out = stream;
        }
        this.timestamp = moment().format("MM-DD-YYYY h:mm:ss");
        this.prefix = this.timestamp + " " + level + " | ";
    }
    Output(msg, encoding) {
        switch (encoding) {
            case 'json': {
                this.out.write("{\"timestamp\":" + "\"" + this.timestamp + "\",\"level\":\"" + this.level + "\",\"message\":" + "\"" + msg + "\"}\n");
                break;
            }
            default: {
                this.out.write(this.prefix + msg + "\n");
            }
        }
    }
}
// for when we don't need to write to a file
class DiscardStream extends stream.Writable {
    _write(chunk, enc, next) {
        // do nothing
    }
}
