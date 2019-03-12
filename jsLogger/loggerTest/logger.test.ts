import Logger from '../logger';
import { expect } from 'chai';

describe('Constructor', () => {
    it('Should create a default logger with utf8 and level = trace', () => {
        const L = new Logger()
        expect(L.GetLogLevel()).to.equal('trace')
        expect(L.GetEncoding()).to.equal('utf8')
    })
    it('Should create a logger with encoding of json and a level of trace', () => {
        const L = new Logger('json', 'error')
        expect(L.GetLogLevel()).to.equal('error')
        expect(L.GetEncoding()).to.equal('json')
    })
    it('Should handle invalid input to constructor', () => {
        const L = new Logger('foo', 'bar')
        expect(L.GetLogLevel()).to.equal('trace')
        expect(L.GetEncoding()).to.equal('utf8')
    })
})
describe('Encoding', () => {
    it('Should change default log level ', () => {
        const L = new Logger()
        L.SetEncoding('json')
        expect(L.GetEncoding()).to.equal('json')
    })
    it('Should handle an invalid request', () => {
        const L = new Logger()
        L.SetEncoding('asdf')
        expect(L.GetEncoding()).to.equal('utf8')
    })
})
describe('LogLevel', () => {
    it('Should change default log level ', () => {
        const L = new Logger()
        L.SetLogStream('error')
        expect(L.GetLogLevel()).to.equal('error')
    })
    it('Should handle an invalid request', () => {
        const L = new Logger()
        L.SetLogStream('asdf')
        expect(L.GetLogLevel()).to.equal('trace')
    })
})

