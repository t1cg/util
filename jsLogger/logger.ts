import * as stream from 'stream';
import * as moment from 'moment';
import * as fs from 'fs';

export default class Logger {
  private Trace: Log;
  private Debug: Log;
  private Info: Log;
  private Warn: Log;
  private Error: Log;

  private level: string;
  private encoding: string;

  constructor(level = 'trace') {
    this.level = level
    this.encoding = ""
    // this.encoding = encoding
    this.SetLogStream()
    // this.SetEncoding('yolo')
    // console.log(encoding)
  }

  // Sets logger output stream to a file
  public SetLogFile(fpath: string, logname: string, level: string): void {
    var logFile: fs.WriteStream = fs.createWriteStream(fpath, {flags: 'a'})
    switch (this.level) {
      case 'trace':
        console.log("made it to trace")
        this.Trace = new Log(false, 'trace', logFile, this.encoding)
        this.Debug = new Log(false, 'debug', logFile, this.encoding)
        this.Info = new Log(false, 'info', logFile, this.encoding)
        this.Warn = new Log(false, 'warn', logFile, this.encoding)
        this.Error = new Log(false, 'error', logFile, this.encoding)
        break;        
      case 'debug':
        this.Trace = new Log(true, 'trace', logFile, this.encoding)
        this.Debug = new Log(false, 'debug', logFile, this.encoding)
        this.Info = new Log(false, 'info', logFile, this.encoding)
        this.Warn = new Log(false, 'warn', logFile, this.encoding)
        this.Error = new Log(false, 'error', logFile, this.encoding)
        break;
      case 'info':
        this.Trace = new Log(true, 'trace', logFile, this.encoding)
        this.Debug = new Log(true, 'debug', logFile, this.encoding)
        this.Info = new Log(false, 'info', logFile, this.encoding)
        this.Warn = new Log(false, 'warn', logFile, this.encoding)
        this.Error = new Log(false, 'error', logFile, this.encoding)
        break;
      case 'warn':
        this.Trace = new Log(true, 'trace', logFile, this.encoding)
        this.Debug = new Log(true, 'debug', logFile, this.encoding)
        this.Info = new Log(true, 'info', logFile, this.encoding)
        this.Warn = new Log(false, 'warn', logFile, this.encoding)
        this.Error = new Log(false, 'error', logFile, this.encoding)
        break;                
      default:
        this.Trace = new Log(true, 'trace', logFile, this.encoding)
        this.Debug = new Log(true, 'debug', logFile, this.encoding)
        this.Info = new Log(true, 'info', logFile, this.encoding)
        this.Warn = new Log(true, 'warn', logFile, this.encoding)
        this.Error = new Log(false, 'error', logFile, this.encoding)
    }
    
  }

  // Set logger output to the IO stream
  public SetLogStream(): void {
    switch (this.level) {
      case 'trace': {
        console.log("made it to trace")
        this.Trace = new Log(false, 'trace', process.stdout, this.encoding)
        this.Debug = new Log(false, 'debug', process.stdout, this.encoding)
        this.Info = new Log(false, 'info', process.stdout, this.encoding)
        this.Warn = new Log(false, 'warn', process.stdout, this.encoding)
        this.Error = new Log(false, 'error', process.stdout, this.encoding)
        break;
      }
      case 'debug': {
        console.log("made it to debug")
        this.Trace = new Log(true, 'trace', process.stdout, this.encoding)
        this.Debug = new Log(false, 'debug', process.stdout, this.encoding)
        this.Info = new Log(false, 'info', process.stdout, this.encoding)
        this.Warn = new Log(false, 'warn', process.stdout, this.encoding)
        this.Error = new Log(false, 'error', process.stdout, this.encoding)
        break;
      }
      case 'info': {
        console.log("made it to info")
        this.Trace = new Log(true, 'trace', process.stdout, this.encoding)
        this.Debug = new Log(true, 'debug', process.stdout, this.encoding)
        this.Info = new Log(false, 'info', process.stdout, this.encoding)
        this.Warn = new Log(false, 'warn', process.stdout, this.encoding)
        this.Error = new Log(false, 'error', process.stdout, this.encoding)
        break;
      }
      case 'warn': {
        console.log("made it to warn")
        this.Trace = new Log(true, 'trace', process.stdout, this.encoding)
        this.Debug = new Log(true, 'debug', process.stdout, this.encoding)
        this.Info = new Log(true, 'info', process.stdout, this.encoding)
        this.Warn = new Log(false, 'warn', process.stdout, this.encoding)
        this.Error = new Log(false, 'error', process.stdout, this.encoding)
        break;
      }
      case 'error': {
        console.log("made it to default")
        this.Trace = new Log(true, 'trace', process.stdout, this.encoding)
        this.Debug = new Log(true, 'debug', process.stdout, this.encoding)
        this.Info = new Log(true, 'info', process.stdout, this.encoding)
        this.Warn = new Log(true, 'warn', process.stdout, this.encoding)
        this.Error = new Log(false, 'error', process.stdout, this.encoding)
        break;
      }
      default: {
        console.log("invalid log level")                                       
      }
    }            
  }

  // Set log level to input
  public SetLogLevel(level: string): void {
    this.level = level
    this.SetLogStream()
  }

  // Set output encoding
  public SetEncoding(encoding: string): void {
    this.encoding = encoding
  }

  public Log(msg: string): void {
    let encoding = this.encoding;
    this.Trace.Output(msg, encoding)
    this.Debug.Output(msg, encoding)
    this.Info.Output(msg, encoding)
    this.Warn.Output(msg, encoding)
    this.Error.Output(msg, encoding)
  }
}

class Log {
  private prefix: string;
  private level: string;
  private out: stream.Writable;
  private encoding: string

  constructor(discard: boolean, level: string, stream: stream.Writable, encoding: string) {
    this.level = level
    this.encoding = encoding

    if (discard) {
      // don't write to file
      this.out = new DiscardStream()
    } else {
      // write to file
      this.out = stream
    }

    this.prefix = moment().format() + " " + level + " | "
  }

  public Output(msg: string, encoding: string) {

    switch (encoding) {
      case 'json': {
        this.out.write("ENCODING: JSON \n")
        this.out.write("{\"prefix\":" + "\"" + this.prefix + "\",\"message\":" + "\"" + msg + "\"}\n")
        break;
      }
      default: {
        this.out.write("ENCODING: " + this.encoding + "\n")
        this.out.write(this.prefix + msg + "\n")
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
