package main
import(
    "os"
	"syscall"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)


//func closeEvent( mainDialog *widgets.QDialog){
//    quit_msg := "你确定要退出吗？"
//    ActivateEvent();
//    reply := widgets.QMessageBox_Question(mainDialog, "Message",
//                     quit_msg, widgets.QMessageBox__Yes, widgets.QMessageBox__No)
//
//    if reply == widgets.QMessageBox__Yes{
//        syscall.Exit(0)
//	}
//}

type AppDoingList struct{
    widgets.QApplication
    _ func() `signal:"quitApp"`
    _ func() `signal:"activateApp"`
}

func main(){
    //var app *widgets.QApplication
    //if(app==nil){
    //    app = widgets.NewQApplication(len(os.Args), os.Args)
    //}
    var app *AppDoingList
    if(app==nil){
        app = NewAppDoingList(len(os.Args), os.Args)
        //app = &AppDoingList{QApplication:*widgets.NewQApplication(len(os.Args), os.Args)}
	}

    if !widgets.QSystemTrayIcon_IsSystemTrayAvailable() {
        widgets.QMessageBox_Critical(nil, "系统托盘", "不支持系统托盘", widgets.QMessageBox__Close, widgets.QMessageBox__Close)
        syscall.Exit(1)
	}

    mainDialog := widgets.NewQDialog(nil, core.Qt__WindowCloseButtonHint | core.Qt__WindowMinimizeButtonHint | core.Qt__MSWindowsFixedSizeDialogHint)
    // | core.Qt__WindowTitleHint
    var ui Ui_Dialog
    ui.SetupUi(mainDialog)
    mainDialog.Show()

    mainWidget := widgets.NewQDesktopWidgetFromPointer(nil)
    wid := new(Ui_SysTray)
    wid.SetupUi( mainWidget, app )
    mainWidget.Show()

    app.SetQuitOnLastWindowClosed(false)

    //app.ConnectQuitApp(closeEvent)
	activateEvent := func(){ui.ShowDialog(app,mainDialog)}
    app.ConnectActivateApp(activateEvent)

    //keyboard.add_hotkey('ctrl+alt+d', activateEvent)
    //keyboard.add_hotkey('shift+alt+c', activateEvent)

    syscall.Exit(app.Exec())
}

