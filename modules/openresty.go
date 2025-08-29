package modules

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/hunter007/gakki/goutils"
)

func configOpenresty(m *Module, debug bool) error {
	zlibMod := GetModule("zlib")
	pcreMod := GetModule("pcre")
	opensslMod := GetModule("openssl")
	ccOpt := `-DNGX_LUA_ABORT_AT_PANIC -I%s/include -I%s/include -I%s/include`
	ccOpt = fmt.Sprintf(ccOpt, zlibMod.Prefix(), pcreMod.Prefix(), opensslMod.Prefix())

	withCcOpt := `--with-cc-opt="-DAPISIX_RUNTIME_VER=%s %s"`
	withCcOpt = fmt.Sprintf(withCcOpt, m.version, ccOpt)

	ldOpt := `-L%s/lib -L%s/lib -L%s/lib -Wl,-rpath,%s/lib:%s/lib:%s/lib`
	ldOpt = fmt.Sprintf(ldOpt, zlibMod.Prefix(), pcreMod.Prefix(), opensslMod.Prefix(), zlibMod.Prefix(), pcreMod.Prefix(), opensslMod.Prefix())
	// TODO: wasmtime-c-api dir
	withLdOpt := `--with-ld-opt="-Wl,-rpath,%s/wasmtime-c-api/lib %s"`
	withLdOpt = fmt.Sprintf(withLdOpt, m.Prefix(), ldOpt)

	debugArgs := ""
	if debug {
		debugArgs = "--with-debug"
	}

	nginxModules := []string{
		"--add-module=../mod_dubbo-1.0.2",
		"--add-module=../ngx_multi_upstream_module-1.3.2",
		"--add-module=../apisix-nginx-module-1.19.2",
		"--add-module=../apisix-nginx-module-1.19.2/src/stream",
		"--add-module=../apisix-nginx-module-1.19.2/src/meta",
		"--add-module=../wasm_nginx_module-0.7.0",
		"--add-module=../lua-var-nginx-module-0.5.3",
		"--add-module=../lua-resty-events-0.3.1",
		"--with-poll_module",
		"--with-pcre-jit",
		"--without-http_rds_json_module",
		"--without-http_rds_csv_module",
		"--without-lua_rds_parser",
		"--with-stream",
		"--with-stream_ssl_module",
		"--with-stream_ssl_preread_module",
		"--with-http_v2_module",
		"--with-http_v3_module",
		"--without-mail_pop3_module",
		"--without-mail_imap_module",
		"--without-mail_smtp_module",
		"--with-http_stub_status_module",
		"--with-http_realip_module",
		"--with-http_addition_module",
		"--with-http_auth_request_module",
		"--with-http_secure_link_module",
		"--with-http_random_index_module",
		"--with-http_gzip_static_module",
		"--with-http_sub_module",
		"--with-http_dav_module",
		"--with-http_flv_module",
		"--with-http_mp4_module",
		"--with-http_gunzip_module",
		"--with-threads",
		"--with-compat",
		"--with-luajit-xcflags=-DLUAJIT_NUMMODE=2 -DLUAJIT_ENABLE_LUA52COMPAT",
	}

	cmd := exec.Command("./configure", "--prefix="+m.Prefix(), withCcOpt, withLdOpt, debugArgs, strings.Join(nginxModules, " "), "-j", strconv.Itoa(int(goutils.Nproc())))

	cmd.Dir = m.Dir(m.version)
	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error(string(output))
		return err
	}
	return nil
}

func mkOpenRestyDir(m *Module) error {
	f, err := os.Stat(m.Prefix())
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			if err = os.Mkdir(m.Prefix(), 0o755); err != nil {
				return err
			}
		}
		return err
	}
	if !f.IsDir() {
		return fmt.Errorf("%s is not directory", m.Prefix())
	}
	return nil
}

func installRestyEvents(openrestyPrefix string) error {
	eventsLuaDir := fmt.Sprintf("%s%c%s%c%s%c", openrestyPrefix, os.PathSeparator, "lualib", os.PathSeparator, "resty", os.PathSeparator, "events")

	eventsMod := GetModule("lua-resty-events")

	cmd := exec.Command("sudo install -d", eventsLuaDir)
	cmd.Dir = eventsMod.Dir(eventsMod.version)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	cmd2 := exec.Command("sudo install -m 664 lualib/resty/events/*.lua", eventsLuaDir)
	cmd2.Dir = eventsMod.Dir(eventsMod.version)
	_, err = cmd2.Output()
	if err != nil {
		return err
	}

	compat := fmt.Sprintf("%s%c%s", eventsLuaDir, os.PathListSeparator, "compat")
	cmd3 := exec.Command("sudo install -d", compat)
	cmd3.Dir = eventsMod.Dir(eventsMod.version)
	_, err = cmd3.Output()
	if err != nil {
		return err
	}

	cmd4 := exec.Command("sudo install -m 664 lualib/resty/events/compat/*.lua", compat)
	cmd4.Dir = eventsMod.Dir(eventsMod.version)
	_, err = cmd4.Output()
	return err
}

func installOpenresty(m *Module) error {
	if err := mkOpenRestyDir(m); err != nil {
		return err
	}
	if err := configOpenresty(m, false); err != nil {
		return err
	}

	cmd := exec.Command("make", "-j", strconv.Itoa(int(goutils.Nproc())))
	cmd.Dir = m.Dir(m.version)
	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error(string(output))
		return err
	}

	cmdInstall := exec.Command("sudo make install")
	cmdInstall.Dir = m.Dir(m.version)
	output, err = cmdInstall.CombinedOutput()
	if err != nil {
		slog.Error(string(output))
		return err
	}

	installRestyEvents(m.Prefix())

	apisixMod := GetModule("apisix_nginx_module")
	if err := apisixMod.Install(apisixMod); err != nil {
		return err
	}

	return nil
}

func setupOpenresty() {
	module := &Module{
		name:   "openresty",
		prefix: "/usr/local/openresty",
		validVersions: []string{
			"1.27.1.2",
			"1.25.3.2",
			"1.25.3.1",
			"1.21.4.4",
			"1.21.4.3",
			"1.21.4.2",
			"1.21.4.1",
			"1.19.9.2",
			"1.19.9.1",
			"1.19.3.2",
			"1.19.3.1",
		},
		downloadTemplate: "https://openresty.org/download/openresty-%s.tar.gz",
		Install:          installOpenresty,
		version:          "1.27.1.2",
	}

	all[module.name] = module
	module.AddDependence(all["pcre2"])
	module.AddDependence(all["zlib"])
}
