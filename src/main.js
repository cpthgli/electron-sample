const electron = require('electron');
const app = electron.app;
const BrowserWindow = electron.BrowserWindow;
let mainWindow;

function createMainWindow() {
  mainWindow = new BrowserWindow({
    width: 800,
    height: 600
  });

  function runServer() {
    var server = require('child_process').spawn('./main', [process.versions.node, process.versions.chrome, process.versions.electron]);
    return server;
  }

  function load() {
    var rq = require('request-promise');
    var mainAddr = 'http://localhost:8080';
    mainWindow.loadURL(mainAddr);
    rq(mainAddr)
      .then(function (htmlString) {
        console.log('server started');
        mainWindow.loadURL(mainAddr);
      })
      .catch(function (err) {
        console.log(err)
        return
      });
  }

  console.log('pre-server');
  var server = runServer();
  console.log('next-server');
  load();
  console.log('next-load');

  // 遅延表示
  // mainWindow.once('ready-to-show', () => {
  //  console.log('main ready-to-show');
  //  mainWindow.show()
  // })

  mainWindow.on('closed', function () {
    mainWindow = null;
    server.kill('SIGINT');
    console.log('closed');
  });
}

app.on('ready', createMainWindow);

app.on('window-all-closed', function () {
  if (process.platform !== 'darwin') {
    app.quit();
  }
  console.log('window all clesed');
})

app.on('activate', function () {
  if (mainWindow === null) {
    createMainWindow();
  }
})
