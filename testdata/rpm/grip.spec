# 'grip.spec'.
# Chris Shiels.


Name:       grip
Version:    3.2.0
Release:    1.mecachis.fc24
BuildArch:  x86_64
Group:      Mecachis
License:    GPLv2
Vendor:     Mecachis
URL:        https://sourceforge.net/projects/grip
Summary:    GTK-based CD-player and CD-ripper / MP3 encoder


Source0:    https://sourceforge.net/projects/grip/files/grip/3.2.0/grip-3.2.0.tar.gz


BuildRequires:  curl-devel
BuildRequires:  libgnomeui-devel
BuildRequires:  vte-devel
Requires:       curl
Requires:       libgnomeui
Requires:       vte


%description
Grip is a GTK-based CD-player and CD-ripper / MP3 encoder.  It has the ripping
capabilities of cdparanoia built in, but can also use external rippers (such as
cdda2wav).  Encoder presets are provided for lame, bladeenc, l3enc, xingmp3enc,
mp3encode, gogo).


%prep
%setup


%build
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
/usr/local/share/applications/*
/usr/local/share/gnome/help/grip
#/usr/local/share/locale/*/LC_MESSAGES/*
/usr/local/share/pixmaps/*
%doc ABOUT-NLS
%doc AUTHORS
%doc ChangeLog
%doc COPYING
%doc CREDITS
%doc doc
%doc INSTALL
%doc NEWS
%doc README
%doc TODO


%changelog
* Thu Jul 7 2016 Chris Shiels <chris@mecachis.net> 3.2.0-1.mecachis.fc24
- Initial release.
