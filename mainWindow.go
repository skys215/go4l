package main
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type DoingList struct{
	widgets.QListWidget
//
//    def keyPressEvent(self, event):
//        if     event.key() == QtCore.Qt.Key_Delete  or event.key() == QtCore.Qt.Key_Enter:
//            self.takeItem( self.currentRow() )
//            return
//        super(DoingList, self).keyPressEvent(event)
}
// func ( s *DoingList ) FocusInEvent(event gui.QFocusEvent_ITF){
// 	cur := s.CurrentRow()
// 	if cur == -1{
// 		s.SetCurrentRow(0)
// 	}
// 	s.SetFocus2()
// 	s.FocusInEvent( event )
// }
//class DoingListItemDelegate(QtWidgets.QItemDelegate):
//    def editorEvent( self, event, model, option, index):
//        print('doinglist item delegate edit event')
//        return False

type Ui_Dialog struct{
	core.QObject
	VerticalLayoutWidget *widgets.QWidget
	VerticalLayout *widgets.QVBoxLayout
	LineEdit *widgets.QLineEdit
	ListWidget *DoingList
	Label *widgets.QLabel
	_ func( *widgets.QApplication, *widgets.QDialog ) `slot:activateEvent`
	_ func( *widgets.QDialog )
}
//Need to wrap with pyqtSlot
// @QtCore.pyqtSlot(QApplication,QDialog)
func (s *Ui_Dialog) ShowDialog (app *AppDoingList, qdialog *widgets.QDialog){
   if qdialog.WindowState() == core.Qt__WindowMinimized{
       qdialog.ShowNormal()
	}

   //print('fired active event\n')

   if !qdialog.IsActiveWindow(){
       app.SetActiveWindow(qdialog)
    }
   qdialog.Raise()
   qdialog.Show()
   //#qdialog.activateWindow()
   //# if qdialog.windowState() == QtCore.Qt.WindowMinimized:
   //#     qdialog.showNormal()

   //# qdialog.show()
   s.LineEdit.SetFocus2()
}

func (s *Ui_Dialog) SetupUi(dialog *widgets.QDialog){
    //self.sig = Communicate()
    //self.sig.showDialog.connect(self.showDialog)

    dialog.SetObjectName("Dialog")
    dialog.Resize2(292, 350)
    dialog.SetFixedSize(dialog.Size())
    dialog.SetWindowIcon(gui.NewQIcon5("logo.png"))

    s.VerticalLayoutWidget = widgets.NewQWidget(dialog, 0)
    s.VerticalLayoutWidget.SetGeometry(core.NewQRect4(10, 10, 271, 320))
    s.VerticalLayoutWidget.SetObjectName("verticalLayoutWidget")
    s.VerticalLayout = widgets.NewQVBoxLayout2(s.VerticalLayoutWidget)
    s.VerticalLayout.SetContentsMargins(0, 0, 0, 0)
    s.VerticalLayout.SetObjectName("verticalLayout")

    s.LineEdit = widgets.NewQLineEdit(s.VerticalLayoutWidget)
    s.LineEdit.SetObjectName("lineEdit")
    s.LineEdit.ConnectReturnPressed(s.FinishEditing)
    s.VerticalLayout.AddWidget(s.LineEdit, 0, core.Qt__AlignTop)

    s.Label = widgets.NewQLabel(s.VerticalLayoutWidget, core.Qt__Widget)
    font := gui.NewQFont()
    font.SetFamily("Calibri")
    font.SetPointSize(12)
    s.Label.SetFont(font)
    s.Label.SetObjectName("label")
    s.VerticalLayout.AddWidget(s.Label, 0, core.Qt__AlignTop)

    s.ListWidget = NewDoingList(s.VerticalLayoutWidget)
    s.ListWidget.SetObjectName("listWidget")
    s.ListWidget.SetStyleSheet("QListWidgetItem")
    s.VerticalLayout.AddWidget(s.ListWidget, 0, core.Qt__AlignTop)

    s.RetranslateUi(dialog)
    //QtCore.QMetaObject.connectSlotsByName(Dialog)

    //closeKeySeq = QtGui.QKeySequence(QtCore.Qt.CTRL+QtCore.Qt.Key_W)
    //closeKeyShortcut = QtWidgets.QShortcut(closeKeySeq, Dialog)
    //closeKeyShortcut.activated.connect( Dialog.close )

    //# MainWindow::changeEvent(QEvent *event) {
    //#   if(event->type() == QEvent::WindowStateChange) {
    //#     if(!isMinimized()) {
    //#       setAttribute(Qt::WA_Mapped);
    //#     }
    //#   }
    //# }

    //quitKeySeq := gui.NewQKeySequence2("CTRL+Q",gui.QKeySequence__PortableText)//core.Qt__CTRL | core.Qt__Key_Q)
    //quitKeyShortcut = QtWidgets.QShortcut(quitKeySeq, Dialog)
    //quitKeyShortcut.activated.connect( QApplication.instance().sig.quitApp )
}

func (s *Ui_Dialog) RetranslateUi(dialog *widgets.QDialog){
    dialog.SetWindowTitle("Doing list")
    s.LineEdit.SetPlaceholderText("I am doing...")
    s.Label.SetText("I was doing:")
}

func (s *Ui_Dialog)FinishEditing( ){
    doing := s.LineEdit.Text()
    if len([]rune(doing)) == 0 {
        return
	}
    s.LineEdit.Clear()
    item := widgets.NewQListWidgetItem2(doing,s.ListWidget,0)
    item.SetFlags(core.Qt__ItemIsSelectable|core.Qt__ItemIsUserCheckable|core.Qt__ItemIsEnabled)
    s.ListWidget.InsertItem(0, item)
}

