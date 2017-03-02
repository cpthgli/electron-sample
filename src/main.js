const electron = require('electron');
const app = electron.app;
const Menu = electron.Menu;
const BrowserWindow = electron.BrowserWindow;
let mainWindow;

const menu = new Menu()

function createMainWindow() {
  mainWindow = new BrowserWindow({
    width: 800,
    height: 600
  });
  // TODO: get random port and set
  function runServer() {
    var server = require('child_process').execFile('./main');
    console.log('server started');
    return server;
  }

  function load() {
    var urlencode = require('urlencode');
    var rq = require('request-promise');
    var mainAddr = 'http://localhost:8080';
    var options = {
      method: 'POST',
      uri: mainAddr,
      form: {
        node: process.versions.node,
        chrome: process.versions.chrome,
        electron: process.versions.electron
      }
    }

    var url = mainAddr + "/?" + urlencode.stringify(options.form);
    console.log(url);
    mainWindow.loadURL(url);
    console.log("url loaded")
    // rq.
    // rq(options)
    //   .then(function (htmlString) {
    //     mainWindow.loadURL(mainAddr);
    //   })
    //   .catch(function (err) {
    //     console.log(err)
    //     return
    //   });
  }

  var server = runServer();
  load();

  // 遅延表示
  // mainWindow.once('ready-to-show', () => {
  //  console.log('main ready-to-show');
  //  mainWindow.show()a
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
  console.log('window all closed');
})

app.on('activate', function () {
  if (mainWindow === null) {
    createMainWindow();
  }
})
