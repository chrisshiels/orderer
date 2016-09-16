# 'mpg123.spec'.
# Chris Shiels.


Name:       mpg123
Version:    1.23.6
Release:    1.mecachis.fc24
BuildArch:  x86_64
Group:      Mecachis
License:    Author
Vendor:     Mecachis
URL:        http://www.mpg123.de/
Summary:    Fast console MPEG Audio Player and decoder library


Source0:    https://www.mpg123.de/download/mpg123-1.23.6.tar.bz2


BuildRequires:  alsa-lib-devel
Requires:       alsa-lib


%description
Real time MPEG 1.0/2.0/2.5 audio player/decoder for layers 1, 2 and 3
(most commonly MPEG 1.0 layer 3 aka MP3), as well as re-usable decoding and
output libraries.


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
/usr/local/include/*
/usr/local/lib/*.la
/usr/local/lib/*.so*
/usr/local/lib/pkgconfig/*
/usr/local/share/man/man1/*
%doc AUTHORS
%doc ChangeLog
%doc COPYING
%doc doc/
%doc INSTALL
%doc NEWS
%doc NEWS.libmpg123
%doc README
%doc TODO


%changelog
* Thu Jul 7 2016 Chris Shiels <chris@mecachis.net> 1.23.6-1.mecachis.fc24
- Initial release.
