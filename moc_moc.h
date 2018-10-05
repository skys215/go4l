/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.11.1)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.11.1. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_DoingList9733cb_t {
    QByteArrayData data[1];
    char stringdata0[16];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_DoingList9733cb_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_DoingList9733cb_t qt_meta_stringdata_DoingList9733cb = {
    {
QT_MOC_LITERAL(0, 0, 15) // "DoingList9733cb"

    },
    "DoingList9733cb"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_DoingList9733cb[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       0,    0, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

       0        // eod
};

void DoingList9733cb::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    Q_UNUSED(_o);
    Q_UNUSED(_id);
    Q_UNUSED(_c);
    Q_UNUSED(_a);
}

QT_INIT_METAOBJECT const QMetaObject DoingList9733cb::staticMetaObject = {
    { &QListWidget::staticMetaObject, qt_meta_stringdata_DoingList9733cb.data,
      qt_meta_data_DoingList9733cb,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *DoingList9733cb::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *DoingList9733cb::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_DoingList9733cb.stringdata0))
        return static_cast<void*>(this);
    return QListWidget::qt_metacast(_clname);
}

int DoingList9733cb::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QListWidget::qt_metacall(_c, _id, _a);
    return _id;
}
struct qt_meta_stringdata_Ui_Dialog9733cb_t {
    QByteArrayData data[7];
    char stringdata0[60];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_Ui_Dialog9733cb_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_Ui_Dialog9733cb_t qt_meta_stringdata_Ui_Dialog9733cb = {
    {
QT_MOC_LITERAL(0, 0, 15), // "Ui_Dialog9733cb"
QT_MOC_LITERAL(1, 16, 13), // "activateEvent"
QT_MOC_LITERAL(2, 30, 0), // ""
QT_MOC_LITERAL(3, 31, 13), // "QApplication*"
QT_MOC_LITERAL(4, 45, 2), // "v0"
QT_MOC_LITERAL(5, 48, 8), // "QDialog*"
QT_MOC_LITERAL(6, 57, 2) // "v1"

    },
    "Ui_Dialog9733cb\0activateEvent\0\0"
    "QApplication*\0v0\0QDialog*\0v1"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_Ui_Dialog9733cb[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       1,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    2,   19,    2, 0x0a /* Public */,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 3, 0x80000000 | 5,    4,    6,

       0        // eod
};

void Ui_Dialog9733cb::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        Ui_Dialog9733cb *_t = static_cast<Ui_Dialog9733cb *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->activateEvent((*reinterpret_cast< QApplication*(*)>(_a[1])),(*reinterpret_cast< QDialog*(*)>(_a[2]))); break;
        default: ;
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject Ui_Dialog9733cb::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_Ui_Dialog9733cb.data,
      qt_meta_data_Ui_Dialog9733cb,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *Ui_Dialog9733cb::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *Ui_Dialog9733cb::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_Ui_Dialog9733cb.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int Ui_Dialog9733cb::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 1)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 1;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 1)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 1;
    }
    return _id;
}
struct qt_meta_stringdata_Ui_SysTray9733cb_t {
    QByteArrayData data[1];
    char stringdata0[17];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_Ui_SysTray9733cb_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_Ui_SysTray9733cb_t qt_meta_stringdata_Ui_SysTray9733cb = {
    {
QT_MOC_LITERAL(0, 0, 16) // "Ui_SysTray9733cb"

    },
    "Ui_SysTray9733cb"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_Ui_SysTray9733cb[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       0,    0, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

       0        // eod
};

void Ui_SysTray9733cb::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    Q_UNUSED(_o);
    Q_UNUSED(_id);
    Q_UNUSED(_c);
    Q_UNUSED(_a);
}

QT_INIT_METAOBJECT const QMetaObject Ui_SysTray9733cb::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_Ui_SysTray9733cb.data,
      qt_meta_data_Ui_SysTray9733cb,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *Ui_SysTray9733cb::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *Ui_SysTray9733cb::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_Ui_SysTray9733cb.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int Ui_SysTray9733cb::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    return _id;
}
struct qt_meta_stringdata_AppDoingList9733cb_t {
    QByteArrayData data[10];
    char stringdata0[104];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_AppDoingList9733cb_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_AppDoingList9733cb_t qt_meta_stringdata_AppDoingList9733cb = {
    {
QT_MOC_LITERAL(0, 0, 18), // "AppDoingList9733cb"
QT_MOC_LITERAL(1, 19, 7), // "quitApp"
QT_MOC_LITERAL(2, 27, 0), // ""
QT_MOC_LITERAL(3, 28, 16), // "Ui_Dialog9733cb*"
QT_MOC_LITERAL(4, 45, 2), // "ui"
QT_MOC_LITERAL(5, 48, 19), // "AppDoingList9733cb*"
QT_MOC_LITERAL(6, 68, 3), // "app"
QT_MOC_LITERAL(7, 72, 8), // "QDialog*"
QT_MOC_LITERAL(8, 81, 10), // "mainDialog"
QT_MOC_LITERAL(9, 92, 11) // "activateApp"

    },
    "AppDoingList9733cb\0quitApp\0\0"
    "Ui_Dialog9733cb*\0ui\0AppDoingList9733cb*\0"
    "app\0QDialog*\0mainDialog\0activateApp"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_AppDoingList9733cb[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       2,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    3,   24,    2, 0x06 /* Public */,
       9,    3,   31,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3, 0x80000000 | 5, 0x80000000 | 7,    4,    6,    8,
    QMetaType::Void, 0x80000000 | 3, 0x80000000 | 5, 0x80000000 | 7,    4,    6,    8,

       0        // eod
};

void AppDoingList9733cb::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        AppDoingList9733cb *_t = static_cast<AppDoingList9733cb *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->quitApp((*reinterpret_cast< Ui_Dialog9733cb*(*)>(_a[1])),(*reinterpret_cast< AppDoingList9733cb*(*)>(_a[2])),(*reinterpret_cast< QDialog*(*)>(_a[3]))); break;
        case 1: _t->activateApp((*reinterpret_cast< Ui_Dialog9733cb*(*)>(_a[1])),(*reinterpret_cast< AppDoingList9733cb*(*)>(_a[2])),(*reinterpret_cast< QDialog*(*)>(_a[3]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 0:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 1:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< AppDoingList9733cb* >(); break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< Ui_Dialog9733cb* >(); break;
            }
            break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 1:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< AppDoingList9733cb* >(); break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< Ui_Dialog9733cb* >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (AppDoingList9733cb::*)(Ui_Dialog9733cb * , AppDoingList9733cb * , QDialog * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AppDoingList9733cb::quitApp)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (AppDoingList9733cb::*)(Ui_Dialog9733cb * , AppDoingList9733cb * , QDialog * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AppDoingList9733cb::activateApp)) {
                *result = 1;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject AppDoingList9733cb::staticMetaObject = {
    { &QApplication::staticMetaObject, qt_meta_stringdata_AppDoingList9733cb.data,
      qt_meta_data_AppDoingList9733cb,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *AppDoingList9733cb::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *AppDoingList9733cb::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_AppDoingList9733cb.stringdata0))
        return static_cast<void*>(this);
    return QApplication::qt_metacast(_clname);
}

int AppDoingList9733cb::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QApplication::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 2)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 2)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    }
    return _id;
}

// SIGNAL 0
void AppDoingList9733cb::quitApp(Ui_Dialog9733cb * _t1, AppDoingList9733cb * _t2, QDialog * _t3)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)), const_cast<void*>(reinterpret_cast<const void*>(&_t3)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void AppDoingList9733cb::activateApp(Ui_Dialog9733cb * _t1, AppDoingList9733cb * _t2, QDialog * _t3)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)), const_cast<void*>(reinterpret_cast<const void*>(&_t3)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
