import Logger from './logger';


function main() {
  const L = new Logger()
  L.SetLogFile('../test.txt', 'info')
  L.Log("hello")
}

main()