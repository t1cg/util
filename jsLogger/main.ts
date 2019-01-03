import Logger from './logger';


function main() {
  const L = new Logger()

  L.Log("hello")

  L.SetLogLevel('warn')
  L.Log("world")
}

main()