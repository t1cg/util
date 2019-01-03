var Logger = /** @class */ (function () {
    function Logger() {
    }
    Logger.prototype.debug = function (msg) {
        var logs = { level: "debug", message: msg };
        this.outputLog(logs);
    };
    Logger.prototype.info = function (msg) {
        var logs = { level: "info", message: msg };
        this.outputLog(logs);
    };
    Logger.prototype.trace = function (msg) {
        var logs = { level: "trace", message: msg };
        this.outputLog(logs);
    };
    Logger.prototype.time = function (msg) {
        var logs = { level: "time", message: msg };
        this.outputLog(logs);
    };
    Logger.prototype.warn = function (msg) {
        var logs = { level: "warn", message: msg };
        this.outputLog(logs);
    };
    Logger.prototype.error = function (msg) {
        var logs = { level: "error", message: msg };
        this.outputLog(logs);
    };
    Logger.prototype.outputLog = function (log) {
        if (log.level == "trace" || log.level == "debug" || log.level == "info" || log.level == "time" || log.level == "warn" || log.level == "error") {
            console.log("LOG LEVEL: " + log.level);
            console.log(log.message);
        }
        else {
            console.error("Invalid Log Level: " + log.level);
        }
    };
    return Logger;
}());
// enum LogLevel {
//   TRACE,
//   DEBUG,
//   INFO,
//   TIME,
//   WARN,
//   ERROR,
// }
var logHandler = new Logger();
logHandler.debug("debug");
logHandler.info("info");
logHandler.warn("warn");
logHandler.trace("trace");
logHandler.error("error");
logHandler.time("time");
