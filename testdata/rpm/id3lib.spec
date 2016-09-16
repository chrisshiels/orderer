# 'id3lib.spec'.
# Chris Shiels.


Name:       id3lib
Version:    3.8.3
Release:    1.mecachis.fc24
BuildArch:  x86_64
Group:      Mecachis
License:    GPLv2
Vendor:     Mecachis
URL:        http://id3lib.sourceforge.net/
Summary:    Library for reading, writing, and manipulating ID3v1 and ID3v2 tags


Source0:    https://sourceforge.net/projects/id3lib/files/id3lib/3.8.3/id3lib-3.8.3.tar.gz


BuildRequires:  compat-gcc-34-c++


%description
id3lib is an open-source, cross-platform software development library for
reading, writing, and manipulating ID3v1 and ID3v2 tags.  It is an on-going
project whose primary goals are full compliance with the ID3v2 standard,
portability across several platforms, and providing a powerful and feature-rich
API with a highly stable and efficient implementation.


%define debug_package %{nil}


%prep
%setup


%build
CXX=g++34 \
	./configure \
		--prefix=/usr/local
make %{?_smp_mflags}


%install
make install DESTDIR=%{buildroot}
# Remove id3lib bundled libz files so these don't clash with the system zlib.
rm -v \
	%{buildroot}/usr/local/lib/libz.a \
	%{buildroot}/usr/local/lib/libz.la \
	%{buildroot}/usr/local/lib/libz.so \
	%{buildroot}/usr/local/lib/libz.so.1 \
	%{buildroot}/usr/local/lib/libz.so.1.0.0


%clean
rm -rf %{buildroot}


%files
%defattr(-, root, root, -)
/usr/local/bin/*
/usr/local/include/*
/usr/local/include/id3
/usr/local/lib/*.a
/usr/local/lib/*.la
/usr/local/lib/*.so*
%doc AUTHORS
%doc ChangeLog
%doc COPYING
%doc doc
%doc examples
%doc HISTORY
%doc INSTALL
%doc NEWS
%doc README
%doc THANKS
%doc TODO


%changelog
* Thu Jul 7 2016 Chris Shiels <chris@mecachis.net> 3.8.3-1.mecachis.fc24
- Initial release.
