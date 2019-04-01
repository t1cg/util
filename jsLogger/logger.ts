import * as stream from 'stream';
import * as moment from 'moment';
import * as fs from 'fs';

const encodings = ['utf8', 'json'];
const levels = ['trace', 'debug', 'info', 'warn', 'error'];

export default class Logger {
  private Trace: Log;
  private Debug: Log;
  private Info: Log;
  private Warn: Log;
  private Error: Log;

  private level: string;
  private encoding: string;

  constructor(encoding = 'utf8', level = 'trace') {
    // check level exists, then set it
    if (levels.indexOf(level) > -1) {
      this.level = level;
    } else {
      console.log('invalid level, set to default (trace)');
      this.level = 'trace';
    }
    // check encoding exists, then set it
    if (encodings.indexOf(encoding) > -1) {
      this.encoding = encoding;
    } else {
      console.log('invalid encoding, set to default (utf8)');
      this.encoding = 'utf8';
    }
    // set log level
    this.SetLogStream(level);
  }

  // Sets logger output stream to a file
  public SetLogFile(fpath: string, level = this.level): void {
    if (levels.indexOf(level) < 0) {
      console.log('invalid level');
      return;
    }

    this.level = level;

    var logFile: fs.WriteStream;

    try {
      fs.accessSync(fpath, fs.constants.W_OK);
      logFile = fs.createWriteStream(fpath, { flags: 'a' });

      // set logs based on level
      switch (this.level) {
        case 'trace': {
          this.Trace = new Log(false, 'trace', logFile, this.encoding);
          this.Debug = new Log(true, 'debug', logFile, this.encoding);
          this.Info = new Log(true, 'info', logFile, this.encoding);
          this.Warn = new Log(true, 'warn', logFile, this.encoding);
          this.Error = new Log(true, 'error', logFile, this.encoding);
          break;
        }
        case 'debug': {
          this.Trace = new Log(true, 'trace', logFile, this.encoding);
          this.Debug = new Log(false, 'debug', logFile, this.encoding);
          this.Info = new Log(true, 'info', logFile, this.encoding);
          this.Warn = new Log(true, 'warn', logFile, this.encoding);
          this.Error = new Log(true, 'error', logFile, this.encoding);
          break;
        }
        case 'info': {
          this.Trace = new Log(true, 'trace', logFile, this.encoding);
          this.Debug = new Log(true, 'debug', logFile, this.encoding);
          this.Info = new Log(false, 'info', logFile, this.encoding);
          this.Warn = new Log(true, 'warn', logFile, this.encoding);
          this.Error = new Log(true, 'error', logFile, this.encoding);
          break;
        }
        case 'warn': {
          this.Trace = new Log(true, 'trace', logFile, this.encoding);
          this.Debug = new Log(true, 'debug', logFile, this.encoding);
          this.Info = new Log(true, 'info', logFile, this.encoding);
          this.Warn = new Log(false, 'warn', logFile, this.encoding);
          this.Error = new Log(true, 'error', logFile, this.encoding);
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
    } catch (err) {
      console.error('no access!');
    }
  }

  // Set logger output to the IO stream
  public SetLogStream(level: string): void {
    if (levels.indexOf(level) < 0) {
      console.log('invalid level');
      return;
    }
    this.level = level;
    // set logs base on level
    switch (this.level) {
      case 'trace': {
        this.Trace = new Log(false, 'trace', process.stdout, this.encoding);
        this.Debug = new Log(true, 'debug', process.stdout, this.encoding);
        this.Info = new Log(true, 'info', process.stdout, this.encoding);
        this.Warn = new Log(true, 'warn', process.stdout, this.encoding);
        this.Error = new Log(true, 'error', process.stdout, this.encoding);
        break;
      }
      case 'debug': {
        this.Trace = new Log(true, 'trace', process.stdout, this.encoding);
        this.Debug = new Log(false, 'debug', process.stdout, this.encoding);
        this.Info = new Log(true, 'info', process.stdout, this.encoding);
        this.Warn = new Log(true, 'warn', process.stdout, this.encoding);
        this.Error = new Log(true, 'error', process.stdout, this.encoding);
        break;
      }
      case 'info': {
        this.Trace = new Log(true, 'trace', process.stdout, this.encoding);
        this.Debug = new Log(true, 'debug', process.stdout, this.encoding);
        this.Info = new Log(false, 'info', process.stdout, this.encoding);
        this.Warn = new Log(true, 'warn', process.stdout, this.encoding);
        this.Error = new Log(true, 'error', process.stdout, this.encoding);
        break;
      }
      case 'warn': {
        this.Trace = new Log(true, 'trace', process.stdout, this.encoding);
        this.Debug = new Log(true, 'debug', process.stdout, this.encoding);
        this.Info = new Log(true, 'info', process.stdout, this.encoding);
        this.Warn = new Log(false, 'warn', process.stdout, this.encoding);
        this.Error = new Log(true, 'error', process.stdout, this.encoding);
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
        console.log('invalid log level');
      }
    }
  }

  // Get encoding
  public GetEncoding(): string {
    return this.encoding;
  }

  // Set encoding
  public SetEncoding(type: string): void {
    if (encodings.indexOf(type) > -1) {
      this.encoding = type;
    } else {
      console.log('invalid encoding');
    }
  }

  // Get log level
  public GetLogLevel(): string {
    return this.level;
  }

  public Log(msg: string, time: number, route: string): void {
    let encoding = this.encoding;
    this.Trace.Output(msg, encoding, time, route);
    this.Debug.Output(msg, encoding, time, route);
    this.Info.Output(msg, encoding, time, route);
    this.Warn.Output(msg, encoding, time, route);
    this.Error.Output(msg, encoding, time, route);
  }
  public MeasureRunTime(time: any, route: string): void {
    const hrend = process.hrtime(time);

    const finalElapsedTime = hrend[0] * 1000 + hrend[1] / 1000000;
    this.Log('This function took:' + Math.round(finalElapsedTime) + 'ms', finalElapsedTime, route);
  }
}
class Log {
  private level: string;
  private out: stream.Writable;
  private encoding: string;

  constructor(discard: boolean, level: string, stream: stream.Writable, encoding: string) {
    this.level = level;
    this.encoding = encoding;

    if (discard) {
      // don't write to file
      this.out = new DiscardStream();
    } else {
      // write to file
      this.out = stream;
    }
  }

  public Output(msg: string, encoding: string, elapsedTime: number, route: string) {
    const timestamp = moment().format('MM-DD-YYYY h:mm:ss');
    const prefix = timestamp + ' ' + this.level + ' | ';

    const logObj = {
      route: route,
      timestamp: timestamp,
      elapsedtime: elapsedTime ? Math.round(elapsedTime) + 'ms' : undefined,
      level: this.level,
      message: msg,
    };

    switch (encoding) {
      case 'json': {
        this.out.write(JSON.stringify(logObj) + '\n');
        break;
      }
      default: {
        this.out.write(prefix + msg + '\n');
      }
    }
  }
}

// for when we don't need to write to a file
class DiscardStream extends stream.Writable {
  _write(chunk: any, enc: any, next: any) {
    // do nothing
  }
}
