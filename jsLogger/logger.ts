class Logger {

  public debug(msg: string): void {
    const logs: Log = <Log>{ level: "debug", message: msg }
    this.outputLog(logs);
  }
  public info(msg: string): void {
    const logs: Log = <Log>{ level: "info", message: msg }
    this.outputLog(logs);
  }
  public trace(msg: string): void {
    const logs: Log = <Log>{ level: "trace", message: msg }
    this.outputLog(logs);
  }
  public time(msg: string): void {
    const logs: Log = <Log>{ level: "time", message: msg }
    this.outputLog(logs);
  }
  public warn(msg: string): void {
    const logs: Log = <Log>{ level: "warn", message: msg }
    this.outputLog(logs);
  }
  public error(msg: string): void {
    const logs: Log = <Log>{ level: "error", message: msg }
    this.outputLog(logs);
  }
  private outputLog(log: Log) {
    if (log.level == "trace" || log.level == "debug" || log.level == "info" || log.level == "time" || log.level == "warn" || log.level == "error") {
      console.log("LOG LEVEL: " + log.level)
      console.log(log.message);
    } else {
      console.error("Invalid Log Level: " + log.level);

    }
  }
}

interface Log {
  level: string;
  message: string;
}



// enum LogLevel {
//   TRACE,
//   DEBUG,
//   INFO,
//   TIME,
//   WARN,
//   ERROR,
// }

let logHandler = new Logger();
logHandler.debug("debug");
logHandler.info("info");
logHandler.warn("warn");
logHandler.trace("trace");
logHandler.error("error");
logHandler.time("time");
