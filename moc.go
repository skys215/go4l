package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
	std_gui "github.com/therecipe/qt/gui"
	std_widgets "github.com/therecipe/qt/widgets"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type Ui_Dialog_ITF interface {
	std_core.QObject_ITF
	Ui_Dialog_PTR() *Ui_Dialog
}

func (ptr *Ui_Dialog) Ui_Dialog_PTR() *Ui_Dialog {
	return ptr
}

func (ptr *Ui_Dialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Ui_Dialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromUi_Dialog(ptr Ui_Dialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Ui_Dialog_PTR().Pointer()
	}
	return nil
}

func NewUi_DialogFromPointer(ptr unsafe.Pointer) (n *Ui_Dialog) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Ui_Dialog)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Ui_Dialog:
			n = deduced

		case *std_core.QObject:
			n = &Ui_Dialog{QObject: *deduced}

		default:
			n = new(Ui_Dialog)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackUi_Dialog9733cb_Constructor
func callbackUi_Dialog9733cb_Constructor(ptr unsafe.Pointer) {
	this := NewUi_DialogFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackUi_Dialog9733cb_ActivateEvent
func callbackUi_Dialog9733cb_ActivateEvent(ptr unsafe.Pointer, v0 unsafe.Pointer, v1 unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activateEvent"); signal != nil {
		signal.(func(*std_widgets.QApplication, *std_widgets.QDialog))(std_widgets.NewQApplicationFromPointer(v0), std_widgets.NewQDialogFromPointer(v1))
	}

}

func (ptr *Ui_Dialog) ConnectActivateEvent(f func(v0 *std_widgets.QApplication, v1 *std_widgets.QDialog)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "activateEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activateEvent", func(v0 *std_widgets.QApplication, v1 *std_widgets.QDialog) {
				signal.(func(*std_widgets.QApplication, *std_widgets.QDialog))(v0, v1)
				f(v0, v1)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activateEvent", f)
		}
	}
}

func (ptr *Ui_Dialog) DisconnectActivateEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "activateEvent")
	}
}

func (ptr *Ui_Dialog) ActivateEvent(v0 std_widgets.QApplication_ITF, v1 std_widgets.QDialog_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb_ActivateEvent(ptr.Pointer(), std_widgets.PointerFromQApplication(v0), std_widgets.PointerFromQDialog(v1))
	}
}

func Ui_Dialog_QRegisterMetaType() int {
	return int(int32(C.Ui_Dialog9733cb_Ui_Dialog9733cb_QRegisterMetaType()))
}

func (ptr *Ui_Dialog) QRegisterMetaType() int {
	return int(int32(C.Ui_Dialog9733cb_Ui_Dialog9733cb_QRegisterMetaType()))
}

func Ui_Dialog_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Ui_Dialog9733cb_Ui_Dialog9733cb_QRegisterMetaType2(typeNameC)))
}

func (ptr *Ui_Dialog) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Ui_Dialog9733cb_Ui_Dialog9733cb_QRegisterMetaType2(typeNameC)))
}

func Ui_Dialog_QmlRegisterType() int {
	return int(int32(C.Ui_Dialog9733cb_Ui_Dialog9733cb_QmlRegisterType()))
}

func (ptr *Ui_Dialog) QmlRegisterType() int {
	return int(int32(C.Ui_Dialog9733cb_Ui_Dialog9733cb_QmlRegisterType()))
}

func Ui_Dialog_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Ui_Dialog9733cb_Ui_Dialog9733cb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Ui_Dialog) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Ui_Dialog9733cb_Ui_Dialog9733cb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Ui_Dialog) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Ui_Dialog9733cb___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Ui_Dialog) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Ui_Dialog) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Ui_Dialog9733cb___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Ui_Dialog) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Ui_Dialog9733cb___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Ui_Dialog) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Ui_Dialog) __findChildren_newList2() unsafe.Pointer {
	return C.Ui_Dialog9733cb___findChildren_newList2(ptr.Pointer())
}

func (ptr *Ui_Dialog) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Ui_Dialog9733cb___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Ui_Dialog) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Ui_Dialog) __findChildren_newList3() unsafe.Pointer {
	return C.Ui_Dialog9733cb___findChildren_newList3(ptr.Pointer())
}

func (ptr *Ui_Dialog) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Ui_Dialog9733cb___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Ui_Dialog) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Ui_Dialog) __findChildren_newList() unsafe.Pointer {
	return C.Ui_Dialog9733cb___findChildren_newList(ptr.Pointer())
}

func (ptr *Ui_Dialog) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Ui_Dialog9733cb___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Ui_Dialog) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Ui_Dialog) __children_newList() unsafe.Pointer {
	return C.Ui_Dialog9733cb___children_newList(ptr.Pointer())
}

func NewUi_Dialog(parent std_core.QObject_ITF) *Ui_Dialog {
	tmpValue := NewUi_DialogFromPointer(C.Ui_Dialog9733cb_NewUi_Dialog(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackUi_Dialog9733cb_DestroyUi_Dialog
func callbackUi_Dialog9733cb_DestroyUi_Dialog(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Ui_Dialog"); signal != nil {
		signal.(func())()
	} else {
		NewUi_DialogFromPointer(ptr).DestroyUi_DialogDefault()
	}
}

func (ptr *Ui_Dialog) ConnectDestroyUi_Dialog(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Ui_Dialog"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Ui_Dialog", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Ui_Dialog", f)
		}
	}
}

func (ptr *Ui_Dialog) DisconnectDestroyUi_Dialog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Ui_Dialog")
	}
}

func (ptr *Ui_Dialog) DestroyUi_Dialog() {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb_DestroyUi_Dialog(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Ui_Dialog) DestroyUi_DialogDefault() {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb_DestroyUi_DialogDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackUi_Dialog9733cb_Event
func callbackUi_Dialog9733cb_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewUi_DialogFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Ui_Dialog) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Ui_Dialog9733cb_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackUi_Dialog9733cb_EventFilter
func callbackUi_Dialog9733cb_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewUi_DialogFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Ui_Dialog) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Ui_Dialog9733cb_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackUi_Dialog9733cb_ChildEvent
func callbackUi_Dialog9733cb_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewUi_DialogFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Ui_Dialog) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackUi_Dialog9733cb_ConnectNotify
func callbackUi_Dialog9733cb_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewUi_DialogFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Ui_Dialog) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackUi_Dialog9733cb_CustomEvent
func callbackUi_Dialog9733cb_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewUi_DialogFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Ui_Dialog) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackUi_Dialog9733cb_DeleteLater
func callbackUi_Dialog9733cb_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewUi_DialogFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Ui_Dialog) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackUi_Dialog9733cb_Destroyed
func callbackUi_Dialog9733cb_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackUi_Dialog9733cb_DisconnectNotify
func callbackUi_Dialog9733cb_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewUi_DialogFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Ui_Dialog) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackUi_Dialog9733cb_ObjectNameChanged
func callbackUi_Dialog9733cb_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackUi_Dialog9733cb_TimerEvent
func callbackUi_Dialog9733cb_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewUi_DialogFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Ui_Dialog) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_Dialog9733cb_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type Ui_SysTray_ITF interface {
	std_core.QObject_ITF
	Ui_SysTray_PTR() *Ui_SysTray
}

func (ptr *Ui_SysTray) Ui_SysTray_PTR() *Ui_SysTray {
	return ptr
}

func (ptr *Ui_SysTray) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Ui_SysTray) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromUi_SysTray(ptr Ui_SysTray_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Ui_SysTray_PTR().Pointer()
	}
	return nil
}

func NewUi_SysTrayFromPointer(ptr unsafe.Pointer) (n *Ui_SysTray) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Ui_SysTray)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Ui_SysTray:
			n = deduced

		case *std_core.QObject:
			n = &Ui_SysTray{QObject: *deduced}

		default:
			n = new(Ui_SysTray)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackUi_SysTray9733cb_Constructor
func callbackUi_SysTray9733cb_Constructor(ptr unsafe.Pointer) {
	this := NewUi_SysTrayFromPointer(ptr)
	qt.Register(ptr, this)
}

func Ui_SysTray_QRegisterMetaType() int {
	return int(int32(C.Ui_SysTray9733cb_Ui_SysTray9733cb_QRegisterMetaType()))
}

func (ptr *Ui_SysTray) QRegisterMetaType() int {
	return int(int32(C.Ui_SysTray9733cb_Ui_SysTray9733cb_QRegisterMetaType()))
}

func Ui_SysTray_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Ui_SysTray9733cb_Ui_SysTray9733cb_QRegisterMetaType2(typeNameC)))
}

func (ptr *Ui_SysTray) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Ui_SysTray9733cb_Ui_SysTray9733cb_QRegisterMetaType2(typeNameC)))
}

func Ui_SysTray_QmlRegisterType() int {
	return int(int32(C.Ui_SysTray9733cb_Ui_SysTray9733cb_QmlRegisterType()))
}

func (ptr *Ui_SysTray) QmlRegisterType() int {
	return int(int32(C.Ui_SysTray9733cb_Ui_SysTray9733cb_QmlRegisterType()))
}

func Ui_SysTray_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Ui_SysTray9733cb_Ui_SysTray9733cb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Ui_SysTray) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Ui_SysTray9733cb_Ui_SysTray9733cb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Ui_SysTray) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Ui_SysTray9733cb___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Ui_SysTray) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Ui_SysTray) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Ui_SysTray9733cb___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Ui_SysTray) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Ui_SysTray9733cb___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Ui_SysTray) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Ui_SysTray) __findChildren_newList2() unsafe.Pointer {
	return C.Ui_SysTray9733cb___findChildren_newList2(ptr.Pointer())
}

func (ptr *Ui_SysTray) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Ui_SysTray9733cb___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Ui_SysTray) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Ui_SysTray) __findChildren_newList3() unsafe.Pointer {
	return C.Ui_SysTray9733cb___findChildren_newList3(ptr.Pointer())
}

func (ptr *Ui_SysTray) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Ui_SysTray9733cb___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Ui_SysTray) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Ui_SysTray) __findChildren_newList() unsafe.Pointer {
	return C.Ui_SysTray9733cb___findChildren_newList(ptr.Pointer())
}

func (ptr *Ui_SysTray) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Ui_SysTray9733cb___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Ui_SysTray) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Ui_SysTray) __children_newList() unsafe.Pointer {
	return C.Ui_SysTray9733cb___children_newList(ptr.Pointer())
}

func NewUi_SysTray(parent std_core.QObject_ITF) *Ui_SysTray {
	tmpValue := NewUi_SysTrayFromPointer(C.Ui_SysTray9733cb_NewUi_SysTray(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackUi_SysTray9733cb_DestroyUi_SysTray
func callbackUi_SysTray9733cb_DestroyUi_SysTray(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Ui_SysTray"); signal != nil {
		signal.(func())()
	} else {
		NewUi_SysTrayFromPointer(ptr).DestroyUi_SysTrayDefault()
	}
}

func (ptr *Ui_SysTray) ConnectDestroyUi_SysTray(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Ui_SysTray"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Ui_SysTray", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Ui_SysTray", f)
		}
	}
}

func (ptr *Ui_SysTray) DisconnectDestroyUi_SysTray() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Ui_SysTray")
	}
}

func (ptr *Ui_SysTray) DestroyUi_SysTray() {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb_DestroyUi_SysTray(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Ui_SysTray) DestroyUi_SysTrayDefault() {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb_DestroyUi_SysTrayDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackUi_SysTray9733cb_Event
func callbackUi_SysTray9733cb_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewUi_SysTrayFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Ui_SysTray) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Ui_SysTray9733cb_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackUi_SysTray9733cb_EventFilter
func callbackUi_SysTray9733cb_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewUi_SysTrayFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Ui_SysTray) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Ui_SysTray9733cb_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackUi_SysTray9733cb_ChildEvent
func callbackUi_SysTray9733cb_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewUi_SysTrayFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Ui_SysTray) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackUi_SysTray9733cb_ConnectNotify
func callbackUi_SysTray9733cb_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewUi_SysTrayFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Ui_SysTray) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackUi_SysTray9733cb_CustomEvent
func callbackUi_SysTray9733cb_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewUi_SysTrayFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Ui_SysTray) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackUi_SysTray9733cb_DeleteLater
func callbackUi_SysTray9733cb_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewUi_SysTrayFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Ui_SysTray) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackUi_SysTray9733cb_Destroyed
func callbackUi_SysTray9733cb_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackUi_SysTray9733cb_DisconnectNotify
func callbackUi_SysTray9733cb_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewUi_SysTrayFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Ui_SysTray) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackUi_SysTray9733cb_ObjectNameChanged
func callbackUi_SysTray9733cb_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackUi_SysTray9733cb_TimerEvent
func callbackUi_SysTray9733cb_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewUi_SysTrayFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Ui_SysTray) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Ui_SysTray9733cb_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type AppDoingList_ITF interface {
	std_widgets.QApplication_ITF
	AppDoingList_PTR() *AppDoingList
}

func (ptr *AppDoingList) AppDoingList_PTR() *AppDoingList {
	return ptr
}

func (ptr *AppDoingList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QApplication_PTR().Pointer()
	}
	return nil
}

func (ptr *AppDoingList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QApplication_PTR().SetPointer(p)
	}
}

func PointerFromAppDoingList(ptr AppDoingList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AppDoingList_PTR().Pointer()
	}
	return nil
}

func NewAppDoingListFromPointer(ptr unsafe.Pointer) (n *AppDoingList) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(AppDoingList)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *AppDoingList:
			n = deduced

		case *std_widgets.QApplication:
			n = &AppDoingList{QApplication: *deduced}

		default:
			n = new(AppDoingList)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackAppDoingList9733cb_Constructor
func callbackAppDoingList9733cb_Constructor(ptr unsafe.Pointer) {
	this := NewAppDoingListFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackAppDoingList9733cb_QuitApp
func callbackAppDoingList9733cb_QuitApp(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "quitApp"); signal != nil {
		signal.(func())()
	}

}

func (ptr *AppDoingList) ConnectQuitApp(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "quitApp") {
			C.AppDoingList9733cb_ConnectQuitApp(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "quitApp"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "quitApp", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "quitApp", f)
		}
	}
}

func (ptr *AppDoingList) DisconnectQuitApp() {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_DisconnectQuitApp(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "quitApp")
	}
}

func (ptr *AppDoingList) QuitApp() {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_QuitApp(ptr.Pointer())
	}
}

//export callbackAppDoingList9733cb_ActivateApp
func callbackAppDoingList9733cb_ActivateApp(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activateApp"); signal != nil {
		signal.(func())()
	}

}

func (ptr *AppDoingList) ConnectActivateApp(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activateApp") {
			C.AppDoingList9733cb_ConnectActivateApp(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activateApp"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activateApp", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activateApp", f)
		}
	}
}

func (ptr *AppDoingList) DisconnectActivateApp() {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_DisconnectActivateApp(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activateApp")
	}
}

func (ptr *AppDoingList) ActivateApp() {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_ActivateApp(ptr.Pointer())
	}
}

func AppDoingList_QRegisterMetaType() int {
	return int(int32(C.AppDoingList9733cb_AppDoingList9733cb_QRegisterMetaType()))
}

func (ptr *AppDoingList) QRegisterMetaType() int {
	return int(int32(C.AppDoingList9733cb_AppDoingList9733cb_QRegisterMetaType()))
}

func AppDoingList_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AppDoingList9733cb_AppDoingList9733cb_QRegisterMetaType2(typeNameC)))
}

func (ptr *AppDoingList) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AppDoingList9733cb_AppDoingList9733cb_QRegisterMetaType2(typeNameC)))
}

func (ptr *AppDoingList) __allWidgets_atList(i int) *std_widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQWidgetFromPointer(C.AppDoingList9733cb___allWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AppDoingList) __allWidgets_setList(i std_widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb___allWidgets_setList(ptr.Pointer(), std_widgets.PointerFromQWidget(i))
	}
}

func (ptr *AppDoingList) __allWidgets_newList() unsafe.Pointer {
	return C.AppDoingList9733cb___allWidgets_newList(ptr.Pointer())
}

func (ptr *AppDoingList) __topLevelWidgets_atList(i int) *std_widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQWidgetFromPointer(C.AppDoingList9733cb___topLevelWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AppDoingList) __topLevelWidgets_setList(i std_widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb___topLevelWidgets_setList(ptr.Pointer(), std_widgets.PointerFromQWidget(i))
	}
}

func (ptr *AppDoingList) __topLevelWidgets_newList() unsafe.Pointer {
	return C.AppDoingList9733cb___topLevelWidgets_newList(ptr.Pointer())
}

func (ptr *AppDoingList) __screens_atList(i int) *std_gui.QScreen {
	if ptr.Pointer() != nil {
		tmpValue := std_gui.NewQScreenFromPointer(C.AppDoingList9733cb___screens_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AppDoingList) __screens_setList(i std_gui.QScreen_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb___screens_setList(ptr.Pointer(), std_gui.PointerFromQScreen(i))
	}
}

func (ptr *AppDoingList) __screens_newList() unsafe.Pointer {
	return C.AppDoingList9733cb___screens_newList(ptr.Pointer())
}

func (ptr *AppDoingList) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AppDoingList9733cb___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AppDoingList) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AppDoingList) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.AppDoingList9733cb___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *AppDoingList) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AppDoingList9733cb___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AppDoingList) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AppDoingList) __findChildren_newList2() unsafe.Pointer {
	return C.AppDoingList9733cb___findChildren_newList2(ptr.Pointer())
}

func (ptr *AppDoingList) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AppDoingList9733cb___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AppDoingList) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AppDoingList) __findChildren_newList3() unsafe.Pointer {
	return C.AppDoingList9733cb___findChildren_newList3(ptr.Pointer())
}

func (ptr *AppDoingList) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AppDoingList9733cb___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AppDoingList) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AppDoingList) __findChildren_newList() unsafe.Pointer {
	return C.AppDoingList9733cb___findChildren_newList(ptr.Pointer())
}

func (ptr *AppDoingList) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AppDoingList9733cb___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AppDoingList) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AppDoingList) __children_newList() unsafe.Pointer {
	return C.AppDoingList9733cb___children_newList(ptr.Pointer())
}

func NewAppDoingList(argc int, argv []string) *AppDoingList {
	argvC := C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	tmpValue := NewAppDoingListFromPointer(C.AppDoingList9733cb_NewAppDoingList(C.int(int32(argc)), argvC))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackAppDoingList9733cb_DestroyAppDoingList
func callbackAppDoingList9733cb_DestroyAppDoingList(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~AppDoingList"); signal != nil {
		signal.(func())()
	} else {
		NewAppDoingListFromPointer(ptr).DestroyAppDoingListDefault()
	}
}

func (ptr *AppDoingList) ConnectDestroyAppDoingList(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~AppDoingList"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~AppDoingList", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~AppDoingList", f)
		}
	}
}

func (ptr *AppDoingList) DisconnectDestroyAppDoingList() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~AppDoingList")
	}
}

func (ptr *AppDoingList) DestroyAppDoingList() {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_DestroyAppDoingList(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *AppDoingList) DestroyAppDoingListDefault() {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_DestroyAppDoingListDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAppDoingList9733cb_Event
func callbackAppDoingList9733cb_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAppDoingListFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *AppDoingList) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.AppDoingList9733cb_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackAppDoingList9733cb_AboutQt
func callbackAppDoingList9733cb_AboutQt(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutQt"); signal != nil {
		signal.(func())()
	} else {
		NewAppDoingListFromPointer(ptr).AboutQtDefault()
	}
}

func (ptr *AppDoingList) AboutQtDefault() {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_AboutQtDefault(ptr.Pointer())
	}
}

//export callbackAppDoingList9733cb_CloseAllWindows
func callbackAppDoingList9733cb_CloseAllWindows(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeAllWindows"); signal != nil {
		signal.(func())()
	} else {
		NewAppDoingListFromPointer(ptr).CloseAllWindowsDefault()
	}
}

func (ptr *AppDoingList) CloseAllWindowsDefault() {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_CloseAllWindowsDefault(ptr.Pointer())
	}
}

//export callbackAppDoingList9733cb_FocusChanged
func callbackAppDoingList9733cb_FocusChanged(ptr unsafe.Pointer, old unsafe.Pointer, now unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusChanged"); signal != nil {
		signal.(func(*std_widgets.QWidget, *std_widgets.QWidget))(std_widgets.NewQWidgetFromPointer(old), std_widgets.NewQWidgetFromPointer(now))
	}

}

//export callbackAppDoingList9733cb_SetAutoSipEnabled
func callbackAppDoingList9733cb_SetAutoSipEnabled(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(ptr, "setAutoSipEnabled"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewAppDoingListFromPointer(ptr).SetAutoSipEnabledDefault(int8(enabled) != 0)
	}
}

func (ptr *AppDoingList) SetAutoSipEnabledDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_SetAutoSipEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackAppDoingList9733cb_SetStyleSheet
func callbackAppDoingList9733cb_SetStyleSheet(ptr unsafe.Pointer, sheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(sheet))
	} else {
		NewAppDoingListFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(sheet))
	}
}

func (ptr *AppDoingList) SetStyleSheetDefault(sheet string) {
	if ptr.Pointer() != nil {
		var sheetC *C.char
		if sheet != "" {
			sheetC = C.CString(sheet)
			defer C.free(unsafe.Pointer(sheetC))
		}
		C.AppDoingList9733cb_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: sheetC, len: C.longlong(len(sheet))})
	}
}

//export callbackAppDoingList9733cb_AutoSipEnabled
func callbackAppDoingList9733cb_AutoSipEnabled(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "autoSipEnabled"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewAppDoingListFromPointer(ptr).AutoSipEnabledDefault())))
}

func (ptr *AppDoingList) AutoSipEnabledDefault() bool {
	if ptr.Pointer() != nil {
		return C.AppDoingList9733cb_AutoSipEnabledDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackAppDoingList9733cb_ApplicationDisplayNameChanged
func callbackAppDoingList9733cb_ApplicationDisplayNameChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "applicationDisplayNameChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackAppDoingList9733cb_ApplicationStateChanged
func callbackAppDoingList9733cb_ApplicationStateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "applicationStateChanged"); signal != nil {
		signal.(func(std_core.Qt__ApplicationState))(std_core.Qt__ApplicationState(state))
	}

}

//export callbackAppDoingList9733cb_CommitDataRequest
func callbackAppDoingList9733cb_CommitDataRequest(ptr unsafe.Pointer, manager unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "commitDataRequest"); signal != nil {
		signal.(func(*std_gui.QSessionManager))(std_gui.NewQSessionManagerFromPointer(manager))
	}

}

//export callbackAppDoingList9733cb_FocusObjectChanged
func callbackAppDoingList9733cb_FocusObjectChanged(ptr unsafe.Pointer, focusObject unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusObjectChanged"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(focusObject))
	}

}

//export callbackAppDoingList9733cb_FocusWindowChanged
func callbackAppDoingList9733cb_FocusWindowChanged(ptr unsafe.Pointer, focusWindow unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusWindowChanged"); signal != nil {
		signal.(func(*std_gui.QWindow))(std_gui.NewQWindowFromPointer(focusWindow))
	}

}

//export callbackAppDoingList9733cb_FontChanged
func callbackAppDoingList9733cb_FontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fontChanged"); signal != nil {
		signal.(func(*std_gui.QFont))(std_gui.NewQFontFromPointer(font))
	}

}

//export callbackAppDoingList9733cb_FontDatabaseChanged
func callbackAppDoingList9733cb_FontDatabaseChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fontDatabaseChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackAppDoingList9733cb_LastWindowClosed
func callbackAppDoingList9733cb_LastWindowClosed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lastWindowClosed"); signal != nil {
		signal.(func())()
	}

}

//export callbackAppDoingList9733cb_LayoutDirectionChanged
func callbackAppDoingList9733cb_LayoutDirectionChanged(ptr unsafe.Pointer, direction C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutDirectionChanged"); signal != nil {
		signal.(func(std_core.Qt__LayoutDirection))(std_core.Qt__LayoutDirection(direction))
	}

}

//export callbackAppDoingList9733cb_PaletteChanged
func callbackAppDoingList9733cb_PaletteChanged(ptr unsafe.Pointer, palette unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paletteChanged"); signal != nil {
		signal.(func(*std_gui.QPalette))(std_gui.NewQPaletteFromPointer(palette))
	}

}

//export callbackAppDoingList9733cb_PrimaryScreenChanged
func callbackAppDoingList9733cb_PrimaryScreenChanged(ptr unsafe.Pointer, screen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "primaryScreenChanged"); signal != nil {
		signal.(func(*std_gui.QScreen))(std_gui.NewQScreenFromPointer(screen))
	}

}

//export callbackAppDoingList9733cb_SaveStateRequest
func callbackAppDoingList9733cb_SaveStateRequest(ptr unsafe.Pointer, manager unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "saveStateRequest"); signal != nil {
		signal.(func(*std_gui.QSessionManager))(std_gui.NewQSessionManagerFromPointer(manager))
	}

}

//export callbackAppDoingList9733cb_ScreenAdded
func callbackAppDoingList9733cb_ScreenAdded(ptr unsafe.Pointer, screen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "screenAdded"); signal != nil {
		signal.(func(*std_gui.QScreen))(std_gui.NewQScreenFromPointer(screen))
	}

}

//export callbackAppDoingList9733cb_ScreenRemoved
func callbackAppDoingList9733cb_ScreenRemoved(ptr unsafe.Pointer, screen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "screenRemoved"); signal != nil {
		signal.(func(*std_gui.QScreen))(std_gui.NewQScreenFromPointer(screen))
	}

}

//export callbackAppDoingList9733cb_AboutToQuit
func callbackAppDoingList9733cb_AboutToQuit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutToQuit"); signal != nil {
		signal.(func())()
	}

}

//export callbackAppDoingList9733cb_Quit
func callbackAppDoingList9733cb_Quit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "quit"); signal != nil {
		signal.(func())()
	} else {
		NewAppDoingListFromPointer(ptr).QuitDefault()
	}
}

func (ptr *AppDoingList) QuitDefault() {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_QuitDefault(ptr.Pointer())
	}
}

//export callbackAppDoingList9733cb_EventFilter
func callbackAppDoingList9733cb_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAppDoingListFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *AppDoingList) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.AppDoingList9733cb_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackAppDoingList9733cb_ChildEvent
func callbackAppDoingList9733cb_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewAppDoingListFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *AppDoingList) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackAppDoingList9733cb_ConnectNotify
func callbackAppDoingList9733cb_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAppDoingListFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AppDoingList) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAppDoingList9733cb_CustomEvent
func callbackAppDoingList9733cb_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewAppDoingListFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *AppDoingList) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackAppDoingList9733cb_DeleteLater
func callbackAppDoingList9733cb_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewAppDoingListFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *AppDoingList) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAppDoingList9733cb_Destroyed
func callbackAppDoingList9733cb_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackAppDoingList9733cb_DisconnectNotify
func callbackAppDoingList9733cb_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAppDoingListFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AppDoingList) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAppDoingList9733cb_ObjectNameChanged
func callbackAppDoingList9733cb_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackAppDoingList9733cb_TimerEvent
func callbackAppDoingList9733cb_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewAppDoingListFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *AppDoingList) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AppDoingList9733cb_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type DoingList_ITF interface {
	std_widgets.QListWidget_ITF
	DoingList_PTR() *DoingList
}

func (ptr *DoingList) DoingList_PTR() *DoingList {
	return ptr
}

func (ptr *DoingList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QListWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *DoingList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QListWidget_PTR().SetPointer(p)
	}
}

func PointerFromDoingList(ptr DoingList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.DoingList_PTR().Pointer()
	}
	return nil
}

func NewDoingListFromPointer(ptr unsafe.Pointer) (n *DoingList) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(DoingList)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *DoingList:
			n = deduced

		case *std_widgets.QListWidget:
			n = &DoingList{QListWidget: *deduced}

		default:
			n = new(DoingList)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackDoingList9733cb_Constructor
func callbackDoingList9733cb_Constructor(ptr unsafe.Pointer) {
	this := NewDoingListFromPointer(ptr)
	qt.Register(ptr, this)
}

func DoingList_QRegisterMetaType() int {
	return int(int32(C.DoingList9733cb_DoingList9733cb_QRegisterMetaType()))
}

func (ptr *DoingList) QRegisterMetaType() int {
	return int(int32(C.DoingList9733cb_DoingList9733cb_QRegisterMetaType()))
}

func DoingList_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.DoingList9733cb_DoingList9733cb_QRegisterMetaType2(typeNameC)))
}

func (ptr *DoingList) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.DoingList9733cb_DoingList9733cb_QRegisterMetaType2(typeNameC)))
}

func DoingList_QmlRegisterType() int {
	return int(int32(C.DoingList9733cb_DoingList9733cb_QmlRegisterType()))
}

func (ptr *DoingList) QmlRegisterType() int {
	return int(int32(C.DoingList9733cb_DoingList9733cb_QmlRegisterType()))
}

func DoingList_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.DoingList9733cb_DoingList9733cb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *DoingList) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.DoingList9733cb_DoingList9733cb_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *DoingList) __findItems_atList(i int) *std_widgets.QListWidgetItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQListWidgetItemFromPointer(C.DoingList9733cb___findItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *DoingList) __findItems_setList(i std_widgets.QListWidgetItem_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___findItems_setList(ptr.Pointer(), std_widgets.PointerFromQListWidgetItem(i))
	}
}

func (ptr *DoingList) __findItems_newList() unsafe.Pointer {
	return C.DoingList9733cb___findItems_newList(ptr.Pointer())
}

func (ptr *DoingList) __items_atList(i int) *std_widgets.QListWidgetItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQListWidgetItemFromPointer(C.DoingList9733cb___items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *DoingList) __items_setList(i std_widgets.QListWidgetItem_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___items_setList(ptr.Pointer(), std_widgets.PointerFromQListWidgetItem(i))
	}
}

func (ptr *DoingList) __items_newList() unsafe.Pointer {
	return C.DoingList9733cb___items_newList(ptr.Pointer())
}

func (ptr *DoingList) __selectedItems_atList(i int) *std_widgets.QListWidgetItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQListWidgetItemFromPointer(C.DoingList9733cb___selectedItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *DoingList) __selectedItems_setList(i std_widgets.QListWidgetItem_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___selectedItems_setList(ptr.Pointer(), std_widgets.PointerFromQListWidgetItem(i))
	}
}

func (ptr *DoingList) __selectedItems_newList() unsafe.Pointer {
	return C.DoingList9733cb___selectedItems_newList(ptr.Pointer())
}

func (ptr *DoingList) __mimeData_items_atList(i int) *std_widgets.QListWidgetItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQListWidgetItemFromPointer(C.DoingList9733cb___mimeData_items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *DoingList) __mimeData_items_setList(i std_widgets.QListWidgetItem_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___mimeData_items_setList(ptr.Pointer(), std_widgets.PointerFromQListWidgetItem(i))
	}
}

func (ptr *DoingList) __mimeData_items_newList() unsafe.Pointer {
	return C.DoingList9733cb___mimeData_items_newList(ptr.Pointer())
}

func (ptr *DoingList) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DoingList9733cb___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *DoingList) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *DoingList) __dataChanged_roles_newList() unsafe.Pointer {
	return C.DoingList9733cb___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *DoingList) __indexesMoved_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.DoingList9733cb___indexesMoved_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *DoingList) __indexesMoved_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___indexesMoved_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *DoingList) __indexesMoved_indexes_newList() unsafe.Pointer {
	return C.DoingList9733cb___indexesMoved_indexes_newList(ptr.Pointer())
}

func (ptr *DoingList) __selectedIndexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.DoingList9733cb___selectedIndexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *DoingList) __selectedIndexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___selectedIndexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *DoingList) __selectedIndexes_newList() unsafe.Pointer {
	return C.DoingList9733cb___selectedIndexes_newList(ptr.Pointer())
}

func (ptr *DoingList) __scrollBarWidgets_atList(i int) *std_widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQWidgetFromPointer(C.DoingList9733cb___scrollBarWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DoingList) __scrollBarWidgets_setList(i std_widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___scrollBarWidgets_setList(ptr.Pointer(), std_widgets.PointerFromQWidget(i))
	}
}

func (ptr *DoingList) __scrollBarWidgets_newList() unsafe.Pointer {
	return C.DoingList9733cb___scrollBarWidgets_newList(ptr.Pointer())
}

func (ptr *DoingList) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.DoingList9733cb___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DoingList) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *DoingList) __addActions_actions_newList() unsafe.Pointer {
	return C.DoingList9733cb___addActions_actions_newList(ptr.Pointer())
}

func (ptr *DoingList) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.DoingList9733cb___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DoingList) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *DoingList) __insertActions_actions_newList() unsafe.Pointer {
	return C.DoingList9733cb___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *DoingList) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.DoingList9733cb___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DoingList) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *DoingList) __actions_newList() unsafe.Pointer {
	return C.DoingList9733cb___actions_newList(ptr.Pointer())
}

func (ptr *DoingList) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.DoingList9733cb___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *DoingList) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *DoingList) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.DoingList9733cb___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *DoingList) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.DoingList9733cb___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DoingList) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DoingList) __findChildren_newList2() unsafe.Pointer {
	return C.DoingList9733cb___findChildren_newList2(ptr.Pointer())
}

func (ptr *DoingList) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.DoingList9733cb___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DoingList) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DoingList) __findChildren_newList3() unsafe.Pointer {
	return C.DoingList9733cb___findChildren_newList3(ptr.Pointer())
}

func (ptr *DoingList) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.DoingList9733cb___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DoingList) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DoingList) __findChildren_newList() unsafe.Pointer {
	return C.DoingList9733cb___findChildren_newList(ptr.Pointer())
}

func (ptr *DoingList) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.DoingList9733cb___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *DoingList) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *DoingList) __children_newList() unsafe.Pointer {
	return C.DoingList9733cb___children_newList(ptr.Pointer())
}

func NewDoingList(parent std_widgets.QWidget_ITF) *DoingList {
	tmpValue := NewDoingListFromPointer(C.DoingList9733cb_NewDoingList(std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackDoingList9733cb_DestroyDoingList
func callbackDoingList9733cb_DestroyDoingList(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~DoingList"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).DestroyDoingListDefault()
	}
}

func (ptr *DoingList) ConnectDestroyDoingList(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~DoingList"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~DoingList", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~DoingList", f)
		}
	}
}

func (ptr *DoingList) DisconnectDestroyDoingList() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~DoingList")
	}
}

func (ptr *DoingList) DestroyDoingList() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_DestroyDoingList(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *DoingList) DestroyDoingListDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_DestroyDoingListDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackDoingList9733cb_DropMimeData
func callbackDoingList9733cb_DropMimeData(ptr unsafe.Pointer, index C.int, data unsafe.Pointer, action C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, *std_core.QMimeData, std_core.Qt__DropAction) bool)(int(int32(index)), std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDoingListFromPointer(ptr).DropMimeDataDefault(int(int32(index)), std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action)))))
}

func (ptr *DoingList) DropMimeDataDefault(index int, data std_core.QMimeData_ITF, action std_core.Qt__DropAction) bool {
	if ptr.Pointer() != nil {
		return C.DoingList9733cb_DropMimeDataDefault(ptr.Pointer(), C.int(int32(index)), std_core.PointerFromQMimeData(data), C.longlong(action)) != 0
	}
	return false
}

//export callbackDoingList9733cb_Event
func callbackDoingList9733cb_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDoingListFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *DoingList) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DoingList9733cb_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackDoingList9733cb_Clear
func callbackDoingList9733cb_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).ClearDefault()
	}
}

func (ptr *DoingList) ClearDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ClearDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_CurrentItemChanged
func callbackDoingList9733cb_CurrentItemChanged(ptr unsafe.Pointer, current unsafe.Pointer, previous unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "currentItemChanged"); signal != nil {
		signal.(func(*std_widgets.QListWidgetItem, *std_widgets.QListWidgetItem))(std_widgets.NewQListWidgetItemFromPointer(current), std_widgets.NewQListWidgetItemFromPointer(previous))
	}

}

//export callbackDoingList9733cb_CurrentRowChanged
func callbackDoingList9733cb_CurrentRowChanged(ptr unsafe.Pointer, currentRow C.int) {
	if signal := qt.GetSignal(ptr, "currentRowChanged"); signal != nil {
		signal.(func(int))(int(int32(currentRow)))
	}

}

//export callbackDoingList9733cb_CurrentTextChanged
func callbackDoingList9733cb_CurrentTextChanged(ptr unsafe.Pointer, currentText C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "currentTextChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(currentText))
	}

}

//export callbackDoingList9733cb_DropEvent
func callbackDoingList9733cb_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*std_gui.QDropEvent))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *DoingList) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackDoingList9733cb_ItemActivated
func callbackDoingList9733cb_ItemActivated(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "itemActivated"); signal != nil {
		signal.(func(*std_widgets.QListWidgetItem))(std_widgets.NewQListWidgetItemFromPointer(item))
	}

}

//export callbackDoingList9733cb_ItemChanged
func callbackDoingList9733cb_ItemChanged(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "itemChanged"); signal != nil {
		signal.(func(*std_widgets.QListWidgetItem))(std_widgets.NewQListWidgetItemFromPointer(item))
	}

}

//export callbackDoingList9733cb_ItemClicked
func callbackDoingList9733cb_ItemClicked(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "itemClicked"); signal != nil {
		signal.(func(*std_widgets.QListWidgetItem))(std_widgets.NewQListWidgetItemFromPointer(item))
	}

}

//export callbackDoingList9733cb_ItemDoubleClicked
func callbackDoingList9733cb_ItemDoubleClicked(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "itemDoubleClicked"); signal != nil {
		signal.(func(*std_widgets.QListWidgetItem))(std_widgets.NewQListWidgetItemFromPointer(item))
	}

}

//export callbackDoingList9733cb_ItemEntered
func callbackDoingList9733cb_ItemEntered(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "itemEntered"); signal != nil {
		signal.(func(*std_widgets.QListWidgetItem))(std_widgets.NewQListWidgetItemFromPointer(item))
	}

}

//export callbackDoingList9733cb_ItemPressed
func callbackDoingList9733cb_ItemPressed(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "itemPressed"); signal != nil {
		signal.(func(*std_widgets.QListWidgetItem))(std_widgets.NewQListWidgetItemFromPointer(item))
	}

}

//export callbackDoingList9733cb_ItemSelectionChanged
func callbackDoingList9733cb_ItemSelectionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "itemSelectionChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackDoingList9733cb_ScrollToItem
func callbackDoingList9733cb_ScrollToItem(ptr unsafe.Pointer, item unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "scrollToItem"); signal != nil {
		signal.(func(*std_widgets.QListWidgetItem, std_widgets.QAbstractItemView__ScrollHint))(std_widgets.NewQListWidgetItemFromPointer(item), std_widgets.QAbstractItemView__ScrollHint(hint))
	} else {
		NewDoingListFromPointer(ptr).ScrollToItemDefault(std_widgets.NewQListWidgetItemFromPointer(item), std_widgets.QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *DoingList) ScrollToItemDefault(item std_widgets.QListWidgetItem_ITF, hint std_widgets.QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ScrollToItemDefault(ptr.Pointer(), std_widgets.PointerFromQListWidgetItem(item), C.longlong(hint))
	}
}

//export callbackDoingList9733cb_SetSelectionModel
func callbackDoingList9733cb_SetSelectionModel(ptr unsafe.Pointer, selectionModel unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setSelectionModel"); signal != nil {
		signal.(func(*std_core.QItemSelectionModel))(std_core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewDoingListFromPointer(ptr).SetSelectionModelDefault(std_core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *DoingList) SetSelectionModelDefault(selectionModel std_core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetSelectionModelDefault(ptr.Pointer(), std_core.PointerFromQItemSelectionModel(selectionModel))
	}
}

//export callbackDoingList9733cb_MimeData
func callbackDoingList9733cb_MimeData(ptr unsafe.Pointer, items C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_widgets.QListWidgetItem) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_widgets.QListWidgetItem {
			out := make([]*std_widgets.QListWidgetItem, int(l.len))
			tmpList := NewDoingListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_items_atList(i)
			}
			return out
		}(items)))
	}

	return std_core.PointerFromQMimeData(NewDoingListFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_widgets.QListWidgetItem {
		out := make([]*std_widgets.QListWidgetItem, int(l.len))
		tmpList := NewDoingListFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_items_atList(i)
		}
		return out
	}(items)))
}

func (ptr *DoingList) MimeDataDefault(items []*std_widgets.QListWidgetItem) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.DoingList9733cb_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewDoingListFromPointer(NewDoingListFromPointer(nil).__mimeData_items_newList())
			for _, v := range items {
				tmpList.__mimeData_items_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackDoingList9733cb_MimeTypes
func callbackDoingList9733cb_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewDoingListFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *DoingList) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.DoingList9733cb_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackDoingList9733cb_SupportedDropActions
func callbackDoingList9733cb_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewDoingListFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *DoingList) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.DoingList9733cb_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackDoingList9733cb_MoveCursor
func callbackDoingList9733cb_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "moveCursor"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(std_widgets.QAbstractItemView__CursorAction, std_core.Qt__KeyboardModifier) *std_core.QModelIndex)(std_widgets.QAbstractItemView__CursorAction(cursorAction), std_core.Qt__KeyboardModifier(modifiers)))
	}

	return std_core.PointerFromQModelIndex(NewDoingListFromPointer(ptr).MoveCursorDefault(std_widgets.QAbstractItemView__CursorAction(cursorAction), std_core.Qt__KeyboardModifier(modifiers)))
}

func (ptr *DoingList) MoveCursorDefault(cursorAction std_widgets.QAbstractItemView__CursorAction, modifiers std_core.Qt__KeyboardModifier) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.DoingList9733cb_MoveCursorDefault(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackDoingList9733cb_CurrentChanged
func callbackDoingList9733cb_CurrentChanged(ptr unsafe.Pointer, current unsafe.Pointer, previous unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "currentChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(current), std_core.NewQModelIndexFromPointer(previous))
	} else {
		NewDoingListFromPointer(ptr).CurrentChangedDefault(std_core.NewQModelIndexFromPointer(current), std_core.NewQModelIndexFromPointer(previous))
	}
}

func (ptr *DoingList) CurrentChangedDefault(current std_core.QModelIndex_ITF, previous std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_CurrentChangedDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(current), std_core.PointerFromQModelIndex(previous))
	}
}

//export callbackDoingList9733cb_DataChanged
func callbackDoingList9733cb_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewDoingListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	} else {
		NewDoingListFromPointer(ptr).DataChangedDefault(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewDoingListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}
}

func (ptr *DoingList) DataChangedDefault(topLeft std_core.QModelIndex_ITF, bottomRight std_core.QModelIndex_ITF, roles []int) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_DataChangedDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(topLeft), std_core.PointerFromQModelIndex(bottomRight), func() unsafe.Pointer {
			tmpList := NewDoingListFromPointer(NewDoingListFromPointer(nil).__dataChanged_roles_newList())
			for _, v := range roles {
				tmpList.__dataChanged_roles_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackDoingList9733cb_DragLeaveEvent
func callbackDoingList9733cb_DragLeaveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*std_gui.QDragLeaveEvent))(std_gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewDoingListFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *DoingList) DragLeaveEventDefault(e std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(e))
	}
}

//export callbackDoingList9733cb_DragMoveEvent
func callbackDoingList9733cb_DragMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*std_gui.QDragMoveEvent))(std_gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewDoingListFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *DoingList) DragMoveEventDefault(e std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(e))
	}
}

//export callbackDoingList9733cb_MouseMoveEvent
func callbackDoingList9733cb_MouseMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewDoingListFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *DoingList) MouseMoveEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackDoingList9733cb_MouseReleaseEvent
func callbackDoingList9733cb_MouseReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewDoingListFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *DoingList) MouseReleaseEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackDoingList9733cb_PaintEvent
func callbackDoingList9733cb_PaintEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*std_gui.QPaintEvent))(std_gui.NewQPaintEventFromPointer(e))
	} else {
		NewDoingListFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *DoingList) PaintEventDefault(e std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(e))
	}
}

//export callbackDoingList9733cb_ResizeEvent
func callbackDoingList9733cb_ResizeEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*std_gui.QResizeEvent))(std_gui.NewQResizeEventFromPointer(e))
	} else {
		NewDoingListFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *DoingList) ResizeEventDefault(e std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(e))
	}
}

//export callbackDoingList9733cb_RowsAboutToBeRemoved
func callbackDoingList9733cb_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewDoingListFromPointer(ptr).RowsAboutToBeRemovedDefault(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *DoingList) RowsAboutToBeRemovedDefault(parent std_core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_RowsAboutToBeRemovedDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackDoingList9733cb_RowsInserted
func callbackDoingList9733cb_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewDoingListFromPointer(ptr).RowsInsertedDefault(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *DoingList) RowsInsertedDefault(parent std_core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_RowsInsertedDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackDoingList9733cb_ScrollTo
func callbackDoingList9733cb_ScrollTo(ptr unsafe.Pointer, index unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "scrollTo"); signal != nil {
		signal.(func(*std_core.QModelIndex, std_widgets.QAbstractItemView__ScrollHint))(std_core.NewQModelIndexFromPointer(index), std_widgets.QAbstractItemView__ScrollHint(hint))
	} else {
		NewDoingListFromPointer(ptr).ScrollToDefault(std_core.NewQModelIndexFromPointer(index), std_widgets.QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *DoingList) ScrollToDefault(index std_core.QModelIndex_ITF, hint std_widgets.QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ScrollToDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

//export callbackDoingList9733cb_SelectionChanged
func callbackDoingList9733cb_SelectionChanged(ptr unsafe.Pointer, selected unsafe.Pointer, deselected unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		signal.(func(*std_core.QItemSelection, *std_core.QItemSelection))(std_core.NewQItemSelectionFromPointer(selected), std_core.NewQItemSelectionFromPointer(deselected))
	} else {
		NewDoingListFromPointer(ptr).SelectionChangedDefault(std_core.NewQItemSelectionFromPointer(selected), std_core.NewQItemSelectionFromPointer(deselected))
	}
}

func (ptr *DoingList) SelectionChangedDefault(selected std_core.QItemSelection_ITF, deselected std_core.QItemSelection_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SelectionChangedDefault(ptr.Pointer(), std_core.PointerFromQItemSelection(selected), std_core.PointerFromQItemSelection(deselected))
	}
}

//export callbackDoingList9733cb_SetSelection
func callbackDoingList9733cb_SetSelection(ptr unsafe.Pointer, rect unsafe.Pointer, command C.longlong) {
	if signal := qt.GetSignal(ptr, "setSelection"); signal != nil {
		signal.(func(*std_core.QRect, std_core.QItemSelectionModel__SelectionFlag))(std_core.NewQRectFromPointer(rect), std_core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewDoingListFromPointer(ptr).SetSelectionDefault(std_core.NewQRectFromPointer(rect), std_core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *DoingList) SetSelectionDefault(rect std_core.QRect_ITF, command std_core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetSelectionDefault(ptr.Pointer(), std_core.PointerFromQRect(rect), C.longlong(command))
	}
}

//export callbackDoingList9733cb_StartDrag
func callbackDoingList9733cb_StartDrag(ptr unsafe.Pointer, supportedActions C.longlong) {
	if signal := qt.GetSignal(ptr, "startDrag"); signal != nil {
		signal.(func(std_core.Qt__DropAction))(std_core.Qt__DropAction(supportedActions))
	} else {
		NewDoingListFromPointer(ptr).StartDragDefault(std_core.Qt__DropAction(supportedActions))
	}
}

func (ptr *DoingList) StartDragDefault(supportedActions std_core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_StartDragDefault(ptr.Pointer(), C.longlong(supportedActions))
	}
}

//export callbackDoingList9733cb_TimerEvent
func callbackDoingList9733cb_TimerEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(e))
	} else {
		NewDoingListFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *DoingList) TimerEventDefault(e std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(e))
	}
}

//export callbackDoingList9733cb_UpdateGeometries
func callbackDoingList9733cb_UpdateGeometries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateGeometries"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *DoingList) UpdateGeometriesDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_UpdateGeometriesDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_WheelEvent
func callbackDoingList9733cb_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*std_gui.QWheelEvent))(std_gui.NewQWheelEventFromPointer(e))
	} else {
		NewDoingListFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *DoingList) WheelEventDefault(e std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(e))
	}
}

//export callbackDoingList9733cb_IndexAt
func callbackDoingList9733cb_IndexAt(ptr unsafe.Pointer, p unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "indexAt"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QPoint) *std_core.QModelIndex)(std_core.NewQPointFromPointer(p)))
	}

	return std_core.PointerFromQModelIndex(NewDoingListFromPointer(ptr).IndexAtDefault(std_core.NewQPointFromPointer(p)))
}

func (ptr *DoingList) IndexAtDefault(p std_core.QPoint_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.DoingList9733cb_IndexAtDefault(ptr.Pointer(), std_core.PointerFromQPoint(p)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackDoingList9733cb_SelectedIndexes
func callbackDoingList9733cb_SelectedIndexes(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "selectedIndexes"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewDoingListFromPointer(NewDoingListFromPointer(nil).__selectedIndexes_newList())
			for _, v := range signal.(func() []*std_core.QModelIndex)() {
				tmpList.__selectedIndexes_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewDoingListFromPointer(NewDoingListFromPointer(nil).__selectedIndexes_newList())
		for _, v := range NewDoingListFromPointer(ptr).SelectedIndexesDefault() {
			tmpList.__selectedIndexes_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *DoingList) SelectedIndexesDefault() []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewDoingListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__selectedIndexes_atList(i)
			}
			return out
		}(C.DoingList9733cb_SelectedIndexesDefault(ptr.Pointer()))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackDoingList9733cb_VisualRect
func callbackDoingList9733cb_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRect"); signal != nil {
		return std_core.PointerFromQRect(signal.(func(*std_core.QModelIndex) *std_core.QRect)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQRect(NewDoingListFromPointer(ptr).VisualRectDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *DoingList) VisualRectDefault(index std_core.QModelIndex_ITF) *std_core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQRectFromPointer(C.DoingList9733cb_VisualRectDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackDoingList9733cb_VisualRegionForSelection
func callbackDoingList9733cb_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRegionForSelection"); signal != nil {
		return std_gui.PointerFromQRegion(signal.(func(*std_core.QItemSelection) *std_gui.QRegion)(std_core.NewQItemSelectionFromPointer(selection)))
	}

	return std_gui.PointerFromQRegion(NewDoingListFromPointer(ptr).VisualRegionForSelectionDefault(std_core.NewQItemSelectionFromPointer(selection)))
}

func (ptr *DoingList) VisualRegionForSelectionDefault(selection std_core.QItemSelection_ITF) *std_gui.QRegion {
	if ptr.Pointer() != nil {
		tmpValue := std_gui.NewQRegionFromPointer(C.DoingList9733cb_VisualRegionForSelectionDefault(ptr.Pointer(), std_core.PointerFromQItemSelection(selection)))
		runtime.SetFinalizer(tmpValue, (*std_gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

//export callbackDoingList9733cb_ViewportSizeHint
func callbackDoingList9733cb_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewportSizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewDoingListFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *DoingList) ViewportSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.DoingList9733cb_ViewportSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackDoingList9733cb_ViewOptions
func callbackDoingList9733cb_ViewOptions(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewOptions"); signal != nil {
		return std_widgets.PointerFromQStyleOptionViewItem(signal.(func() *std_widgets.QStyleOptionViewItem)())
	}

	return std_widgets.PointerFromQStyleOptionViewItem(NewDoingListFromPointer(ptr).ViewOptionsDefault())
}

func (ptr *DoingList) ViewOptionsDefault() *std_widgets.QStyleOptionViewItem {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQStyleOptionViewItemFromPointer(C.DoingList9733cb_ViewOptionsDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

//export callbackDoingList9733cb_IsIndexHidden
func callbackDoingList9733cb_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDoingListFromPointer(ptr).IsIndexHiddenDefault(std_core.NewQModelIndexFromPointer(index)))))
}

func (ptr *DoingList) IsIndexHiddenDefault(index std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DoingList9733cb_IsIndexHiddenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

//export callbackDoingList9733cb_HorizontalOffset
func callbackDoingList9733cb_HorizontalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "horizontalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewDoingListFromPointer(ptr).HorizontalOffsetDefault()))
}

func (ptr *DoingList) HorizontalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.DoingList9733cb_HorizontalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackDoingList9733cb_VerticalOffset
func callbackDoingList9733cb_VerticalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "verticalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewDoingListFromPointer(ptr).VerticalOffsetDefault()))
}

func (ptr *DoingList) VerticalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.DoingList9733cb_VerticalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackDoingList9733cb_Edit2
func callbackDoingList9733cb_Edit2(ptr unsafe.Pointer, index unsafe.Pointer, trigger C.longlong, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "edit2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, std_widgets.QAbstractItemView__EditTrigger, *std_core.QEvent) bool)(std_core.NewQModelIndexFromPointer(index), std_widgets.QAbstractItemView__EditTrigger(trigger), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDoingListFromPointer(ptr).Edit2Default(std_core.NewQModelIndexFromPointer(index), std_widgets.QAbstractItemView__EditTrigger(trigger), std_core.NewQEventFromPointer(event)))))
}

func (ptr *DoingList) Edit2Default(index std_core.QModelIndex_ITF, trigger std_widgets.QAbstractItemView__EditTrigger, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DoingList9733cb_Edit2Default(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.longlong(trigger), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackDoingList9733cb_EventFilter
func callbackDoingList9733cb_EventFilter(ptr unsafe.Pointer, object unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(object), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDoingListFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(object), std_core.NewQEventFromPointer(event)))))
}

func (ptr *DoingList) EventFilterDefault(object std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DoingList9733cb_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(object), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackDoingList9733cb_FocusNextPrevChild
func callbackDoingList9733cb_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDoingListFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *DoingList) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.DoingList9733cb_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackDoingList9733cb_ViewportEvent
func callbackDoingList9733cb_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDoingListFromPointer(ptr).ViewportEventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *DoingList) ViewportEventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.DoingList9733cb_ViewportEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackDoingList9733cb_Activated
func callbackDoingList9733cb_Activated(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(index))
	}

}

//export callbackDoingList9733cb_ClearSelection
func callbackDoingList9733cb_ClearSelection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearSelection"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).ClearSelectionDefault()
	}
}

func (ptr *DoingList) ClearSelectionDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ClearSelectionDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_Clicked
func callbackDoingList9733cb_Clicked(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(index))
	}

}

//export callbackDoingList9733cb_CloseEditor
func callbackDoingList9733cb_CloseEditor(ptr unsafe.Pointer, editor unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "closeEditor"); signal != nil {
		signal.(func(*std_widgets.QWidget, std_widgets.QAbstractItemDelegate__EndEditHint))(std_widgets.NewQWidgetFromPointer(editor), std_widgets.QAbstractItemDelegate__EndEditHint(hint))
	} else {
		NewDoingListFromPointer(ptr).CloseEditorDefault(std_widgets.NewQWidgetFromPointer(editor), std_widgets.QAbstractItemDelegate__EndEditHint(hint))
	}
}

func (ptr *DoingList) CloseEditorDefault(editor std_widgets.QWidget_ITF, hint std_widgets.QAbstractItemDelegate__EndEditHint) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_CloseEditorDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(editor), C.longlong(hint))
	}
}

//export callbackDoingList9733cb_CommitData
func callbackDoingList9733cb_CommitData(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "commitData"); signal != nil {
		signal.(func(*std_widgets.QWidget))(std_widgets.NewQWidgetFromPointer(editor))
	} else {
		NewDoingListFromPointer(ptr).CommitDataDefault(std_widgets.NewQWidgetFromPointer(editor))
	}
}

func (ptr *DoingList) CommitDataDefault(editor std_widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_CommitDataDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(editor))
	}
}

//export callbackDoingList9733cb_DoubleClicked
func callbackDoingList9733cb_DoubleClicked(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(index))
	}

}

//export callbackDoingList9733cb_DragEnterEvent
func callbackDoingList9733cb_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*std_gui.QDragEnterEvent))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *DoingList) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackDoingList9733cb_Edit
func callbackDoingList9733cb_Edit(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "edit"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(index))
	} else {
		NewDoingListFromPointer(ptr).EditDefault(std_core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *DoingList) EditDefault(index std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_EditDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index))
	}
}

//export callbackDoingList9733cb_EditorDestroyed
func callbackDoingList9733cb_EditorDestroyed(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "editorDestroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(editor))
	} else {
		NewDoingListFromPointer(ptr).EditorDestroyedDefault(std_core.NewQObjectFromPointer(editor))
	}
}

func (ptr *DoingList) EditorDestroyedDefault(editor std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_EditorDestroyedDefault(ptr.Pointer(), std_core.PointerFromQObject(editor))
	}
}

//export callbackDoingList9733cb_Entered
func callbackDoingList9733cb_Entered(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "entered"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(index))
	}

}

//export callbackDoingList9733cb_FocusInEvent
func callbackDoingList9733cb_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *DoingList) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackDoingList9733cb_FocusOutEvent
func callbackDoingList9733cb_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *DoingList) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackDoingList9733cb_IconSizeChanged
func callbackDoingList9733cb_IconSizeChanged(ptr unsafe.Pointer, size unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconSizeChanged"); signal != nil {
		signal.(func(*std_core.QSize))(std_core.NewQSizeFromPointer(size))
	}

}

//export callbackDoingList9733cb_InputMethodEvent
func callbackDoingList9733cb_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*std_gui.QInputMethodEvent))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *DoingList) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackDoingList9733cb_KeyPressEvent
func callbackDoingList9733cb_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *DoingList) KeyPressEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackDoingList9733cb_KeyboardSearch
func callbackDoingList9733cb_KeyboardSearch(ptr unsafe.Pointer, search C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "keyboardSearch"); signal != nil {
		signal.(func(string))(cGoUnpackString(search))
	} else {
		NewDoingListFromPointer(ptr).KeyboardSearchDefault(cGoUnpackString(search))
	}
}

func (ptr *DoingList) KeyboardSearchDefault(search string) {
	if ptr.Pointer() != nil {
		var searchC *C.char
		if search != "" {
			searchC = C.CString(search)
			defer C.free(unsafe.Pointer(searchC))
		}
		C.DoingList9733cb_KeyboardSearchDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: searchC, len: C.longlong(len(search))})
	}
}

//export callbackDoingList9733cb_MouseDoubleClickEvent
func callbackDoingList9733cb_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *DoingList) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackDoingList9733cb_MousePressEvent
func callbackDoingList9733cb_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *DoingList) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackDoingList9733cb_Pressed
func callbackDoingList9733cb_Pressed(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(index))
	}

}

//export callbackDoingList9733cb_Reset
func callbackDoingList9733cb_Reset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).ResetDefault()
	}
}

func (ptr *DoingList) ResetDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ResetDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_ScrollToBottom
func callbackDoingList9733cb_ScrollToBottom(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scrollToBottom"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).ScrollToBottomDefault()
	}
}

func (ptr *DoingList) ScrollToBottomDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ScrollToBottomDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_ScrollToTop
func callbackDoingList9733cb_ScrollToTop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scrollToTop"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).ScrollToTopDefault()
	}
}

func (ptr *DoingList) ScrollToTopDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ScrollToTopDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_SelectAll
func callbackDoingList9733cb_SelectAll(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectAll"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).SelectAllDefault()
	}
}

func (ptr *DoingList) SelectAllDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SelectAllDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_SetCurrentIndex
func callbackDoingList9733cb_SetCurrentIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setCurrentIndex"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(index))
	} else {
		NewDoingListFromPointer(ptr).SetCurrentIndexDefault(std_core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *DoingList) SetCurrentIndexDefault(index std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetCurrentIndexDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index))
	}
}

//export callbackDoingList9733cb_SetModel
func callbackDoingList9733cb_SetModel(ptr unsafe.Pointer, model unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setModel"); signal != nil {
		signal.(func(*std_core.QAbstractItemModel))(std_core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewDoingListFromPointer(ptr).SetModelDefault(std_core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *DoingList) SetModelDefault(model std_core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetModelDefault(ptr.Pointer(), std_core.PointerFromQAbstractItemModel(model))
	}
}

//export callbackDoingList9733cb_SetRootIndex
func callbackDoingList9733cb_SetRootIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setRootIndex"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(index))
	} else {
		NewDoingListFromPointer(ptr).SetRootIndexDefault(std_core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *DoingList) SetRootIndexDefault(index std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetRootIndexDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index))
	}
}

//export callbackDoingList9733cb_Update
func callbackDoingList9733cb_Update(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(index))
	} else {
		NewDoingListFromPointer(ptr).UpdateDefault(std_core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *DoingList) UpdateDefault(index std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_UpdateDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index))
	}
}

//export callbackDoingList9733cb_ViewportEntered
func callbackDoingList9733cb_ViewportEntered(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "viewportEntered"); signal != nil {
		signal.(func())()
	}

}

//export callbackDoingList9733cb_SelectionCommand
func callbackDoingList9733cb_SelectionCommand(ptr unsafe.Pointer, index unsafe.Pointer, event unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "selectionCommand"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex, *std_core.QEvent) std_core.QItemSelectionModel__SelectionFlag)(std_core.NewQModelIndexFromPointer(index), std_core.NewQEventFromPointer(event)))
	}

	return C.longlong(NewDoingListFromPointer(ptr).SelectionCommandDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQEventFromPointer(event)))
}

func (ptr *DoingList) SelectionCommandDefault(index std_core.QModelIndex_ITF, event std_core.QEvent_ITF) std_core.QItemSelectionModel__SelectionFlag {
	if ptr.Pointer() != nil {
		return std_core.QItemSelectionModel__SelectionFlag(C.DoingList9733cb_SelectionCommandDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQEvent(event)))
	}
	return 0
}

//export callbackDoingList9733cb_InputMethodQuery
func callbackDoingList9733cb_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewDoingListFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *DoingList) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.DoingList9733cb_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackDoingList9733cb_SizeHintForColumn
func callbackDoingList9733cb_SizeHintForColumn(ptr unsafe.Pointer, column C.int) C.int {
	if signal := qt.GetSignal(ptr, "sizeHintForColumn"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(column)))))
	}

	return C.int(int32(NewDoingListFromPointer(ptr).SizeHintForColumnDefault(int(int32(column)))))
}

func (ptr *DoingList) SizeHintForColumnDefault(column int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DoingList9733cb_SizeHintForColumnDefault(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

//export callbackDoingList9733cb_SizeHintForRow
func callbackDoingList9733cb_SizeHintForRow(ptr unsafe.Pointer, row C.int) C.int {
	if signal := qt.GetSignal(ptr, "sizeHintForRow"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(row)))))
	}

	return C.int(int32(NewDoingListFromPointer(ptr).SizeHintForRowDefault(int(int32(row)))))
}

func (ptr *DoingList) SizeHintForRowDefault(row int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DoingList9733cb_SizeHintForRowDefault(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

//export callbackDoingList9733cb_ContextMenuEvent
func callbackDoingList9733cb_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*std_gui.QContextMenuEvent))(std_gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewDoingListFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *DoingList) ContextMenuEventDefault(e std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackDoingList9733cb_ScrollContentsBy
func callbackDoingList9733cb_ScrollContentsBy(ptr unsafe.Pointer, dx C.int, dy C.int) {
	if signal := qt.GetSignal(ptr, "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(int32(dx)), int(int32(dy)))
	} else {
		NewDoingListFromPointer(ptr).ScrollContentsByDefault(int(int32(dx)), int(int32(dy)))
	}
}

func (ptr *DoingList) ScrollContentsByDefault(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ScrollContentsByDefault(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

//export callbackDoingList9733cb_SetupViewport
func callbackDoingList9733cb_SetupViewport(ptr unsafe.Pointer, viewport unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setupViewport"); signal != nil {
		signal.(func(*std_widgets.QWidget))(std_widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewDoingListFromPointer(ptr).SetupViewportDefault(std_widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *DoingList) SetupViewportDefault(viewport std_widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetupViewportDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(viewport))
	}
}

//export callbackDoingList9733cb_MinimumSizeHint
func callbackDoingList9733cb_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewDoingListFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *DoingList) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.DoingList9733cb_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackDoingList9733cb_SizeHint
func callbackDoingList9733cb_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewDoingListFromPointer(ptr).SizeHintDefault())
}

func (ptr *DoingList) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.DoingList9733cb_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackDoingList9733cb_ChangeEvent
func callbackDoingList9733cb_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(ev))
	} else {
		NewDoingListFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(ev))
	}
}

func (ptr *DoingList) ChangeEventDefault(ev std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(ev))
	}
}

//export callbackDoingList9733cb_Close
func callbackDoingList9733cb_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewDoingListFromPointer(ptr).CloseDefault())))
}

func (ptr *DoingList) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.DoingList9733cb_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackDoingList9733cb_ActionEvent
func callbackDoingList9733cb_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*std_gui.QActionEvent))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *DoingList) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackDoingList9733cb_CloseEvent
func callbackDoingList9733cb_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*std_gui.QCloseEvent))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *DoingList) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackDoingList9733cb_CustomContextMenuRequested
func callbackDoingList9733cb_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*std_core.QPoint))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackDoingList9733cb_EnterEvent
func callbackDoingList9733cb_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *DoingList) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackDoingList9733cb_Hide
func callbackDoingList9733cb_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).HideDefault()
	}
}

func (ptr *DoingList) HideDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_HideDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_HideEvent
func callbackDoingList9733cb_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*std_gui.QHideEvent))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *DoingList) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackDoingList9733cb_KeyReleaseEvent
func callbackDoingList9733cb_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *DoingList) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackDoingList9733cb_LeaveEvent
func callbackDoingList9733cb_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *DoingList) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackDoingList9733cb_Lower
func callbackDoingList9733cb_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).LowerDefault()
	}
}

func (ptr *DoingList) LowerDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_LowerDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_MoveEvent
func callbackDoingList9733cb_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*std_gui.QMoveEvent))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *DoingList) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackDoingList9733cb_Raise
func callbackDoingList9733cb_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *DoingList) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_RaiseDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_Repaint
func callbackDoingList9733cb_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *DoingList) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_RepaintDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_SetDisabled
func callbackDoingList9733cb_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewDoingListFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *DoingList) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackDoingList9733cb_SetEnabled
func callbackDoingList9733cb_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewDoingListFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *DoingList) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackDoingList9733cb_SetFocus2
func callbackDoingList9733cb_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *DoingList) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_SetHidden
func callbackDoingList9733cb_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewDoingListFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *DoingList) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackDoingList9733cb_SetStyleSheet
func callbackDoingList9733cb_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewDoingListFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *DoingList) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.DoingList9733cb_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackDoingList9733cb_SetVisible
func callbackDoingList9733cb_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewDoingListFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *DoingList) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackDoingList9733cb_SetWindowModified
func callbackDoingList9733cb_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewDoingListFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *DoingList) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackDoingList9733cb_SetWindowTitle
func callbackDoingList9733cb_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewDoingListFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *DoingList) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.DoingList9733cb_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackDoingList9733cb_Show
func callbackDoingList9733cb_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).ShowDefault()
	}
}

func (ptr *DoingList) ShowDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ShowDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_ShowEvent
func callbackDoingList9733cb_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*std_gui.QShowEvent))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *DoingList) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackDoingList9733cb_ShowFullScreen
func callbackDoingList9733cb_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *DoingList) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_ShowMaximized
func callbackDoingList9733cb_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *DoingList) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_ShowMinimized
func callbackDoingList9733cb_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *DoingList) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_ShowNormal
func callbackDoingList9733cb_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *DoingList) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_TabletEvent
func callbackDoingList9733cb_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*std_gui.QTabletEvent))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *DoingList) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackDoingList9733cb_UpdateMicroFocus
func callbackDoingList9733cb_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *DoingList) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackDoingList9733cb_WindowIconChanged
func callbackDoingList9733cb_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*std_gui.QIcon))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackDoingList9733cb_WindowTitleChanged
func callbackDoingList9733cb_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackDoingList9733cb_PaintEngine
func callbackDoingList9733cb_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine(signal.(func() *std_gui.QPaintEngine)())
	}

	return std_gui.PointerFromQPaintEngine(NewDoingListFromPointer(ptr).PaintEngineDefault())
}

func (ptr *DoingList) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.DoingList9733cb_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackDoingList9733cb_HasHeightForWidth
func callbackDoingList9733cb_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewDoingListFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *DoingList) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.DoingList9733cb_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackDoingList9733cb_HeightForWidth
func callbackDoingList9733cb_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewDoingListFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *DoingList) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DoingList9733cb_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackDoingList9733cb_Metric
func callbackDoingList9733cb_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(std_gui.QPaintDevice__PaintDeviceMetric) int)(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewDoingListFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *DoingList) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.DoingList9733cb_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackDoingList9733cb_InitPainter
func callbackDoingList9733cb_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*std_gui.QPainter))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewDoingListFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *DoingList) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackDoingList9733cb_ChildEvent
func callbackDoingList9733cb_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *DoingList) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackDoingList9733cb_ConnectNotify
func callbackDoingList9733cb_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDoingListFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *DoingList) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDoingList9733cb_CustomEvent
func callbackDoingList9733cb_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewDoingListFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *DoingList) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackDoingList9733cb_DeleteLater
func callbackDoingList9733cb_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewDoingListFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *DoingList) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackDoingList9733cb_Destroyed
func callbackDoingList9733cb_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackDoingList9733cb_DisconnectNotify
func callbackDoingList9733cb_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDoingListFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *DoingList) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.DoingList9733cb_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDoingList9733cb_ObjectNameChanged
func callbackDoingList9733cb_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}
