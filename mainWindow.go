package main
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)
//type Communicate struct{
//	widgets.QWidget
//	ShowDialog 
//}
//class Communicate(QtCore.QObject):
//    showDialog = QtCore.pyqtSignal(QApplication,QDialog)
//
//class DoingList(QtWidgets.QListWidget):
//    def focusInEvent(self, event):
//        cur = self.currentRow()
//        if cur == -1:
//            self.setCurrentRow(0)
//
//        self.setFocus()
//        super(DoingList, self).focusInEvent(event)
//
//
//    def keyPressEvent(self, event):
//        if     event.key() == QtCore.Qt.Key_Delete  or event.key() == QtCore.Qt.Key_Enter:
//            self.takeItem( self.currentRow() )
//            return
//        super(DoingList, self).keyPressEvent(event)
//
//class DoingListItemDelegate(QtWidgets.QItemDelegate):
//    def editorEvent( self, event, model, option, index):
//        print('doinglist item delegate edit event')
//        return False

type Ui_Dialog struct{
    //@QtCore.pyqtSlot(QApplication,QDialog)
	core.QObject
	_ func( *widgets.QApplication, *widgets.QDialog )
	_ func( *widgets.QDialog )
}
func (s *Ui_Dialog) showDialog (app *widgets.QApplication, qdialog *widgetse.QDialog){
    if qdialog.WindowState() == core.Qt__WindowMinimized{
        qdialog.ShowNormal()
	}

    //print('fired active event\n')

    if !app.IsActiveWindow():
        app.SetActiveWindow(qdialog)
    qdialog.Raise_()
    qdialog.Show()
    #qdialog.activateWindow()
    # if qdialog.windowState() == QtCore.Qt.WindowMinimized:
    #     qdialog.showNormal()

    # qdialog.show()
    s.LineEdit.setFocus()
}

func (s *Ui_Dialog) setupUi(Dialog *widgets.QDialog){
    Dialog.SetWindowFlags(
        core.Qt__WindowCloseButtonHint
        | core.Qt__WindowMinimizeButtonHint
        | core.Qt__MSWindowsFixedSizeDialogHint
        # | core.Qt__WindowTitleHint
    )
    //self.sig = Communicate()
    //self.sig.showDialog.connect(self.showDialog)

    Dialog.SetObjectName("Dialog")
    Dialog.Resize(292, 350)
    Dialog.SetFixedSize(Dialog.Size())
    Dialog.SetWindowIcon(gui.QIcon("logo.png"))

    //self.verticalLayoutWidget = QtWidgets.QWidget(Dialog)
    //self.verticalLayoutWidget.setGeometry(QtCore.QRect(10, 10, 271, 320))
    //self.verticalLayoutWidget.setObjectName("verticalLayoutWidget")
    //self.verticalLayout = QtWidgets.QVBoxLayout(self.verticalLayoutWidget)
    //self.verticalLayout.setContentsMargins(0, 0, 0, 0)
    //self.verticalLayout.setObjectName("verticalLayout")

    //self.lineEdit = QtWidgets.QLineEdit(self.verticalLayoutWidget)
    //self.lineEdit.setObjectName("lineEdit")
    //self.lineEdit.returnPressed.connect(self.finishEditing)
    //self.verticalLayout.addWidget(self.lineEdit)

    //self.label = QtWidgets.QLabel(self.verticalLayoutWidget)
    //font = QtGui.QFont()
    //font.setFamily("Calibri")
    //font.setPointSize(12)
    //self.label.setFont(font)
    //self.label.setObjectName("label")
    //self.verticalLayout.addWidget(self.label)

    //self.listWidget = DoingList(self.verticalLayoutWidget)
    //self.listWidget.setObjectName("listWidget")
    //self.listWidget.setStyleSheet("QListWidgetItem")
    //self.verticalLayout.addWidget(self.listWidget)

    //self.retranslateUi(Dialog)
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

    //quitKeySeq = QtGui.QKeySequence(QtCore.Qt.CTRL+QtCore.Qt.Key_Q)
    //quitKeyShortcut = QtWidgets.QShortcut(quitKeySeq, Dialog)
    //quitKeyShortcut.activated.connect( QApplication.instance().sig.quitApp )
}

//def retranslateUi(self, Dialog):
//    _translate = QtCore.QCoreApplication.translate
//    Dialog.setWindowTitle(_translate("Dialog", "Doing list"))
//    self.lineEdit.setPlaceholderText(_translate("Dialog", "I am doing..."))
//    self.label.setText(_translate("Dialog", "I was doing:"))
//
//def finishEditing( self ):
//    doing = self.lineEdit.text()
//    if len(doing) == 0 :
//        return
//    self.lineEdit.clear()
//    item = QtWidgets.QListWidgetItem()
//    item.setFlags(QtCore.Qt.ItemIsSelectable|QtCore.Qt.ItemIsUserCheckable|QtCore.Qt.ItemIsEnabled)
//    item.setText(doing)
//    self.listWidget.insertItem(0, item)

