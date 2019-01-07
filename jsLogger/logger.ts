import * as stream from 'stream';
import * as moment from 'moment';
import * as fs from 'fs';

const encodings = ['utf8', 'json']
const levels = ['trace', 'debug', 'info', 'warn', 'error']

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
      this.level = level
    } else {
      console.log('invalid level, set to default (trace)')
      this.level = 'trace'
    }
    // check encoding exists, then set it
    if (encodings.indexOf(encoding) > -1) {
      this.encoding = encoding
    } else {
      console.log('invalid encoding, set to default (utf8)')
      this.encoding = 'utf8'
    }
    // set log level
    this.SetLogStream(level)
  }

  // Sets logger output stream to a file
  public SetLogFile(fpath: string, level = this.level): void {
    this.level = level

    var logFile: fs.WriteStream;

    try {
      fs.accessSync('path', fs.constants.W_OK);
      logFile = fs.createWriteStream(fpath, {flags: 'a'})

      // set logs based on level
      switch (this.level) {
        case 'trace': { 
          this.Trace = new Log(false, 'trace', logFile)
          this.Debug = new Log(false, 'debug', logFile)
          this.Info = new Log(false, 'info', logFile)
          this.Warn = new Log(false, 'warn', logFile)
          this.Error = new Log(false, 'error', logFile)
          break;        }
        case 'debug': {
          this.Trace = new Log(true, 'trace', logFile)
          this.Debug = new Log(false, 'debug', logFile)
          this.Info = new Log(false, 'info', logFile)
          this.Warn = new Log(false, 'warn', logFile)
          this.Error = new Log(false, 'error', logFile)
          break;
        }
        case 'info': {
          this.Trace = new Log(true, 'trace', logFile)
          this.Debug = new Log(true, 'debug', logFile)
          this.Info = new Log(false, 'info', logFile)
          this.Warn = new Log(false, 'warn', logFile)
          this.Error = new Log(false, 'error', logFile)
          break;
        }
        case 'warn': {
          this.Trace = new Log(true, 'trace', logFile)
          this.Debug = new Log(true, 'debug', logFile)
          this.Info = new Log(true, 'info', logFile)
          this.Warn = new Log(false, 'warn', logFile)
          this.Error = new Log(false, 'error', logFile)
          break;
        }         
        case 'error': {
          this.Trace = new Log(true, 'trace', logFile)
          this.Debug = new Log(true, 'debug', logFile)
          this.Info = new Log(true, 'info', logFile)
          this.Warn = new Log(true, 'warn', logFile)
          this.Error = new Log(false, 'error', logFile)
        }
        default: {
          console.log('invalid log level')
        }
      }
    } catch (err) {
      console.error('no access!');
    }
  }

  // Set logger output to the IO stream
  public SetLogStream(level: string): void {
    // set logs base on level
    switch (this.level) {
      case 'trace': {
        this.Trace = new Log(false, 'trace', process.stdout)
        this.Debug = new Log(false, 'debug', process.stdout)
        this.Info = new Log(false, 'info', process.stdout)
        this.Warn = new Log(false, 'warn', process.stdout)
        this.Error = new Log(false, 'error', process.stdout)
        break;
      }
      case 'debug': {
        this.Trace = new Log(true, 'trace', process.stdout)
        this.Debug = new Log(false, 'debug', process.stdout)
        this.Info = new Log(false, 'info', process.stdout)
        this.Warn = new Log(false, 'warn', process.stdout)
        this.Error = new Log(false, 'error', process.stdout)
        break;
      }
      case 'info': {
        this.Trace = new Log(true, 'trace', process.stdout)
        this.Debug = new Log(true, 'debug', process.stdout)
        this.Info = new Log(false, 'info', process.stdout)
        this.Warn = new Log(false, 'warn', process.stdout)
        this.Error = new Log(false, 'error', process.stdout)
        break;
      }
      case 'warn': {
        this.Trace = new Log(true, 'trace', process.stdout)
        this.Debug = new Log(true, 'debug', process.stdout)
        this.Info = new Log(true, 'info', process.stdout)
        this.Warn = new Log(false, 'warn', process.stdout)
        this.Error = new Log(false, 'error', process.stdout)
        break;
      }
      case 'error': {
        this.Trace = new Log(true, 'trace', process.stdout)
        this.Debug = new Log(true, 'debug', process.stdout)
        this.Info = new Log(true, 'info', process.stdout)
        this.Warn = new Log(true, 'warn', process.stdout)
        this.Error = new Log(false, 'error', process.stdout)
        break;
      }
      default: {
        console.log("invalid log level")                                       
      }
    }            
  }

  // Get encoding
  public GetEncoding(): string {
    return this.encoding
  }

  // Set encoding
  public SetEncoding(type: string): void {
    if (encodings.indexOf(type) > -1) {
      this.encoding = type
    } else {
      console.log('invalid encoding')
    }
  }

  // Get log level
  public GetLogLevel(): string {
    return this.level;
  }

  // Set log level to input
  public SetLogLevel(level: string): void {
    if (levels.indexOf(level) > -1) {
      this.level = level
      this.SetLogStream(level)
    } else {
      console.log('invalid level')
    }
  }

  public Log(msg: string): void {
    this.Trace.Output(msg)
    this.Debug.Output(msg)
    this.Info.Output(msg)
    this.Warn.Output(msg)
    this.Error.Output(msg)
  }
}

class Log {
  private prefix: string;
  private level: string;
  private out: stream.Writable;

  constructor(discard: boolean, level: string, stream: stream.Writable) {
    this.level = level

    if (discard) {
      // don't write to file
      this.out = new DiscardStream()
    } else {
      // write to file
      this.out = stream
    }

    this.prefix = moment().format() + " " + level + " | "
  }

  public Output(msg: string) {
    this.out.write(this.prefix + msg + "\n")
  }

}

// for when we don't need to write to a file
class DiscardStream extends stream.Writable {
  _write(chunk: any, enc: any, next: any) {
    // do nothing
  }
}
