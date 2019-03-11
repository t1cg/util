Object.defineProperty(exports, "__esModule", { value: true });
const logger_1 = require("../logger");
const chai_1 = require("chai");
describe('Constructor', () => {
    it('Should create a default logger with utf8 and level = trace', () => {
        const L = new logger_1.default();
        chai_1.expect(L.GetLogLevel()).to.equal('trace');
        chai_1.expect(L.GetEncoding()).to.equal('utf8');
    });
    it('Should create a logger with encoding of json and a level of trace', () => {
        const L = new logger_1.default('json', 'error');
        chai_1.expect(L.GetLogLevel()).to.equal('error');
        chai_1.expect(L.GetEncoding()).to.equal('json');
    });
    it('Should handle invalid input to constructor', () => {
        const L = new logger_1.default('foo', 'bar');
        chai_1.expect(L.GetLogLevel()).to.equal('trace');
        chai_1.expect(L.GetEncoding()).to.equal('utf8');
    });
});
describe('Encoding', () => {
    it('Should change default log level ', () => {
        const L = new logger_1.default();
        L.SetEncoding('json');
        chai_1.expect(L.GetEncoding()).to.equal('json');
    });
    it('Should handle an invalid request', () => {
        const L = new logger_1.default();
        L.SetEncoding('asdf');
        chai_1.expect(L.GetEncoding()).to.equal('utf8');
    });
});
describe('LogLevel', () => {
    it('Should change default log level and ', () => {
        const L = new logger_1.default();
        L.SetLogLevel('error');
        chai_1.expect(L.GetLogLevel()).to.equal('error');
    });
    it('Should handle an invalid request', () => {
        const L = new logger_1.default();
        L.SetLogLevel('asdf');
        chai_1.expect(L.GetLogLevel()).to.equal('trace');
    });
});
