package main
import(
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)
type Ui_SysTray struct{
	core.QObject
}
func (s *Ui_SysTray) SetupUi( widget *widgets.QDesktopWidget, app *AppDoingList){
    systray := widgets.NewQSystemTrayIcon2( gui.NewQIcon5("logo.png"), widget)
    systray.SetToolTip("B4L")
    //systray.Activated.connect(lambda reason: self.iconTriggered(reason, app))

    tray_icon_menu := widgets.NewQMenu(nil)
    //openAction := tray_icon_menu.AddAction("Open")
    //openAction.Triggered.Connect(app.Sig.ActivateApp.Emit)
    //exitAction := tray_icon_menu.AddAction("Exit")
    //exitAction.Triggered.Connect(app.Sig.QuitApp.Emit)

    systray.SetContextMenu(tray_icon_menu)
    systray.Show()
}

func (self *Ui_SysTray) IconTriggered(reason widgets.QSystemTrayIcon__ActivationReason, app *widgets.QApplication){
    if reason == widgets.QSystemTrayIcon__DoubleClick{
		fmt.Sprintf("%s","Double clicked system trsy icon")
        //app.Sig.ActivateApp.Emit()
	}
}
