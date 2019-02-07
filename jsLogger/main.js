Object.defineProperty(exports, "__esModule", { value: true });
const logger_1 = require("./logger");
function main() {
    const L = new logger_1.default();
    // L.SetEncoding('json');
    // L.SetLogLevel('info')
    L.Log("hello");
    L.SetEncoding('json');
    L.SetLogLevel('warn');
    L.Log("world");
    L.SetLogFile('../test.txt', 'info');
    L.Log("hello");
}
main();
