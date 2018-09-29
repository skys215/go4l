# -*- mode: python -*-

block_cipher = None


a = Analysis(['main.py'],
    pathex=['./'],
    binaries=[],
    datas=[( '/Users/pires/codes/pyqt/B4L/images/logo.png','./')],
    hiddenimports=[],
    hookspath=[],
    runtime_hooks=[],
    excludes=[],
    win_no_prefer_redirects=False,
    win_private_assemblies=False,
    cipher=block_cipher)
pyz = PYZ(a.pure, a.zipped_data,
             cipher=block_cipher)
exe = EXE(pyz,
     a.scripts,
     a.binaries,
     a.zipfiles,
     a.datas,
     name='b4l',
     debug=False,
     strip=False,
     upx=False,
     console=False , icon='/Users/pires/codes/pyqt/B4L/images/logo.ico', resources=[])

app = BUNDLE(exe,
        name='B4L.app',
	    bundle_identifier='com.skys215.b4l',
        info_plist={
            'CFBundleVersion' : '1.0.0',
        }
)
