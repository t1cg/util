Object.defineProperty(exports, "__esModule", { value: true });
const logger_1 = require("./logger");
function main() {
    const L = new logger_1.default();
    L.Log("hello");
    L.SetLogLevel('warn');
    L.Log("world");
}
main();
