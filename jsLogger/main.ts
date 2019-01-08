import Logger from './logger';


function main() {
  const L = new Logger()
  // L.SetEncoding('json');
  // L.SetLogLevel('info')
  L.Log("hello")
  L.SetEncoding('json');
  L.SetLogLevel('warn')
  L.Log("world")
  L.SetLogFile('../test.txt', 'info')
  L.Log("hello")
}

main()