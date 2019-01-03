Object.defineProperty(exports, "__esModule", { value: true });
const logger_1 = require("./logger");
function main() {
    const L = new logger_1.default();
    L.SetLogFile('../test.txt', 'info');
    L.Log("hello");
}
main();
