/*
 * @Author: sh06 hehe_0108@outlook.com
 * @Date: 2023-12-25 22:42:17
 * @LastEditors: sh06 hehe_0108@outlook.com
 * @LastEditTime: 2023-12-25 23:01:20
 * @FilePath: /kratos-demo/internal/pkg/util/crypto.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package util

import "golang.org/x/crypto/bcrypt"

// 加密密码
// @param pwd 要加密的密码
// @return 加密后的密码
func HashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// 验证密码
// @param hash 加密后的密码
// @param pwd 要验证的密码
// @return 验证结果
func VerifyPassword(hash, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd)); err != nil {
		return false
	}
	return true
}
