import Logger from './logger';


function main() {
  const L = new Logger()

  L.Log("hello")
  L.SetEncoding('json');
  L.SetLogStream('warn')
  L.Log("world")
  L.SetLogFile('../test.txt', 'info')
  L.Log("hello")
}

main()