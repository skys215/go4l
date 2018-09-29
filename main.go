package main
import(
    "os"
	"syscall"

	//"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

//class Communicate(QtCore.QObject):
//    quitApp = QtCore.pyqtSignal()
//    activateApp = QtCore.pyqtSignal()

//func closeEvent(){
//    quit_msg = "你确定要退出吗？"
//    activateEvent()
//    reply = QtWidgets.QMessageBox.question(mainDialog, 'Message',
//                     quit_msg, QtWidgets.QMessageBox.Yes, QtWidgets.QMessageBox.No)
//
//    if reply == QtWidgets.QMessageBox.Yes:
//        sys.exit()
//}

//func activateEvent():
//    ui.sig.showDialog.emit(app,mainDialog)
//}

func main(){
	//ok
    //var app *core.QCoreApplication
    //app = core.QCoreApplication_Instance( )

    var app *widgets.QApplication
    if(app==nil){
        app = widgets.NewQApplication(len(os.Args), os.Args)
	}
    //app.sig = Communicate()
    //app.sig.quitApp.connect(closeEvent)
    //app.sig.activateApp.connect(activateEvent)


    if widgets.QSystemTrayIcon_IsSystemTrayAvailable() {
        widgets.QMessageBox_Critical(nil, "系统托盘", "不支持系统托盘", widgets.QMessageBox__Close, widgets.QMessageBox__Close)
        syscall.Exit(1)
	}

    mainDialog := widgets.NewQDialogFromPointer(nil)
    ui = go4l.mainWindow.Ui_Dialog()
    //ui.setupUi(mainDialog)
    // mainDialog.show()

    //mainWidget = QDesktopWidget()
    //wid = Ui_SysTray()
    //wid.setupUi( mainWidget, app )
    //mainWidget.show()

    //QApplication.setQuitOnLastWindowClosed(False)

    //keyboard.add_hotkey('ctrl+alt+d', activateEvent)
    //keyboard.add_hotkey('shift+alt+c', activateEvent)

    syscall.Exit(app.Exec())
}

