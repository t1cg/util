Object.defineProperty(exports, "__esModule", { value: true });
const stream = require("stream");
const moment = require("moment");
const fs = require("fs");
class Logger {
    // public debug(msg: string): void {
    //   const logs: Log = <Log>{ level: "debug", message: msg }
    //   this.outputLog(logs);
    // }
    // public info(msg: string): void {
    //   const logs: Log = <Log>{ level: "info", message: msg }
    //   this.outputLog(logs);
    // }
    // public trace(msg: string): void {
    //   const logs: Log = <Log>{ level: "trace", message: msg }
    //   this.outputLog(logs);
    // }
    // public time(msg: string): void {
    //   const logs: Log = <Log>{ level: "time", message: msg }
    //   this.outputLog(logs);
    // }
    // public warn(msg: string): void {
    //   const logs: Log = <Log>{ level: "warn", message: msg }
    //   this.outputLog(logs);
    // }
    // public error(msg: string): void {
    //   const logs: Log = <Log>{ level: "error", message: msg }
    //   this.outputLog(logs);
    // }
    // private outputLog(log: Log) {
    //   if (log.level == "trace" || log.level == "debug" || log.level == "info" || log.level == "time" || log.level == "warn" || log.level == "error") {
    //     console.log("LOG LEVEL: " + log.level)
    //     console.log(log.message);
    //   } else {
    //     console.error("Invalid Log Level: " + log.level);
    //   }
    // }
    constructor(encoding = 'utf8', level = 'trace') {
        this.level = level;
        this.encoding = encoding;
        this.SetLogStream(level);
    }
    // Sets logger output stream to a file
    SetLogFile(fpath, logname, level) {
        var logFile = fs.createWriteStream(fpath, { flags: 'a' });
        switch (this.level) {
            case 'trace':
                console.log("made it to trace");
                this.Trace = new Log(false, 'trace', logFile);
                this.Debug = new Log(false, 'debug', logFile);
                this.Info = new Log(false, 'info', logFile);
                this.Warn = new Log(false, 'warn', logFile);
                this.Error = new Log(false, 'error', logFile);
                break;
            case 'debug':
                this.Trace = new Log(true, 'trace', logFile);
                this.Debug = new Log(false, 'debug', logFile);
                this.Info = new Log(false, 'info', logFile);
                this.Warn = new Log(false, 'warn', logFile);
                this.Error = new Log(false, 'error', logFile);
                break;
            case 'info':
                this.Trace = new Log(true, 'trace', logFile);
                this.Debug = new Log(true, 'debug', logFile);
                this.Info = new Log(false, 'info', logFile);
                this.Warn = new Log(false, 'warn', logFile);
                this.Error = new Log(false, 'error', logFile);
                break;
            case 'warn':
                this.Trace = new Log(true, 'trace', logFile);
                this.Debug = new Log(true, 'debug', logFile);
                this.Info = new Log(true, 'info', logFile);
                this.Warn = new Log(false, 'warn', logFile);
                this.Error = new Log(false, 'error', logFile);
                break;
            default:
                this.Trace = new Log(true, 'trace', logFile);
                this.Debug = new Log(true, 'debug', logFile);
                this.Info = new Log(true, 'info', logFile);
                this.Warn = new Log(true, 'warn', logFile);
                this.Error = new Log(false, 'error', logFile);
        }
    }
    // Set logger output to the IO stream
    SetLogStream(level) {
        //TODO
        switch (this.level) {
            case 'trace': {
                console.log("made it to trace");
                this.Trace = new Log(false, 'trace', process.stdout);
                this.Debug = new Log(false, 'debug', process.stdout);
                this.Info = new Log(false, 'info', process.stdout);
                this.Warn = new Log(false, 'warn', process.stdout);
                this.Error = new Log(false, 'error', process.stdout);
                break;
            }
            case 'debug': {
                console.log("made it to debug");
                this.Trace = new Log(true, 'trace', process.stdout);
                this.Debug = new Log(false, 'debug', process.stdout);
                this.Info = new Log(false, 'info', process.stdout);
                this.Warn = new Log(false, 'warn', process.stdout);
                this.Error = new Log(false, 'error', process.stdout);
                break;
            }
            case 'info': {
                console.log("made it to info");
                this.Trace = new Log(true, 'trace', process.stdout);
                this.Debug = new Log(true, 'debug', process.stdout);
                this.Info = new Log(false, 'info', process.stdout);
                this.Warn = new Log(false, 'warn', process.stdout);
                this.Error = new Log(false, 'error', process.stdout);
                break;
            }
            case 'warn': {
                console.log("made it to warn");
                this.Trace = new Log(true, 'trace', process.stdout);
                this.Debug = new Log(true, 'debug', process.stdout);
                this.Info = new Log(true, 'info', process.stdout);
                this.Warn = new Log(false, 'warn', process.stdout);
                this.Error = new Log(false, 'error', process.stdout);
                break;
            }
            case 'error': {
                console.log("made it to default");
                this.Trace = new Log(true, 'trace', process.stdout);
                this.Debug = new Log(true, 'debug', process.stdout);
                this.Info = new Log(true, 'info', process.stdout);
                this.Warn = new Log(true, 'warn', process.stdout);
                this.Error = new Log(false, 'error', process.stdout);
                break;
            }
            default: {
                console.log("invalid log level");
            }
        }
    }
    // Set log level to input
    SetLogLevel(level) {
        this.level = level;
        this.SetLogStream(level);
    }
    Log(msg) {
        this.Trace.Output(msg);
        this.Debug.Output(msg);
        this.Info.Output(msg);
        this.Warn.Output(msg);
        this.Error.Output(msg);
    }
}
exports.default = Logger;
// code in getLogger from ../logger/logger.go
class Log {
    constructor(discard, level, stream) {
        this.level = level;
        if (discard) {
            // don't write to file
            this.out = new DiscardStream();
        }
        else {
            // write to file
            this.out = stream;
        }
        this.prefix = moment().format() + " " + level + " | ";
    }
    Output(msg) {
        this.out.write(this.prefix + msg + "\n");
    }
}
// for when we don't need to write to a file
class DiscardStream extends stream.Writable {
    _write(chunk, enc, next) {
        // do nothing
    }
}
// interface Log {
//   level: string;
//   message: string;
// }
// enum LogLevel {
//   TRACE,
//   DEBUG,
//   INFO,
//   TIME,
//   WARN,
//   ERROR,
// }
// let logHandler = new Logger();
// logHandler.debug("debug");
// logHandler.info("info");
// logHandler.warn("warn");
// logHandler.trace("trace");
// logHandler.error("error");
// logHandler.time("time");
