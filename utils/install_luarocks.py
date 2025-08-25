import os

openresty_prefix = "/usr/local/openresty"
openresty_prefix = "/opt/homebrew"
luarocks_ver = "3.12.0"
WITH_LUA_OPT = f"--with-lua=${openresty_prefix}/luajit"
openresty_bin = f"${openresty_prefix}/bin/openresty"
openssl_prefix = "/opt/homebrew/Cellar/openssl@3/3.5.1"

def install_luarocks():
    '''
    By default, `make install' will install all the files in `/usr/local', `/usr/local/lib' etc.
    '''

    path = "./deps"
    script = f'''
    cd {path}
    tar xzvf v3.12.0.tar.gz
    cd luarocks-3.12.0
    # 使用openresty使用的luajit
    ./configure --with-lua={openresty_prefix}/luajit --prefix=/usr/local
    make build
    sudo make install
    mkdir ~/.luarocks
    /usr/local/bin/luarocks config variables.OPENSSL_LIBDIR {openssl_prefix}/lib
    /usr/local/bin/luarocks config variables.OPENSSL_INCDIR {openssl_prefix}/include
    /usr/local/bin/luarocks config variables.YAML_DIR /usr
    cd ..
    rm -rf luarocks-3.12.0
    '''
    print(script)

install_luarocks()
