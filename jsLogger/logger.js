"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
exports.__esModule = true;
var stream = require("stream");
var moment = require("moment");
var fs = require("fs");
var Logger = /** @class */ (function () {
    function Logger(level) {
        if (level === void 0) { level = 'trace'; }
        this.level = level;
        this.encoding = "";
        // this.encoding = encoding
        this.SetLogStream();
        // this.SetEncoding('yolo')
        // console.log(encoding)
    }
    // Sets logger output stream to a file
    Logger.prototype.SetLogFile = function (fpath, logname, level) {
        var logFile = fs.createWriteStream(fpath, { flags: 'a' });
        switch (this.level) {
            case 'trace':
                console.log("made it to trace");
                this.Trace = new Log(false, 'trace', logFile, this.encoding);
                this.Debug = new Log(false, 'debug', logFile, this.encoding);
                this.Info = new Log(false, 'info', logFile, this.encoding);
                this.Warn = new Log(false, 'warn', logFile, this.encoding);
                this.Error = new Log(false, 'error', logFile, this.encoding);
                break;
            case 'debug':
                this.Trace = new Log(true, 'trace', logFile, this.encoding);
                this.Debug = new Log(false, 'debug', logFile, this.encoding);
                this.Info = new Log(false, 'info', logFile, this.encoding);
                this.Warn = new Log(false, 'warn', logFile, this.encoding);
                this.Error = new Log(false, 'error', logFile, this.encoding);
                break;
            case 'info':
                this.Trace = new Log(true, 'trace', logFile, this.encoding);
                this.Debug = new Log(true, 'debug', logFile, this.encoding);
                this.Info = new Log(false, 'info', logFile, this.encoding);
                this.Warn = new Log(false, 'warn', logFile, this.encoding);
                this.Error = new Log(false, 'error', logFile, this.encoding);
                break;
            case 'warn':
                this.Trace = new Log(true, 'trace', logFile, this.encoding);
                this.Debug = new Log(true, 'debug', logFile, this.encoding);
                this.Info = new Log(true, 'info', logFile, this.encoding);
                this.Warn = new Log(false, 'warn', logFile, this.encoding);
                this.Error = new Log(false, 'error', logFile, this.encoding);
                break;
            default:
                this.Trace = new Log(true, 'trace', logFile, this.encoding);
                this.Debug = new Log(true, 'debug', logFile, this.encoding);
                this.Info = new Log(true, 'info', logFile, this.encoding);
                this.Warn = new Log(true, 'warn', logFile, this.encoding);
                this.Error = new Log(false, 'error', logFile, this.encoding);
        }
    };
    // Set logger output to the IO stream
    Logger.prototype.SetLogStream = function () {
        switch (this.level) {
            case 'trace': {
                console.log("made it to trace");
                this.Trace = new Log(false, 'trace', process.stdout, this.encoding);
                this.Debug = new Log(false, 'debug', process.stdout, this.encoding);
                this.Info = new Log(false, 'info', process.stdout, this.encoding);
                this.Warn = new Log(false, 'warn', process.stdout, this.encoding);
                this.Error = new Log(false, 'error', process.stdout, this.encoding);
                break;
            }
            case 'debug': {
                console.log("made it to debug");
                this.Trace = new Log(true, 'trace', process.stdout, this.encoding);
                this.Debug = new Log(false, 'debug', process.stdout, this.encoding);
                this.Info = new Log(false, 'info', process.stdout, this.encoding);
                this.Warn = new Log(false, 'warn', process.stdout, this.encoding);
                this.Error = new Log(false, 'error', process.stdout, this.encoding);
                break;
            }
            case 'info': {
                console.log("made it to info");
                this.Trace = new Log(true, 'trace', process.stdout, this.encoding);
                this.Debug = new Log(true, 'debug', process.stdout, this.encoding);
                this.Info = new Log(false, 'info', process.stdout, this.encoding);
                this.Warn = new Log(false, 'warn', process.stdout, this.encoding);
                this.Error = new Log(false, 'error', process.stdout, this.encoding);
                break;
            }
            case 'warn': {
                console.log("made it to warn");
                this.Trace = new Log(true, 'trace', process.stdout, this.encoding);
                this.Debug = new Log(true, 'debug', process.stdout, this.encoding);
                this.Info = new Log(true, 'info', process.stdout, this.encoding);
                this.Warn = new Log(false, 'warn', process.stdout, this.encoding);
                this.Error = new Log(false, 'error', process.stdout, this.encoding);
                break;
            }
            case 'error': {
                console.log("made it to default");
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
    };
    // Set log level to input
    Logger.prototype.SetLogLevel = function (level) {
        this.level = level;
        this.SetLogStream();
    };
    // Set output encoding
    Logger.prototype.SetEncoding = function (encoding) {
        this.encoding = encoding;
    };
    Logger.prototype.Log = function (msg) {
        var encoding = this.encoding;
        this.Trace.Output(msg, encoding);
        this.Debug.Output(msg, encoding);
        this.Info.Output(msg, encoding);
        this.Warn.Output(msg, encoding);
        this.Error.Output(msg, encoding);
    };
    return Logger;
}());
exports["default"] = Logger;
var Log = /** @class */ (function () {
    function Log(discard, level, stream, encoding) {
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
        this.prefix = moment().format() + " " + level + " | ";
    }
    Log.prototype.Output = function (msg, encoding) {
        switch (encoding) {
            case 'json': {
                this.out.write("ENCODING: JSON \n");
                this.out.write("{\"prefix\":" + "\"" + this.prefix + "\",\"message\":" + "\"" + msg + "\"}\n");
                break;
            }
            default: {
                this.out.write("ENCODING: " + this.encoding + "\n");
                this.out.write(this.prefix + msg + "\n");
            }
        }
    };
    return Log;
}());
// for when we don't need to write to a file
var DiscardStream = /** @class */ (function (_super) {
    __extends(DiscardStream, _super);
    function DiscardStream() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    DiscardStream.prototype._write = function (chunk, enc, next) {
        // do nothing
    };
    return DiscardStream;
}(stream.Writable));
