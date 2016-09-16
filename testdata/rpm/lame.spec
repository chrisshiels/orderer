# 'lame.spec'.
# Chris Shiels.


Name:       lame
Version:    3.99.5
Release:    1.mecachis.fc24
BuildArch:  x86_64
Group:      Mecachis
License:    GPLv2
Vendor:     Mecachis
URL:        http://lame.sourceforge.net/
Summary:    High quality MPEG Audio Layer III (MP3) encoder


Source0:    https://sourceforge.net/projects/lame/files/lame/3.99/lame-3.99.5.tar.gz


%description
High quality MPEG Audio Layer III (MP3) encoder.


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
/usr/local/include/lame
/usr/local/lib/*.a
/usr/local/lib/*.la
/usr/local/lib/*.so*
/usr/local/share/doc/lame
/usr/local/share/man/man1/*
%doc API
%doc ChangeLog
%doc COPYING
%doc DEFINES
%doc doc
%doc HACKING
%doc INSTALL
%doc INSTALL.configure
%doc LICENSE
%doc README
%doc README.WINGTK
%doc STYLEGUIDE
%doc TODO
%doc USAGE


%changelog
* Thu Jul 7 2016 Chris Shiels <chris@mecachis.net> 3.99.5-1.mecachis.fc24
- Initial release.
