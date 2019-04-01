'use strict';
var __extends =
  (this && this.__extends) ||
  (function() {
    var extendStatics = function(d, b) {
      extendStatics =
        Object.setPrototypeOf ||
        ({ __proto__: [] } instanceof Array &&
          function(d, b) {
            d.__proto__ = b;
          }) ||
        function(d, b) {
          for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
        };
      return extendStatics(d, b);
    };
    return function(d, b) {
      extendStatics(d, b);
      function __() {
        this.constructor = d;
      }
      d.prototype = b === null ? Object.create(b) : ((__.prototype = b.prototype), new __());
    };
  })();
exports.__esModule = true;
var stream = require('stream');
var moment = require('moment');
var fs = require('fs');
var encodings = ['utf8', 'json'];
var levels = ['trace', 'debug', 'info', 'warn', 'error'];
var Logger = /** @class */ (function() {
  function Logger(encoding, level) {
    if (encoding === void 0) {
      encoding = 'utf8';
    }
    if (level === void 0) {
      level = 'trace';
    }
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
  Logger.prototype.SetLogFile = function(fpath, level) {
    if (level === void 0) {
      level = this.level;
    }
    if (levels.indexOf(level) < 0) {
      console.log('invalid level');
      return;
    }
    this.level = level;
    var logFile;
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
  };
  // Set logger output to the IO stream
  Logger.prototype.SetLogStream = function(level) {
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
  };
  // Get encoding
  Logger.prototype.GetEncoding = function() {
    return this.encoding;
  };
  // Set encoding
  Logger.prototype.SetEncoding = function(type) {
    if (encodings.indexOf(type) > -1) {
      this.encoding = type;
    } else {
      console.log('invalid encoding');
    }
  };
  // Get log level
  Logger.prototype.GetLogLevel = function() {
    return this.level;
  };
  Logger.prototype.Log = function(msg, time, route) {
    var encoding = this.encoding;
    this.Trace.Output(msg, encoding, time, route);
    this.Debug.Output(msg, encoding, time, route);
    this.Info.Output(msg, encoding, time, route);
    this.Warn.Output(msg, encoding, time, route);
    this.Error.Output(msg, encoding, time, route);
  };

  Logger.prototype.MeasureRunTime = function(time, route) {
    const hrend = process.hrtime(time);

    const finalElapsedTime = hrend[0] * 1000 + hrend[1] / 1000000;

    this.Log('This function took:' + Math.round(finalElapsedTime) + 'ms', finalElapsedTime, route);
  };
  return Logger;
})();
exports['default'] = Logger;
var Log = /** @class */ (function() {
  function Log(discard, level, stream, encoding) {
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
  Log.prototype.Output = function(msg, encoding, elapsedTime, route) {
    var timestamp = moment().format('MM-DD-YYYY h:mm:ss');
    var prefix = timestamp + ' ' + this.level + ' | ';

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
  };

  return Log;
})();
// for when we don't need to write to a file
var DiscardStream = /** @class */ (function(_super) {
  __extends(DiscardStream, _super);
  function DiscardStream() {
    return (_super !== null && _super.apply(this, arguments)) || this;
  }
  DiscardStream.prototype._write = function(chunk, enc, next) {
    // do nothing
  };
  return DiscardStream;
})(stream.Writable);
