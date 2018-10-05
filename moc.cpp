

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemDelegate>
#include <QAbstractItemModel>
#include <QAbstractItemView>
#include <QAction>
#include <QActionEvent>
#include <QApplication>
#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QDial>
#include <QDialog>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFocusEvent>
#include <QFont>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QKeyEvent>
#include <QLayout>
#include <QListWidget>
#include <QListWidgetItem>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPalette>
#include <QPdfWriter>
#include <QPoint>
#include <QQuickItem>
#include <QRadioData>
#include <QRect>
#include <QRegion>
#include <QResizeEvent>
#include <QScreen>
#include <QSessionManager>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionViewItem>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWheelEvent>
#include <QWidget>
#include <QWindow>


class DoingList9733cb: public QListWidget
{
Q_OBJECT
public:
	DoingList9733cb(QWidget *parent = Q_NULLPTR) : QListWidget(parent) {qRegisterMetaType<quintptr>("quintptr");DoingList9733cb_DoingList9733cb_QRegisterMetaType();DoingList9733cb_DoingList9733cb_QRegisterMetaTypes();callbackDoingList9733cb_Constructor(this);};
	 ~DoingList9733cb() { callbackDoingList9733cb_DestroyDoingList(this); };
	bool dropMimeData(int index, const QMimeData * data, Qt::DropAction action) { return callbackDoingList9733cb_DropMimeData(this, index, const_cast<QMimeData*>(data), action) != 0; };
	bool event(QEvent * e) { return callbackDoingList9733cb_Event(this, e) != 0; };
	
	void clear() { callbackDoingList9733cb_Clear(this); };
	void Signal_CurrentItemChanged(QListWidgetItem * current, QListWidgetItem * previous) { callbackDoingList9733cb_CurrentItemChanged(this, current, previous); };
	void Signal_CurrentRowChanged(int currentRow) { callbackDoingList9733cb_CurrentRowChanged(this, currentRow); };
	void Signal_CurrentTextChanged(const QString & currentText) { QByteArray t5a0ada = currentText.toUtf8(); Moc_PackedString currentTextPacked = { const_cast<char*>(t5a0ada.prepend("WHITESPACE").constData()+10), t5a0ada.size()-10 };callbackDoingList9733cb_CurrentTextChanged(this, currentTextPacked); };
	void dropEvent(QDropEvent * event) { callbackDoingList9733cb_DropEvent(this, event); };
	void Signal_ItemActivated(QListWidgetItem * item) { callbackDoingList9733cb_ItemActivated(this, item); };
	void Signal_ItemChanged(QListWidgetItem * item) { callbackDoingList9733cb_ItemChanged(this, item); };
	void Signal_ItemClicked(QListWidgetItem * item) { callbackDoingList9733cb_ItemClicked(this, item); };
	void Signal_ItemDoubleClicked(QListWidgetItem * item) { callbackDoingList9733cb_ItemDoubleClicked(this, item); };
	void Signal_ItemEntered(QListWidgetItem * item) { callbackDoingList9733cb_ItemEntered(this, item); };
	void Signal_ItemPressed(QListWidgetItem * item) { callbackDoingList9733cb_ItemPressed(this, item); };
	void Signal_ItemSelectionChanged() { callbackDoingList9733cb_ItemSelectionChanged(this); };
	void scrollToItem(const QListWidgetItem * item, QAbstractItemView::ScrollHint hint) { callbackDoingList9733cb_ScrollToItem(this, const_cast<QListWidgetItem*>(item), hint); };
	void setSelectionModel(QItemSelectionModel * selectionModel) { callbackDoingList9733cb_SetSelectionModel(this, selectionModel); };
	QMimeData * mimeData(const QList<QListWidgetItem *> items) const { return static_cast<QMimeData*>(callbackDoingList9733cb_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QListWidgetItem *>* tmpValue = new QList<QListWidgetItem *>(items); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackDoingList9733cb_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackDoingList9733cb_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, Qt::KeyboardModifiers modifiers) { return *static_cast<QModelIndex*>(callbackDoingList9733cb_MoveCursor(this, cursorAction, modifiers)); };
	void currentChanged(const QModelIndex & current, const QModelIndex & previous) { callbackDoingList9733cb_CurrentChanged(this, const_cast<QModelIndex*>(&current), const_cast<QModelIndex*>(&previous)); };
	void dataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackDoingList9733cb_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void dragLeaveEvent(QDragLeaveEvent * e) { callbackDoingList9733cb_DragLeaveEvent(this, e); };
	void dragMoveEvent(QDragMoveEvent * e) { callbackDoingList9733cb_DragMoveEvent(this, e); };
	void mouseMoveEvent(QMouseEvent * e) { callbackDoingList9733cb_MouseMoveEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackDoingList9733cb_MouseReleaseEvent(this, e); };
	void paintEvent(QPaintEvent * e) { callbackDoingList9733cb_PaintEvent(this, e); };
	void resizeEvent(QResizeEvent * e) { callbackDoingList9733cb_ResizeEvent(this, e); };
	void rowsAboutToBeRemoved(const QModelIndex & parent, int start, int end) { callbackDoingList9733cb_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), start, end); };
	void rowsInserted(const QModelIndex & parent, int start, int end) { callbackDoingList9733cb_RowsInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void scrollTo(const QModelIndex & index, QAbstractItemView::ScrollHint hint) { callbackDoingList9733cb_ScrollTo(this, const_cast<QModelIndex*>(&index), hint); };
	void selectionChanged(const QItemSelection & selected, const QItemSelection & deselected) { callbackDoingList9733cb_SelectionChanged(this, const_cast<QItemSelection*>(&selected), const_cast<QItemSelection*>(&deselected)); };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags command) { callbackDoingList9733cb_SetSelection(this, const_cast<QRect*>(&rect), command); };
	void startDrag(Qt::DropActions supportedActions) { callbackDoingList9733cb_StartDrag(this, supportedActions); };
	void timerEvent(QTimerEvent * e) { callbackDoingList9733cb_TimerEvent(this, e); };
	void updateGeometries() { callbackDoingList9733cb_UpdateGeometries(this); };
	void wheelEvent(QWheelEvent * e) { callbackDoingList9733cb_WheelEvent(this, e); };
	QModelIndex indexAt(const QPoint & p) const { return *static_cast<QModelIndex*>(callbackDoingList9733cb_IndexAt(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPoint*>(&p))); };
	QList<QModelIndex> selectedIndexes() const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackDoingList9733cb_SelectedIndexes(const_cast<void*>(static_cast<const void*>(this)))); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QRect visualRect(const QModelIndex & index) const { return *static_cast<QRect*>(callbackDoingList9733cb_VisualRect(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QRegion visualRegionForSelection(const QItemSelection & selection) const { return *static_cast<QRegion*>(callbackDoingList9733cb_VisualRegionForSelection(const_cast<void*>(static_cast<const void*>(this)), const_cast<QItemSelection*>(&selection))); };
	QSize viewportSizeHint() const { return *static_cast<QSize*>(callbackDoingList9733cb_ViewportSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QStyleOptionViewItem viewOptions() const { return *static_cast<QStyleOptionViewItem*>(callbackDoingList9733cb_ViewOptions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isIndexHidden(const QModelIndex & index) const { return callbackDoingList9733cb_IsIndexHidden(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index)) != 0; };
	int horizontalOffset() const { return callbackDoingList9733cb_HorizontalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	int verticalOffset() const { return callbackDoingList9733cb_VerticalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	bool edit(const QModelIndex & index, QAbstractItemView::EditTrigger trigger, QEvent * event) { return callbackDoingList9733cb_Edit2(this, const_cast<QModelIndex*>(&index), trigger, event) != 0; };
	bool eventFilter(QObject * object, QEvent * event) { return callbackDoingList9733cb_EventFilter(this, object, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackDoingList9733cb_FocusNextPrevChild(this, next) != 0; };
	bool viewportEvent(QEvent * event) { return callbackDoingList9733cb_ViewportEvent(this, event) != 0; };
	void Signal_Activated(const QModelIndex & index) { callbackDoingList9733cb_Activated(this, const_cast<QModelIndex*>(&index)); };
	void clearSelection() { callbackDoingList9733cb_ClearSelection(this); };
	void Signal_Clicked(const QModelIndex & index) { callbackDoingList9733cb_Clicked(this, const_cast<QModelIndex*>(&index)); };
	void closeEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint) { callbackDoingList9733cb_CloseEditor(this, editor, hint); };
	void commitData(QWidget * editor) { callbackDoingList9733cb_CommitData(this, editor); };
	void Signal_DoubleClicked(const QModelIndex & index) { callbackDoingList9733cb_DoubleClicked(this, const_cast<QModelIndex*>(&index)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackDoingList9733cb_DragEnterEvent(this, event); };
	void edit(const QModelIndex & index) { callbackDoingList9733cb_Edit(this, const_cast<QModelIndex*>(&index)); };
	void editorDestroyed(QObject * editor) { callbackDoingList9733cb_EditorDestroyed(this, editor); };
	void Signal_Entered(const QModelIndex & index) { callbackDoingList9733cb_Entered(this, const_cast<QModelIndex*>(&index)); };
	void focusInEvent(QFocusEvent * event) { callbackDoingList9733cb_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackDoingList9733cb_FocusOutEvent(this, event); };
	void Signal_IconSizeChanged(const QSize & size) { callbackDoingList9733cb_IconSizeChanged(this, const_cast<QSize*>(&size)); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackDoingList9733cb_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackDoingList9733cb_KeyPressEvent(this, event); };
	void keyboardSearch(const QString & search) { QByteArray t3559d7 = search.toUtf8(); Moc_PackedString searchPacked = { const_cast<char*>(t3559d7.prepend("WHITESPACE").constData()+10), t3559d7.size()-10 };callbackDoingList9733cb_KeyboardSearch(this, searchPacked); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackDoingList9733cb_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackDoingList9733cb_MousePressEvent(this, event); };
	void Signal_Pressed(const QModelIndex & index) { callbackDoingList9733cb_Pressed(this, const_cast<QModelIndex*>(&index)); };
	void reset() { callbackDoingList9733cb_Reset(this); };
	void scrollToBottom() { callbackDoingList9733cb_ScrollToBottom(this); };
	void scrollToTop() { callbackDoingList9733cb_ScrollToTop(this); };
	void selectAll() { callbackDoingList9733cb_SelectAll(this); };
	void setCurrentIndex(const QModelIndex & index) { callbackDoingList9733cb_SetCurrentIndex(this, const_cast<QModelIndex*>(&index)); };
	void setModel(QAbstractItemModel * model) { callbackDoingList9733cb_SetModel(this, model); };
	void setRootIndex(const QModelIndex & index) { callbackDoingList9733cb_SetRootIndex(this, const_cast<QModelIndex*>(&index)); };
	void update(const QModelIndex & index) { callbackDoingList9733cb_Update(this, const_cast<QModelIndex*>(&index)); };
	void Signal_ViewportEntered() { callbackDoingList9733cb_ViewportEntered(this); };
	QItemSelectionModel::SelectionFlags selectionCommand(const QModelIndex & index, const QEvent * event) const { return static_cast<QItemSelectionModel::SelectionFlag>(callbackDoingList9733cb_SelectionCommand(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), const_cast<QEvent*>(event))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackDoingList9733cb_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	int sizeHintForColumn(int column) const { return callbackDoingList9733cb_SizeHintForColumn(const_cast<void*>(static_cast<const void*>(this)), column); };
	int sizeHintForRow(int row) const { return callbackDoingList9733cb_SizeHintForRow(const_cast<void*>(static_cast<const void*>(this)), row); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackDoingList9733cb_ContextMenuEvent(this, e); };
	void scrollContentsBy(int dx, int dy) { callbackDoingList9733cb_ScrollContentsBy(this, dx, dy); };
	void setupViewport(QWidget * viewport) { callbackDoingList9733cb_SetupViewport(this, viewport); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackDoingList9733cb_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackDoingList9733cb_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void changeEvent(QEvent * ev) { callbackDoingList9733cb_ChangeEvent(this, ev); };
	bool close() { return callbackDoingList9733cb_Close(this) != 0; };
	void actionEvent(QActionEvent * event) { callbackDoingList9733cb_ActionEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackDoingList9733cb_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackDoingList9733cb_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void enterEvent(QEvent * event) { callbackDoingList9733cb_EnterEvent(this, event); };
	void hide() { callbackDoingList9733cb_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackDoingList9733cb_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackDoingList9733cb_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackDoingList9733cb_LeaveEvent(this, event); };
	void lower() { callbackDoingList9733cb_Lower(this); };
	void moveEvent(QMoveEvent * event) { callbackDoingList9733cb_MoveEvent(this, event); };
	void raise() { callbackDoingList9733cb_Raise(this); };
	void repaint() { callbackDoingList9733cb_Repaint(this); };
	void setDisabled(bool disable) { callbackDoingList9733cb_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackDoingList9733cb_SetEnabled(this, vbo); };
	void setFocus() { callbackDoingList9733cb_SetFocus2(this); };
	void setHidden(bool hidden) { callbackDoingList9733cb_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackDoingList9733cb_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackDoingList9733cb_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackDoingList9733cb_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackDoingList9733cb_SetWindowTitle(this, vqsPacked); };
	void show() { callbackDoingList9733cb_Show(this); };
	void showEvent(QShowEvent * event) { callbackDoingList9733cb_ShowEvent(this, event); };
	void showFullScreen() { callbackDoingList9733cb_ShowFullScreen(this); };
	void showMaximized() { callbackDoingList9733cb_ShowMaximized(this); };
	void showMinimized() { callbackDoingList9733cb_ShowMinimized(this); };
	void showNormal() { callbackDoingList9733cb_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackDoingList9733cb_TabletEvent(this, event); };
	void updateMicroFocus() { callbackDoingList9733cb_UpdateMicroFocus(this); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackDoingList9733cb_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackDoingList9733cb_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackDoingList9733cb_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackDoingList9733cb_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackDoingList9733cb_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackDoingList9733cb_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void initPainter(QPainter * painter) const { callbackDoingList9733cb_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	void childEvent(QChildEvent * event) { callbackDoingList9733cb_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackDoingList9733cb_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackDoingList9733cb_CustomEvent(this, event); };
	void deleteLater() { callbackDoingList9733cb_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackDoingList9733cb_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackDoingList9733cb_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackDoingList9733cb_ObjectNameChanged(this, objectNamePacked); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(DoingList9733cb*)


void DoingList9733cb_DoingList9733cb_QRegisterMetaTypes() {
}

class Ui_Dialog9733cb: public QObject
{
Q_OBJECT
public:
	Ui_Dialog9733cb(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Ui_Dialog9733cb_Ui_Dialog9733cb_QRegisterMetaType();Ui_Dialog9733cb_Ui_Dialog9733cb_QRegisterMetaTypes();callbackUi_Dialog9733cb_Constructor(this);};
	 ~Ui_Dialog9733cb() { callbackUi_Dialog9733cb_DestroyUi_Dialog(this); };
	bool event(QEvent * e) { return callbackUi_Dialog9733cb_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackUi_Dialog9733cb_EventFilter(this, watched, event) != 0; };
	
	void childEvent(QChildEvent * event) { callbackUi_Dialog9733cb_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackUi_Dialog9733cb_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackUi_Dialog9733cb_CustomEvent(this, event); };
	void deleteLater() { callbackUi_Dialog9733cb_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackUi_Dialog9733cb_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackUi_Dialog9733cb_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackUi_Dialog9733cb_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackUi_Dialog9733cb_TimerEvent(this, event); };
	
signals:
public slots:
	void activateEvent(QApplication* v0, QDialog* v1) { callbackUi_Dialog9733cb_ActivateEvent(this, v0, v1); };
private:
};

Q_DECLARE_METATYPE(Ui_Dialog9733cb*)


void Ui_Dialog9733cb_Ui_Dialog9733cb_QRegisterMetaTypes() {
}

class Ui_SysTray9733cb: public QObject
{
Q_OBJECT
public:
	Ui_SysTray9733cb(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Ui_SysTray9733cb_Ui_SysTray9733cb_QRegisterMetaType();Ui_SysTray9733cb_Ui_SysTray9733cb_QRegisterMetaTypes();callbackUi_SysTray9733cb_Constructor(this);};
	 ~Ui_SysTray9733cb() { callbackUi_SysTray9733cb_DestroyUi_SysTray(this); };
	bool event(QEvent * e) { return callbackUi_SysTray9733cb_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackUi_SysTray9733cb_EventFilter(this, watched, event) != 0; };
	
	void childEvent(QChildEvent * event) { callbackUi_SysTray9733cb_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackUi_SysTray9733cb_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackUi_SysTray9733cb_CustomEvent(this, event); };
	void deleteLater() { callbackUi_SysTray9733cb_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackUi_SysTray9733cb_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackUi_SysTray9733cb_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackUi_SysTray9733cb_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackUi_SysTray9733cb_TimerEvent(this, event); };
	
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(Ui_SysTray9733cb*)


void Ui_SysTray9733cb_Ui_SysTray9733cb_QRegisterMetaTypes() {
}

class AppDoingList9733cb: public QApplication
{
Q_OBJECT
public:
	AppDoingList9733cb(int &argc, char **argv) : QApplication(argc, argv) {qRegisterMetaType<quintptr>("quintptr");AppDoingList9733cb_AppDoingList9733cb_QRegisterMetaType();AppDoingList9733cb_AppDoingList9733cb_QRegisterMetaTypes();callbackAppDoingList9733cb_Constructor(this);};
	void Signal_QuitApp(Ui_Dialog9733cb* ui, AppDoingList9733cb* app, QDialog* mainDialog) { callbackAppDoingList9733cb_QuitApp(this, ui, app, mainDialog); };
	void Signal_ActivateApp(Ui_Dialog9733cb* ui, AppDoingList9733cb* app, QDialog* mainDialog) { callbackAppDoingList9733cb_ActivateApp(this, ui, app, mainDialog); };
	 ~AppDoingList9733cb() { callbackAppDoingList9733cb_DestroyAppDoingList(this); };
	bool event(QEvent * e) { return callbackAppDoingList9733cb_Event(this, e) != 0; };
	
	void aboutQt() { callbackAppDoingList9733cb_AboutQt(this); };
	void closeAllWindows() { callbackAppDoingList9733cb_CloseAllWindows(this); };
	void Signal_FocusChanged(QWidget * old, QWidget * now) { callbackAppDoingList9733cb_FocusChanged(this, old, now); };
	void setAutoSipEnabled(const bool enabled) { callbackAppDoingList9733cb_SetAutoSipEnabled(this, enabled); };
	void setStyleSheet(const QString & sheet) { QByteArray t542ebc = sheet.toUtf8(); Moc_PackedString sheetPacked = { const_cast<char*>(t542ebc.prepend("WHITESPACE").constData()+10), t542ebc.size()-10 };callbackAppDoingList9733cb_SetStyleSheet(this, sheetPacked); };
	bool autoSipEnabled() const { return callbackAppDoingList9733cb_AutoSipEnabled(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	
	void Signal_ApplicationDisplayNameChanged() { callbackAppDoingList9733cb_ApplicationDisplayNameChanged(this); };
	void Signal_ApplicationStateChanged(Qt::ApplicationState state) { callbackAppDoingList9733cb_ApplicationStateChanged(this, state); };
	void Signal_CommitDataRequest(QSessionManager & manager) { callbackAppDoingList9733cb_CommitDataRequest(this, static_cast<QSessionManager*>(&manager)); };
	void Signal_FocusObjectChanged(QObject * focusObject) { callbackAppDoingList9733cb_FocusObjectChanged(this, focusObject); };
	void Signal_FocusWindowChanged(QWindow * focusWindow) { callbackAppDoingList9733cb_FocusWindowChanged(this, focusWindow); };
	void Signal_FontChanged(const QFont & font) { callbackAppDoingList9733cb_FontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_FontDatabaseChanged() { callbackAppDoingList9733cb_FontDatabaseChanged(this); };
	void Signal_LastWindowClosed() { callbackAppDoingList9733cb_LastWindowClosed(this); };
	void Signal_LayoutDirectionChanged(Qt::LayoutDirection direction) { callbackAppDoingList9733cb_LayoutDirectionChanged(this, direction); };
	void Signal_PaletteChanged(const QPalette & palette) { callbackAppDoingList9733cb_PaletteChanged(this, const_cast<QPalette*>(&palette)); };
	void Signal_PrimaryScreenChanged(QScreen * screen) { callbackAppDoingList9733cb_PrimaryScreenChanged(this, screen); };
	void Signal_SaveStateRequest(QSessionManager & manager) { callbackAppDoingList9733cb_SaveStateRequest(this, static_cast<QSessionManager*>(&manager)); };
	void Signal_ScreenAdded(QScreen * screen) { callbackAppDoingList9733cb_ScreenAdded(this, screen); };
	void Signal_ScreenRemoved(QScreen * screen) { callbackAppDoingList9733cb_ScreenRemoved(this, screen); };
	void Signal_AboutToQuit() { callbackAppDoingList9733cb_AboutToQuit(this); };
	void quit() { callbackAppDoingList9733cb_Quit(this); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAppDoingList9733cb_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackAppDoingList9733cb_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAppDoingList9733cb_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAppDoingList9733cb_CustomEvent(this, event); };
	void deleteLater() { callbackAppDoingList9733cb_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAppDoingList9733cb_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAppDoingList9733cb_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAppDoingList9733cb_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAppDoingList9733cb_TimerEvent(this, event); };
signals:
	void quitApp(Ui_Dialog9733cb* ui, AppDoingList9733cb* app, QDialog* mainDialog);
	void activateApp(Ui_Dialog9733cb* ui, AppDoingList9733cb* app, QDialog* mainDialog);
public slots:
private:
};

Q_DECLARE_METATYPE(AppDoingList9733cb*)


void AppDoingList9733cb_AppDoingList9733cb_QRegisterMetaTypes() {
}

int DoingList9733cb_DoingList9733cb_QRegisterMetaType()
{
	return qRegisterMetaType<DoingList9733cb*>();
}

int DoingList9733cb_DoingList9733cb_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<DoingList9733cb*>(const_cast<const char*>(typeName));
}

int DoingList9733cb_DoingList9733cb_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<DoingList9733cb>();
#else
	return 0;
#endif
}

int DoingList9733cb_DoingList9733cb_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<DoingList9733cb>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* DoingList9733cb___findItems_atList(void* ptr, int i)
{
	return ({QListWidgetItem * tmp = static_cast<QList<QListWidgetItem *>*>(ptr)->at(i); if (i == static_cast<QList<QListWidgetItem *>*>(ptr)->size()-1) { static_cast<QList<QListWidgetItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___findItems_setList(void* ptr, void* i)
{
	static_cast<QList<QListWidgetItem *>*>(ptr)->append(static_cast<QListWidgetItem*>(i));
}

void* DoingList9733cb___findItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QListWidgetItem *>();
}

void* DoingList9733cb___items_atList(void* ptr, int i)
{
	return ({QListWidgetItem * tmp = static_cast<QList<QListWidgetItem *>*>(ptr)->at(i); if (i == static_cast<QList<QListWidgetItem *>*>(ptr)->size()-1) { static_cast<QList<QListWidgetItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___items_setList(void* ptr, void* i)
{
	static_cast<QList<QListWidgetItem *>*>(ptr)->append(static_cast<QListWidgetItem*>(i));
}

void* DoingList9733cb___items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QListWidgetItem *>();
}

void* DoingList9733cb___selectedItems_atList(void* ptr, int i)
{
	return ({QListWidgetItem * tmp = static_cast<QList<QListWidgetItem *>*>(ptr)->at(i); if (i == static_cast<QList<QListWidgetItem *>*>(ptr)->size()-1) { static_cast<QList<QListWidgetItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___selectedItems_setList(void* ptr, void* i)
{
	static_cast<QList<QListWidgetItem *>*>(ptr)->append(static_cast<QListWidgetItem*>(i));
}

void* DoingList9733cb___selectedItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QListWidgetItem *>();
}

void* DoingList9733cb___mimeData_items_atList(void* ptr, int i)
{
	return ({QListWidgetItem * tmp = static_cast<QList<QListWidgetItem *>*>(ptr)->at(i); if (i == static_cast<QList<QListWidgetItem *>*>(ptr)->size()-1) { static_cast<QList<QListWidgetItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___mimeData_items_setList(void* ptr, void* i)
{
	static_cast<QList<QListWidgetItem *>*>(ptr)->append(static_cast<QListWidgetItem*>(i));
}

void* DoingList9733cb___mimeData_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QListWidgetItem *>();
}

int DoingList9733cb___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void DoingList9733cb___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* DoingList9733cb___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* DoingList9733cb___indexesMoved_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void DoingList9733cb___indexesMoved_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* DoingList9733cb___indexesMoved_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* DoingList9733cb___selectedIndexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void DoingList9733cb___selectedIndexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* DoingList9733cb___selectedIndexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* DoingList9733cb___scrollBarWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___scrollBarWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* DoingList9733cb___scrollBarWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* DoingList9733cb___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* DoingList9733cb___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* DoingList9733cb___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* DoingList9733cb___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* DoingList9733cb___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* DoingList9733cb___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* DoingList9733cb___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void DoingList9733cb___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* DoingList9733cb___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* DoingList9733cb___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DoingList9733cb___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* DoingList9733cb___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DoingList9733cb___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* DoingList9733cb___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DoingList9733cb___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* DoingList9733cb___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void DoingList9733cb___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* DoingList9733cb___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* DoingList9733cb_NewDoingList(void* parent)
{
		return new DoingList9733cb(static_cast<QWidget*>(parent));
}

void DoingList9733cb_DestroyDoingList(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->~DoingList9733cb();
}

void DoingList9733cb_DestroyDoingListDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char DoingList9733cb_DropMimeDataDefault(void* ptr, int index, void* data, long long action)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::dropMimeData(index, static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action));
}

char DoingList9733cb_EventDefault(void* ptr, void* e)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::event(static_cast<QEvent*>(e));
}

void DoingList9733cb_ClearDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::clear();
}

void DoingList9733cb_DropEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void DoingList9733cb_ScrollToItemDefault(void* ptr, void* item, long long hint)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::scrollToItem(static_cast<QListWidgetItem*>(item), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void DoingList9733cb_SetSelectionModelDefault(void* ptr, void* selectionModel)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void* DoingList9733cb_MimeDataDefault(void* ptr, void* items)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::mimeData(*static_cast<QList<QListWidgetItem *>*>(items));
}

struct Moc_PackedString DoingList9733cb_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray tcde693 = static_cast<DoingList9733cb*>(ptr)->QListWidget::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(tcde693.prepend("WHITESPACE").constData()+10), tcde693.size()-10 }; });
}

long long DoingList9733cb_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::supportedDropActions();
}



void* DoingList9733cb_MoveCursorDefault(void* ptr, long long cursorAction, long long modifiers)
{
	return new QModelIndex(static_cast<DoingList9733cb*>(ptr)->QListWidget::moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

void DoingList9733cb_CurrentChangedDefault(void* ptr, void* current, void* previous)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::currentChanged(*static_cast<QModelIndex*>(current), *static_cast<QModelIndex*>(previous));
}

void DoingList9733cb_DataChangedDefault(void* ptr, void* topLeft, void* bottomRight, void* roles)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::dataChanged(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight), *static_cast<QVector<int>*>(roles));
}

void DoingList9733cb_DragLeaveEventDefault(void* ptr, void* e)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void DoingList9733cb_DragMoveEventDefault(void* ptr, void* e)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void DoingList9733cb_MouseMoveEventDefault(void* ptr, void* e)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void DoingList9733cb_MouseReleaseEventDefault(void* ptr, void* e)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void DoingList9733cb_PaintEventDefault(void* ptr, void* e)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::paintEvent(static_cast<QPaintEvent*>(e));
}

void DoingList9733cb_ResizeEventDefault(void* ptr, void* e)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::resizeEvent(static_cast<QResizeEvent*>(e));
}

void DoingList9733cb_RowsAboutToBeRemovedDefault(void* ptr, void* parent, int start, int end)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::rowsAboutToBeRemoved(*static_cast<QModelIndex*>(parent), start, end);
}

void DoingList9733cb_RowsInsertedDefault(void* ptr, void* parent, int start, int end)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::rowsInserted(*static_cast<QModelIndex*>(parent), start, end);
}

void DoingList9733cb_ScrollToDefault(void* ptr, void* index, long long hint)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void DoingList9733cb_SelectionChangedDefault(void* ptr, void* selected, void* deselected)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
}

void DoingList9733cb_SetSelectionDefault(void* ptr, void* rect, long long command)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void DoingList9733cb_StartDragDefault(void* ptr, long long supportedActions)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::startDrag(static_cast<Qt::DropAction>(supportedActions));
}

void DoingList9733cb_TimerEventDefault(void* ptr, void* e)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::timerEvent(static_cast<QTimerEvent*>(e));
}

void DoingList9733cb_UpdateGeometriesDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::updateGeometries();
}

void DoingList9733cb_WheelEventDefault(void* ptr, void* e)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::wheelEvent(static_cast<QWheelEvent*>(e));
}

void* DoingList9733cb_IndexAtDefault(void* ptr, void* p)
{
	return new QModelIndex(static_cast<DoingList9733cb*>(ptr)->QListWidget::indexAt(*static_cast<QPoint*>(p)));
}

struct Moc_PackedList DoingList9733cb_SelectedIndexesDefault(void* ptr)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<DoingList9733cb*>(ptr)->QListWidget::selectedIndexes()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* DoingList9733cb_VisualRectDefault(void* ptr, void* index)
{
	return ({ QRect tmpValue = static_cast<DoingList9733cb*>(ptr)->QListWidget::visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* DoingList9733cb_VisualRegionForSelectionDefault(void* ptr, void* selection)
{
	return new QRegion(static_cast<DoingList9733cb*>(ptr)->QListWidget::visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

void* DoingList9733cb_ViewportSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<DoingList9733cb*>(ptr)->QListWidget::viewportSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* DoingList9733cb_ViewOptionsDefault(void* ptr)
{
	return new QStyleOptionViewItem(static_cast<DoingList9733cb*>(ptr)->QListWidget::viewOptions());
}

char DoingList9733cb_IsIndexHiddenDefault(void* ptr, void* index)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::isIndexHidden(*static_cast<QModelIndex*>(index));
}

int DoingList9733cb_HorizontalOffsetDefault(void* ptr)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::horizontalOffset();
}

int DoingList9733cb_VerticalOffsetDefault(void* ptr)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::verticalOffset();
}

char DoingList9733cb_Edit2Default(void* ptr, void* index, long long trigger, void* event)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::edit(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::EditTrigger>(trigger), static_cast<QEvent*>(event));
}

char DoingList9733cb_EventFilterDefault(void* ptr, void* object, void* event)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::eventFilter(static_cast<QObject*>(object), static_cast<QEvent*>(event));
}

char DoingList9733cb_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::focusNextPrevChild(next != 0);
}

char DoingList9733cb_ViewportEventDefault(void* ptr, void* event)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::viewportEvent(static_cast<QEvent*>(event));
}

void DoingList9733cb_ClearSelectionDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::clearSelection();
}

void DoingList9733cb_CloseEditorDefault(void* ptr, void* editor, long long hint)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::closeEditor(static_cast<QWidget*>(editor), static_cast<QAbstractItemDelegate::EndEditHint>(hint));
}

void DoingList9733cb_CommitDataDefault(void* ptr, void* editor)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::commitData(static_cast<QWidget*>(editor));
}

void DoingList9733cb_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void DoingList9733cb_EditDefault(void* ptr, void* index)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::edit(*static_cast<QModelIndex*>(index));
}

void DoingList9733cb_EditorDestroyedDefault(void* ptr, void* editor)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::editorDestroyed(static_cast<QObject*>(editor));
}

void DoingList9733cb_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void DoingList9733cb_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void DoingList9733cb_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void DoingList9733cb_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void DoingList9733cb_KeyboardSearchDefault(void* ptr, struct Moc_PackedString search)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::keyboardSearch(QString::fromUtf8(search.data, search.len));
}

void DoingList9733cb_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void DoingList9733cb_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void DoingList9733cb_ResetDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::reset();
}

void DoingList9733cb_ScrollToBottomDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::scrollToBottom();
}

void DoingList9733cb_ScrollToTopDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::scrollToTop();
}

void DoingList9733cb_SelectAllDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::selectAll();
}

void DoingList9733cb_SetCurrentIndexDefault(void* ptr, void* index)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setCurrentIndex(*static_cast<QModelIndex*>(index));
}

void DoingList9733cb_SetModelDefault(void* ptr, void* model)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setModel(static_cast<QAbstractItemModel*>(model));
}

void DoingList9733cb_SetRootIndexDefault(void* ptr, void* index)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setRootIndex(*static_cast<QModelIndex*>(index));
}

void DoingList9733cb_UpdateDefault(void* ptr, void* index)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::update(*static_cast<QModelIndex*>(index));
}

long long DoingList9733cb_SelectionCommandDefault(void* ptr, void* index, void* event)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::selectionCommand(*static_cast<QModelIndex*>(index), static_cast<QEvent*>(event));
}

void* DoingList9733cb_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<DoingList9733cb*>(ptr)->QListWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

int DoingList9733cb_SizeHintForColumnDefault(void* ptr, int column)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::sizeHintForColumn(column);
}

int DoingList9733cb_SizeHintForRowDefault(void* ptr, int row)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::sizeHintForRow(row);
}

void DoingList9733cb_ContextMenuEventDefault(void* ptr, void* e)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void DoingList9733cb_ScrollContentsByDefault(void* ptr, int dx, int dy)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::scrollContentsBy(dx, dy);
}

void DoingList9733cb_SetupViewportDefault(void* ptr, void* viewport)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setupViewport(static_cast<QWidget*>(viewport));
}

void* DoingList9733cb_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<DoingList9733cb*>(ptr)->QListWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* DoingList9733cb_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<DoingList9733cb*>(ptr)->QListWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void DoingList9733cb_ChangeEventDefault(void* ptr, void* ev)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::changeEvent(static_cast<QEvent*>(ev));
}

char DoingList9733cb_CloseDefault(void* ptr)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::close();
}

void DoingList9733cb_ActionEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void DoingList9733cb_CloseEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void DoingList9733cb_EnterEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::enterEvent(static_cast<QEvent*>(event));
}

void DoingList9733cb_HideDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::hide();
}

void DoingList9733cb_HideEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void DoingList9733cb_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void DoingList9733cb_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::leaveEvent(static_cast<QEvent*>(event));
}

void DoingList9733cb_LowerDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::lower();
}

void DoingList9733cb_MoveEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void DoingList9733cb_RaiseDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::raise();
}

void DoingList9733cb_RepaintDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::repaint();
}

void DoingList9733cb_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setDisabled(disable != 0);
}

void DoingList9733cb_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setEnabled(vbo != 0);
}

void DoingList9733cb_SetFocus2Default(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setFocus();
}

void DoingList9733cb_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setHidden(hidden != 0);
}

void DoingList9733cb_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void DoingList9733cb_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setVisible(visible != 0);
}

void DoingList9733cb_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setWindowModified(vbo != 0);
}

void DoingList9733cb_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void DoingList9733cb_ShowDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::show();
}

void DoingList9733cb_ShowEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::showEvent(static_cast<QShowEvent*>(event));
}

void DoingList9733cb_ShowFullScreenDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::showFullScreen();
}

void DoingList9733cb_ShowMaximizedDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::showMaximized();
}

void DoingList9733cb_ShowMinimizedDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::showMinimized();
}

void DoingList9733cb_ShowNormalDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::showNormal();
}

void DoingList9733cb_TabletEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void DoingList9733cb_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::updateMicroFocus();
}

void* DoingList9733cb_PaintEngineDefault(void* ptr)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::paintEngine();
}

char DoingList9733cb_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::hasHeightForWidth();
}

int DoingList9733cb_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::heightForWidth(w);
}

int DoingList9733cb_MetricDefault(void* ptr, long long m)
{
	return static_cast<DoingList9733cb*>(ptr)->QListWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void DoingList9733cb_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::initPainter(static_cast<QPainter*>(painter));
}

void DoingList9733cb_ChildEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::childEvent(static_cast<QChildEvent*>(event));
}

void DoingList9733cb_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void DoingList9733cb_CustomEventDefault(void* ptr, void* event)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::customEvent(static_cast<QEvent*>(event));
}

void DoingList9733cb_DeleteLaterDefault(void* ptr)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::deleteLater();
}

void DoingList9733cb_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<DoingList9733cb*>(ptr)->QListWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Ui_Dialog9733cb_ActivateEvent(void* ptr, void* v0, void* v1)
{
	QMetaObject::invokeMethod(static_cast<Ui_Dialog9733cb*>(ptr), "activateEvent", Q_ARG(QApplication*, static_cast<QApplication*>(v0)), Q_ARG(QDialog*, static_cast<QDialog*>(v1)));
}

int Ui_Dialog9733cb_Ui_Dialog9733cb_QRegisterMetaType()
{
	return qRegisterMetaType<Ui_Dialog9733cb*>();
}

int Ui_Dialog9733cb_Ui_Dialog9733cb_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Ui_Dialog9733cb*>(const_cast<const char*>(typeName));
}

int Ui_Dialog9733cb_Ui_Dialog9733cb_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Ui_Dialog9733cb>();
#else
	return 0;
#endif
}

int Ui_Dialog9733cb_Ui_Dialog9733cb_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Ui_Dialog9733cb>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Ui_Dialog9733cb___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Ui_Dialog9733cb___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Ui_Dialog9733cb___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Ui_Dialog9733cb___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Ui_Dialog9733cb___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Ui_Dialog9733cb___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Ui_Dialog9733cb___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Ui_Dialog9733cb___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Ui_Dialog9733cb___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Ui_Dialog9733cb___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Ui_Dialog9733cb___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Ui_Dialog9733cb___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Ui_Dialog9733cb___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Ui_Dialog9733cb___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Ui_Dialog9733cb___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Ui_Dialog9733cb_NewUi_Dialog(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Ui_Dialog9733cb(static_cast<QWindow*>(parent));
	} else {
		return new Ui_Dialog9733cb(static_cast<QObject*>(parent));
	}
}

void Ui_Dialog9733cb_DestroyUi_Dialog(void* ptr)
{
	static_cast<Ui_Dialog9733cb*>(ptr)->~Ui_Dialog9733cb();
}

void Ui_Dialog9733cb_DestroyUi_DialogDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Ui_Dialog9733cb_EventDefault(void* ptr, void* e)
{
	return static_cast<Ui_Dialog9733cb*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Ui_Dialog9733cb_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Ui_Dialog9733cb*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Ui_Dialog9733cb_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Ui_Dialog9733cb*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Ui_Dialog9733cb_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Ui_Dialog9733cb*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Ui_Dialog9733cb_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Ui_Dialog9733cb*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Ui_Dialog9733cb_DeleteLaterDefault(void* ptr)
{
	static_cast<Ui_Dialog9733cb*>(ptr)->QObject::deleteLater();
}

void Ui_Dialog9733cb_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Ui_Dialog9733cb*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Ui_Dialog9733cb_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Ui_Dialog9733cb*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



int Ui_SysTray9733cb_Ui_SysTray9733cb_QRegisterMetaType()
{
	return qRegisterMetaType<Ui_SysTray9733cb*>();
}

int Ui_SysTray9733cb_Ui_SysTray9733cb_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Ui_SysTray9733cb*>(const_cast<const char*>(typeName));
}

int Ui_SysTray9733cb_Ui_SysTray9733cb_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Ui_SysTray9733cb>();
#else
	return 0;
#endif
}

int Ui_SysTray9733cb_Ui_SysTray9733cb_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Ui_SysTray9733cb>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Ui_SysTray9733cb___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Ui_SysTray9733cb___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Ui_SysTray9733cb___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Ui_SysTray9733cb___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Ui_SysTray9733cb___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Ui_SysTray9733cb___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Ui_SysTray9733cb___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Ui_SysTray9733cb___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Ui_SysTray9733cb___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Ui_SysTray9733cb___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Ui_SysTray9733cb___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Ui_SysTray9733cb___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Ui_SysTray9733cb___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Ui_SysTray9733cb___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Ui_SysTray9733cb___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Ui_SysTray9733cb_NewUi_SysTray(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Ui_SysTray9733cb(static_cast<QWindow*>(parent));
	} else {
		return new Ui_SysTray9733cb(static_cast<QObject*>(parent));
	}
}

void Ui_SysTray9733cb_DestroyUi_SysTray(void* ptr)
{
	static_cast<Ui_SysTray9733cb*>(ptr)->~Ui_SysTray9733cb();
}

void Ui_SysTray9733cb_DestroyUi_SysTrayDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Ui_SysTray9733cb_EventDefault(void* ptr, void* e)
{
	return static_cast<Ui_SysTray9733cb*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Ui_SysTray9733cb_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Ui_SysTray9733cb*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Ui_SysTray9733cb_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Ui_SysTray9733cb*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Ui_SysTray9733cb_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Ui_SysTray9733cb*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Ui_SysTray9733cb_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Ui_SysTray9733cb*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Ui_SysTray9733cb_DeleteLaterDefault(void* ptr)
{
	static_cast<Ui_SysTray9733cb*>(ptr)->QObject::deleteLater();
}

void Ui_SysTray9733cb_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Ui_SysTray9733cb*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Ui_SysTray9733cb_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Ui_SysTray9733cb*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void AppDoingList9733cb_ConnectQuitApp(void* ptr)
{
	QObject::connect(static_cast<AppDoingList9733cb*>(ptr), static_cast<void (AppDoingList9733cb::*)(Ui_Dialog9733cb*, AppDoingList9733cb*, QDialog*)>(&AppDoingList9733cb::quitApp), static_cast<AppDoingList9733cb*>(ptr), static_cast<void (AppDoingList9733cb::*)(Ui_Dialog9733cb*, AppDoingList9733cb*, QDialog*)>(&AppDoingList9733cb::Signal_QuitApp));
}

void AppDoingList9733cb_DisconnectQuitApp(void* ptr)
{
	QObject::disconnect(static_cast<AppDoingList9733cb*>(ptr), static_cast<void (AppDoingList9733cb::*)(Ui_Dialog9733cb*, AppDoingList9733cb*, QDialog*)>(&AppDoingList9733cb::quitApp), static_cast<AppDoingList9733cb*>(ptr), static_cast<void (AppDoingList9733cb::*)(Ui_Dialog9733cb*, AppDoingList9733cb*, QDialog*)>(&AppDoingList9733cb::Signal_QuitApp));
}

void AppDoingList9733cb_QuitApp(void* ptr, void* ui, void* app, void* mainDialog)
{
	static_cast<AppDoingList9733cb*>(ptr)->quitApp(static_cast<Ui_Dialog9733cb*>(ui), static_cast<AppDoingList9733cb*>(app), static_cast<QDialog*>(mainDialog));
}

void AppDoingList9733cb_ConnectActivateApp(void* ptr)
{
	QObject::connect(static_cast<AppDoingList9733cb*>(ptr), static_cast<void (AppDoingList9733cb::*)(Ui_Dialog9733cb*, AppDoingList9733cb*, QDialog*)>(&AppDoingList9733cb::activateApp), static_cast<AppDoingList9733cb*>(ptr), static_cast<void (AppDoingList9733cb::*)(Ui_Dialog9733cb*, AppDoingList9733cb*, QDialog*)>(&AppDoingList9733cb::Signal_ActivateApp));
}

void AppDoingList9733cb_DisconnectActivateApp(void* ptr)
{
	QObject::disconnect(static_cast<AppDoingList9733cb*>(ptr), static_cast<void (AppDoingList9733cb::*)(Ui_Dialog9733cb*, AppDoingList9733cb*, QDialog*)>(&AppDoingList9733cb::activateApp), static_cast<AppDoingList9733cb*>(ptr), static_cast<void (AppDoingList9733cb::*)(Ui_Dialog9733cb*, AppDoingList9733cb*, QDialog*)>(&AppDoingList9733cb::Signal_ActivateApp));
}

void AppDoingList9733cb_ActivateApp(void* ptr, void* ui, void* app, void* mainDialog)
{
	static_cast<AppDoingList9733cb*>(ptr)->activateApp(static_cast<Ui_Dialog9733cb*>(ui), static_cast<AppDoingList9733cb*>(app), static_cast<QDialog*>(mainDialog));
}

int AppDoingList9733cb_AppDoingList9733cb_QRegisterMetaType()
{
	return qRegisterMetaType<AppDoingList9733cb*>();
}

int AppDoingList9733cb_AppDoingList9733cb_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AppDoingList9733cb*>(const_cast<const char*>(typeName));
}

void* AppDoingList9733cb___allWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AppDoingList9733cb___allWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* AppDoingList9733cb___allWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* AppDoingList9733cb___topLevelWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AppDoingList9733cb___topLevelWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* AppDoingList9733cb___topLevelWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* AppDoingList9733cb___screens_atList(void* ptr, int i)
{
	return ({QScreen * tmp = static_cast<QList<QScreen *>*>(ptr)->at(i); if (i == static_cast<QList<QScreen *>*>(ptr)->size()-1) { static_cast<QList<QScreen *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AppDoingList9733cb___screens_setList(void* ptr, void* i)
{
	static_cast<QList<QScreen *>*>(ptr)->append(static_cast<QScreen*>(i));
}

void* AppDoingList9733cb___screens_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QScreen *>();
}

void* AppDoingList9733cb___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AppDoingList9733cb___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AppDoingList9733cb___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AppDoingList9733cb___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AppDoingList9733cb___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AppDoingList9733cb___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AppDoingList9733cb___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AppDoingList9733cb___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AppDoingList9733cb___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AppDoingList9733cb___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AppDoingList9733cb___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AppDoingList9733cb___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AppDoingList9733cb___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AppDoingList9733cb___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AppDoingList9733cb___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AppDoingList9733cb_NewAppDoingList(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return new AppDoingList9733cb(argcs, argvs);
}

void AppDoingList9733cb_DestroyAppDoingList(void* ptr)
{
	static_cast<AppDoingList9733cb*>(ptr)->~AppDoingList9733cb();
}

void AppDoingList9733cb_DestroyAppDoingListDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char AppDoingList9733cb_EventDefault(void* ptr, void* e)
{
	return static_cast<AppDoingList9733cb*>(ptr)->QApplication::event(static_cast<QEvent*>(e));
}

void AppDoingList9733cb_AboutQtDefault(void* ptr)
{
	static_cast<AppDoingList9733cb*>(ptr)->QApplication::aboutQt();
}

void AppDoingList9733cb_CloseAllWindowsDefault(void* ptr)
{
	static_cast<AppDoingList9733cb*>(ptr)->QApplication::closeAllWindows();
}

void AppDoingList9733cb_SetAutoSipEnabledDefault(void* ptr, char enabled)
{
	static_cast<AppDoingList9733cb*>(ptr)->QApplication::setAutoSipEnabled(enabled != 0);
}

void AppDoingList9733cb_SetStyleSheetDefault(void* ptr, struct Moc_PackedString sheet)
{
	static_cast<AppDoingList9733cb*>(ptr)->QApplication::setStyleSheet(QString::fromUtf8(sheet.data, sheet.len));
}

char AppDoingList9733cb_AutoSipEnabledDefault(void* ptr)
{
	return static_cast<AppDoingList9733cb*>(ptr)->QApplication::autoSipEnabled();
}



void AppDoingList9733cb_QuitDefault(void* ptr)
{
	static_cast<AppDoingList9733cb*>(ptr)->QApplication::quit();
}

char AppDoingList9733cb_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AppDoingList9733cb*>(ptr)->QApplication::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AppDoingList9733cb_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AppDoingList9733cb*>(ptr)->QApplication::childEvent(static_cast<QChildEvent*>(event));
}

void AppDoingList9733cb_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AppDoingList9733cb*>(ptr)->QApplication::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AppDoingList9733cb_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AppDoingList9733cb*>(ptr)->QApplication::customEvent(static_cast<QEvent*>(event));
}

void AppDoingList9733cb_DeleteLaterDefault(void* ptr)
{
	static_cast<AppDoingList9733cb*>(ptr)->QApplication::deleteLater();
}

void AppDoingList9733cb_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AppDoingList9733cb*>(ptr)->QApplication::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void AppDoingList9733cb_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AppDoingList9733cb*>(ptr)->QApplication::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
