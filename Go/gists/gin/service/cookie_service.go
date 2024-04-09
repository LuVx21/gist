package service

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/sha1"
    "database/sql"
    "errors"
    "fmt"
    "golang.org/x/crypto/pbkdf2"
    lcommon "luvx/common"
    "luvx/gin/common"
    "luvx/gin/db"
    _ "modernc.org/sqlite"
    "os/exec"
)

var (
    passwordByte = []byte("")
    client       *sql.DB
)

func GetCookieStrByHost(hosts ...string) string {
    resultMap := GetCookieByHost(hosts...)
    var result string
    for k, v := range resultMap {
        result += k + "=" + v
    }
    return result
}

func GetCookieByHost(hosts ...string) map[string]string {
    if client == nil {
        home, _ := common.Dir()
        client, _ = db.GetDataSource(home + "/data/sqlite/Cookies")
    }

    sql := `
select name, encrypted_value, host_key, path, creation_utc, expires_utc, is_secure, is_httponly, has_expires, is_persistent
from cookies
where true
and host_key in (%s)
order by host_key, name
-- limit 1
;
`
    var args string
    for i := 0; i < len(hosts); i++ {
        args += lcommon.IfThen(i == 0, "", ", ") + fmt.Sprintf("'%s'", hosts[i])
    }
    sql = fmt.Sprintf(sql, args)

    key := masterKey()

    result := make(map[string]string)
    rows, _ := client.Query(sql)
    defer rows.Close()
    for rows.Next() {
        var (
            name, host, path                              string
            isSecure, isHTTPOnly, hasExpire, isPersistent int
            createDate, expireDate                        int64
            encryptValue                                  []byte
        )
        if err := rows.Scan(&name, &encryptValue, &host, &path, &createDate, &expireDate, &isSecure, &isHTTPOnly, &hasExpire, &isPersistent); err != nil {
        }

        value, _ := DecryptWithChromium(key, encryptValue)
        //fmt.Println(host, "|", name, "|", string(value))
        result[name] = string(value)
    }
    return result
}

func masterKey() []byte {
    if len(passwordByte) == 0 {
        var (
            stdout, stderr bytes.Buffer
        )
        cmd := exec.Command("security", "find-generic-password", "-wa", "Microsoft Edge")
        cmd.Stdout = &stdout
        cmd.Stderr = &stderr
        if err := cmd.Run(); err != nil {
            fmt.Println(cmd.String())
            //fmt.Printf("run security command failed: %w, message %s", err, stderr.String())
        }
        passwordByte = stdout.Bytes()
    }

    secret := bytes.TrimSpace(passwordByte)
    key := pbkdf2.Key(secret, []byte("saltysalt"), 1003, 16, sha1.New)
    return key
}

func DecryptWithChromium(key, password []byte) ([]byte, error) {
    if len(password) <= 3 {
        return nil, errors.New("ciphertext length is invalid")
    }
    var iv = []byte{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32}
    return AES128CBCDecrypt(key, iv, password[3:])
}

func AES128CBCDecrypt(key, iv, ciphertext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    // Check ciphertext length
    if len(ciphertext) < aes.BlockSize {
        return nil, errors.New("AES128CBCDecrypt: ciphertext too short")
    }
    if len(ciphertext)%aes.BlockSize != 0 {
        return nil, errors.New("AES128CBCDecrypt: ciphertext is not a multiple of the block size")
    }

    decryptedData := make([]byte, len(ciphertext))
    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(decryptedData, ciphertext)

    decryptedData, err = pkcs5UnPadding(decryptedData)
    if err != nil {
        return nil, fmt.Errorf("AES128CBCDecrypt: %w", err)
    }

    return decryptedData, nil
}

func pkcs5UnPadding(src []byte) ([]byte, error) {
    length := len(src)
    if length == 0 {
        return nil, errors.New("pkcs5UnPadding: src should not be empty")
    }
    padding := int(src[length-1])
    if padding < 1 || padding > aes.BlockSize {
        return nil, errors.New("pkcs5UnPadding: invalid padding size")
    }
    return src[:length-padding], nil
}
