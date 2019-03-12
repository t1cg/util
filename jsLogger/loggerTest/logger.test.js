"use strict";
exports.__esModule = true;
var logger_1 = require("../logger");
var chai_1 = require("chai");
describe('Constructor', function () {
    it('Should create a default logger with utf8 and level = trace', function () {
        var L = new logger_1["default"]();
        chai_1.expect(L.GetLogLevel()).to.equal('trace');
        chai_1.expect(L.GetEncoding()).to.equal('utf8');
    });
    it('Should create a logger with encoding of json and a level of trace', function () {
        var L = new logger_1["default"]('json', 'error');
        chai_1.expect(L.GetLogLevel()).to.equal('error');
        chai_1.expect(L.GetEncoding()).to.equal('json');
    });
    it('Should handle invalid input to constructor', function () {
        var L = new logger_1["default"]('foo', 'bar');
        chai_1.expect(L.GetLogLevel()).to.equal('trace');
        chai_1.expect(L.GetEncoding()).to.equal('utf8');
    });
});
describe('Encoding', function () {
    it('Should change default log level ', function () {
        var L = new logger_1["default"]();
        L.SetEncoding('json');
        chai_1.expect(L.GetEncoding()).to.equal('json');
    });
    it('Should handle an invalid request', function () {
        var L = new logger_1["default"]();
        L.SetEncoding('asdf');
        chai_1.expect(L.GetEncoding()).to.equal('utf8');
    });
});
describe('LogLevel', function () {
    it('Should change default log level ', function () {
        var L = new logger_1["default"]();
        L.SetLogStream('error');
        chai_1.expect(L.GetLogLevel()).to.equal('error');
    });
    it('Should handle an invalid request', function () {
        var L = new logger_1["default"]();
        L.SetLogStream('asdf');
        chai_1.expect(L.GetLogLevel()).to.equal('trace');
    });
});
