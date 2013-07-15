/*
 * Copyright (C) 2013  Alexey Gladkov <gladkov.alexey@gmail.com>
 *
 * This file is covered by the GNU General Public License,
 * which should be included with go-passwdqc as the file COPYING.
 */
package passwdqc

/*
#cgo LDFLAGS: -lpasswdqc

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/param.h>
#include <passwdqc.h>

#define MESSAGE_GENERATION_FAILED "Unable to generate a passphrase"
#define MESSAGE_INVALID_CONFIG    "Invalid config value or config file format"
#define MESSAGE_NOMEMORY          "Out of memory"

static const char _empty[] = "";

static const char *
_go_passwdqc_check(const char *login, const char *newpass, const char *oldpass, const char *cfg)
{
	int ac = 0;
	char *av[] = { NULL, NULL };

	char *reason;
	struct passwd pw;

	passwdqc_params_t params;
	passwdqc_params_reset(&params);

	if (cfg && strlen(cfg) > 0) {
		char config[MAXPATHLEN + 7];
		sprintf(config, "config=%s", cfg);
		av[ac++] = config;
	}

	if (passwdqc_params_parse(&params, &reason, ac, (const char *const *)av)) {
		return (reason ? MESSAGE_INVALID_CONFIG : MESSAGE_NOMEMORY);
	}

	memset(&pw, 0, sizeof(struct passwd));

	pw.pw_name  = (char *)login;
	pw.pw_gecos = (char *)_empty;
	pw.pw_dir   = (char *)_empty;

	return passwdqc_check(&params.qc, newpass, oldpass, &pw);
}

static char *
_go_passwdqc_random(const char *cfg, char **error)
{
	int ac = 0;
	char *av[] = { NULL, NULL };
	char *reason, *result;

	passwdqc_params_t params;
	passwdqc_params_reset(&params);

	if (cfg && strlen(cfg) > 0) {
		char config[MAXPATHLEN + 7];
		sprintf(config, "config=%s", cfg);
		av[ac++] = config;
	}

	if (passwdqc_params_parse(&params, &reason, ac, (const char *const *)av)) {
		*error = (reason ? MESSAGE_INVALID_CONFIG : MESSAGE_NOMEMORY);
		return NULL;
	}

	result = passwdqc_random(&params.qc);
	if (!result) {
		*error = MESSAGE_GENERATION_FAILED;
	}
	return result;
}

*/
import "C"
import "unsafe"
import "errors"

func CheckCustom(login, newpass, oldpass, config string) error {
	var pw_config *C.char

	if len(config) > 0 {
		pw_config = C.CString(config)
		defer C.free(unsafe.Pointer(pw_config))
	}

	/* if memory allocation failed:
	 * panic: runtime error: invalid memory address or nil pointer dereference
	 */
	pw_login   := C.CString(login)
	pw_newpass := C.CString(newpass)
	pw_oldpass := C.CString(oldpass)

	defer C.free(unsafe.Pointer(pw_login))
	defer C.free(unsafe.Pointer(pw_newpass))
	defer C.free(unsafe.Pointer(pw_oldpass))

	reason := C._go_passwdqc_check(pw_login, pw_newpass, pw_oldpass, pw_config)

	if reason != nil {
		return errors.New(C.GoString(reason))
	}
	return nil
}

func GenerateCustom(config string) (string, error) {
	var pw_config *C.char
	var pw_error  *C.char

	if len(config) > 0 {
		pw_config = C.CString(config)
		defer C.free(unsafe.Pointer(pw_config))
	}

	res := C._go_passwdqc_random(pw_config, &pw_error);
	defer C.free(unsafe.Pointer(res))

	if res != nil {
		return C.GoString(res), nil
	}

	if pw_error != nil {
		return "", errors.New(C.GoString(pw_error))
	}

	return "", errors.New("Failed to generate a passphrase")
}

func CheckPassword(newpass string) error {
	return CheckCustom("qwerty", newpass, "", "")
}

func CheckLoginPassword(login, newpass string) error {
	return CheckCustom(login, newpass, "", "")
}

func CheckAccount(login, newpass, oldpass string) error {
	return CheckCustom(login, newpass, oldpass, "")
}

func Generate() (string, error) {
	return GenerateCustom("");
}
