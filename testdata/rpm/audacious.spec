# 'audacious.spec'.
# Chris Shiels.


Name:       audacious
Version:    3.7.2gtk3
Release:    1.mecachis.fc24
BuildArch:  x86_64
Group:      Mecachis
License:    Author
Vendor:     Mecachis
URL:        http://audacious-media-player.org/
Summary:    Open source audio player descended from xmms


Source0:    http://distfiles.audacious-media-player.org/audacious-3.7.2-gtk3.tar.bz2


BuildRequires:  alsa-lib-devel
BuildRequires:  glib2-devel
BuildRequires:  gtk3-devel
BuildRequires:  gettext
Requires:       alsa-lib
Requires:       glib2
Requires:       gtk3


%description
Audacious is an open source audio player.  A descendant of XMMS, Audacious
plays your music how you want it, without stealing away your computer’s
resources from other tasks.  Drag and drop folders and individual song files,
search for artists and albums in your entire music library, or create and edit
your own custom playlists.  Listen to CD’s or stream music from the Internet.
Tweak the sound with the graphical equalizer or experiment with LADSPA effects.
Enjoy the modern GTK-themed interface or change things up with Winamp Classic
skins.  Use the plugins included with Audacious to fetch lyrics for your music,
to set an alarm in the morning, and more.


%prep
%setup -n audacious-3.7.2-gtk3


%build
LDFLAGS=-Wl,-rpath=/usr/local/lib \
	./configure \
		--prefix=/usr/local
make %{?_smp_mflags}


%install
make install DESTDIR=%{buildroot}


%clean
rm -rf %{buildroot}


%files
%defattr(-, root, root, -)
/usr/local/bin/*
/usr/local/include/audacious
/usr/local/include/libaudcore
/usr/local/include/libaudgui
/usr/local/lib/*.so*
/usr/local/lib/pkgconfig/*
/usr/local/share/applications/*
/usr/local/share/audacious
/usr/local/share/icons/hicolor/48x48/apps/*
/usr/local/share/icons/hicolor/scalable/apps/*
/usr/local/share/locale/*/LC_MESSAGES/*
/usr/local/share/man/man1/*
%doc AUTHORS
%doc COPYING
%doc INSTALL


%changelog
* Thu Jul 7 2016 Chris Shiels <chris@mecachis.net> 3.7.2gtk3-1.mecachis.fc24
- Initial release.
