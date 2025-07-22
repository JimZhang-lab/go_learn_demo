/*
 * @Author: JimZhang
 * @Date: 2025-07-22 12:06:32
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-22 12:27:08
 * @FilePath: /server/utils/crypto.go
 * @Description:
 *
 */
package utils

import "golang.org/x/crypto/bcrypt"

// Encrpypt 加密函数，用于对输入的字符串进行 bcrypt 加密
// 参数:
//
//	stText: 需要加密的明文字符串
//
// 返回值:
//
//	string: 加密后的字符串
//	error: 如果加密过程中出现错误，返回相应的错误信息
func Encrpypt(stText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(stText), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), err
}

// CompareHashAndPassword 用于比较加密后的哈希值和原始文本是否匹配
// 参数:
//
//	stHash: 加密后的哈希字符串
//	stText: 用于比较的原始明文字符串
//
// 返回值:
//
//	error: 如果比较过程中出现错误，返回相应的错误信息；如果匹配成功，返回 nil
func CompareHashAndPassword(stHash, stText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(stHash), []byte(stText))
	return err == nil
}
